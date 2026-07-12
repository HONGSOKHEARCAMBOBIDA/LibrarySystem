<script setup>
// Global overlay shown whenever the loading store has in-flight requests.
// Mounted once in App.vue — nothing else needs to render this manually.
import { useLoadingStore } from '@/stores/loading'

const loadingStore = useLoadingStore()
</script>

<template>
  <transition name="app-fade">
    <div v-if="loadingStore.isLoading" class="loading-overlay">
      <div class="loading-box">
        <el-icon class="is-loading" :size="28"><Loading /></el-icon>
      </div>
    </div>
  </transition>
</template>

<style scoped>
.loading-overlay {
  position: fixed;
  inset: 0;
  z-index: 9999;
  background: rgba(255, 255, 255, 0.35);
  backdrop-filter: blur(1px);
  display: flex;
  align-items: center;
  justify-content: center;
  pointer-events: none;
}
.loading-box {
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  box-shadow: var(--shadow-card);
  border-radius: 12px;
  padding: 14px;
  color: var(--color-primary);
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
