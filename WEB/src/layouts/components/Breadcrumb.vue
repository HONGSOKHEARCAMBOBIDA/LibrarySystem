<script setup>
// Auto-generates a breadcrumb trail from the matched route records.
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const crumbs = computed(() =>
  route.matched
    .filter((r) => r.meta?.title)
    .map((r) => ({ title: r.meta.title, path: r.path }))
)
</script>

<template>
  <el-breadcrumb v-if="crumbs.length" separator-icon="ArrowRight">
    <el-breadcrumb-item :to="{ name: 'Dashboard' }">Home</el-breadcrumb-item>
    <el-breadcrumb-item v-for="crumb in crumbs" :key="crumb.path">
      {{ crumb.title }}
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>
