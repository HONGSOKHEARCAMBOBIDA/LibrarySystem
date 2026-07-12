import { ElNotification, ElMessage } from 'element-plus'

/**
 * Toast / notification helper.
 * Use `notify.success('Saved!')` etc. anywhere in the app — no need to
 * import Element Plus directly in every component.
 */
export const notify = {
  success(message, title = 'Success') {
    ElNotification({ title, message, type: 'success', duration: 3000 })
  },
  error(message, title = 'Error') {
    ElNotification({ title, message, type: 'error', duration: 4000 })
  },
  warning(message, title = 'Warning') {
    ElNotification({ title, message, type: 'warning', duration: 3500 })
  },
  info(message, title = 'Info') {
    ElNotification({ title, message, type: 'info', duration: 3000 })
  },
  // Lightweight inline message variants (top-of-page, no title)
  message: {
    success: (msg) => ElMessage({ message: msg, type: 'success' }),
    error: (msg) => ElMessage({ message: msg, type: 'error' }),
    warning: (msg) => ElMessage({ message: msg, type: 'warning' }),
    info: (msg) => ElMessage({ message: msg, type: 'info' }),
  },
}
