<template>
  <div class="clients-view">
    <div class="clients-container">
      <h1>Клиенты и занятые номера</h1>
      <div class="filters">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Поиск по имени клиента или номеру паспорта"
          class="search-input"
        />
      </div>
      <div v-if="loading" class="loading">Загрузка...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else class="client-list">
        <div
          v-for="client in filteredClients"
          :key="client.id"
          class="client-card"
        >
          <div class="card-content">
            <h3>{{ client.full_name }}</h3>
            <div class="info-row">
              <span class="label">Паспорт:</span>
              <span class="value">{{ client.passport_number || '—' }}</span>
            </div>
            <div class="info-row">
              <span class="label">Город:</span>
              <span class="value">{{ client.city_from || '—' }}</span>
            </div>
            <div class="info-row">
              <span class="label">Номер:</span>
              <span class="value">{{ client.room_id ? `№${client.room_id}` : '—' }}</span>
            </div>
            <div class="info-row">
              <span class="label">Место:</span>
              <span class="value">{{ client.place_number }}</span>
            </div>
            <div class="info-row">
              <span class="label">Прибытие:</span>
              <span class="value">{{ formatDate(client.arrival_date) }}</span>
            </div>
            <div class="info-row">
              <span class="label">Дней оплачено:</span>
              <span class="value">{{ client.payment_days }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const clients = ref([])
const rooms = ref([])
const loading = ref(true)
const error = ref(null)
const searchQuery = ref('')

async function loadData() {
  try {
    const [clientsRes, roomsRes] = await Promise.all([
      axios.get('http://localhost:8000/api/clients'),
      axios.get('http://localhost:8000/api/rooms')
    ])
    clients.value = clientsRes.data
    rooms.value = roomsRes.data
    loading.value = false
  } catch (err) {
    error.value = 'Ошибка загрузки: ' + (err.response?.data?.message || err.message)
    loading.value = false
    console.error('Ошибка загрузки:', err)
  }
}

const filteredClients = computed(() => {
  if (!searchQuery.value) return clients.value
  const query = searchQuery.value.toLowerCase()
  return clients.value.filter(client =>
    client.full_name.toLowerCase().includes(query) ||
    (client.passport_number && client.passport_number.toLowerCase().includes(query)) ||
    (client.room_id && client.room_id.toString().includes(query))
  )
})

function formatDate(date) {
  if (!date) return '—'
  if (typeof date === 'string') {
    return new Date(date).toLocaleDateString('ru-RU')
  }
  if (date instanceof Date) {
    return date.toLocaleDateString('ru-RU')
  }
  return date
}

onMounted(loadData)
</script>

<style scoped>
.clients-view {
  min-height: 100vh;
  background: #f9f9f9;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 0;
}
.clients-container {
  width: 80%;
  max-width: 1200px;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 16px 40px rgba(0, 0, 0, 0.08);
  padding: 40px;
  margin: 0 auto;
}
h1 {
  color: #a83b3b;
  font-size: 2.5rem;
  margin-bottom: 32px;
  text-align: center;
  font-weight: 700;
  font-family: 'Montserrat', sans-serif;
}
.filters {
  margin-bottom: 32px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.search-input {
  width: 100%;
  max-width: 400px;
  padding: 12px 20px;
  border: 2px solid #e1e4e8;
  border-radius: 12px;
  font-size: 1rem;
  font-family: 'Montserrat', sans-serif;
  outline: none;
  transition: all 0.3s ease;
}
.search-input:focus {
  border-color: #a83b3b;
  box-shadow: 0 0 8px rgba(168, 59, 59, 0.2);
}
.loading, .error {
  text-align: center;
  margin: 40px 0;
  font-size: 1.2rem;
}
.error {
  color: #e74c3c;
  font-weight: bold;
}
.client-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 24px;
}
.client-card {
  background: #f8f8f8;
  border-radius: 16px;
  padding: 24px;
  border: 2px solid #e1e4e8;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}
.client-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 16px 32px rgba(168, 59, 59, 0.1);
  border-color: #a83b3b;
}
.card-content {
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.card-content h3 {
  color: #a83b3b;
  font-size: 1.5rem;
  margin-bottom: 12px;
  font-weight: 600;
}
.info-row {
  display: flex;
  justify-content: space-between;
  font-size: 1.1rem;
}
.info-row .label {
  font-weight: 600;
  color: #6d6d6d;
}
.info-row .value {
  color: #333;
  font-weight: 500;
}
</style>
