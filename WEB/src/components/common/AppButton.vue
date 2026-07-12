<script setup>
// Thin wrapper around el-button that standardizes our button styles
// and forwards all native/EP props + the default slot.
defineProps({
  type: { type: String, default: 'primary' },
  size: { type: String, default: 'default' },
  loading: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  plain: { type: Boolean, default: false },
  round: { type: Boolean, default: false },
  icon: { type: String, default: '' },
})
defineEmits(['click'])
</script>

<template>
  <el-button
    :type="type"
    :size="size"
    :loading="loading"
    :disabled="disabled"
    :plain="plain"
    :round="round"
    class="app-button"
    @click="$emit('click', $event)"
  >
    <el-icon v-if="icon && !loading" class="mr-1"><component :is="icon" /></el-icon>
    <slot />
  </el-button>
</template>

<style scoped>
.app-button {
  font-weight: 500;
}
</style>
