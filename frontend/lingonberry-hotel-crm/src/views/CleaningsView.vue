<template>
  <div class="cleaning-view">
    <div class="cleaning-container">
      <h1>Уборки</h1>
      <div class="controls">
        <button class="add-button" @click="openAddModal">+ Добавить уборку</button>
      </div>
      <div class="calendar">
        <div class="calendar-header">
          <button class="month-button" @click="prevMonth">←</button>
          <h2>{{ currentMonthName }} {{ currentYear }}</h2>
          <button class="month-button" @click="nextMonth">→</button>
        </div>
        <div class="weekdays">
          <span v-for="day in ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс']" :key="day">{{ day }}</span>
        </div>
        <div class="days-grid">
          <div
            v-for="day in daysInMonth"
            :key="day.date"
            class="day-cell"
            :class="{ 
              'selected': selectedDate === day.date, 
              'today': day.isToday,
              'empty': !day.number
            }"
            @click="selectDate(day.date)"
          >
            <span class="day-number">{{ day.number }}</span>
          </div>
        </div>
      </div>
      <div class="room-select-panel" v-if="selectedDate">
        <div class="panel-header">
          <h3>Выберите номер для {{ selectedDateFormatted }}</h3>
        </div>
        <div class="rooms-list">
          <button
            v-for="room in rooms"
            :key="room.id"
            class="room-button"
            :class="{ 'selected': selectedRoom?.id === room.id }"
            @click="selectRoom(room)"
          >
            №{{ room.id }}
          </button>
        </div>
      </div>
      <div class="cleaner-panel" v-if="selectedRoom && selectedDate">
        <div class="panel-header">
          <h3>Уборка номера №{{ selectedRoom.id }} — {{ selectedDateFormatted }}</h3>
        </div>
        <div class="cleaner-info" v-if="roomHasCleaning">
          <div class="info-row">
            <span class="label">Сотрудник:</span>
            <span class="value">{{ getEmployeeById(currentCleaning.employee_id).full_name }}</span>
          </div>
          <div class="info-row">
            <label>Изменить сотрудника:</label>
            <select v-model="currentCleaning.employee_id" @change="updateCleaningEmployee">
              <option v-for="employee in employees" :key="employee.id" :value="employee.id">{{ employee.full_name }}</option>
            </select>
          </div>
          <div class="info-row">
            <button class="edit-button" @click="openEditModal(currentCleaning)">✏️</button>
            <button class="delete-button" @click="deleteCleaning(currentCleaning.id)">🗑️</button>
          </div>
        </div>
        <div class="cleaner-info" v-else>
          <p>Уборка не назначена</p>
          <button class="add-button" @click="openAddModalForRoom(selectedRoom.id)">Добавить уборку</button>
        </div>
      </div>
    </div>

    <!-- Модальное окно добавления/редактирования уборки -->
    <div v-if="showAddEditModal" class="modal-overlay" @click="closeAddEditModal">
      <div class="modal" @click.stop>
        <button class="close-modal" @click="closeAddEditModal">✕</button>
        <div class="modal-header">
          <h2>{{ editingCleaning ? 'Редактировать уборку' : 'Добавить уборку' }}</h2>
        </div>
        <div class="modal-content">
          <form @submit.prevent="saveCleaning">
            <div class="form-group">
              <label>Дата:</label>
              <input v-model="cleaningForm.cleaning_date" type="date" required />
            </div>
            <div class="form-group">
              <label>Номер:</label>
              <select v-model="cleaningForm.room_id" required>
                <option v-for="room in rooms" :key="room.id" :value="room.id">№{{ room.id }}</option>
              </select>
            </div>
            <div class="form-group">
              <label>Сотрудник:</label>
              <select v-model="cleaningForm.employee_id" required>
                <option v-for="employee in employees" :key="employee.id" :value="employee.id">{{ employee.full_name }}</option>
              </select>
            </div>
            <div class="form-actions">
              <button type="submit" class="save-button">
                {{ editingCleaning ? 'Сохранить' : 'Добавить' }}
              </button>
              <button type="button" class="cancel-button" @click="closeAddEditModal">Отмена</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const cleanings = ref([])
const rooms = ref([])
const employees = ref([])
const loading = ref(true)
const error = ref(null)
const selectedDate = ref(null)
const selectedRoom = ref(null)
const showAddEditModal = ref(false)
const editingCleaning = ref(null)
const currentYear = ref(new Date().getFullYear())
const currentMonth = ref(new Date().getMonth())
const cleaningForm = ref({
  cleaning_date: '',
  room_id: '',
  employee_id: ''
})

