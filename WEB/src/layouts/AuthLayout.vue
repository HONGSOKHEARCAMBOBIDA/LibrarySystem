<script setup>
// Simple centered layout for unauthenticated pages (login, forgot password).
import { APP_TITLE } from '@/config/constants'
</script>

<template>
  <div
    class="min-h-screen w-full flex items-center justify-center p-4"
  >
    <div class="w-full max-w-md">

      <div class="app-card p-10">
        <router-view />
      </div>
    </div>
  </div>
</template>
