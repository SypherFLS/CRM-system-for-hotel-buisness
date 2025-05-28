<template>
  <nav class="navbar">
    <div class="navbar-container">
      <div class="logo">
        <img :src="icon" alt="Логотип" />
        <span>LGBR-HTL</span>
      </div>
      <div class="nav-buttons">
        <router-link
          v-for="route in routesWithIcon"
          :key="route.path"
          :to="route.path"
          class="nav-button"
          :title="route.name"
          @mouseenter="hover(route.path)"
          @mouseleave="stopHover"
          :class="{ 'active': activeRoute === route.path }"
        >
          <img :src="route.iconUrl" :alt="route.name" />
          <span class="nav-button-text">{{ route.name }}</span>
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'

import icon from '@/assets/icon.png'
import emp from '@/assets/employees.png'
import main from '@/assets/main.png'
import rooms from '@/assets/rooms.png'
import clients from '@/assets/clients.png'
import mop from '@/assets/cleaning.png'

const route = useRoute()
const activeRoute = computed(() => route.path)

const routes = [
  { path: '/', name: 'Главная', icon: main },
  { path: '/rooms', name: 'Номера', icon: rooms },
  { path: '/clients', name: 'Клиенты', icon: clients },
  { path: '/employees', name: 'Сотрудники', icon: emp },
  { path: '/cleanings', name: 'Уборки', icon: mop },
]

const routesWithIcon = computed(() => routes.map(r => ({
  ...r,
  iconUrl: typeof r.icon === 'string' ? r.icon : r.icon.default || r.icon
})))

const hover = (path) => { /* для анимации */ }
const stopHover = () => { /* для анимации */ }
</script>


<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background: #fff;
  box-shadow: 0 16px 40px rgba(0, 0, 0, 0.08);
  border-radius: 0 0 18px 18px;
  z-index: 1000;
  padding: 10px 0;
  font-family: 'Montserrat', sans-serif;
}

.navbar-container {
  width: 90%;
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 700;
  font-size: 1.5rem;
  color: #a83b3b;
}

.logo img {
  width: 36px;
  height: 36px;
}

.nav-buttons {
  display: flex;
  gap: 10px;
}

.nav-button {
  background: #f8f8f8;
  border: 2px solid #e1e4e8;
  border-radius: 12px;
  padding: 10px 20px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
  color: #6d6d6d;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.nav-button:hover {
  background: #a83b3b;
  color: #fff;
  border-color: #a83b3b;
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(168, 59, 59, 0.1);
}

.nav-button.active {
  background: #a83b3b;
  color: #fff;
  border-color: #a83b3b;
}

.nav-button img {
  width: 24px;
  height: 24px;
}

.nav-button-text {
  display: inline-block;
}
</style>
