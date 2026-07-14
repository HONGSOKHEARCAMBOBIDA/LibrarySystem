import api from './api'

export default {

  async login(credentials) {
    const { data } = await api.post('/v1/login', credentials)
    return data.data
  },

  async logout() {
    if (USE_MOCK) return mockDelay(true, 200)
    const { data } = await api.post('/auth/logout')
    return data
  },

  async getProfile() {
    if (USE_MOCK) return mockDelay(MOCK_USER, 300)
    const { data } = await api.get('/auth/profile')
    return data
  },

  async refreshToken(refreshtoken) {
    const { data } = await api.post('/v1/refresh', { refresh_token: refreshtoken })
    return data.data
  },


  async register(payload) {
    const { data } = await api.post('/v1/register', payload)
    return data.data
  },

  async update(id,data) {
    const {data: res} = await api.put(`/v1/user.update/${id}`,data)
    return res
  },

  async getUser(params) {
    const { data } = await api.get('/v1/user.view', { params })

    return {
      list: data.data.map((u) => ({
        id: u.id,
        name: u.name_en,
        nameKh: u.name_kh,
        email: u.email,
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
