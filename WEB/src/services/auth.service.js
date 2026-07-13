import api from './api'

export default {
  /**
   * POST /auth/login
   * @param {{email: string, password: string}} credentials
   * @param {{refresh_token: string}} refreshtoken
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
  async refreshToken(refreshtoken) {
    const { data } = await api.post('/v1/refresh', { refresh_token: refreshtoken })
    return data.data
  },

  /**
   * POST /v1/register
   * @param {{name_kh: string, name_en: string, role_id: number, gender: number, dob: string}} payload
   */
  async register(payload) {
    const { data } = await api.post('/v1/register', payload)
    return data.data
  },

  /**
   * GET /v1/user
   * @param {{page: number, pageSize: number, name?: string, role_id?: number|string}} params
   * @returns {Promise<{ list: any[], pagination: object }>}
   */
  async getUser(params) {
    const { data } = await api.get('/v1/user.view', { params })

    return {
      list: data.data.map((u) => ({
        id: u.id,
        name: u.name_en,
        nameKh: u.name_kh,
        role: u.role_name,
        roleId: u.role_id,
        gender: u.gender,
        dob: u.dob,
        status: u.is_active ? 'active' : 'suspended',
      })),
      pagination: data.pagination,
    }
  }
}
