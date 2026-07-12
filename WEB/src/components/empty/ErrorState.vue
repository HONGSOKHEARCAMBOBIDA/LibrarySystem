<script setup>
defineProps({
  description: { type: String, default: 'Something went wrong. Please try again.' },
  retryText: { type: String, default: 'Retry' },
})
defineEmits(['retry'])
</script>

<template>
  <div class="flex flex-col items-center justify-center py-14 px-4 text-center">
    <el-icon :size="52" class="mb-3 text-red-400">
      <CircleCloseFilled />
    </el-icon>
    <p class="text-sm" style="color: var(--color-text-secondary)">{{ description }}</p>
    <el-button class="mt-4" size="small" type="primary" plain @click="$emit('retry')">
      {{ retryText }}
    </el-button>
  </div>
</template>
