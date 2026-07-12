<script setup>
// Top navbar: sidebar/drawer toggle, notification bell, dark-mode switch,
// and the user menu (profile / logout).
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAppStore } from '@/stores/app'
import { useAuthStore } from '@/stores/auth'
import { useConfirm } from '@/composables/useConfirm'
import UserAvatar from '@/components/common/UserAvatar.vue'
import { notify } from '@/utils/notify'

const appStore = useAppStore()
const authStore = useAuthStore()
const router = useRouter()
const { confirm } = useConfirm()

const isMobile = computed(() => appStore.device === 'mobile')
const isDark = computed(() => appStore.theme === 'dark')

const notifications = ref([
  { id: 1, title: 'New user registered', time: '5m ago', read: false },
  { id: 2, title: 'Server backup completed', time: '2h ago', read: false },
  { id: 3, title: 'Monthly report is ready', time: '1d ago', read: true },
])
const unreadCount = computed(() => notifications.value.filter((n) => !n.read).length)

function handleToggleSidebar() {
  if (isMobile.value) {
    appStore.toggleDrawer()
  } else {
    appStore.toggleSidebar()
  }
}

async function handleLogout() {
  const ok = await confirm('Are you sure you want to log out?', { title: 'Log out' })
  if (!ok) return
  await authStore.logout()
  notify.success('You have been logged out.')
  router.push({ name: 'Login' })
}

function handleCommand(command) {
  if (command === 'profile') router.push({ name: 'Profile' })
  if (command === 'settings') router.push({ name: 'Settings' })
  if (command === 'logout') handleLogout()
}
</script>

<template>
  <header
    class="sticky top-0 z-30 flex items-center justify-between h-16 px-4 md:px-6 border-b"
    style="background: var(--color-surface); border-color: var(--color-border)"
  >
    <div class="flex items-center gap-3">
      <el-button text circle @click="handleToggleSidebar">
        <el-icon :size="20"><Expand /></el-icon>
      </el-button>
    </div>

    <div class="flex items-center gap-1 md:gap-3">
      <!-- Dark mode toggle -->
      <el-tooltip :content="isDark ? 'Switch to light mode' : 'Switch to dark mode'">
        <el-button text circle @click="appStore.toggleTheme()">
          <el-icon :size="18"><component :is="isDark ? 'Sunny' : 'Moon'" /></el-icon>
        </el-button>
      </el-tooltip>

      <!-- Notifications -->
      <el-popover placement="bottom-end" width="320" trigger="click">
        <template #reference>
          <el-badge :value="unreadCount" :hidden="unreadCount === 0" :max="9">
            <el-button text circle>
              <el-icon :size="18"><Bell /></el-icon>
            </el-button>
          </el-badge>
        </template>
        <div>
          <div class="font-semibold px-1 pb-2 border-b" style="border-color: var(--color-border)">
            Notifications
          </div>
          <div v-if="notifications.length === 0" class="text-sm text-gray-400 py-6 text-center">
            No notifications
          </div>
          <div
            v-for="n in notifications"
            :key="n.id"
            class="flex items-start gap-2 py-2.5 px-1 border-b last:border-0 cursor-pointer hover:bg-black/[.02] rounded"
            style="border-color: var(--color-border)"
          >
            <span
              class="w-2 h-2 rounded-full mt-1.5 shrink-0"
              :class="n.read ? 'bg-transparent' : 'bg-primary-500'"
            />
            <div>
              <p class="text-sm" style="color: var(--color-text-primary)">{{ n.title }}</p>
              <p class="text-xs" style="color: var(--color-text-secondary)">{{ n.time }}</p>
            </div>
          </div>
        </div>
      </el-popover>

      <!-- User menu -->
      <el-dropdown trigger="click" @command="handleCommand">
        <div class="flex items-center gap-2 pl-2 cursor-pointer">
          <UserAvatar :src="authStore.userAvatar" :name="authStore.userName" :size="34" />
          <div class="hidden md:flex flex-col leading-tight text-left">
            <span class="text-sm font-medium" style="color: var(--color-text-primary)">
              {{ authStore.userName }}
            </span>
            <span class="text-xs" style="color: var(--color-text-secondary)">
             <span class="text-xs" style="color: var(--color-text-secondary)">
                {{ authStore.userRoles?.[0]?.name || 'Member' }}
              </span>
            </span>
          </div>
          <el-icon class="hidden md:block text-gray-400"><ArrowDown /></el-icon>
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item command="profile">
              <el-icon><User /></el-icon> Profile
            </el-dropdown-item>
            <el-dropdown-item command="settings">
              <el-icon><Setting /></el-icon> Settings
            </el-dropdown-item>
            <el-dropdown-item command="logout" divided>
              <el-icon><SwitchButton /></el-icon> Logout
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </header>
</template>
