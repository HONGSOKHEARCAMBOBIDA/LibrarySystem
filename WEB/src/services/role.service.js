import api from './api'

export default {

    async getrole() {
        const { data } = await api.get('/v1/role.view')
        return data.data
    },

    async updaterole(id, data) {
        const { data: res } = await api.put(`/v1/role.update/${id}`, data)
        return res
    },

    async getrolepermission(id) {
        const { data } = await api.get(`/v1/role.permission.view/${id}`)
        return data.data
    },

    async createrolepermission(payload) {
        const { data } = await api.post('/v1/role.permission.create', payload)
        return data
    },

    async deleterolepermission(payload) {
        const { data } = await api.delete('/v1/role.permission.delete', { data: payload })
        return data
    }
}