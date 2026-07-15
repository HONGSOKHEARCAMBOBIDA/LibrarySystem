<script setup>
import { ref, reactive, watch, onMounted } from 'vue';
import { getauthor } from '../../services/author.service';
import { useNotification } from '../../composables/useNotification';
import TableCustom from '../../components/tables/TableCustom.vue';
import AppButton from '../../components/button/AppButton.vue';
import AppInput from '../../components/input/AppInput.vue';
import { debounce } from "lodash-es";
import AppSelect from '../../components/common/AppSelect.vue';
import AppFilterBar from '../../components/common/AppFilterBar.vue';

const notify = useNotification();

const authors = ref([]);
const loading = ref(false);
const total = ref(0);
const page = ref(1);
const pageSize = ref(10);

const genderOptions = [
  { label: 'ប្រុស', value: 1 },
  { label: 'ស្រី', value: 2 },
];

const statusOptions = [
  { label: 'សកម្ម', value: 1 },
  { label: 'អសកម្ម', value: 0 },
];

const filters = reactive({ name: "", gender: null, is_active: null });

const columns = [
  { prop: "id", label: "ល.រ", width: 70 },
  { prop: "name", label: "ឈ្មោះ" },
  { prop: "gender", label: "ភេទ", slot: 'gender', width: 100 },
  { prop: "is_active", label: "ស្ថានភាព", slot: "isActive", width: 100 },
  { prop: "note", label: "ចំណាំ" },
  { prop: "faculty", label: "ក្នុងមហាវិទ្យាល័យ", slot: "faculty", width: 300 }
];

function buildParams() {
  const params = { page: page.value, pageSize: pageSize.value };
  if (filters.name) params.name = filters.name;
  if (filters.gender !== null && filters.gender !== "") params.gender = filters.gender;
  if (filters.is_active !== null && filters.is_active !== "") params.is_active = filters.is_active;
  return params;
}

async function fetchauthor() {
  loading.value = true;
  try {
    const res = await getauthor(buildParams());
    authors.value = res.data.data || [];
    total.value = res.data.pagination?.totalCount || 0;
  } catch (err) {
    notify.error(err?.response?.data?.message || err.message || "Failed to load authors");
  } finally {
    loading.value = false;
  }
}

function onReset() {
  filters.name = "";
  filters.gender = null;
  filters.is_active = null;
  page.value = 1;
  fetchauthor();
}

const debouncedFetch = debounce(() => {
  page.value = 1; // reset to first page whenever filters change
  fetchauthor();
}, 500);

// watch the actual filter values, not a function reference
watch(
  () => ({ ...filters }),
  debouncedFetch,
  { deep: true }
);

onMounted(fetchauthor);
</script>

<template>
  <div class="author-page">
    <AppFilterBar :fields="[
      { slot: 'name', span: 10 },
      { slot: 'gender', span: 6 },
      { slot: 'status', span: 5 },
    ]" :action-span="3">
      <template #name>
        <AppInput v-model="filters.name" placeholder="Search by name" />
      </template>

      <template #gender>
        <AppSelect label="ជ្រេីសរេីសភេទ" v-model="filters.gender" :options="genderOptions" placeholder="ប្រុស-ស្រី"
          clearable />
      </template>

      <template #status>
        <AppSelect label="ជ្រេីសរេីសស្ថានភាព" v-model="filters.is_active" :options="statusOptions"
          placeholder="សកម្ម-មិនសកម្ម" clearable />
      </template>

      <!-- <template #actions>
        <AppButton @click="onReset" size="default">សម្អាត</AppButton>
      </template> -->
    </AppFilterBar>

    <TableCustom :data="authors" :columns="columns" :loading="loading" :total="total" v-model:current-page="page"
      v-model:page-size="pageSize" @page-change="fetchauthor">
      <template #isActive="{ row }">
        <el-tag :type="row.is_active ? 'success' : 'info'">
          {{ row.is_active ? 'Active' : 'Inactive' }}
        </el-tag>
      </template>
      <template #gender="{ row }">
        <el-text>
          {{ row.gender === 1 ? 'ប្រុស' : 'ស្រី' }}
        </el-text>
      </template>
      <template #faculty="{ row }">
        <el-tag v-for="f in row.faculty" :key="f.id" class="faculty-tag" type="warning" effect="plain">
          {{ f.name }}
        </el-tag>
        <span v-if="!row.faculty || row.faculty.length === 0" class="text-muted">—</span>
      </template>
    </TableCustom>
  </div>
</template>

<style scoped>
.author-filters {
  display: flex;
  gap: 10px;
  margin-bottom: 16px;
  flex-wrap: wrap;
  align-items: center;
}

.faculty-tag {
  margin-right: 4px;
  margin-bottom: 4px;
}
</style>