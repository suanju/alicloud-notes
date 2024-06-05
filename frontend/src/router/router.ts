import { createRouter, createWebHistory } from 'vue-router'
  
import Login from '@/view/login/login.vue'

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Login },
    { path: '/login', component: Login },
  ]
})    