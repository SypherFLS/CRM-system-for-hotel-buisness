<template>
  <div class="rooms-page">
    <div class="rooms-container">
      <h1>Номера отеля</h1>
      <div class="filters">
        <input v-model="searchQuery" type="text" placeholder="Поиск по номеру или этажу" class="search-input" />
        <div class="filter-buttons">
          <button @click="filterOccupied = null" :class="{ active: filterOccupied === null }">Все</button>
          <button @click="filterOccupied = true" :class="{ active: filterOccupied === true }">Занятые</button>
          <button @click="filterOccupied = false" :class="{ active: filterOccupied === false }">Свободные</button>
          <button @click="filterHasCleaning = !filterHasCleaning" :class="{ active: filterHasCleaning }">С уборкой</button>
        </div>
      </div>
      <button class="add-button" @click="openAddModal">+ Добавить номер</button>
      <div v-if="loading" class="loading">Загрузка...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else class="room-grid">
        <div
          v-for="room in filteredRooms"
          :key="room.id"
          class="room-card"
          @click="selectRoom(room)"
          :class="{ 'occupied': isRoomOccupied(room.id), 'has-cleaning': hasCleaning(room.id) }"
        >
          <div class="card-content">
            <h3>Этаж {{ room.floor }}, №{{ room.id }}</h3>
            <p>{{ room.room_type }}</p>
            <div class="price">{{ room.price_per_day }}₽/ночь</div>
            <div class="capacity">{{ room.capacity }} мест</div>
            <div class="status-badge">
              <span v-if="isRoomOccupied(room.id)" class="occupied-badge">Занят</span>
              <span v-else class="free-badge">Свободен</span>
              <span v-if="hasCleaning(room.id)" class="cleaning-badge">Уборка</span>
            </div>
            <div class="room-actions">
              <button class="edit-button" @click.stop="openEditModal(room)">✏️</button>
              <button class="delete-button" @click.stop="deleteRoom(room.id)">🗑️</button>
            </div>
          </div>
        </div>
      </div>
      <!-- Модальное окно добавления/редактирования -->
      <div v-if="showAddEditModal" class="modal-overlay" @click="closeAddEditModal">
        <div class="modal" @click.stop>
          <button class="close-modal" @click="closeAddEditModal">✕</button>
          <div class="modal-header">
            <h2>{{ editingRoom ? 'Редактировать номер' : 'Добавить номер' }}</h2>
          </div>
          <div class="modal-content">
            <form @submit.prevent="saveRoom">
              <div v-if="!editingRoom" class="form-group">
                <label>Номер:</label>
                <input v-model="roomForm.id" type="number" min="1" :disabled="autoId" />
                <label>
                  <input type="checkbox" v-model="autoId" />
                  Автоматически (следующий свободный)
                </label>
              </div>
              <div v-if="editingRoom" class="form-group">
                <label>Номер:</label>
                <input v-model="roomForm.id" type="number" min="1" disabled />
              </div>
              <div class="form-group">
                <label>Этаж:</label>
                <input v-model="roomForm.floor" type="number" min="1" required />
              </div>
              <div class="form-group">
                <label>Тип:</label>
                <input v-model="roomForm.room_type" type="text" required />
              </div>
              <div class="form-group">
                <label>Цена за ночь:</label>
                <input v-model="roomForm.price_per_day" type="number" min="0" required />
              </div>
              <div class="form-group">
                <label>Мест:</label>
                <input v-model="roomForm.capacity" type="number" min="1" required />
              </div>
              <div class="form-actions">
                <button type="submit" class="save-button">{{ editingRoom ? 'Сохранить' : 'Добавить' }}</button>
                <button type="button" class="cancel-button" @click="closeAddEditModal">Отмена</button>
              </div>
            </form>
          </div>
        </div>
      </div>
      <!-- Модальное окно просмотра -->
      <div v-if="selectedRoom" class="modal-overlay" @click="selectedRoom = null">
        <div class="modal">
          <button class="close-modal" @click="selectedRoom = null">✕</button>
          <div class="modal-header">
            <h2>Этаж {{ selectedRoom.floor }}, №{{ selectedRoom.id }}</h2>
            <p class="room-type">{{ selectedRoom.room_type }}</p>
          </div>
          <div class="modal-content">
            <div class="info-row">
              <span class="label">Цена:</span>
              <span class="value">{{ selectedRoom.price_per_day }}₽/ночь</span>
            </div>
            <div class="info-row">
              <span class="label">Мест:</span>
              <span class="value">{{ selectedRoom.capacity }}</span>
            </div>
            <div class="info-row">
              <span class="label">Статус:</span>
              <span class="value">
                <span v-if="isRoomOccupied(selectedRoom.id)" class="occupied-badge">Занят</span>
                <span v-else class="free-badge">Свободен</span>
              </span>
            </div>
            <div v-if="isRoomOccupied(selectedRoom.id)" class="info-row">
              <span class="label">Клиент:</span>
              <span class="value">{{ getClientByRoom(selectedRoom.id).full_name }}</span>
            </div>
            <div v-if="isRoomOccupied(selectedRoom.id)" class="info-row">
              <span class="label">Прибытие:</span>
              <span class="value">{{ formatDate(getClientByRoom(selectedRoom.id).arrival_date) }}</span>
            </div>
            <div v-if="isRoomOccupied(selectedRoom.id)" class="info-row">
              <span class="label">Дней оплачено:</span>
              <span class="value">{{ getClientByRoom(selectedRoom.id).payment_days }}</span>
            </div>
            <div v-if="isRoomOccupied(selectedRoom.id)" class="info-row">
              <button class="release-button" @click.stop="releaseRoom(selectedRoom.id)">Освободить номер</button>
            </div>
            <div class="cleanings">
              <h3>Уборки</h3>
              <ul>
                <li v-for="cleaning in getCleaningsByRoom(selectedRoom.id)" :key="cleaning.id">
                  {{ formatDate(cleaning.cleaning_date) }} — {{ getEmployeeById(cleaning.employee_id).full_name }}
                </li>
              </ul>
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

