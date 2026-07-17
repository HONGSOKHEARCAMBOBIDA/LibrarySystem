import api from "./api";

export const getcategory = () => api.get('/v1/category.view')
export const createcategory = (data) => api.post('/v1/category.create',data)
export const updatecategory = (id,data) => api.put(`/v1/category.update/${id}`,data)
export const togglestatuscategory = (id) => api.put(`/v1/category.toggle.status/${id}`)