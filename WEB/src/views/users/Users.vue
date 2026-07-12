<script setup>
// Users management page — a complete example of AppTable + useTable +
// AppDialog + DeleteDialog wired together. Swap `fetchUsers`/`saveUser`/
// `deleteUser` for real calls to a users.service.js once your API exists.
import { reactive, ref } from 'vue'
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

// ---------------------------------------------------------------------
// Mock data source — replace with services/users.service.js later.
// The fetcher signature (page, pageSize, search, sort) already matches
// what a real paginated REST API would expect.
// ---------------------------------------------------------------------
const MOCK_ROLES = ['Admin', 'Editor', 'Viewer']
const MOCK_STATUSES = ['active', 'invited', 'suspended']
let mockUsers = Array.from({ length: 57 }, (_, i) => ({
  id: i + 1,
  name: ['Sophia Carter', 'Liam Johnson', 'Olivia Smith', 'Noah Williams', 'Emma Brown', 'James Davis'][i % 6] + (i >= 6 ? ` ${i}` : ''),
  email: `user${i + 1}@example.com`,
  avatar: `https://i.pravatar.cc/150?img=${(i % 70) + 1}`,
  role: MOCK_ROLES[i % MOCK_ROLES.length],
  status: MOCK_STATUSES[i % MOCK_STATUSES.length],
  createdAt: Date.now() - i * 1000 * 60 * 60 * 24,
}))

function fetchUsers({ page, pageSize, search, sort }) {
  return new Promise((resolve) => {
    setTimeout(() => {
      let rows = [...mockUsers]
      if (search) {
        const q = search.toLowerCase()
        rows = rows.filter(
          (u) => u.name.toLowerCase().includes(q) || u.email.toLowerCase().includes(q)
        )
      }
      if (sort.prop && sort.order) {
        const dir = sort.order === 'ascending' ? 1 : -1
        rows.sort((a, b) => (a[sort.prop] > b[sort.prop] ? dir : -dir))
      }
      const total = rows.length
      const start = (page - 1) * pageSize
      resolve({ rows: rows.slice(start, start + pageSize), total })
    }, 400)
  })
}

const {
  rows,
  loading,
  search,
  pagination,
  reload,
  handleSortChange,
  handlePageChange,
  handlePageSizeChange,
} = useTable(fetchUsers)

const columns = [
  { prop: 'name', label: 'Name', slot: 'name', sortable: true, minWidth: 220 },
  { prop: 'role', label: 'Role', width: 120 },
  { prop: 'status', label: 'Status', slot: 'status', width: 120 },
  { prop: 'createdAt', label: 'Joined', slot: 'createdAt', sortable: true, width: 140 },
]

const statusType = { active: 'success', invited: 'info', suspended: 'danger' }

// ---- Create / Edit dialog ----
const dialogVisible = ref(false)
const dialogLoading = ref(false)
const isEditing = ref(false)
const formRef = ref(null)
const form = reactive({ id: null, name: '', email: '', role: 'Viewer' })
const rules = {
  name: [{ required: true, message: 'Name is required', trigger: 'blur' }],
  email: [
    { required: true, message: 'Email is required', trigger: 'blur' },
    { type: 'email', message: 'Enter a valid email', trigger: 'blur' },
  ],
}

function openCreate() {
  isEditing.value = false
  Object.assign(form, { id: null, name: '', email: '', role: 'Viewer' })
  dialogVisible.value = true
}
function openEdit(row) {
  isEditing.value = true
  Object.assign(form, row)
  dialogVisible.value = true
}
async function handleSave() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }
  dialogLoading.value = true
  await new Promise((r) => setTimeout(r, 500)) // simulate request
  if (isEditing.value) {
    mockUsers = mockUsers.map((u) => (u.id === form.id ? { ...u, ...form } : u))
    notify.success('User updated successfully.')
  } else {
    mockUsers.unshift({
      ...form,
      id: Math.max(...mockUsers.map((u) => u.id)) + 1,
      status: 'invited',
      avatar: '',
      createdAt: Date.now(),
    })
    notify.success('User invited successfully.')
  }
  dialogLoading.value = false
  dialogVisible.value = false
  reload()
}

