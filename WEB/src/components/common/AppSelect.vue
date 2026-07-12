<script setup>
// Reusable select built on el-select. Pass `options` as
// [{ label, value }] or a plain array of strings.
defineProps({
  modelValue: { type: [String, Number, Array], default: '' },
  label: { type: String, default: '' },
  placeholder: { type: String, default: 'Select' },
  options: { type: Array, default: () => [] },
  multiple: { type: Boolean, default: false },
  clearable: { type: Boolean, default: true },
  filterable: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  errorMessage: { type: String, default: '' },
})
defineEmits(['update:modelValue', 'change'])

function normalize(opt) {
  return typeof opt === 'object' ? opt : { label: opt, value: opt }
}
</script>

<template>
  <div class="app-select">
    <label v-if="label" class="block text-sm font-medium mb-1.5" style="color: var(--color-text-primary)">
      {{ label }}
    </label>
    <el-select
      :model-value="modelValue"
      :placeholder="placeholder"
      :multiple="multiple"
      :clearable="clearable"
      :filterable="filterable"
      :disabled="disabled"
      class="w-full"
      @update:model-value="$emit('update:modelValue', $event)"
      @change="$emit('change', $event)"
    >
      <el-option
        v-for="opt in options.map(normalize)"
        :key="opt.value"
        :label="opt.label"
        :value="opt.value"
      />
    </el-select>
    <p v-if="errorMessage" class="text-xs text-red-500 mt-1">{{ errorMessage }}</p>
  </div>
</template>
