<script setup>
// Consistent max-width + spacing wrapper for page bodies. Optional — pages
// can also lay out content directly since DashboardLayout already pads.
defineProps({
  maxWidth: { type: String, default: '100%' },
})
</script>

<template>
  <div class="w-full mx-auto" :style="{ maxWidth }">
    <slot />
  </div>
</template>
