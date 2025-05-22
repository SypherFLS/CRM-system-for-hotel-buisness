import { createRouter, createWebHistory } from 'vue-router';
import Home from '@/views/Home.vue'; 
import Hotel from '@/views/hotel.vue'; 
import Brone from '@/views/redaction.vue'; 
import Rew from '@/views/rew.vue';

const routes = [
  { path: '/', name: 'home', component: Home },
  { path: '/hotel', name: 'hotel', component: Hotel },
  { path: '/brone', name: 'brone', component: Brone },
  { path: '/rew', name: 'rew', component: Rew },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
