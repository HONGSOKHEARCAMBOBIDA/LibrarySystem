import { ElMessageBox } from 'element-plus'

/**
 * useConfirm — promise-based confirmation dialogs.
 *
 * const { confirm, confirmDelete } = useConfirm()
 * const ok = await confirm('Are you sure you want to publish this post?')
 * if (ok) { ... }
 */
export function useConfirm() {
  async function confirm(message, options = {}) {
    try {
      await ElMessageBox.confirm(message, options.title || 'Please Confirm', {
        confirmButtonText: options.confirmText || 'Confirm',
        cancelButtonText: options.cancelText || 'Cancel',
        type: options.type || 'warning',
        ...options,
      })
      return true
    } catch {
      return false
    }
  }

  function confirmDelete(itemName = 'this item') {
    return confirm(`This will permanently delete ${itemName}. This action cannot be undone.`, {
      title: 'Delete Confirmation',
      confirmText: 'Delete',
      type: 'error',
      confirmButtonClass: 'el-button--danger',
    })
  }

  return { confirm, confirmDelete }
}
