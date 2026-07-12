import { defineStore } from 'pinia'

/**
 * App store
 * Holds global UI state that is not specific to any one page:
 * sidebar collapse state, current theme, and device breakpoint.
 */
export const useAppStore = defineStore('app', {
  state: () => ({
    sidebarCollapsed: false,
    sidebarDrawerOpen: false, // mobile drawer
    theme: 'light', // 'light' | 'dark'
    device: 'desktop', // 'desktop' | 'tablet' | 'mobile'
  }),

  actions: {
    toggleSidebar() {
      this.sidebarCollapsed = !this.sidebarCollapsed
    },
    setSidebarCollapsed(value) {
      this.sidebarCollapsed = value
    },
    toggleDrawer() {
      this.sidebarDrawerOpen = !this.sidebarDrawerOpen
    },
    closeDrawer() {
      this.sidebarDrawerOpen = false
    },
    setDevice(device) {
      this.device = device
      // Auto-collapse the sidebar on tablet, hide it (drawer mode) on mobile
      if (device === 'tablet') this.sidebarCollapsed = true
    },
    toggleTheme() {
      this.theme = this.theme === 'light' ? 'dark' : 'light'
      this.applyTheme()
    },
    setTheme(theme) {
      this.theme = theme
      this.applyTheme()
    },
    /** Sync `<html class="dark">` with the current theme value. */
    applyTheme() {
      const root = document.documentElement
      if (this.theme === 'dark') {
        root.classList.add('dark')
      } else {
        root.classList.remove('dark')
      }
    },
  },

  persist: {
    key: 'admin-dashboard-app',
    storage: localStorage,
    pick: ['theme', 'sidebarCollapsed'],
  },
})
