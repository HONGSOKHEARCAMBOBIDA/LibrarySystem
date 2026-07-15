import api from "./api";
export const getauthor = (params) => api.get('/v1/author.view',{params})