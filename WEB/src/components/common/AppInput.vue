<script setup>
// Reusable text input with label + error-message slot, built on el-input.
// Works standalone or inside an el-form (validation still handled by the form).
defineProps({
  modelValue: { type: [String, Number], default: '' },
  label: { type: String, default: '' },
  placeholder: { type: String, default: '' },
  type: { type: String, default: 'text' },
  disabled: { type: Boolean, default: false },
  clearable: { type: Boolean, default: true },
  errorMessage: { type: String, default: '' },
  prefixIcon: { type: String, default: '' },
  size: { type: String, default: 'default' },
})
defineEmits(['update:modelValue', 'blur', 'focus'])
</script>

<template>
  <div class="app-input">
    <label v-if="label" class="block text-sm font-medium mb-1.5" style="color: var(--color-text-primary)">
      {{ label }}
    </label>
    <el-input
      :model-value="modelValue"
      :type="type"
      :placeholder="placeholder"
      :disabled="disabled"
      :clearable="clearable"
      :size="size"
      :class="{ 'is-error': errorMessage }"
      @update:model-value="$emit('update:modelValue', $event)"
      @blur="$emit('blur', $event)"
      @focus="$emit('focus', $event)"
    >
      <template v-if="prefixIcon" #prefix>
        <el-icon><component :is="prefixIcon" /></el-icon>
      </template>
    </el-input>
    <p v-if="errorMessage" class="text-xs text-red-500 mt-1">{{ errorMessage }}</p>
  </div>
</template>
