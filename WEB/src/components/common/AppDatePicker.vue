<script setup>
// Reusable date picker wrapping el-date-picker (single date or range).
defineProps({
  modelValue: { type: [String, Date, Array], default: null },
  label: { type: String, default: '' },
  type: { type: String, default: 'date' }, // date | daterange | datetime | month
  placeholder: { type: String, default: 'Select date' },
  format: { type: String, default: 'YYYY-MM-DD' },
  valueFormat: { type: String, default: 'YYYY-MM-DD' },
  clearable: { type: Boolean, default: true },
  disabled: { type: Boolean, default: false },
})
defineEmits(['update:modelValue', 'change'])
</script>

<template>
  <div class="app-date-picker">
    <label v-if="label" class="block text-sm font-medium mb-1.5" style="color: var(--color-text-primary)">
      {{ label }}
    </label>
    <el-date-picker
      :model-value="modelValue"
      :type="type"
      :placeholder="placeholder"
      :start-placeholder="type.includes('range') ? 'Start date' : undefined"
      :end-placeholder="type.includes('range') ? 'End date' : undefined"
      :format="format"
      :value-format="valueFormat"
      :clearable="clearable"
      :disabled="disabled"
      class="w-full"
      @update:model-value="$emit('update:modelValue', $event)"
      @change="$emit('change', $event)"
    />
  </div>
</template>
