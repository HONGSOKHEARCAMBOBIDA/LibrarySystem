<script setup>
// Generic modal wrapper. Use v-model:visible, put form content in the
// default slot and buttons in the #footer slot (or use the built-in
// confirm/cancel buttons via showDefaultFooter).
defineProps({
  visible: { type: Boolean, required: true },
  title: { type: String, default: '' },
  width: { type: [String, Number], default: '520px' },
  loading: { type: Boolean, default: false },
  showDefaultFooter: { type: Boolean, default: true },
  confirmText: { type: String, default: 'Save' },
  cancelText: { type: String, default: 'Cancel' },
})
const emit = defineEmits(['update:visible', 'confirm', 'cancel', 'close'])

function close() {
  emit('update:visible', false)
  emit('close')
}
function cancel() {
  close()
  emit('cancel')
}
function confirm() {
  emit('confirm')
}
</script>

<template>
  <el-dialog
    :model-value="visible"
    :title="title"
    :width="width"
    destroy-on-close
    @update:model-value="$emit('update:visible', $event)"
    @close="close"
  >
    <slot />

    <template v-if="showDefaultFooter" #footer>
      <slot name="footer">
        <el-button @click="cancel">{{ cancelText }}</el-button>
        <el-button type="primary" :loading="loading" @click="confirm">
          {{ confirmText }}
        </el-button>
      </slot>
    </template>
  </el-dialog>
</template>