const rooms = ref([])
const clients = ref([])
const employees = ref([])
const cleanings = ref([])
const selectedRoom = ref(null)
const loading = ref(true)
const error = ref(null)
const searchQuery = ref('')
const filterOccupied = ref(null)
const filterHasCleaning = ref(false)
const showAddEditModal = ref(false)
const editingRoom = ref(null)
const autoId = ref(true)
const roomForm = ref({
  id: '',
  floor: '',
  room_type: '',
  price_per_day: '',
  capacity: ''
})

function isRoomOccupied(roomId) {
  return clients.value.some(client => client.room_id === roomId)
}

function hasCleaning(roomId) {
  return cleanings.value.some(cleaning => cleaning.room_id === roomId)
}

function getClientByRoom(roomId) {
  return clients.value.find(client => client.room_id === roomId) || {}
}

function getCleaningsByRoom(roomId) {
  return cleanings.value.filter(cleaning => cleaning.room_id === roomId)
}

function getEmployeeById(employeeId) {
  return employees.value.find(emp => emp.id === employeeId) || { full_name: 'Неизвестно' }
}

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

const filteredRooms = computed(() => {
  let result = rooms.value
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    result = result.filter(room =>
      room.id.toString().includes(query) ||
      room.floor.toString().includes(query)
    )
  }
  if (filterOccupied.value !== null) {
    result = result.filter(room => isRoomOccupied(room.id) === filterOccupied.value)
  }
  if (filterHasCleaning.value) {
    result = result.filter(room => hasCleaning(room.id))
  }
  return result.sort((a, b) => a.id - b.id)
})

async function loadData() {
  try {
    const [roomsRes, clientsRes, employeesRes, cleaningsRes] = await Promise.all([
      axios.get('http://localhost:8000/api/rooms'),
      axios.get('http://localhost:8000/api/clients'),
      axios.get('http://localhost:8000/api/employees'),
      axios.get('http://localhost:8000/api/cleanings')
    ])
    rooms.value = roomsRes.data
    clients.value = clientsRes.data
    employees.value = employeesRes.data
    cleanings.value = cleaningsRes.data
    loading.value = false
  } catch (err) {
    error.value = 'Ошибка загрузки: ' + (err.response?.data?.message || err.message)
    loading.value = false
    console.error('Ошибка загрузки:', err)
  }
}

onMounted(loadData)

function selectRoom(room) {
  selectedRoom.value = room
}

function openAddModal() {
  roomForm.value = { id: '', floor: '', room_type: '', price_per_day: '', capacity: '' }
  editingRoom.value = null
  autoId.value = true
  showAddEditModal.value = true
}

function openEditModal(room) {
  roomForm.value = { ...room }
  editingRoom.value = room
  autoId.value = false
  showAddEditModal.value = true
}

