import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  // /device and /channel are wrapped in layout so the sidebar is always visible.
  // Other business pages are registered dynamically by the backend menu system.
  {
    path: '/device',
    component: () => import('@/view/layout/index.vue'),
    children: [
      {
        path: '',
        name: 'Device',
        meta: { title: '设备管理', requireAuth: true },
        component: () => import('@/view/device/index.vue')
      }
    ]
  },
  {
    path: '/channel',
    component: () => import('@/view/layout/index.vue'),
    children: [
      {
        path: '',
        name: 'Channel',
        meta: { title: '通道管理', requireAuth: true },
        component: () => import('@/view/channel/index.vue')
      }
    ]
  },
  {
    path: '/:catchAll(.*)',
    meta: { closeTab: true },
    component: () => import('@/view/error/index.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