async function loadData() {
  try {
    const [cleaningsRes, roomsRes, employeesRes] = await Promise.all([
      axios.get('http://localhost:8000/api/cleanings'),
      axios.get('http://localhost:8000/api/rooms'),
      axios.get('http://localhost:8000/api/employees')
    ])
    cleanings.value = cleaningsRes.data
    rooms.value = roomsRes.data
    employees.value = employeesRes.data
    loading.value = false
  } catch (err) {
    error.value = 'Ошибка загрузки: ' + (err.response?.data?.message || err.message)
    loading.value = false
    console.error('Ошибка загрузки:', err)
  }
}

const daysInMonth = computed(() => {
  const days = []
  const date = new Date(currentYear.value, currentMonth.value, 1)
  const daysCount = new Date(currentYear.value, currentMonth.value + 1, 0).getDate()
  const firstDay = (date.getDay() + 6) % 7 // Пн = 0, Вс = 6

  // Пустые дни в начале месяца
  for (let i = 0; i < firstDay; i++) {
    days.push({ number: '', date: '' })
  }
  for (let i = 1; i <= daysCount; i++) {
    const dateStr = `${currentYear.value}-${String(currentMonth.value + 1).padStart(2, '0')}-${String(i).padStart(2, '0')}`
    days.push({
      number: i,
      date: dateStr,
      isToday: isToday(dateStr)
    })
  }
  return days
})

function isToday(dateStr) {
  const today = new Date().toISOString().split('T')[0]
  return dateStr === today
}

const currentMonthName = computed(() => {
  const monthNames = ['Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь', 'Июль', 'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь']
  return monthNames[currentMonth.value]
})

function nextMonth() {
  currentMonth.value++
  if (currentMonth.value > 11) {
    currentMonth.value = 0
    currentYear.value++
  }
  selectedDate.value = null
  selectedRoom.value = null
}

function prevMonth() {
  currentMonth.value--
  if (currentMonth.value < 0) {
    currentMonth.value = 11
    currentYear.value--
  }
  selectedDate.value = null
  selectedRoom.value = null
}

function selectDate(date) {
  if (!date) return // Если день пустой (не принадлежит месяцу)
  selectedDate.value = date
  selectedRoom.value = null
}

const selectedDateFormatted = computed(() => {
  if (!selectedDate.value) return ''
  const [year, month, day] = selectedDate.value.split('-')
  return `${day}.${month}.${year}`
})

function selectRoom(room) {
  selectedRoom.value = room
}

const getCleaningsForDate = computed(() => {
  if (!selectedDate.value) return []
  return cleanings.value.filter(c => {
    const cleaningDate = new Date(c.cleaning_date).toISOString().split('T')[0]
    return cleaningDate === selectedDate.value
  })
})

const roomHasCleaning = computed(() => {
  if (!selectedRoom.value || !selectedDate.value) return false
  return getCleaningsForDate.value.some(c => c.room_id === selectedRoom.value.id)
})

const currentCleaning = computed(() => {
  if (!selectedRoom.value || !selectedDate.value) return null
  return getCleaningsForDate.value.find(c => c.room_id === selectedRoom.value.id)
})

function getEmployeeById(id) {
  return employees.value.find(e => e.id === id) || { full_name: '—' }
}

function openAddModal() {
  cleaningForm.value = {
    cleaning_date: selectedDate.value || new Date().toISOString().split('T')[0],
    room_id: '',
    employee_id: ''
  }
  editingCleaning.value = null
  showAddEditModal.value = true
}

function openAddModalForRoom(roomId) {
  cleaningForm.value = {
    cleaning_date: selectedDate.value || new Date().toISOString().split('T')[0],
    room_id: roomId,
    employee_id: ''
  }
  editingCleaning.value = null
  showAddEditModal.value = true
}

function openEditModal(cleaning) {
  cleaningForm.value = {
    cleaning_date: new Date(cleaning.cleaning_date).toISOString().split('T')[0],
    room_id: cleaning.room_id,
    employee_id: cleaning.employee_id
  }
  editingCleaning.value = cleaning
  showAddEditModal.value = true
}

function closeAddEditModal() {
  showAddEditModal.value = false
}

async function updateCleaningEmployee(event) {
  if (!currentCleaning.value) return
  const employeeId = event.target.value
  try {
    await axios.put(
      `http://localhost:8000/api/cleanings?id=${currentCleaning.value.id}`,
      {
        cleaning_date: new Date(currentCleaning.value.cleaning_date).toISOString().split('T')[0],
        room_id: currentCleaning.value.room_id,
        employee_id: employeeId
      }
    )
    await loadData()
  } catch (err) {
    error.value = 'Ошибка сохранения: ' + (err.response?.data?.message || err.message)
    console.error('Ошибка сохранения:', err)
  }
}

async function saveCleaning() {
  try {
    const data = {
      cleaning_date: cleaningForm.value.cleaning_date,
      room_id: cleaningForm.value.room_id,
      employee_id: cleaningForm.value.employee_id
    }
    if (editingCleaning.value) {
      await axios.put(`http://localhost:8000/api/cleanings?id=${editingCleaning.value.id}`, data)
    } else {
      await axios.post('http://localhost:8000/api/cleanings', data)
    }
    await loadData()
    closeAddEditModal()
  } catch (err) {
    error.value = 'Ошибка сохранения: ' + (err.response?.data?.message || err.message)
    console.error('Ошибка сохранения:', err)
  }
}