function closeAddEditModal() {
  showAddEditModal.value = false
}

async function saveRoom() {
  try {
    const data = { ...roomForm.value }
    // Если автоматический ID — удаляем поле id
    if (autoId.value) delete data.id
    // Проверка обязательных полей
    if (!data.floor || !data.room_type || !data.price_per_day || !data.capacity) {
      error.value = 'Заполните все обязательные поля'
      return
    }
    if (editingRoom.value) {
      await axios.put(`http://localhost:8000/api/rooms/${data.id}`, data)
    } else {
      await axios.post('http://localhost:8000/api/rooms', data)
    }
    await loadData()
    closeAddEditModal()
  } catch (err) {
    error.value = 'Ошибка сохранения: ' + (err.response?.data?.message || err.message)
    console.error('Ошибка сохранения:', err)
  }
}

async function deleteRoom(id) {
  if (!confirm('Удалить номер?')) return
  try {
    await axios.delete(`http://localhost:8000/api/rooms/${id}`)
    await loadData()
  } catch (err) {
    error.value = 'Ошибка удаления: ' + (err.response?.data?.message || err.message)
    console.error('Ошибка удаления:', err)
  }
}

async function releaseRoom(roomId) {
  if (!confirm('Освободить номер?')) return
  try {
    const client = clients.value.find(c => c.room_id === roomId)
    if (!client) {
      error.value = 'Клиент не найден'
      return
    }
    await axios.delete(`http://localhost:8000/api/clients/${client.id}`)
    await loadData()
    selectedRoom.value = null
  } catch (err) {
    error.value = 'Ошибка освобождения: ' + (err.response?.data?.message || err.message || 'Network Error')
    console.error('Ошибка освобождения:', err)
  }
}
</script>

