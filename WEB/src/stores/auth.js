import { defineStore } from 'pinia'
import authService from '@/services/auth.service'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: null,
    refreshToken: null,
    user: null, 
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
    userRoles: (state) => state.user?.roles || [],
    userPermissions: (state) => state.user?.permissions || [],
    userName: (state) => state.user?.name || 'Guest',
    userAvatar: (state) => state.user?.avatar || '',
  },
  
  actions: {
    async login(credentials) {
      const res = await authService.login(credentials)
      const user = {
        name: res.name,
        email: res.email,
        roles: res.roles,
        permissions: res.permissions,
      }
      this.setSession({ token: res.token, refreshToken: res.refreshToken, user })
      return user
    },

    async fetchProfile() {
      const user = await authService.getProfile()
      this.user = user
      return user
    },

    async refreshAccessToken() {
      const { token, refreshToken } = await authService.refreshToken(this.refreshToken)
      this.token = token
      if (refreshToken) this.refreshToken = refreshToken
      return token
    },

    setSession({ token, refreshToken, user }) {
      this.token = token
      this.refreshToken = refreshToken ?? this.refreshToken
      this.user = user ?? this.user
    },

    hasPermission(permission) {
      if (!permission) return true
      return this.userPermissions.includes(permission)
    },

    hasRole(role) {
      if (!role) return true
      return this.userRoles.includes(role) 
    },

    async logout() {
      try {
        await authService.logout()
      } catch {
      } finally {
        this.$reset()
      }
    },
  },

  persist: {
    key: 'admin-dashboard-auth',
    storage: localStorage,
  },
})
