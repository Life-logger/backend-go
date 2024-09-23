import { createRouter, createWebHistory } from 'vue-router'
// import Timer from '@/components/Timer.vue';
// import Timer from '/web/src/components/Timer.vue';
import Timer from '@/components/Timer.vue';
// import Home from 'S:/LifeLogger/backend-go/web/src/components/HomePage.vue';
import Calendar from 'S:/LifeLogger/backend-go/web/src/components/Calendar.vue';

// 라우트 컴포넌트를 동적으로 로드합니다.
// const Home = () => import('@/pages/Home.vue')
const Login = () => import('@/pages/Login.vue')
const HomePage = () => import('@/components/HomePage.vue')
// const Home = () => import('../pages/Home.vue')


const routes = [
  { path: '/', name: 'HomePage', component: HomePage },
  { path: '/login', name: 'Login', component: Login },
  { path: '/timer', name: 'Timer', component: Timer },
  { path: '/Calendar', name: 'Calendar', component: Calendar },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
