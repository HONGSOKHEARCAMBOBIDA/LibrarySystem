import api from "./api";
export const getauthor = (params) => api.get('/v1/author.view',{params})
export const createauthor = (data) => api.post('/v1/author.create',data)
export const updateauthor = (id,data) => api.put(`/v1/author.update/${id}`,data)
export const togglestatusauthor = (id) => api.put(`/v1/author.toggle.status/${id}`)