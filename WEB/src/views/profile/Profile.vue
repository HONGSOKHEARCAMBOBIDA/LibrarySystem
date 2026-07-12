<script setup>
// User profile page — edit personal info + change password, both using
// AppForm for consistent validate/submit/loading behavior.
import { reactive, ref } from 'vue'
import { useAuthStore } from '@/stores/auth'
import PageHeader from '@/components/common/PageHeader.vue'
import DashboardCard from '@/components/cards/DashboardCard.vue'
import AppForm from '@/components/forms/AppForm.vue'
import UserAvatar from '@/components/common/UserAvatar.vue'
import { notify } from '@/utils/notify'

const authStore = useAuthStore()

const profileForm = reactive({
  name: authStore.userName,
  email: authStore.user?.email || '',
  phone: '+1 (555) 123-4567',
  bio: 'Product designer focused on building clean, usable admin tools.',
})
const profileRules = {
  name: [{ required: true, message: 'Name is required', trigger: 'blur' }],
  email: [
    { required: true, message: 'Email is required', trigger: 'blur' },
    { type: 'email', message: 'Enter a valid email', trigger: 'blur' },
  ],
}
const profileSaving = ref(false)
async function handleProfileSubmit() {
  profileSaving.value = true
  await new Promise((r) => setTimeout(r, 600))
  profileSaving.value = false
  notify.success('Profile updated successfully.')
}

const passwordForm = reactive({ current: '', next: '', confirm: '' })
const passwordRules = {
  current: [{ required: true, message: 'Current password is required', trigger: 'blur' }],
  next: [
    { required: true, message: 'New password is required', trigger: 'blur' },
    { min: 8, message: 'Must be at least 8 characters', trigger: 'blur' },
  ],
  confirm: [
    { required: true, message: 'Please confirm your new password', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== passwordForm.next) callback(new Error('Passwords do not match'))
        else callback()
      },
      trigger: 'blur',
    },
  ],
}
const passwordSaving = ref(false)
async function handlePasswordSubmit() {
  passwordSaving.value = true
  await new Promise((r) => setTimeout(r, 600))
  passwordSaving.value = false
  Object.assign(passwordForm, { current: '', next: '', confirm: '' })
  notify.success('Password changed successfully.')
}
</script>

<template>
  <div>
    <PageHeader title="Profile" subtitle="Manage your personal information and account security." />

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <DashboardCard class="lg:col-span-1">
        <div class="flex flex-col items-center text-center py-4">
          <UserAvatar :src="authStore.userAvatar" :name="authStore.userName" :size="88" />
          <h3 class="mt-4 font-semibold text-[15px]" style="color: var(--color-text-primary)">
            {{ authStore.userName }}
          </h3>
          <p class="text-sm" style="color: var(--color-text-secondary)">
            {{ authStore.user?.email }}
          </p>
          <el-tag class="mt-3" size="small" round>{{ authStore.user?.roles?.[0] || 'Member' }}</el-tag>
          <el-button class="mt-5" size="small">Change Avatar</el-button>
        </div>
      </DashboardCard>

      <DashboardCard title="Personal Information" class="lg:col-span-2">
        <AppForm
          :model="profileForm"
          :rules="profileRules"
          :loading="profileSaving"
          submit-text="Save Changes"
          @submit="handleProfileSubmit"
        >
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-x-4">
            <el-form-item label="Full Name" prop="name">
              <el-input v-model="profileForm.name" />
            </el-form-item>
            <el-form-item label="Email" prop="email">
              <el-input v-model="profileForm.email" />
            </el-form-item>
            <el-form-item label="Phone" prop="phone">
              <el-input v-model="profileForm.phone" />
            </el-form-item>
          </div>
          <el-form-item label="Bio" prop="bio">
            <el-input v-model="profileForm.bio" type="textarea" :rows="3" />
          </el-form-item>
        </AppForm>
      </DashboardCard>

      <DashboardCard title="Change Password" class="lg:col-span-3">
        <AppForm
          :model="passwordForm"
          :rules="passwordRules"
          :loading="passwordSaving"
          submit-text="Update Password"
          @submit="handlePasswordSubmit"
        >
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-x-4">
            <el-form-item label="Current Password" prop="current">
              <el-input v-model="passwordForm.current" type="password" show-password />
            </el-form-item>
            <el-form-item label="New Password" prop="next">
              <el-input v-model="passwordForm.next" type="password" show-password />
            </el-form-item>
            <el-form-item label="Confirm Password" prop="confirm">
              <el-input v-model="passwordForm.confirm" type="password" show-password />
            </el-form-item>
          </div>
        </AppForm>
      </DashboardCard>
    </div>
  </div>
</template>
