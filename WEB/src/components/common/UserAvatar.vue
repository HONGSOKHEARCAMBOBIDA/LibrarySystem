<script setup>
// Avatar with automatic fallback to initials when no image is available.
import { computed } from 'vue'

const props = defineProps({
  src: { type: String, default: '' },
  name: { type: String, default: '' },
  size: { type: Number, default: 36 },
})

const initials = computed(() =>
  props.name
    .split(' ')
    .filter(Boolean)
    .slice(0, 2)
    .map((w) => w[0]?.toUpperCase())
    .join('')
)
</script>

<template>
  <el-avatar :size="size" :src="src || undefined" class="shrink-0" style="background: var(--color-primary)">
    <span v-if="!src" class="text-xs font-semibold">{{ initials || 'U' }}</span>
  </el-avatar>
</template>
