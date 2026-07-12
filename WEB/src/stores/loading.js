import { defineStore } from 'pinia'

/**
 * Loading store
 * Tracks the number of in-flight requests so multiple simultaneous calls
 * don't hide the overlay too early. Axios interceptors call start()/stop()
 * automatically — see services/api.js.
 *
 * Usage:
 *   const loading = useLoadingStore()
 *   loading.start()
 *   ...
 *   loading.stop()
 */
export const useLoadingStore = defineStore('loading', {
  state: () => ({
    requestCount: 0,
    fullScreen: true, 
  }),

  getters: {
    isLoading: (state) => state.requestCount > 0,
  },

  actions: {
    start() {
      this.requestCount++
    },
    stop() {
      this.requestCount = Math.max(0, this.requestCount - 1)
    },
    reset() {
      this.requestCount = 0
    },
  },
})
