<script setup>
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { getcategory, createcategory, updatecategory, togglestatuscategory } from '../../services/category.service'
import { useNotification } from '../../composables/useNotification'
import TableCustom from '../../components/tables/TableCustom.vue'
import AppButton from '../../components/button/AppButton.vue'
import AppInput from '../../components/input/AppInput.vue'
import AppSelect from '../../components/common/AppSelect.vue'
import AppFilterBar from '../../components/common/AppFilterBar.vue'
import { debounce } from 'lodash-es'
import AppDialog from '@/components/dialogs/AppDialog.vue'
import AppForm from '@/components/forms/AppForm.vue'

const notify = useNotification()

const categories = ref([])
const loading = ref(false)
const submitting = ref(false)
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

const dialogVisible = ref(false)
const isEditing = ref(false)
const editingId = ref(null)
const formRef = ref(null)

const dialogTitle = computed(() => (isEditing.value ? 'កែប្រែប្រភេទ' : 'បង្កេីតប្រភេទ'))

const form = reactive({
  name: '',
  description: '',
})

const rules = {
  name: [{ required: true, message: 'ត្រូវបញ្ចូលឈ្មោះ', trigger: 'blur' }],
}

const statusOptions = [
  { label: 'សកម្ម', value: 1 },
  { label: 'អសកម្ម', value: 0 },
]

const filters = reactive({ name: '', is_active: null })

const columns = [
  { prop: 'id', label: 'ល.រ', width: 70 },
  { prop: 'name', label: 'ឈ្មោះ' },
  { prop: 'description', label: 'ការពិពណ៌នា' },
  { prop: 'is_active', label: 'ស្ថានភាព', slot: 'isActive', width: 100 },
]

function resetForm() {
  Object.assign(form, {
    name: '',
    description: '',
  })
}

function openCreate() {
  isEditing.value = false
  editingId.value = null
  resetForm()
  dialogVisible.value = true
}

function openEdit(row) {
  isEditing.value = true
  editingId.value = row.id

  Object.assign(form, {
    name: row.name,
    description: row.description,
  })

  dialogVisible.value = true
}

function closeDialog() {
  dialogVisible.value = false
}

// keep the form in sync whenever the dialog closes, however it was closed
// (submit, cancel button, backdrop click, or the dialog's own close icon)
watch(dialogVisible, (visible) => {
  if (!visible) resetForm()
})

function buildParams() {
  const params = { page: page.value, pageSize: pageSize.value }
  if (filters.name) params.name = filters.name
  if (filters.is_active !== null && filters.is_active !== '') params.is_active = filters.is_active
  return params
}

async function fetchCategory() {
  loading.value = true
  try {
    const res = await getcategory(buildParams())
    categories.value = res.data.data || []
    total.value = res.data.data.length  
  } catch (e) {
    notify.error(e?.response?.data?.message || e.message || 'Failed to load categories')
  } finally {
    loading.value = false
  }
}

async function handleSubmit() {
  const payload = {
    name: form.name,
    description: form.description,
  }

  submitting.value = true
  try {
    if (isEditing.value) {
      await updatecategory(editingId.value, payload)
      notify.success('កែប្រែជោគជ័យ')
    } else {
      await createcategory(payload)
      notify.success('បង្កេីតជោគជ័យ')
    }
    dialogVisible.value = false
    await fetchCategory()
  } catch (e) {
    notify.error(e?.response?.data?.error || e.message || 'Failed to save category')
  } finally {
    submitting.value = false
  }
}

async function togggle(row) {
  try {
    await togglestatuscategory(row.id)
    notify.success('កែប្រែជោគជ័យ')
    fetchCategory()
  } catch (e) {
    notify.error(e?.response?.data?.error || e.message || 'Failed to save category')
  }
}

function onReset() {
  filters.name = ''
  filters.is_active = null
  page.value = 1
  fetchCategory()
}

const debouncedFetch = debounce(() => {
  page.value = 1 // reset to first page whenever filters change
  fetchCategory()
}, 500)

// watch the actual filter values, not a function reference
watch(() => ({ ...filters }), debouncedFetch, { deep: true })

onMounted(() => {
  fetchCategory()
})
</script>

<template>
  <div class="category-page">
    <AppFilterBar
      :fields="[
        { slot: 'name', span: 10 },
        { slot: 'status', span: 5 },
        { slot: 'add', span: 1 },
      ]"
      :action-span="4"
    >
      <template #name>
        <AppInput v-model="filters.name" placeholder="ស្វែងរក..." clearable />
      </template>

      <template #status>
        <AppSelect
          label="ជ្រេីសរេីសស្ថានភាព"
          v-model="filters.is_active"
          :options="statusOptions"
          placeholder="សកម្ម-មិនសកម្ម"
          clearable
        />
      </template>

      <template #add>
        <AppButton @click="openCreate" size="default">បន្ថែម</AppButton>
      </template>
    </AppFilterBar>

    <TableCustom
      :data="categories"
      :columns="columns"
      :loading="loading"
      :total="total"
      v-model:current-page="page"
      v-model:page-size="pageSize"
      @page-change="fetchCategory"
    >
      <template #isActive="{ row }">
        <el-tag :type="row.is_active ? 'success' : 'danger'">
          {{ row.is_active ? 'សកម្ម' : 'អសកម្ម' }}
        </el-tag>
      </template>

      <template #actions="{ row }">
        <el-tooltip content="កែប្រែ" placement="top">
          <AppButton icon="Edit" circle size="small" type="warning" @click="openEdit(row)" />
        </el-tooltip>
        <el-tooltip content="បិទ/បេីក" placement="top">
          <AppButton :icon="row.is_active ? 'CircleClose' : 'CircleCheck'" circle size="small" type="danger" @click="togggle(row)" />
        </el-tooltip>
      </template>
    </TableCustom>

    <AppDialog v-model:visible="dialogVisible" :title="dialogTitle" :showDefaultFooter="false">
      <AppForm
        ref="formRef"
        :model="form"
        :rules="rules"
        :loading="submitting"
        :show-actions="true"
        @submit="handleSubmit"
        @reset="closeDialog"
        submitText="រក្សាទុក"
        resetText="ចាកចេញ"
      >
        <AppInput v-model="form.name" placeholder="បញ្ចូលឈ្មោះ" clearable prop="name" label="ឈ្មោះ" />
        <AppInput
          v-model="form.description"
          placeholder="បញ្ចូលការពិពណ៌នា"
          clearable
          prop="description"
          label="ការពិពណ៌នា"
          type="textarea"
        />
      </AppForm>
    </AppDialog>
  </div>
</template>

<style scoped>
.category-filters {
  display: flex;
  gap: 10px;
  margin-bottom: 16px;
  flex-wrap: wrap;
  align-items: center;
}
</style>