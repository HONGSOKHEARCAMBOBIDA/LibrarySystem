import { useAuthStore } from '@/stores/auth'
import { APP_TITLE } from '@/config/constants'
import { notify } from '@/utils/notify'

/**
 * Registers global navigation guards on the router instance:
 *  - requiresAuth: redirect unauthenticated users to /login
 *  - guestOnly: redirect already-authenticated users away from /login
 *  - permission: block access to routes the user lacks permission for
 *  - sets document.title from route meta
 */
export function setupGuards(router) {
  router.beforeEach((to, from, next) => {
    const authStore = useAuthStore()

    document.title = to.meta?.title ? `${to.meta.title} · ${APP_TITLE}` : APP_TITLE

    const requiresAuth = to.matched.some((record) => record.meta.requiresAuth)
    const guestOnly = to.matched.some((record) => record.meta.guestOnly)

    if (requiresAuth && !authStore.isAuthenticated) {
      next({ name: 'Login', query: { redirect: to.fullPath } })
      return
    }

    if (guestOnly && authStore.isAuthenticated) {
      next({ name: 'Dashboard' })
      return
    }

    // Permission guard — route.meta.permission is an optional string like
    // 'users.view'. Extend to check meta.role the same way if needed.
    const requiredPermission = to.matched.find((r) => r.meta.permission)?.meta.permission
    if (requiredPermission && !authStore.hasPermission(requiredPermission)) {
      notify.error("You don't have permission to access that page.")
      next(from.fullPath && from.name ? false : { name: 'Dashboard' })
      return
    }

    next()
  })
}
