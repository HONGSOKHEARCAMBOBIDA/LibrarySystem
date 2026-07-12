<script setup>
// Declarative confirmation dialog (alternative to the useConfirm composable
// when you want it embedded in the template rather than promise-based).
defineProps({
  visible: { type: Boolean, required: true },
  title: { type: String, default: 'Please Confirm' },
  message: { type: String, required: true },
  confirmText: { type: String, default: 'Confirm' },
  cancelText: { type: String, default: 'Cancel' },
  loading: { type: Boolean, default: false },
  type: { type: String, default: 'warning' }, // warning | danger
})
const emit = defineEmits(['update:visible', 'confirm', 'cancel'])
</script>

<template>
  <el-dialog
    :model-value="visible"
    :title="title"
    width="420px"
    @update:model-value="$emit('update:visible', $event)"
  >
    <div class="flex items-start gap-3">
      <el-icon :size="22" :class="type === 'danger' ? 'text-red-500' : 'text-amber-500'">
        <WarningFilled />
      </el-icon>
      <p style="color: var(--color-text-primary)">{{ message }}</p>
    </div>

    <template #footer>
      <el-button @click="$emit('update:visible', false); $emit('cancel')">
        {{ cancelText }}
      </el-button>
      <el-button
        :type="type === 'danger' ? 'danger' : 'primary'"
        :loading="loading"
        @click="$emit('confirm')"
      >
        {{ confirmText }}
      </el-button>
    </template>
  </el-dialog>
</template>
