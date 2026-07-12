<script setup>
import { useRouter } from 'vue-router'
const router = useRouter()
</script>

<template>
  <div class="min-h-screen flex flex-col items-center justify-center p-6 text-center" style="background: var(--color-bg)">
    <p class="text-7xl font-bold text-primary-500 mb-2">404</p>
    <h1 class="text-xl font-semibold mb-2" style="color: var(--color-text-primary)">
      Page not found
    </h1>
    <p class="text-sm mb-6 max-w-sm" style="color: var(--color-text-secondary)">
      The page you're looking for doesn't exist or has been moved.
    </p>
    <el-button type="primary" @click="router.push({ name: 'Dashboard' })">
      Back to Dashboard
    </el-button>
  </div>
</template>
