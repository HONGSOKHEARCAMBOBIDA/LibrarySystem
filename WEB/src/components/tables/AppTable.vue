<script setup>
// Generic data table. Pass `columns` to define dynamic columns and pair
// with the `useTable` composable for search/sort/pagination/loading state.
//
// columns: [
//   { prop: 'name', label: 'Name', sortable: true, width: 200 },
//   { prop: 'status', label: 'Status', slot: 'status' },   // use #status slot for custom render
// ]
import EmptyState from '@/components/empty/EmptyState.vue'

defineProps({
  columns: { type: Array, required: true },
  data: { type: Array, default: () => [] },
  loading: { type: Boolean, default: false },
  rowKey: { type: String, default: 'id' },
  showSelection: { type: Boolean, default: false },
  emptyText: { type: String, default: 'No data found' },
  stripe: { type: Boolean, default: false },
})
defineEmits(['sort-change', 'selection-change', 'row-click'])
</script>

<template>
  <div class="app-table app-card overflow-hidden">
    <el-table
      v-loading="loading"
      :data="data"
      :row-key="rowKey"
      :stripe="stripe"
      class="w-full"
      style="min-width: 100%"
      @sort-change="$emit('sort-change', $event)"
      @selection-change="$emit('selection-change', $event)"
      @row-click="$emit('row-click', $event)"
    >
      <el-table-column v-if="showSelection" type="selection" width="46" />

      <el-table-column
        v-for="col in columns"
        :key="col.prop || col.label"
        :prop="col.prop"
        :label="col.label"
        :width="col.width"
        :min-width="col.minWidth"
        :sortable="col.sortable ? 'custom' : false"
        :fixed="col.fixed"
        :align="col.align || 'left'"
        show-overflow-tooltip
      >
        <template v-if="col.slot" #default="scope">
          <slot :name="col.slot" v-bind="scope" />
        </template>
      </el-table-column>

      <el-table-column
        v-if="$slots.actions"
        label="Actions"
        :width="180"
        fixed="right"
        align="right"
      >
        <template #default="scope">
          <slot name="actions" v-bind="scope" />
        </template>
      </el-table-column>

      <template #empty>
        <EmptyState :description="emptyText" />
      </template>
    </el-table>
  </div>
</template>

<style scoped>
.app-table :deep(.el-table) {
  --el-table-border-color: var(--color-border);
  --el-table-header-bg-color: var(--color-surface-hover);
  background: transparent;
}
.app-table :deep(.el-table__inner-wrapper::before) {
  display: none;
}
</style>
