import { ref, reactive, watch } from 'vue'
import { DEFAULT_PAGE_SIZE } from '@/config/constants'

/**
 * useTable — shared logic for any list page (Users, Products, Orders, ...).
 *
 * @param {Function} fetcher async ({ page, pageSize, search, sort, filters }) => { rows, total }
 * @param {Object} options { immediate = true, pageSize = DEFAULT_PAGE_SIZE }
 *
 * Usage:
 *   const { rows, loading, pagination, search, reload } = useTable(fetchUsers)
 */
export function useTable(fetcher, options = {}) {
  const rows = ref([])
  const loading = ref(false)
  const error = ref(null)
  const search = ref('')

  const pagination = reactive({
    page: 1,
    pageSize: options.pageSize || DEFAULT_PAGE_SIZE,
    total: 0,
  })

  const sort = reactive({ prop: null, order: null }) // order: 'ascending' | 'descending'
  const filters = reactive({})

  async function load() {
    loading.value = true
    error.value = null
    try {
      const result = await fetcher({
        page: pagination.page,
        pageSize: pagination.pageSize,
        search: search.value,
        sort: { ...sort },
        filters: { ...filters },
      })
      rows.value = result.rows
      pagination.total = result.total
    } catch (err) {
      error.value = err
      rows.value = []
    } finally {
      loading.value = false
    }
  }

  function handleSortChange({ prop, order }) {
    sort.prop = prop
    sort.order = order
    pagination.page = 1
    load()
  }

  function handlePageChange(page) {
    pagination.page = page
    load()
  }

  function handlePageSizeChange(pageSize) {
    pagination.pageSize = pageSize
    pagination.page = 1
    load()
  }

  function setFilter(key, value) {
    filters[key] = value
    pagination.page = 1
    load()
  }

  function reload() {
    pagination.page = 1
    load()
  }

  // Debounced-ish: reload whenever the search text changes
  let searchTimer = null
  watch(search, () => {
    clearTimeout(searchTimer)
    searchTimer = setTimeout(() => {
      pagination.page = 1
      load()
    }, 350)
  })

  if (options.immediate !== false) {
    load()
  }

  return {
    rows,
    loading,
    error,
    search,
    pagination,
    sort,
    filters,
    load,
    reload,
    setFilter,
    handleSortChange,
    handlePageChange,
    handlePageSizeChange,
  }
}
