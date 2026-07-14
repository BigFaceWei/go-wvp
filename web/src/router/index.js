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
  {
    path: '/device',
    name: 'Device',
    meta: {
      title: '设备管理',
      requireAuth: true
    },
    component: () => import('@/view/device/index.vue')
  },
  {
    path: '/channel',
    name: 'Channel',
    meta: {
      title: '通道管理',
      requireAuth: true
    },
    component: () => import('@/view/channel/index.vue')
  },
  {
    path: '/platform',
    name: 'Platform',
    meta: {
      title: '平台级联',
      requireAuth: true
    },
    component: () => import('@/view/platform/index.vue')
  },
  {
    path: '/alarm',
    name: 'Alarm',
    meta: {
      title: '报警管理',
      requireAuth: true
    },
    component: () => import('@/view/alarm/index.vue')
  },
  {
    path: '/position',
    name: 'Position',
    meta: {
      title: '移动位置',
      requireAuth: true
    },
    component: () => import('@/view/position/index.vue')
  },
  {
    path: '/record',
    name: 'Record',
    meta: {
      title: '录像管理',
      requireAuth: true
    },
    component: () => import('@/view/record/index.vue')
  },
  {
    path: '/streamProxy',
    name: 'StreamProxy',
    meta: {
      title: '拉流代理',
      requireAuth: true
    },
    component: () => import('@/view/streamProxy/index.vue')
  },
  {
    path: '/streamPush',
    name: 'StreamPush',
    meta: {
      title: '推流管理',
      requireAuth: true
    },
    component: () => import('@/view/streamPush/index.vue')
  },
  {
    path: '/group',
    name: 'Group',
    meta: {
      title: '业务分组',
      requireAuth: true
    },
    component: () => import('@/view/group/index.vue')
  },
  {
    path: '/region',
    name: 'Region',
    meta: {
      title: '行政区域',
      requireAuth: true
    },
    component: () => import('@/view/region/index.vue')
  },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true
    },
    component: () => import('@/view/error/index.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
