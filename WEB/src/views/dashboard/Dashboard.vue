<script setup>
// Dashboard overview page: KPI cards + a recent-activity table.
// Data is mocked inline here — swap for a real API call in services/ later.
import { ref } from 'vue'
import StatisticCard from '@/components/cards/StatisticCard.vue'
import DashboardCard from '@/components/cards/DashboardCard.vue'
import PageHeader from '@/components/common/PageHeader.vue'
import AppTable from '@/components/tables/AppTable.vue'
import UserAvatar from '@/components/common/UserAvatar.vue'
import { formatCurrency, formatRelativeTime } from '@/utils'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const stats = [
  { label: 'Total Revenue', value: formatCurrency(84250), icon: 'Money', iconColor: '#10b981'},
  { label: 'Active Users', value: '3,482', icon: 'User', iconColor: '#6366f1'},
  { label: 'New Orders', value: '642', icon: 'ShoppingCart', iconColor: '#f59e0b'},
  { label: 'Conversion Rate', value: '4.8%', icon: 'TrendCharts', iconColor: '#ec4899' },
]

const activityColumns = [
  { prop: 'user', label: 'User', slot: 'user', minWidth: 200 },
  { prop: 'action', label: 'Action', minWidth: 200 },
  { prop: 'status', label: 'Status', slot: 'status', width: 130 },
  { prop: 'date', label: 'Date', slot: 'date', width: 140 },
]

const activity = ref([
  { id: 1, user: 'Sophia Carter', avatar: 'https://i.pravatar.cc/150?img=47', action: 'Created a new project', status: 'completed', date: Date.now() - 1000 * 60 * 12 },
  { id: 2, user: 'Liam Johnson', avatar: 'https://i.pravatar.cc/150?img=12', action: 'Updated billing information', status: 'completed', date: Date.now() - 1000 * 60 * 60 * 3 },
  { id: 3, user: 'Olivia Smith', avatar: 'https://i.pravatar.cc/150?img=32', action: 'Requested a refund', status: 'pending', date: Date.now() - 1000 * 60 * 60 * 8 },
  { id: 4, user: 'Noah Williams', avatar: 'https://i.pravatar.cc/150?img=15', action: 'Cancelled subscription', status: 'failed', date: Date.now() - 1000 * 60 * 60 * 26 },
  { id: 5, user: 'Emma Brown', avatar: 'https://i.pravatar.cc/150?img=25', action: 'Upgraded to Pro plan', status: 'completed', date: Date.now() - 1000 * 60 * 60 * 50 },
])

const statusType = { completed: 'success', pending: 'warning', failed: 'danger' }
</script>

<template>
  <div>
    <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4 mb-6">
      <StatisticCard v-for="s in stats" :key="s.label" v-bind="s" />
    </div>

    <div class="grid grid-cols-1 xl:grid-cols-3 gap-6">
      <DashboardCard title="Revenue Overview" subtitle="Last 6 months" class="xl:col-span-2">
        <div
          class="h-64 rounded-lg flex items-center justify-center text-sm"
          style="background: var(--color-surface-hover); color: var(--color-text-secondary)"
        >
          Chart placeholder — plug in Recharts / ECharts once your backend is connected.
        </div>
      </DashboardCard>

      <DashboardCard title="Quick Stats" subtitle="This month">
        <div class="space-y-4 pt-1">
          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-secondary)">Task completion</span>
            <span class="text-sm font-medium">78%</span>
          </div>
          <el-progress :percentage="78" :stroke-width="8" color="#6366f1" />

          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-secondary)">Server uptime</span>
            <span class="text-sm font-medium">99.9%</span>
          </div>
          <el-progress :percentage="99.9" :stroke-width="8" color="#10b981" />

          <div class="flex items-center justify-between">
            <span class="text-sm" style="color: var(--color-text-secondary)">Storage used</span>
            <span class="text-sm font-medium">42%</span>
          </div>
          <el-progress :percentage="42" :stroke-width="8" color="#f59e0b" />
        </div>
      </DashboardCard>
    </div>

    <DashboardCard title="Recent Activity" subtitle="Latest actions across your account" class="mt-6" body-class="!px-0 !pb-0">
      <AppTable :columns="activityColumns" :data="activity" row-key="id">
        <template #user="{ row }">
          <div class="flex items-center gap-2.5">
            <UserAvatar :src="row.avatar" :name="row.user" :size="30" />
            <span class="text-sm font-medium">{{ row.user }}</span>
          </div>
        </template>
        <template #status="{ row }">
          <el-tag :type="statusType[row.status]" size="small" effect="light" round>
            {{ row.status }}
          </el-tag>
        </template>
        <template #date="{ row }">
          <span class="text-sm" style="color: var(--color-text-secondary)">
            {{ formatRelativeTime(row.date) }}
          </span>
        </template>
      </AppTable>
    </DashboardCard>
  </div>
</template>
