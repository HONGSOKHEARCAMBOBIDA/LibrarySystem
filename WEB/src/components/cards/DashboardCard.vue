<script setup>
// Generic content card: title + optional actions slot + body slot.
// Used to wrap charts, lists, tables anywhere on a dashboard-style page.
defineProps({
  title: { type: String, default: '' },
  subtitle: { type: String, default: '' },
  bodyClass: { type: String, default: '' },
})
</script>

<template>
  <div class="app-card overflow-hidden">
    <div
      v-if="title || $slots.actions"
      class="flex items-center justify-between px-5 pt-5 pb-3"
    >
      <div>
        <h3 class="font-semibold text-[15px]" style="color: var(--color-text-primary)">
          {{ title }}
        </h3>
        <p v-if="subtitle" class="text-xs mt-0.5" style="color: var(--color-text-secondary)">
          {{ subtitle }}
        </p>
      </div>
      <div v-if="$slots.actions">
        <slot name="actions" />
      </div>
    </div>
    <div class="px-5 pb-5" :class="bodyClass">
      <slot />
    </div>
  </div>
</template>
