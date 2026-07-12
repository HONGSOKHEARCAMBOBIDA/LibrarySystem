import { onMounted, onUnmounted } from 'vue'
import { useAppStore } from '@/stores/app'
import { BREAKPOINTS } from '@/config/constants'

/**
 * useDevice — watches window width and keeps app.device in sync so
 * DashboardLayout can switch between sidebar / drawer automatically.
 */
export function useDevice() {
  const appStore = useAppStore()

  function computeDevice() {
    const width = window.innerWidth
    if (width < BREAKPOINTS.mobile) return 'mobile'
    if (width < BREAKPOINTS.tablet) return 'tablet'
    return 'desktop'
  }

  function handleResize() {
    appStore.setDevice(computeDevice())
  }

  onMounted(() => {
    handleResize()
    window.addEventListener('resize', handleResize)
  })

  onUnmounted(() => {
    window.removeEventListener('resize', handleResize)
  })
}
