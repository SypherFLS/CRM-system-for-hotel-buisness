import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import RoomsView from '@/views/RoomsView.vue'
import ClientsView from '@/views/ClientsView.vue'
import EmployeesView from '@/views/EmployeesView.vue'
import CleaningsView from '@/views/CleaningsView.vue'

const routes = [
  { path: '/', name: 'home', component: HomeView },
  { path: '/rooms', name: 'rooms', component: RoomsView },
  { path: '/clients', name: 'clients', component: ClientsView },
  { path: '/employees', name: 'employees', component: EmployeesView },
  { path: '/cleanings', name: 'cleanings', component: CleaningsView },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
