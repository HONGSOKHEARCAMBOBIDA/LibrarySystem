<script setup>
import { reactive, ref, onMounted,computed } from 'vue'
import PageHeader from '@/components/common/PageHeader.vue'
import AppTable from '@/components/tables/AppTable.vue'
import AppSearch from '@/components/tables/AppSearch.vue'
import AppPagination from '@/components/tables/AppPagination.vue'
import AppDialog from '@/components/dialogs/AppDialog.vue'
import { useTable } from '@/composables/useTable'
import { notify } from '@/utils/notify'
import authApi from '@/services/auth.service'
import roleApi from '@/services/role.service'
import AppButton from '../../components/button/AppButton.vue'
import AppForm from '@/components/forms/AppForm.vue'
import AppInput from '../../components/input/AppInput.vue'
import AppSelect from '../../components/common/AppSelect.vue'
import AppDatePicker from '../../components/common/AppDatePicker.vue'
async function fetchUsers({ page, pageSize, search }) {
  const { list, pagination } = await authApi.getUser({
    page,
    pageSize,
    name: search || undefined,
  })

  return {
    rows: list,
    total: pagination.totalCount,
  }
}

const { rows, loading, search, pagination, reload, handlePageChange, handlePageSizeChange } =
  useTable(fetchUsers)

const columns = [
  { prop: 'name', label: 'ឈ្មោះ', slot: 'name', sortable: true, minWidth: 220 },
  { slot: 'gender', label: 'ភេទ', width: 130 },
  { prop: 'role', label: 'តួនាទី', width: 140 },
  { prop: 'status', label: 'ស្ថានភាព', slot: 'status', width: 120 },
]

const statusType = { active: 'success', invited: 'info', suspended: 'danger' }

const roles = ref([])
onMounted(async () => {
  try {
    roles.value = await roleApi.getrole()
  } catch (e) {
    notify.error('Failed to load roles.')
  }
})

const roleOptions = computed(() =>
  roles.value.map(role => ({
    label: role.display_name || role.name,
    value: role.id,
  }))
)

function handleReset() {
  Object.assign(form, {
    name_kh: '',
    name_en: '',
    role_id: null,
    gender: 1,
    dob: '',
  })
}

const genderOptions = [
  { label: 'ប្រុស', value: 1 },
  { label: 'ស្រី', value: 2 },
]

const dialogVisible = ref(false)
const dialogLoading = ref(false)
const form = reactive({
  name_kh: '',
  name_en: '',
  role_id: null,
  gender: null,
  dob: '',
})
const rules = {
  name_kh: [{ required: true, message: 'Khmer name is required', trigger: 'blur' }],
  name_en: [{ required: true, message: 'English name is required', trigger: 'blur' }],
  role_id: [{ required: true, message: 'Role is required', trigger: 'change' }],
  dob: [{ required: true, message: 'Date of birth is required', trigger: 'change' }],
}

function openCreate() {
  Object.assign(form, { name_kh: '', name_en: '', role_id: null, gender: 1, dob: '' })
  dialogVisible.value = true
}

async function handleSave(payload) {
  dialogLoading.value = true

  try {
    await authApi.register(payload)

    notify.success('User registered successfully.')

    dialogVisible.value = false

    reload()
  } catch (err) {
    notify.error(err?.response?.data?.message || 'Failed to register user.')
  } finally {
    dialogLoading.value = false
  }
}
</script>

<template>
  <div>
    <PageHeader title="អ្នកប្រេីប្រាស់" subtitle="គ្រប់គ្រងសមាជិកក្រុមរបស់អ្នក និងការអនុញ្ញាតគណនីរបស់ពួកគេ">
      <template #actions>
        <AppButton type="primary" @click="openCreate">
          បង្កេីតថ្មី
        </AppButton>
      </template>
    </PageHeader>

    <div class="flex items-center justify-between mb-4">
      <AppSearch v-model="search" placeholder="ស្វែងរក..." />
      <span class="text-sm hidden sm:block" style="color: var(--color-text-secondary)">
        អ្នកប្រេីប្រាស់សរុប {{ pagination.total }} នាក់
      </span>
    </div>

    <AppTable :columns="columns" :data="rows" :loading="loading" row-key="id" empty-text="No users match your search">
      <template #name="{ row }">
        <div class="flex flex-col gap-1">
          <span class="font-medium text-gray-900 leading-none">
            {{ row.name }}
          </span>
          <span class="text-xs text-green-500 leading-none">
            {{ row.nameKh }}
          </span>
        </div>
      </template>
      <template #gender="{ row }">
        <el-text>
          {{ row.gender === 1 ? 'ប្រុស' : 'ស្រី' }}
        </el-text>
      </template>
      <template #status="{ row }">
        <el-tag :type="statusType[row.status]" size="small" effect="light">
          {{ row.status }}
        </el-tag>
      </template>
    </AppTable>

    <AppPagination :page="pagination.page" :page-size="pagination.pageSize" :total="pagination.total"
      @update:page="handlePageChange" @update:page-size="handlePageSizeChange" />

    <AppDialog v-model:visible="dialogVisible" title="បង្កេីតអ្នកប្រេីប្រាស់" :showDefaultFooter="false">
      <AppForm :model="form" :rules="rules" :loading="dialogLoading" :show-actions="true" @submit="handleSave" @reset="handleReset"
      submitText="កែប្រែ" resetText="ថយក្រោយ"
      >
        <AppInput v-model="form.name_en" prop="name_en" label="ឈ្មោះអង់គ្លេស" clearable></AppInput>
        <AppInput v-model="form.name_kh" prop="name_kh" label="ឈ្មោះខ្មែរ" clearable></AppInput>

        <AppSelect label="ជ្រេីសតួនាទី" prop="role_id" v-model="form.role_id" :options="roleOptions" placeholder="ជ្រេីសរេីសតួនាទី"
          clearable />


        <AppSelect label="ជ្រេីសភេទ" v-model="form.gender" :options="genderOptions" placeholder="ជ្រេីសរេីសភេទ"
          clearable />
        <AppDatePicker v-model="form.dob" label="ថ្ងៃ-ខែ-ឆ្នាំកំណេីត" clearable>

        </AppDatePicker>
      </AppForm>
    </AppDialog>
  </div>
</template>
