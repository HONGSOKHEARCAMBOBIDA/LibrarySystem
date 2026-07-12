<script setup>
// Debounced search box, typically placed above an AppTable.
// v-model updates immediately (for controlled input) but you should read
// the debounced value from useTable's `search` ref for actual fetching —
// useTable already debounces internally, so this component just forwards.
defineProps({
  modelValue: { type: String, default: '' },
  placeholder: { type: String, default: 'Search...' },
})
defineEmits(['update:modelValue'])
</script>

<template>
  <el-input
    :model-value="modelValue"
    :placeholder="placeholder"
    clearable
    class="app-search max-w-xs"
    @update:model-value="$emit('update:modelValue', $event)"
  >
    <template #prefix>
      <el-icon><Search /></el-icon>
    </template>
  </el-input>
</template>
