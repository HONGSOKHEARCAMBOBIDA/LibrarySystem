<script setup>
import { reactive, ref, onMounted } from 'vue'
import PageHeader from '@/components/common/PageHeader.vue'
import AppTable from '@/components/tables/AppTable.vue'
import AppSearch from '@/components/tables/AppSearch.vue'
import AppPagination from '@/components/tables/AppPagination.vue'
import AppDialog from '@/components/dialogs/AppDialog.vue'
import DeleteDialog from '@/components/dialogs/DeleteDialog.vue'
import UserAvatar from '@/components/common/UserAvatar.vue'
import { useTable } from '@/composables/useTable'
import { formatDate } from '@/utils'
import { notify } from '@/utils/notify'
import authApi from '@/services/auth.service'
import roleApi from '@/services/role.service'
import AppButton from '../../components/button/AppButton.vue'

async function fetchUsers({ page, pageSize, search, sortProp, sortOrder } = {}) {
  const { list, pagination } = await authApi.getUser({
    page,
    pageSize,
    name: search || undefined,
  })
  return {
    rows: list.map((u) => ({
      id: u.id,
      name: u.name_en,
      nameKh: u.name_kh,
      role: u.role_name,
      roleId: u.role_id,
      gender: u.gender,
      dob: u.dob,
      status: u.is_active ? 'active' : 'suspended',
    })),
    total: pagination?.totalCount ?? 0,
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

const dialogVisible = ref(false)
const dialogLoading = ref(false)
const formRef = ref(null)
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

async function handleSave() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }
  dialogLoading.value = true
  try {
    await authApi.register({ ...form })
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
    <PageHeader
      title="អ្នកប្រេីប្រាស់"
      subtitle="គ្រប់គ្រងសមាជិកក្រុមរបស់អ្នក និងការអនុញ្ញាតគណនីរបស់ពួកគេ"
    >
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

    <AppTable
      :columns="columns"
      :data="rows"
      :loading="loading"
      row-key="id"
      empty-text="No users match your search"
    >
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

    <AppPagination
      :page="pagination.page"
      :page-size="pagination.pageSize"
      :total="pagination.total"
      @update:page="handlePageChange"
      @update:page-size="handlePageSizeChange"
    />

    <AppDialog
      v-model:visible="dialogVisible"
      title="បង្កេីតអ្នកប្រេីប្រាស់"
      :loading="dialogLoading"
      confirm-text="Register"
      @confirm="handleSave"
    >
      <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
        <el-form-item label="Name (English)" prop="name_en">
          <el-input v-model="form.name_en" placeholder="e.g. Jane Doe" />
        </el-form-item>
        <el-form-item label="Name (Khmer)" prop="name_kh">
          <el-input v-model="form.name_kh" placeholder="ឈ្មោះ" />
        </el-form-item>
        <el-form-item label="Role" prop="role_id">
          <el-select v-model="form.role_id" class="w-full">
            <el-option
              v-for="r in roles"
              :key="r.id"
              :label="r.display_name || r.name"
              :value="r.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Gender" prop="gender">
          <el-select v-model="form.gender" class="w-full">
            <el-option label="Male" :value="1" />
            <el-option label="Female" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="Date of Birth" prop="dob">
          <el-date-picker
            v-model="form.dob"
            type="date"
            value-format="YYYY-MM-DD"
            class="w-full"
            placeholder="Select date"
          />
        </el-form-item>
      </el-form>
    </AppDialog>
  </div>
</template>
