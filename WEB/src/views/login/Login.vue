<script setup>
// Login page — rendered inside AuthLayout.
import { reactive, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { notify } from '@/utils/notify'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const formRef = ref(null)
const loading = ref(false)
const showPassword = ref(false)

const form = reactive({
  email: '',
  password: '',
})

const rules = {
  email: [
    { required: true, message: 'Email is required', trigger: 'blur' },
    { type: 'email', message: 'Enter a valid email address', trigger: ['blur', 'change'] },
  ],
  password: [
    { required: true, message: 'Password is required', trigger: 'blur' },
    { min: 6, message: 'Password must be at least 6 characters', trigger: 'blur' },
  ],
}

async function handleSubmit() {
  if (!formRef.value) return
  try {
    await formRef.value.validate()
  } catch {
    return
  }

  loading.value = true
  try {
    await authStore.login({ email: form.email, password: form.password })
    notify.success(`Welcome back, ${authStore.userName}!`)
    router.push(route.query.redirect || { name: 'Dashboard' })
  } catch (err) {
    notify.error(err?.message || 'Invalid email or password.')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div>
    <h2 class="text-xl font-semibold mb-1" style="color: var(--color-text-primary)">
      សូមស្វាគមន៍ការត្រឡប់មកវិញ
    </h2>
    <p class="text-sm mb-6" style="color: var(--color-text-secondary)">
      ចូលដើម្បីចូលប្រើផ្ទាំងគ្រប់គ្រងរបស់អ្នក
    </p>

    <el-form ref="formRef" :model="form" :rules="rules" label-position="top" @submit.prevent>
      <el-form-item label="Email" prop="email">
        <el-input v-model="form.email" placeholder="you@example.com" size="large">
          <template #prefix><el-icon><Message /></el-icon></template>
        </el-input>
      </el-form-item>

      <el-form-item label="Password" prop="password">
        <el-input
          v-model="form.password"
          :type="showPassword ? 'text' : 'password'"
          placeholder="••••••••"
          size="large"
          @keyup.enter="handleSubmit"
        >
          <template #prefix><el-icon><Lock /></el-icon></template>
          <template #suffix>
            <el-icon class="cursor-pointer" @click="showPassword = !showPassword">
              <component :is="showPassword ? 'View' : 'Hide'" />
            </el-icon>
          </template>
        </el-input>
      </el-form-item>

      <el-button type="primary" size="large" class="w-full" :loading="loading" @click="handleSubmit">
        Sign In
      </el-button>
    </el-form>
  </div>
</template>
