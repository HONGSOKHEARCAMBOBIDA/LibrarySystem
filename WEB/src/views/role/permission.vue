<script setup>
import { reactive, ref, computed, onMounted } from 'vue'
import PageHeader from '@/components/common/PageHeader.vue'
import AppTable from '@/components/tables/AppTable.vue'
import AppButton from '../../components/button/AppButton.vue'
import AppSearch from '@/components/tables/AppSearch.vue'
import AppPagination from '@/components/tables/AppPagination.vue'
import AppDialog from '@/components/dialogs/AppDialog.vue'
import { formatDate } from '@/utils'
import { notify } from '@/utils/notify'
import roleService from '@/services/role.service.js'
import { watch } from 'vue'
const role = ref([])
const rolehaspermission = ref([])
const loading = ref(false)
const searchKeyword = ref('')
const isAddMode = ref(false)
const selectroleID = ref(null)
const permissiondialogVisible = ref(false)
const submitting = ref(false)
const pagination = reactive({
  page: 1,
  pageSize: 10,
})

async function fetchrole() {
  loading.value = true
  try {
    const res = await roleService.getrole()
    role.value = res || []
  } catch (e) {
    notify.error(e.response?.data?.error || 'មានបញ្ហាក្នុងការទាញយកទិន្នន័យ')
  } finally {
    loading.value = false
  }
}

async function fetchrolehaspermission() {
  loading.value = true
  try {
    const res = await roleService.getrolepermission(selectroleID.value)
    rolehaspermission.value = res || []
  } catch (e) {
    notify.error(e.response?.data?.error || 'មានបញ្ហាក្នុងការទាញយកទិន្នន័យ')
  } finally {
    loading.value = false
  }
}

function openPermissionDialog(row) {
  selectroleID.value = row.id
  isAddMode.value = true
  permissiondialogVisible.value = true
  fetchrolehaspermission()
}

async function handleToggleAssigned(row) {
  const payload = {
    role_id: selectroleID.value,
    permission_ids: [row.id],
  }

  submitting.value = true
  try {
    if (row.assigned) {
      await roleService.createrolepermission(payload)
      notify.success('បន្ថែមសិទ្ធិដោយជោគជ័យ')
    } else {
      await roleService.deleterolepermission(payload)
      notify.success('ដកសិទ្ធិដោយជោគជ័យ')
    }
  } catch (e) {
    row.assigned = !row.assigned
    notify.error(e.response?.data?.error || 'មានបញ្ហាក្នុងការធ្វើបច្ចុប្បន្នភាព')
  } finally {
    submitting.value = false
  }
}

onMounted(fetchrole)

const filteredRoles = computed(() => {
  if (!searchKeyword.value) return role.value
  const kw = searchKeyword.value.toLowerCase()
  return role.value.filter(
    (r) => r.name?.toLowerCase().includes(kw) || r.module_name?.toLowerCase().includes(kw)
  )
})

const total = computed(() => filteredRoles.value.length)

const pagedRoles = computed(() => {
  const start = (pagination.page - 1) * pagination.pageSize
  return filteredRoles.value.slice(start, start + pagination.pageSize)
})

function handlePageChange(page) {
  pagination.page = page
}

function handlePageSizeChange(size) {
  pagination.pageSize = size
  pagination.page = 1
}

watch(searchKeyword, () => {
  pagination.page = 1
})

const columns = [
  { prop: 'name', label: 'ឈ្មោះ', sortable: false, minWidth: 200 },
  { prop: 'module_name', label: 'ឈ្មោះបង្ហាញ', slot: 'module_name', minWidth: 200 },
]

const permissioncolumns = [
  { prop: 'name', label: 'ឈ្មោះ', minWidth: 110 },
  { prop: 'module_name', label: 'ឈ្មោះបង្ហាញ', minWidth: 120 },
  { label: 'សិទ្ធិ', slot: 'assigned', minWidth: 80, align: 'center' },
]

const dialogVisible = ref(false)
const dialogLoading = ref(false)
const formRef = ref(null)
const roleForm = reactive({
  id: null,
  name: '',
  module_name: '',
})

const rules = {
  name: [{ required: true, message: 'Name is required', trigger: 'blur' }],
  module_name: [{ required: true, message: 'Module name is required', trigger: 'blur' }],
}

function openEdit(row) {
  Object.assign(roleForm, row)
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
    await roleService.updaterole(roleForm.id, {
      name: roleForm.name,
      module_name: roleForm.module_name,
    })
    notify.success('Role updated successfully.')
    dialogVisible.value = false
    await fetchrole()
  } catch (e) {
    notify.error(e.response?.data?.error || 'Failed to update role.')
  } finally {
    dialogLoading.value = false
  }
}
</script>

<template>
  <div>
    <PageHeader title="Roles" subtitle="Manage roles and the modules they grant access to." />

    <div class="flex items-center justify-between mb-4">
      <AppSearch v-model="searchKeyword" placeholder="Search by name or module..." />
      <span class="text-sm hidden sm:block" style="color: var(--color-text-secondary)">
        {{ total }} total roles
      </span>
    </div>

    <AppTable
      :columns="columns"
      :data="pagedRoles"
      :loading="loading"
      row-key="id"
      empty-text="No roles match your search"
    >
      <template #module_name="{ row }">
        <el-tag size="small" effect="light">
          {{ row.module_name }}
        </el-tag>
      </template>
      <template #actions="{ row }">
        <AppButton @click="openEdit(row)" icon="Edit" circle size="small"> </AppButton>
        <AppButton
          @click="openPermissionDialog(row)"
          icon="Lock"
          circle
          size="small"
          type="warning"
        ></AppButton>
      </template>
    </AppTable>

    <AppPagination
      :page="pagination.page"
      :page-size="pagination.pageSize"
      :total="total"
      @update:page="handlePageChange"
      @update:page-size="handlePageSizeChange"
    />

    <!-- Edit dialog -->
    <AppDialog
      v-model:visible="dialogVisible"
      title="កែប្រែតួនាទី"
      :loading="dialogLoading"
      confirm-text="កែប្រែ"
      @confirm="handleSave"
    >
      <el-form ref="formRef" :model="roleForm" :rules="rules" label-position="top">
        <el-form-item label="Role Name" prop="name">
          <el-input v-model="roleForm.name" placeholder="e.g. Editor" />
        </el-form-item>
        <el-form-item label="Module Name" prop="module_name">
          <el-input v-model="roleForm.module_name" placeholder="e.g. users" />
        </el-form-item>
      </el-form>
    </AppDialog>

    <!-- Permission Dialog -->

    <AppDialog
      v-model:visible="permissiondialogVisible"
      title="គ្រប់គ្រងសិទ្ធិ"
      :showDefaultFooter="false"
    >
      <AppTable
        :columns="permissioncolumns"
        :data="rolehaspermission"
        :loading="loading"
        row-key="id"
        empty-text="No roles match your search"
      >
        <template #assigned="{ row }">
          <el-switch
            v-model="row.assigned"
            style="--el-switch-on-color: #13ce66; --el-switch-off-color: #dcdfe6"
            :loading="submitting"
            @change="handleToggleAssigned(row)"
          />
        </template>
      </AppTable>
    </AppDialog>
  </div>
</template>