async function deleteCleaning(id) {
  if (!confirm('Удалить уборку?')) return
  try {
    await axios.delete(`http://localhost:8000/api/cleanings?id=${id}`)
    await loadData()
  } catch (err) {
    error.value = 'Ошибка удаления: ' + (err.response?.data?.message || err.message)
    console.error('Ошибка удаления:', err)
  }
}

onMounted(loadData)
</script>

<style scoped>
.cleaning-view {
  min-height: 100vh;
  background: #f9f9f9;
  padding: 40px 0;
}
.cleaning-container {
  max-width: 1000px;
  margin: 0 auto;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  padding: 32px;
}
h1 {
  color: #a83b3b;
  font-size: 2rem;
  margin-bottom: 24px;
  text-align: center;
  font-weight: 700;
}
.controls {
  margin-bottom: 24px;
  display: flex;
  justify-content: flex-end;
}
.add-button {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  background: #a83b3b;
  color: #fff;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}
.add-button:hover {
  background: #8c2b2b;
  transform: translateY(-2px);
}
.calendar {
  margin-bottom: 32px;
}
.calendar-header {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 20px;
  margin-bottom: 16px;
}
.month-button {
  background: none;
  border: none;
  font-size: 1.2rem;
  cursor: pointer;
  color: #a83b3b;
}
.weekdays {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  background: #f8f8f8;
  font-weight: 600;
  text-align: center;
  padding: 8px 0;
  border-radius: 8px 8px 0 0;
}
.weekdays > span {
  padding: 4px;
}
.days-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  background: #fff;
  border-radius: 0 0 8px 8px;
  border: 2px solid #e1e4e8;
  overflow: hidden;
}
.day-cell {
  min-height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: all 0.2s ease;
  border: 1px solid #e1e4e8;
}
.day-cell:hover {
  background: #e1e4e8;
}
.day-cell.selected {
  background: #a83b3b;
  color: #fff;
}
.day-cell.today {
  border-color: #4caf50;
}
.day-cell.empty {
  background: transparent;
  border: none;
}
.day-number {
  font-weight: 600;
}
.room-select-panel, .cleaner-panel {
  margin-top: 24px;
  background: #fff;
  border-radius: 12px;
  padding: 20px;
  border: 2px solid #e1e4e8;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}
.panel-header {
  margin-bottom: 16px;
}
.rooms-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 12px;
}
.room-button {
  background: #fff;
  border: 2px solid #e1e4e8;
  border-radius: 8px;
  padding: 12px;
  text-align: center;
  cursor: pointer;
  transition: all 0.2s ease;
}
.room-button:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 8px rgba(168, 59, 59, 0.1);
}
.room-button.selected {
  background: rgba(76, 175, 80, 0.15);
  border-color: #4caf50;
}
.cleaner-info {
  padding: 16px;
  background: #f8f8f8;
  border-radius: 8px;
  border: 2px solid #e1e4e8;
}
.info-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 12px;
  align-items: center;
}
.info-row label {
  font-weight: 600;
  color: #6d6d6d;
}
.info-row select {
  padding: 8px 12px;
  border: 2px solid #e1e4e8;
  border-radius: 8px;
  font-size: 1rem;
  outline: none;
  transition: all 0.2s ease;
}
.info-row select:focus {
  border-color: #a83b3b;
  box-shadow: 0 0 8px rgba(168, 59, 59, 0.2);
}
.info-row .value {
  color: #333;
  font-weight: 500;
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
}
.modal {
  background: #fff;
  border-radius: 16px;
  padding: 24px;
  max-width: 500px;
  width: 90%;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  position: relative;
}
.modal::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 8px;
  background: #a83b3b;
  border-radius: 16px 16px 0 0;
}
.modal-header {
  margin-bottom: 16px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 8px;
}
.close-modal {
  position: absolute;
  top: 16px;
  right: 16px;
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #a83b3b;
  transition: all 0.2s ease;
}
.close-modal:hover {
  color: #e74c3c;
  transform: rotate(90deg);
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
.form-group input, .form-group select {
  width: 100%;
  padding: 10px 12px;
  border: 2px solid #e1e4e8;
  border-radius: 8px;
  font-size: 1rem;
  outline: none;
  transition: all 0.2s ease;
}
.form-group input:focus, .form-group select:focus {
  border-color: #a83b3b;
  box-shadow: 0 0 8px rgba(168, 59, 59, 0.2);
}
.form-actions {
  margin-top: 24px;
  display: flex;
  gap: 12px;
}
.save-button, .cancel-button {
  padding: 10px 20px;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
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
</style>
