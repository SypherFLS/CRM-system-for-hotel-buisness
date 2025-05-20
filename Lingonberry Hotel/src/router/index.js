import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/views/Home.vue'; 
import Hotel from '@/views/hotel.vue'; 

const routes = [
  { path: '/', name: 'home', component: Home },
  { path: '/hotel', name: 'hotel', component: Hotel },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
