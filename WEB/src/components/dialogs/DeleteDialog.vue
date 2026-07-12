<script setup>
// Specialized delete-confirmation dialog with a "type item name to confirm"
// pattern is overkill for a starter, so this stays simple + reusable.
defineProps({
  visible: { type: Boolean, required: true },
  itemName: { type: String, default: 'this item' },
  loading: { type: Boolean, default: false },
})
const emit = defineEmits(['update:visible', 'confirm', 'cancel'])
</script>

<template>
  <el-dialog
    :model-value="visible"
    title="Delete Confirmation"
    width="420px"
    @update:model-value="$emit('update:visible', $event)"
  >
    <div class="flex items-start gap-3">
      <el-icon :size="22" class="text-red-500"><Delete /></el-icon>
      <div>
        <p style="color: var(--color-text-primary)">
          Are you sure you want to delete <strong>{{ itemName }}</strong>?
        </p>
        <p class="text-sm mt-1" style="color: var(--color-text-secondary)">
          This action cannot be undone.
        </p>
      </div>
    </div>

    <template #footer>
      <el-button @click="$emit('update:visible', false); $emit('cancel')">Cancel</el-button>
      <el-button type="danger" :loading="loading" @click="$emit('confirm')">Delete</el-button>
    </template>
  </el-dialog>
</template>
