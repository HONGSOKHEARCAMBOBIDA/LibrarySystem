<script setup>
// Standard page header used at the top of every view: title + optional
// subtitle on the left, action buttons (slot) on the right.
defineProps({
  title: { type: String, required: true },
  subtitle: { type: String, default: '' },
})
</script>

<template>
  <div class="flex flex-col md:flex-row md:items-center md:justify-between gap-3 mb-6">
    <div>
      <h1 class="text-xl md:text-2xl font-semibold" style="color: var(--color-text-primary)">
        {{ title }}
      </h1>
      <p v-if="subtitle" class="text-sm mt-1" style="color: var(--color-text-secondary)">
        {{ subtitle }}
      </p>
    </div>
    <div class="flex items-center gap-2 shrink-0">
      <slot name="actions" />
    </div>
  </div>
</template>
