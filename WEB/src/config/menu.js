/**
 * Sidebar menu configuration.
 * Add/remove entries here — Sidebar.vue renders this list recursively so
 * nested menus (children) work out of the box.
 *
 * Fields:
 *  - title: label shown in the sidebar
 *  - icon: Element Plus icon component name (auto-registered globally)
 *  - path: route path (omit for a parent-only item with children)
 *  - permission: optional permission string checked via auth store
 *  - children: optional nested items
 */
export const menuConfig = [
  {
    title: 'Dashboard',
    icon: 'Odometer',
    path: '/dashboard',
  },
  {
    title: 'Settings',
    icon: 'Setting',
    children: [
      { title: 'General', icon: 'Tools', path: '/settings' },
      {title: 'អ្នកប្រេីប្រាស់', icon: 'User',path: '/users'},
      {title: 'សិទ្ធ',icon: 'UserFilled',path: '/role'},
      {title: 'អ្នកនិពន្ធ', icon: 'UserFilled',path: '/author'},
      {title: 'ប្រភេទសៀវភៅ', icon: 'Document',path: '/category'}
    ],
  },
]
