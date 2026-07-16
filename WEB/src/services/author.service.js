import api from "./api";
export const getauthor = (params) => api.get('/v1/author.view',{params})
export const createauthor = (data) => api.post('/v1/author.create',data)
export const updateauthor = (data,id) => api.post(`/v1/author.create/${id}`,data)