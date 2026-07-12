<script setup>
// Main authenticated shell: Sidebar + Navbar + routed page content.
// Handles responsive behavior by switching the sidebar to a drawer on mobile.
import { computed } from 'vue'
import { useAppStore } from '@/stores/app'
import { useDevice } from '@/composables/useDevice'
import Sidebar from './components/Sidebar.vue'
import Navbar from './components/Navbar.vue'
import Breadcrumb from './components/Breadcrumb.vue'
import AppFooter from './components/Footer.vue'

const appStore = useAppStore()
useDevice()

const isMobile = computed(() => appStore.device === 'mobile')
const contentMargin = computed(() => {
  if (isMobile.value) return '0'
  return appStore.sidebarCollapsed ? 'var(--sidebar-width-collapsed)' : 'var(--sidebar-width)'
})
</script>

<template>
  <div class="min-h-screen" style="background: var(--color-bg)">
    <!-- Desktop / tablet sidebar -->
    <Sidebar v-if="!isMobile" />

    <!-- Mobile drawer sidebar -->
    <el-drawer
      v-else
      v-model="appStore.sidebarDrawerOpen"
      direction="ltr"
      size="260px"
      :with-header="false"
      class="mobile-sidebar-drawer"
    >
      <Sidebar mobile @navigate="appStore.closeDrawer()" />
    </el-drawer>

    <div
      class="flex flex-col min-h-screen transition-all duration-200"
      :style="{ marginLeft: contentMargin }"
    >
      <Navbar />

      <main class="flex-1 p-4 md:p-6">
        <Breadcrumb class="mb-4" />
        <router-view v-slot="{ Component }">
          <transition name="app-fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </main>

      <AppFooter />
    </div>
  </div>
</template>

<style scoped>
:deep(.mobile-sidebar-drawer .el-drawer__body) {
  padding: 0;
  background: var(--sidebar-bg);
}
</style>
