import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'

Vue.use(VueRouter)

  const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Index',
    component: () => import('../views/Index.vue')
  },
  {
    path: '/login',
    name: 'Auth/Login',
    component: () => import('../views/Auth/Login.vue')
  },
  {
    path: '/register',
    name: 'Auth/Register',
    component: () => import('../views/Auth/Register.vue')
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import(/* webpackChunkName: "dashboard" */ '../views/Dashboard.vue')
  },
  {
    path: '/sessions',
    name: 'Sessions',
    component: () => import(/* webpackChunkName: "sessions" */ '../views/Sessions.vue')
  },
  {
    path: '/applications',
    name: 'Applications',
    component: () => import(/* webpackChunkName: "applications" */ '../views/Applications.vue')
  },
  {
    path: '/storage',
    name: 'Storage',
    component: () => import(/* webpackChunkName: "storage" */ '../views/Storage.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
