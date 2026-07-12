<script setup>
// Generic form wrapper around el-form that standardizes validate -> submit,
// reset, loading and disabled behavior. Put your fields (AppInput,
// AppSelect, or raw el-form-item) in the default slot.
//
// Usage:
//   <AppForm ref="formRef" :model="form" :rules="rules" :loading="saving"
//            @submit="handleSubmit">
//     <el-form-item label="Name" prop="name">
//       <el-input v-model="form.name" />
//     </el-form-item>
//   </AppForm>
import { ref } from 'vue'

const props = defineProps({
  model: { type: Object, required: true },
  rules: { type: Object, default: () => ({}) },
  loading: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  labelWidth: { type: [String, Number], default: '120px' },
  labelPosition: { type: String, default: 'top' },
  submitText: { type: String, default: 'Submit' },
  resetText: { type: String, default: 'Reset' },
  showActions: { type: Boolean, default: true },
})
const emit = defineEmits(['submit', 'reset'])

const formRef = ref(null)

async function handleSubmit() {
  if (!formRef.value) return
  try {
    await formRef.value.validate()
    emit('submit', props.model)
  } catch {
    // validation failed — el-form already highlights the offending fields
  }
}

function handleReset() {
  formRef.value?.resetFields()
  emit('reset')
}

defineExpose({ validate: () => formRef.value?.validate(), resetFields: handleReset, formRef })
</script>

<template>
  <el-form
    ref="formRef"
    :model="model"
    :rules="rules"
    :label-width="labelWidth"
    :label-position="labelPosition"
    :disabled="disabled || loading"
    @submit.prevent
  >
    <slot />

    <div v-if="showActions" class="flex items-center gap-2 mt-2">
      <el-button type="primary" :loading="loading" @click="handleSubmit">
        {{ submitText }}
      </el-button>
      <el-button :disabled="loading" @click="handleReset">{{ resetText }}</el-button>
    </div>
  </el-form>
</template>
