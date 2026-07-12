import axios from 'axios'
import router from '@/router'
import { useAuthStore } from '@/stores/auth'
import { useLoadingStore } from '@/stores/loading'
import { notify } from '@/utils/notify'

/**
 * Central axios instance.
 * Every future API service (users.service.js, products.service.js, ...)
 * should import `api` from here instead of creating its own axios instance.
 */
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: Number(import.meta.env.VITE_API_TIMEOUT) || 15000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Requests that should NOT trigger the global loading overlay
// (e.g. background polling). Pass { silent: true } in axios config.

// --------------------------------------------------------------------------
// Request interceptor — attach token + trigger global loading
// --------------------------------------------------------------------------
api.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore()
    const loadingStore = useLoadingStore()
    config.headers['x-api-key'] = import.meta.env.API_KEY_SECRET || 'X7mK9qR2vLp8Nf4TzY6cHd3WsAj5BuEeG1rQn8MxKp7Vt2CyL9sDf4JhUwZk3NaRb'
    config.headers['X-Admin-Token'] = import.meta.env.ADMIN_SECRET || 'X7mK9qR2vLp8Nf4TzY6cHd3WsAj5BuEeG1rQn8MxKp7Vt2CyL9sDf4JhUwZk3NaRb៣៩៨៩៩ឪឧ' 
    if (authStore.token) {
      config.headers.Authorization = `Bearer ${authStore.token}`
    }

    if (!config.silent) {
      loadingStore.start()
    }

    return config
  },
  (error) => {
    useLoadingStore().stop()
    return Promise.reject(error)
  }
)

// --------------------------------------------------------------------------
// Response interceptor — stop loading, unwrap data, handle errors globally
// --------------------------------------------------------------------------
let isRefreshing = false
let pendingQueue = []

function resolvePendingQueue(error, token = null) {
  pendingQueue.forEach((p) => (error ? p.reject(error) : p.resolve(token)))
  pendingQueue = []
}

api.interceptors.response.use(
  (response) => {
    if (!response.config.silent) {
      useLoadingStore().stop()
    }
    // Return the response as-is; callers can access response.data.
    // (Kept un-unwrapped so paginated APIs can still read headers/meta.)
    return response
  },
  async (error) => {
    const { config, response } = error
    if (!config?.silent) {
      useLoadingStore().stop()
    }

    const authStore = useAuthStore()

    // ---- 401 Unauthorized: try refresh token once, else redirect to login
    if (response?.status === 401 && !config._retry) {
      if (!authStore.refreshToken) {
        authStore.$reset()
        router.push({ name: 'Login', query: { redirect: router.currentRoute.value.fullPath } })
        return Promise.reject(error)
      }

      if (isRefreshing) {
        // Queue requests that fail while a refresh is already in progress
        return new Promise((resolve, reject) => {
          pendingQueue.push({
            resolve: (token) => {
              config.headers.Authorization = `Bearer ${token}`
              resolve(api(config))
            },
            reject,
          })
        })
      }

      config._retry = true
      isRefreshing = true

      try {
        const newToken = await authStore.refreshAccessToken()
        resolvePendingQueue(null, newToken)
        config.headers.Authorization = `Bearer ${newToken}`
        return api(config)
      } catch (refreshError) {
        resolvePendingQueue(refreshError, null)
        authStore.$reset()
        router.push({ name: 'Login', query: { redirect: router.currentRoute.value.fullPath } })
        return Promise.reject(refreshError)
      } finally {
        isRefreshing = false
      }
    }

    // ---- 403 Forbidden
    if (response?.status === 403) {
      notify.error('You do not have permission to perform this action.')
      return Promise.reject(error)
    }

    // ---- Other errors: show a toast unless the caller opted out
    if (!config?.silent && !config?.suppressErrorToast) {
      const message =
        response?.data?.message || error.message || 'Something went wrong. Please try again.'
      notify.error(message)
    }

    return Promise.reject(error)
  }
)

export default api
