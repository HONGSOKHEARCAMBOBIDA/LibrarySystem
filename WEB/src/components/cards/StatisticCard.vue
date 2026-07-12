<script setup>
// KPI card used on the Dashboard page — number + label + trend + icon.
import { computed } from 'vue'

const props = defineProps({
  label: { type: String, required: true },
  value: { type: [String, Number], required: true },
  icon: { type: String, default: 'TrendCharts' },
  iconColor: { type: String, default: '#6366f1' },
  trend: { type: Number, default: null }, // e.g. 12.5 or -4.2 (percent)
})

const isPositive = computed(() => (props.trend ?? 0) >= 0)
</script>

<template>
  <div class="app-card p-5 flex items-start justify-between hover:shadow-md transition-shadow">
    <div>
      <p class="text-sm mb-1.5" style="color: var(--color-text-secondary)">{{ label }}</p>
      <p class="text-2xl font-semibold" style="color: var(--color-text-primary)">{{ value }}</p>
    </div>
    <div
      class="w-11 h-11 rounded-xl flex items-center justify-center shrink-0"
      :style="{ background: iconColor + '1A', color: iconColor }"
    >
      <el-icon :size="20"><component :is="icon" /></el-icon>
    </div>
  </div>
</template>
