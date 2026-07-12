<script setup>
// General settings page — theme, locale, and notification preferences.
// All state is local/mock; wire to a settings.service.js when ready.
import { reactive } from 'vue'
import { useAppStore } from '@/stores/app'
import PageHeader from '@/components/common/PageHeader.vue'
import DashboardCard from '@/components/cards/DashboardCard.vue'
import { notify } from '@/utils/notify'

const appStore = useAppStore()

const prefs = reactive({
  language: 'en',
  timezone: 'UTC',
  emailNotifications: true,
  pushNotifications: false,
  weeklyDigest: true,
})

function handleThemeChange(value) {
  appStore.setTheme(value ? 'dark' : 'light')
}

function handleSave() {
  notify.success('Settings saved.')
}
</script>

<template>
  <div>
    <PageHeader title="Settings" subtitle="Configure application preferences and notifications." />

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <DashboardCard title="Appearance">
        <div class="flex items-center justify-between py-2">
          <div>
            <p class="text-sm font-medium" style="color: var(--color-text-primary)">Dark Mode</p>
            <p class="text-xs" style="color: var(--color-text-secondary)">
              Switch between light and dark theme.
            </p>
          </div>
          <el-switch :model-value="appStore.theme === 'dark'" @change="handleThemeChange" />
        </div>

        <el-divider />

        <el-form label-position="top">
          <el-form-item label="Language">
            <el-select v-model="prefs.language" class="w-full">
              <el-option label="English" value="en" />
              <el-option label="Spanish" value="es" />
              <el-option label="French" value="fr" />
            </el-select>
          </el-form-item>
          <el-form-item label="Timezone">
            <el-select v-model="prefs.timezone" class="w-full">
              <el-option label="UTC" value="UTC" />
              <el-option label="Eastern Time (US)" value="ET" />
              <el-option label="Pacific Time (US)" value="PT" />
            </el-select>
          </el-form-item>
        </el-form>
      </DashboardCard>

      <DashboardCard title="Notifications">
        <div class="space-y-5 py-1">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium" style="color: var(--color-text-primary)">
                Email Notifications
              </p>
              <p class="text-xs" style="color: var(--color-text-secondary)">
                Receive updates about your account via email.
              </p>
            </div>
            <el-switch v-model="prefs.emailNotifications" />
          </div>
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium" style="color: var(--color-text-primary)">
                Push Notifications
              </p>
              <p class="text-xs" style="color: var(--color-text-secondary)">
                Receive push notifications on this device.
              </p>
            </div>
            <el-switch v-model="prefs.pushNotifications" />
          </div>
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm font-medium" style="color: var(--color-text-primary)">
                Weekly Digest
              </p>
              <p class="text-xs" style="color: var(--color-text-secondary)">
                A weekly summary of activity on your account.
              </p>
            </div>
            <el-switch v-model="prefs.weeklyDigest" />
          </div>
        </div>
      </DashboardCard>

      <div class="lg:col-span-2 flex justify-end">
        <el-button type="primary" @click="handleSave">Save Settings</el-button>
      </div>
    </div>
  </div>
</template>
