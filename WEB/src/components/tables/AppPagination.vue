<script setup>
import { PAGE_SIZE_OPTIONS } from '@/config/constants'

defineProps({
  page: { type: Number, required: true },
  pageSize: { type: Number, required: true },
  total: { type: Number, required: true },
})
defineEmits(['update:page', 'update:pageSize'])
</script>

<template>
  <div class="flex justify-end pt-4">
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="PAGE_SIZE_OPTIONS"
      :total="total"
      layout="total, sizes, prev, pager, next, jumper"
      background
      @current-change="$emit('update:page', $event)"
      @size-change="$emit('update:pageSize', $event)"
    />
  </div>
</template>
