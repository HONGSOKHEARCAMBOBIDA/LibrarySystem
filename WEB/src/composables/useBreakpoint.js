import { computed } from 'vue'
import { useWindowSize } from '@vueuse/core'
import { BREAKPOINTS } from '@/config/constants'

export function useBreakpoint() {
  const { width } = useWindowSize()

  const isMobile = computed(() => width.value <= BREAKPOINTS.mobile)
  const isTablet = computed(
    () =>
      width.value > BREAKPOINTS.mobile &&
      width.value <= BREAKPOINTS.tablet
  )
  const isDesktop = computed(() => width.value > BREAKPOINTS.tablet)

  return {
    width,
    isMobile,
    isTablet,
    isDesktop,
  }
}