<style scoped>
/* Стили остаются такими же, как в предыдущем примере */
.rooms-page {
  min-height: 100vh;
  background: #f9f9f9;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 0;
}
.rooms-container {
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
.filter-buttons {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}
.filter-buttons button {
  padding: 8px 16px;
  border: 2px solid #e1e4e8;
  border-radius: 8px;
  font-size: 0.9rem;
  font-family: 'Montserrat', sans-serif;
  background: #fff;
  cursor: pointer;
  transition: all 0.2s ease;
}
.filter-buttons button.active {
  background: #a83b3b;
  color: #fff;
  border-color: #a83b3b;
}
.add-button {
  margin-bottom: 24px;
  padding: 12px 24px;
  border: none;
  border-radius: 12px;
  background: #a83b3b;
  color: #fff;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  font-family: 'Montserrat', sans-serif;
}
.add-button:hover {
  background: #8c2b2b;
  transform: translateY(-2px);
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
.room-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 24px;
}
.room-card {
  background: #f8f8f8;
  border-radius: 16px;
  padding: 24px;
  transition: all 0.3s ease;
  cursor: pointer;
  border: 2px solid #e1e4e8;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
  position: relative;
  overflow: hidden;
}
.room-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 16px 32px rgba(168, 59, 59, 0.1);
  border-color: #a83b3b;
}
.room-card.occupied {
  border-color: #e74c3c;
  background: #fadbd8;
}
.room-card.has-cleaning::after {
  content: '';
  position: absolute;
  top: 8px;
  right: 8px;
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: #4caf50;
  box-shadow: 0 0 8px rgba(76, 175, 80, 0.5);
}
.room-card .card-content {
  height: 100%;
  display: flex;
  flex-direction: column;
}
.room-card h3 {
  color: #a83b3b;
  font-size: 1.5rem;
  margin-bottom: 12px;
  font-weight: 600;
}
.room-card p {
  color: #6d6d6d;
  font-size: 1rem;
  margin-bottom: 8px;
  flex-grow: 1;
}
.price {
  font-weight: 700;
  color: #a83b3b;
  font-size: 1.2rem;
}
.capacity {
  color: #6d6d6d;
}
.status-badge {
  margin-top: 8px;
  display: flex;
  gap: 8px;
}
.status-badge span {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.8rem;
  font-weight: 500;
}
.occupied-badge {
  background: #e74c3c;
  color: #fff;
}
.free-badge {
  background: #4caf50;
  color: #fff;
}
.cleaning-badge {
  background: #2196f3;
  color: #fff;
}
.room-actions {
  margin-top: 16px;
  display: flex;
  gap: 10px;
}
.edit-button, .delete-button {
  padding: 8px 12px;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.2s ease;
}
.edit-button {
  background: #2196f3;
  color: #fff;
}
.edit-button:hover {
  background: #1976d2;
  transform: translateY(-2px);
}
.delete-button {
  background: #e74c3c;
  color: #fff;
}
.delete-button:hover {
  background: #c0392b;
  transform: translateY(-2px);
}
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
  backdrop-filter: blur(4px);
  animation: fadeIn 0.3s ease;
}
.modal {
  background: #fff;
  border-radius: 20px;
  padding: 36px;
  max-width: 600px;
  width: 90%;
  box-shadow: 0 24px 48px rgba(0, 0, 0, 0.12);
  position: relative;
  border: 1px solid rgba(0, 0, 0, 0.08);
  animation: slideUp 0.3s ease;
  overflow: hidden;
}
.modal::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 6px;
  background: #a83b3b;
  border-radius: 20px 20px 0 0;
}
.modal-header {
  margin-bottom: 24px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 8px;
}
.modal-header h2 {
  color: #a83b3b;
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 4px;
  font-family: 'Montserrat', sans-serif;
}
.modal-header .room-type {
  color: #6d6d6d;
  font-size: 1.2rem;
  font-weight: 500;
}
.modal-content {
  margin-top: 16px;
}
.info-row {
  display: flex;
  justify-content: space-between;
  padding: 16px 0;
  border-bottom: 1px solid rgba(0, 0, 0, 0.08);
  font-size: 1.1rem;
}
.info-row:last-child {
  border-bottom: none;
}
.info-row .label {
  font-weight: 600;
  color: #6d6d6d;
}
.info-row .value {
  color: #333;
  font-weight: 500;
}
.info-row .value .occupied-badge,
.info-row .value .free-badge {
  display: inline-block;
  padding: 4px 10px;
  border-radius: 50px;
  font-size: 0.9rem;
  font-weight: 600;
}
.info-row .value .occupied-badge {
  background: #e74c3c;
  color: #fff;
}
.info-row .value .free-badge {
  background: #4caf50;
  color: #fff;
}
.form-group {
  margin-bottom: 16px;
}
.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 600;
  color: #6d6d6d;
}
.form-group input {
  width: 100%;
  padding: 12px 16px;
  border: 2px solid #e1e4e8;
  border-radius: 10px;
  font-size: 1rem;
  font-family: 'Montserrat', sans-serif;
  outline: none;
  transition: all 0.3s ease;
}
.form-group input:focus {
  border-color: #a83b3b;
  box-shadow: 0 0 8px rgba(168, 59, 59, 0.2);
}
.form-actions {
  margin-top: 24px;
  display: flex;
  gap: 12px;
}
.save-button, .cancel-button, .release-button {
  padding: 12px 24px;
  border: none;
  border-radius: 10px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  font-family: 'Montserrat', sans-serif;
}
.save-button {
  background: #a83b3b;
  color: #fff;
}
.save-button:hover {
  background: #8c2b2b;
  transform: translateY(-2px);
}
.cancel-button {
  background: #e1e4e8;
  color: #333;
}
.cancel-button:hover {
  background: #d1d5d8;
  transform: translateY(-2px);
}
.release-button {
  background: #e74c3c;
  color: #fff;
  width: 100%;
  margin-top: 12px;
}
.release-button:hover {
  background: #c0392b;
  transform: translateY(-2px);
}
.cleanings {
  margin-top: 24px;
  border-top: 1px solid rgba(0, 0, 0, 0.08);
  padding-top: 16px;
}
.cleanings h3 {
  color: #a83b3b;
  font-size: 1.3rem;
  margin-bottom: 12px;
  font-family: 'Montserrat', sans-serif;
}
.cleanings ul {
  list-style: none;
  padding: 0;
}
.cleanings li {
  padding: 10px 0;
  border-bottom: 1px solid rgba(0, 0, 0, 0.05);
  display: flex;
  justify-content: space-between;
}
.close-modal {
  position: absolute;
  top: 16px;
  right: 16px;
  background: none;
  border: none;
  font-size: 1.8rem;
  cursor: pointer;
  color: #a83b3b;
  transition: all 0.2s ease;
}
.close-modal:hover {
  color: #e74c3c;
  transform: rotate(90deg);
}
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
@keyframes slideUp {
  from { transform: translateY(20px); opacity: 0; }
  to { transform: translateY(0); opacity: 1; }
}
</style>
