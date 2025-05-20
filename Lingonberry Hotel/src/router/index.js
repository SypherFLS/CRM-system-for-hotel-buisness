import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/views/Home.vue'; // ваша главная страница
import Hotel from '@/views/hotel.vue'; // страница отелей (если есть)

const routes = [
  { path: '/', name: 'home', component: Home },
  { path: '/hotel', name: 'hotel', component: Hotel },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
