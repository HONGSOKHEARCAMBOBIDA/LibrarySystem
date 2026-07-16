<script setup>
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { getauthor, createauthor, updateauthor,togglestatusauthor } from '../../services/author.service'
import { getfaculty } from '../../services/faculty.service.js'
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

const authors = ref([])
const faculty = ref([])
const loading = ref(false)
const submitting = ref(false)
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)

const dialogVisible = ref(false)
const isEditing = ref(false)
const editingId = ref(null)
const formRef = ref(null)

const dialogTitle = computed(() => (isEditing.value ? 'កែប្រែអ្នកនិពន្ធ' : 'បង្កេីតអ្នកនិពន្ធ'))

const form = reactive({
  name: '',
  gender: null,
  note: '',
  faculty_ids: [],
})

const rules = {
  name: [{ required: true, message: 'ត្រូវបញ្ចូលឈ្មោះ', trigger: 'blur' }],
  gender: [{ required: true, message: 'ត្រូវរេីសភេទ', trigger: 'change' }],
}

const genderOptions = [
  { label: 'ប្រុស', value: 1 },
  { label: 'ស្រី', value: 2 },
]

const statusOptions = [
  { label: 'សកម្ម', value: 1 },
  { label: 'អសកម្ម', value: 0 },
]

const facultyOptions = computed(() =>
  faculty.value.map((f) => ({
    label: f.name,
    value: f.id,
  }))
)

const filters = reactive({ name: '', gender: null, is_active: null })

const columns = [
  { prop: 'id', label: 'ល.រ', width: 70 },
  { prop: 'name', label: 'ឈ្មោះ' },
  { prop: 'gender', label: 'ភេទ', slot: 'gender', width: 100 },
  { prop: 'is_active', label: 'ស្ថានភាព', slot: 'isActive', width: 100 },
  { prop: 'note', label: 'ចំណាំ' },
  { prop: 'faculty', label: 'ក្នុងមហាវិទ្យាល័យ', slot: 'faculty', width: 300 },
]

function resetForm() {
  Object.assign(form, {
    name: '',
    gender: null,
    note: '',
    faculty_ids: [],
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
    gender: row.gender,
    note: row.note,
    faculty_ids: (row.faculty || []).map((f) => f.id),
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
  if (filters.gender !== null && filters.gender !== '') params.gender = filters.gender
  if (filters.is_active !== null && filters.is_active !== '') params.is_active = filters.is_active
  return params
}

async function fetchFaculty() {
  try {
    const res = await getfaculty()
    faculty.value = res.data.data || []
  } catch (e) {
    notify.error(e?.response?.data?.message || e.message || 'Failed to load faculty')
  }
}

async function fetchAuthor() {
  loading.value = true
  try {
    const res = await getauthor(buildParams())
    authors.value = res.data.data || []
    total.value = res.data.pagination?.totalCount || 0
  } catch (e) {
    notify.error(e?.response?.data?.message || e.message || 'Failed to load authors')
  } finally {
    loading.value = false
  }
}

async function handleSubmit() {
  const payload = {
    name: form.name,
    gender: form.gender,
    note: form.note,
    faculty_ids: form.faculty_ids,
  }

  submitting.value = true
  try {
    if (isEditing.value) {
      await updateauthor(editingId.value, payload)
      notify.success('កែប្រែជោគជ័យ')
    } else {
      await createauthor(payload)
      notify.success('បង្កេីតជោគជ័យ')
    }
    dialogVisible.value = false
    await fetchAuthor()
  } catch (e) {
    notify.error(e?.response?.data?.error || e.message || 'Failed to save author')
  } finally {
    submitting.value = false
  }
}

async function togggle(row) {
  try {
    await togglestatusauthor(row.id);
     notify.success('កែប្រែជោគជ័យ')
    fetchAuthor();
  } catch (e) {
     notify.error(e?.response?.data?.error || e.message || 'Failed to save author')
  }
}

function onReset() {
  filters.name = ''
  filters.gender = null
  filters.is_active = null
  page.value = 1
  fetchAuthor()
}

const debouncedFetch = debounce(() => {
  page.value = 1 // reset to first page whenever filters change
  fetchAuthor()
}, 500)

// watch the actual filter values, not a function reference
watch(() => ({ ...filters }), debouncedFetch, { deep: true })

onMounted(() => {
  fetchAuthor()
  fetchFaculty()
})
</script>

<template>
  <div class="author-page">
    <AppFilterBar
      :fields="[
        { slot: 'name', span: 9 },
        { slot: 'gender', span: 5 },
        { slot: 'status', span: 5 },
        { slot: 'add', span: 1 },
      ]"
      :action-span="4"
    >
      <template #name>
        <AppInput v-model="filters.name" placeholder="ស្វែងរក..." clearable />
      </template>

      <template #gender>
        <AppSelect
          label="ជ្រេីសរេីសភេទ"
          v-model="filters.gender"
          :options="genderOptions"
          placeholder="ប្រុស-ស្រី"
          clearable
        />
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
      :data="authors"
      :columns="columns"
      :loading="loading"
      :total="total"
      v-model:current-page="page"
      v-model:page-size="pageSize"
      @page-change="fetchAuthor"
    >
      <template #isActive="{ row }">
        <el-tag :type="row.is_active ? 'success' : 'danger'">
          {{ row.is_active ? 'សកម្ម' : 'អសកម្ម' }}
        </el-tag>
      </template>

      <template #gender="{ row }">
        <el-text>
          {{ row.gender === 1 ? 'ប្រុស' : 'ស្រី' }}
        </el-text>
      </template>

      <template #faculty="{ row }">
        <el-tag
          v-for="f in row.faculty"
          :key="f.id"
          class="faculty-tag"
          type="info"
          effect="plain"
        >
          {{ f.name }}
        </el-tag>
        <span v-if="!row.faculty || row.faculty.length === 0" class="text-muted">—</span>
      </template>

      <template #actions="{row}">
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
        <AppSelect
          label="ភេទ"
          v-model="form.gender"
          :options="genderOptions"
          placeholder="ជ្រេីសរេីសភេទ"
          clearable
          prop="gender"
        />
        <AppSelect
         label="មហាវិទ្យាល័យ"
          v-model="form.faculty_ids"
          :options="facultyOptions"
          placeholder="ជ្រេីសរេីសមហាវិទ្យាល័យ"
          clearable
          prop="faculty_ids"
          multiple
        ></AppSelect>
        <AppInput v-model="form.note" placeholder="ចំណាំ" clearable prop="note"  label="ចំណាំ"></AppInput>
      </AppForm>
    </AppDialog>
  </div>
</template>

<style scoped>
.author-filters {
  display: flex;
  gap: 10px;
  margin-bottom: 16px;
  flex-wrap: wrap;
  align-items: center;
}

.faculty-tag {
  margin-right: 8px;
  margin-bottom: 4px;
}
</style>