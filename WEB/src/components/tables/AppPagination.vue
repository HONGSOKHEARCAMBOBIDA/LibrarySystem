<script setup>
import { PAGE_SIZE_OPTIONS } from '@/config/constants'
import { BREAKPOINTS } from '../../config/constants';
import { onMounted } from 'vue';
defineProps({
  page: { type: Number, required: true },
  pageSize: { type: Number, required: true },
  total: { type: Number, required: true },
})
defineEmits(['update:page', 'update:pageSize'])
const isMobile = ref(false)
const checkScreen = () => {
  isMobile.value = window.innerWidth < BREAKPOINTS.mobile
}
onMounted(()=>{
  checkScreen()
  window.addEventListener('resize', checkScreen)
})
onUnmounted(() => {
  window.removeEventListener('resize', checkScreen)
})
const paginationLayout = computed(() =>
  isMobile.value
    ? 'prev, pager, next'
    : 'total, sizes, prev, pager, next, jumper'
)
</script>

<template>
  <div class="flex justify-center md:justify-end pt-4 overflow-x-auto">
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="PAGE_SIZE_OPTIONS"
      :total="total"
      :layout="paginationLayout"
      background
      @current-change="$emit('update:page', $event)"
      @size-change="$emit('update:pageSize', $event)"
    />
  </div>
</template>
