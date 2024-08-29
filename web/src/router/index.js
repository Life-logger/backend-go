import { createRouter, createWebHistory } from 'vue-router'

// 라우트 컴포넌트를 동적으로 로드합니다.
// const Home = () => import('@/pages/Home.vue')
const Login = () => import('../pages/Login.vue')
const Home = () => import('../pages/Home.vue')

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
