import api from './api'

export default {
  /**
   * POST /auth/login
   * @param {{email: string, password: string}} credentials
   */
  async login(credentials) {
    const { data } = await api.post('/v1/login', credentials)
    return data.data
  },

  /** POST /auth/logout */
  async logout() {
    if (USE_MOCK) return mockDelay(true, 200)
    const { data } = await api.post('/auth/logout')
    return data
  },

  /** GET /auth/profile */
  async getProfile() {
    if (USE_MOCK) return mockDelay(MOCK_USER, 300)
    const { data } = await api.get('/auth/profile')
    return data
  },

  /** POST /auth/refresh */
  async refreshToken(refreshToken) {
    const { data } = await api.post('/v1/refresh', { refreshToken })
    return data
  },
}
