# Admin Dashboard Starter

A production-ready Vue 3 + Vite admin dashboard starter. Everything is wired
up and working with mock data — you only need to write API services to
connect it to your real backend.

## Tech Stack

Vue 3 (Composition API + `<script setup>`) · Vite · Vue Router · Pinia ·
Axios · Element Plus · Tailwind CSS · VueUse · Day.js · SheetJS (xlsx) ·
ESLint + Prettier

## Getting Started

```bash
npm install
npm run dev
```

Open http://localhost:5173. The login page is pre-filled with demo
credentials — just click **Sign In** (auth is mocked, see below).

```bash
npm run build      # production build to dist/
npm run preview    # preview the production build
npm run lint        # eslint --fix
npm run format       # prettier --write src/
```

## Project Structure

```
src/
  assets/            static assets
  components/
    common/          AppButton, AppInput, AppSelect, AppDatePicker, UserAvatar, PageHeader
    forms/            AppForm (validation/loading/reset/submit wrapper)
    tables/            AppTable, AppPagination, AppSearch
    dialogs/           AppDialog, ConfirmDialog, DeleteDialog
    cards/             StatisticCard, DashboardCard
    layout/            PageContainer
    loading/           LoadingOverlay (global, driven by the loading store)
    empty/             EmptyState, ErrorState
  composables/        useTable, useConfirm, useDevice
  config/             menu.js (sidebar config), constants.js
  layouts/            AuthLayout, DashboardLayout (+ Navbar/Sidebar/Breadcrumb/Footer)
  router/             index.js (routes), guards.js (auth/permission guards)
  services/           api.js (axios instance), auth.service.js
  stores/             auth.js, app.js, loading.js (Pinia)
  styles/             index.scss (design tokens, dark mode, Tailwind layers)
  utils/              formatDate, formatCurrency, debounce, throttle, copyText,
                       downloadFile, exportExcel, notify.js
  views/
    dashboard/        Dashboard.vue
    users/             Users.vue (full CRUD example: search/sort/paginate/create/edit/delete)
    profile/           Profile.vue
    settings/          Settings.vue
    login/              Login.vue
    error/              NotFound.vue
```

## Connecting Your Real Backend

Everything currently runs against **mock data** so the app works out of the
box. To connect your API:

1. **Set the base URL** in `.env`:
   ```
   VITE_API_BASE_URL=https://your-api.com/api
   ```
2. **Auth**: open `src/services/auth.service.js` and flip `USE_MOCK = false`
   (or delete the mock branches). The methods already call the shared
   `api` axios instance from `src/services/api.js`, which has request/response
   interceptors that:
   - Attach the Bearer token automatically
   - Show the global loading overlay
   - Refresh the access token on `401` and retry the original request
   - Redirect to `/login` when the refresh token is invalid/expired
   - Show an error toast for failed requests (unless `{ silent: true }` or
     `{ suppressErrorToast: true }` is passed in the request config)
3. **New resources** (e.g. Users, Products, Orders): create a new file such
   as `src/services/users.service.js` that exports functions calling `api`,
   then swap the mock `fetchUsers` function inside `src/views/users/Users.vue`
   for your service call. The `useTable` composable already expects the
   `{ rows, total }` shape a paginated list endpoint typically returns.

## Theming

Design tokens (colors, spacing, sidebar width, etc.) live in
`src/styles/index.scss` as CSS variables, with a `.dark` class override for
dark mode. Toggle dark mode from the navbar (moon/sun icon) or
`useAppStore().toggleTheme()`.

## Auth & Permissions

- `src/stores/auth.js` — token, refresh token, user, persisted to
  localStorage via `pinia-plugin-persistedstate`.
- `src/router/guards.js` — redirects unauthenticated users to `/login`,
  redirects logged-in users away from `/login`, and blocks routes whose
  `meta.permission` the user doesn't have (see the `Users` route in
  `src/router/index.js` for an example, `permission: 'users.view'`).
- `authStore.hasPermission('some.permission')` / `authStore.hasRole('admin')`
  can be used anywhere (menu config, buttons, etc.) to conditionally show UI.

## Responsive Behavior

`useDevice()` (used in `DashboardLayout.vue`) watches the viewport width and
switches the sidebar to an `el-drawer` below the `tablet` breakpoint
(configurable in `src/config/constants.js`). Tables scroll horizontally on
small screens automatically via Element Plus.
