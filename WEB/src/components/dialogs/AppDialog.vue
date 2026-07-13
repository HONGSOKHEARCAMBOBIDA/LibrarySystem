<script setup>
import { useBreakpoint } from '@/composables/useBreakpoint'

const { isMobile } = useBreakpoint()

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
    v-if="!isMobile"
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

  <el-drawer
    v-else
    :model-value="visible"
    :title="title"
    direction="btt"
    size="auto"
    class="app-drawer-sheet"
    destroy-on-close
    @update:model-value="$emit('update:visible', $event)"
    @close="close"
  >
    <div class="max-h-[70vh] overflow-y-auto">
      <slot />
    </div>

    <template v-if="showDefaultFooter" #footer>
      <slot name="footer">
        <div class="flex gap-2">
          <el-button class="flex-1" @click="cancel">{{ cancelText }}</el-button>
          <el-button
            class="flex-1"
            type="primary"
            :loading="loading"
            @click="confirm"
          >
            {{ confirmText }}
          </el-button>
        </div>
      </slot>
    </template>
  </el-drawer>
</template>

<style scoped>
:deep(.app-drawer-sheet) {
  border-radius: 16px 16px 0 0;
  overflow: hidden;
}

:deep(.app-drawer-sheet .el-drawer__header) {
  margin-bottom: 0;
  padding: 16px 20px 12px;
  border-bottom: 1px solid var(--el-border-color-lighter);
}

:deep(.app-drawer-sheet .el-drawer__body) {
  padding: 16px 20px;
}

:deep(.app-drawer-sheet .el-drawer__footer) {
  padding: 12px 20px 20px;
  border-top: 1px solid var(--el-border-color-lighter);
}
</style>