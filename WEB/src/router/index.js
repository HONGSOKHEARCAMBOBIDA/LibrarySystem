import { createRouter, createWebHistory } from 'vue-router'
import { setupGuards } from './guards'

const DashboardLayout = () => import('@/layouts/DashboardLayout.vue')
const AuthLayout = () => import('@/layouts/AuthLayout.vue')

const routes = [
  {
    path: '/',
    component: AuthLayout,
    children: [
      {
        path: '',
        redirect: '/login',
      },
      {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/Login.vue'),
        meta: { title: 'Login', guestOnly: true },
      },
    ],
  },
  {
    path: '/',
    component: DashboardLayout,
    meta: { requiresAuth: true },
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/Dashboard.vue'),
        meta: { title: 'Dashboard', icon: 'Odometer' },
      },
      {
        path: 'users',
        name: 'Users',
        component: () => import('@/views/users/Users.vue'),
        meta: { title: 'Users', icon: 'User', },
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('@/views/profile/Profile.vue'),
        meta: { title: 'Profile', icon: 'UserFilled' },
      },
      {
        path: 'settings',
        name: 'Settings',
        component: () => import('@/views/settings/Settings.vue'),
        meta: { title: 'Settings', icon: 'Setting' },
      },
      {
        path: 'role',
        name: 'សិទ្ធ',
        component: () => import('@/views/role/role.vue'),
        meta: {title: 'សិទ្ធ', icone: 'UserFilled'}
      }
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/error/NotFound.vue'),
    meta: { title: 'Page Not Found' },
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  },
})

setupGuards(router)

export default router