// ---- Delete ----
const deleteVisible = ref(false)
const deleteLoading = ref(false)
const rowToDelete = ref(null)
function openDelete(row) {
  rowToDelete.value = row
  deleteVisible.value = true
}
async function handleDelete() {
  deleteLoading.value = true
  await new Promise((r) => setTimeout(r, 500))
  mockUsers = mockUsers.filter((u) => u.id !== rowToDelete.value.id)
  deleteLoading.value = false
  deleteVisible.value = false
  notify.success('User deleted.')
  reload()
}
</script>

<template>
  <div>
    <PageHeader title="Users" subtitle="Manage your team members and their account permissions.">
      <template #actions>
        <el-button type="primary" @click="openCreate">
          <el-icon class="mr-1"><Plus /></el-icon>
          Invite User
        </el-button>
      </template>
    </PageHeader>

    <div class="flex items-center justify-between mb-4">
      <AppSearch v-model="search" placeholder="Search by name or email..." />
      <span class="text-sm hidden sm:block" style="color: var(--color-text-secondary)">
        {{ pagination.total }} total users
      </span>
    </div>

    <AppTable
      :columns="columns"
      :data="rows"
      :loading="loading"
      row-key="id"
      empty-text="No users match your search"
      @sort-change="handleSortChange"
    >
      <template #name="{ row }">
        <div class="flex items-center gap-2.5">
          <UserAvatar :src="row.avatar" :name="row.name" :size="32" />
          <div>
            <p class="text-sm font-medium leading-tight">{{ row.name }}</p>
            <p class="text-xs" style="color: var(--color-text-secondary)">{{ row.email }}</p>
          </div>
        </div>
      </template>
      <template #status="{ row }">
        <el-tag :type="statusType[row.status]" size="small" effect="light" round>
          {{ row.status }}
        </el-tag>
      </template>
      <template #createdAt="{ row }">
        <span class="text-sm" style="color: var(--color-text-secondary)">
          {{ formatDate(row.createdAt) }}
        </span>
      </template>
      <template #actions="{ row }">
        <el-button text size="small" @click="openEdit(row)">
          <el-icon><Edit /></el-icon>
        </el-button>
        <el-button text size="small" type="danger" @click="openDelete(row)">
          <el-icon><Delete /></el-icon>
        </el-button>
      </template>
    </AppTable>

    <AppPagination
      :page="pagination.page"
      :page-size="pagination.pageSize"
      :total="pagination.total"
      @update:page="handlePageChange"
      @update:page-size="handlePageSizeChange"
    />

    <!-- Create / Edit dialog -->
    <AppDialog
      v-model:visible="dialogVisible"
      :title="isEditing ? 'Edit User' : 'Invite User'"
      :loading="dialogLoading"
      :confirm-text="isEditing ? 'Save Changes' : 'Send Invite'"
      @confirm="handleSave"
    >
      <el-form ref="formRef" :model="form" :rules="rules" label-position="top">
        <el-form-item label="Full Name" prop="name">
          <el-input v-model="form.name" placeholder="e.g. Jane Doe" />
        </el-form-item>
        <el-form-item label="Email" prop="email">
          <el-input v-model="form.email" placeholder="jane@example.com" />
        </el-form-item>
        <el-form-item label="Role" prop="role">
          <el-select v-model="form.role" class="w-full">
            <el-option v-for="r in MOCK_ROLES" :key="r" :label="r" :value="r" />
          </el-select>
        </el-form-item>
      </el-form>
    </AppDialog>

    <!-- Delete confirmation -->
    <DeleteDialog
      v-model:visible="deleteVisible"
      :item-name="rowToDelete?.name"
      :loading="deleteLoading"
      @confirm="handleDelete"
    />
  </div>
</template>
