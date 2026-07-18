<script setup>
import EmptyState from '@/components/empty/EmptyState.vue'
import { useBreakpoint } from '@/composables/useBreakpoint'

const { isMobile } = useBreakpoint()

defineProps({
  columns: { type: Array, required: true },
  data: { type: Array, default: () => [] },
  loading: { type: Boolean, default: false },
  rowKey: { type: String, default: 'id' },
  showSelection: { type: Boolean, default: false },
  emptyText: { type: String, default: 'No data found' },
  stripe: { type: Boolean, default: true },
})

defineEmits(['sort-change', 'selection-change', 'row-click'])
</script>

<template>
  <div class="app-table app-card overflow-hidden">
    <!-- Desktop -->
    <el-table
      border
      v-if="!isMobile"
   
      :data="data"
      :row-key="rowKey"
      :stripe="stripe"
      class="w-full"
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
        :fixed="col.fixed"
        :align="col.align || 'left'"
      >
        <template v-if="col.slot" #default="scope">
          <slot :name="col.slot" v-bind="scope" />
        </template>
      </el-table-column>

      <el-table-column v-if="$slots.actions" label="សកម្មភាព" width="180" align="center">
        <template #default="scope">
          <slot name="actions" v-bind="scope" />
        </template>
      </el-table-column>

      <template #empty>
        <EmptyState :description="emptyText" />
      </template>
    </el-table>

    <!-- Mobile -->
    <!-- Mobile -->
    <div v-else v-loading="loading" class="flex flex-col gap-3">
      <EmptyState v-if="!loading && !data.length" :description="emptyText" />

      <div
        v-for="(row, index) in data"
        :key="row[rowKey] || index"
        class="rounded-xl border border-[var(--el-border-color-lighter)] bg-[var(--el-bg-color)] px-4 py-4 shadow-sm transition-shadow"
        :class="{ 'cursor-pointer hover:shadow-md active:shadow-none': $attrs.onRowClick || true }"
        @click="$emit('row-click', row)"
      >
        <div class="flex items-start justify-between gap-3">
          <el-checkbox
            v-if="showSelection"
            class="mt-0.5"
            @click.stop
            @change="$emit('selection-change', $event ? [row] : [])"
          />

          <div class="flex-1 divide-y divide-[var(--el-border-color-lighter)]">
            <div
              v-for="col in columns"
              :key="col.prop || col.label"
              class="flex items-center justify-between gap-4 py-2 first:pt-0 last:pb-0"
            >
              <span class="shrink-0 text-sm font-medium text-[var(--el-text-color-secondary)]">
                {{ col.label }}
              </span>

              <div class="break-words text-right text-sm text-[var(--el-text-color-primary)]">
                <slot v-if="col.slot" :name="col.slot" :row="row" :$index="index" />
                <template v-else>
                  {{ row[col.prop] }}
                </template>
              </div>
            </div>
          </div>
        </div>

        <div
          v-if="$slots.actions"
          class="mt-3 flex justify-end gap-2 border-t border-[var(--el-border-color-lighter)] pt-3"
          @click.stop
        >
          <slot name="actions" :row="row" :$index="index" />
        </div>
      </div>
    </div>
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

:deep(.el-table__header-wrapper th) {
  color: #030303 !important;
}
</style>
