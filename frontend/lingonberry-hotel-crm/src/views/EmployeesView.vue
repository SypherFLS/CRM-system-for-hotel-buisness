<template>
  <div class="employees-view">
    <div class="employees-container">
      <h1>Сотрудники</h1>
      <div class="filters">
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Поиск по имени сотрудника"
          class="search-input"
        />
      </div>
      <button class="add-button" @click="openAddModal">+ Добавить сотрудника</button>
      <div v-if="loading" class="loading">Загрузка...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else class="employee-list">
        <div
          v-for="employee in filteredEmployees"
          :key="employee.id"
          class="employee-card"
        >
          <div class="card-content">
            <h3>{{ employee.full_name }}</h3>
            <div class="info-row">
              <span class="label">Этажи:</span>
              <span class="value">{{ employee.floors || '—' }}</span>
            </div>
            <div class="info-row">
              <span class="label">Дни недели:</span>
              <span class="value">{{ employee.week_days || '—' }}</span>
            </div>
            <div class="employee-actions">
              <button class="edit-button" @click.stop="openEditModal(employee)">✏️</button>
              <button class="delete-button" @click.stop="deleteEmployee(employee.id)">🗑️</button>
            </div>
          </div>
        </div>
      </div>
      <!-- Модальное окно добавления/редактирования -->
      <div v-if="showAddEditModal" class="modal-overlay" @click="closeAddEditModal">
        <div class="modal" @click.stop>
          <button class="close-modal" @click="closeAddEditModal">✕</button>
          <div class="modal-header">
            <h2>{{ editingEmployee ? 'Редактировать сотрудника' : 'Добавить сотрудника' }}</h2>
          </div>
          <div class="modal-content">
            <form @submit.prevent="saveEmployee">
              <div class="form-group">
                <label>ФИО:</label>
                <input v-model="employeeForm.full_name" type="text" required />
              </div>
              <div class="form-group">
                <label>Этажи:</label>
                <input v-model="employeeForm.floors" type="text" />
              </div>
              <div class="form-group">
                <label>Дни недели:</label>
                <input v-model="employeeForm.week_days" type="text" />
              </div>
              <div class="form-actions">
                <button type="submit" class="save-button">
                  {{ editingEmployee ? 'Сохранить' : 'Добавить' }}
                </button>
                <button type="button" class="cancel-button" @click="closeAddEditModal">
                  Отмена
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const employees = ref([])
const loading = ref(true)
const error = ref(null)
const searchQuery = ref('')
const showAddEditModal = ref(false)
const editingEmployee = ref(null)
const employeeForm = ref({
  full_name: '',
  floors: '',
  week_days: ''
})

async function loadData() {
  try {
    const response = await axios.get('http://localhost:8000/api/employees')
    employees.value = response.data
    loading.value = false
  } catch (err) {
    error.value = 'Ошибка загрузки: ' + (err.response?.data?.message || err.message)
    loading.value = false
    console.error('Ошибка загрузки:', err)
  }
}

const filteredEmployees = computed(() => {
  if (!searchQuery.value) return employees.value
  const query = searchQuery.value.toLowerCase()
  return employees.value.filter(employee =>
    employee.full_name.toLowerCase().includes(query)
  )
})

function openAddModal() {
  employeeForm.value = { full_name: '', floors: '', week_days: '' }
  editingEmployee.value = null
  showAddEditModal.value = true
}

function openEditModal(employee) {
  employeeForm.value = { ...employee }
  editingEmployee.value = employee
  showAddEditModal.value = true
}

function closeAddEditModal() {
  showAddEditModal.value = false
}

async function saveEmployee() {
  try {
    if (editingEmployee.value) {
      await axios.put(`http://localhost:8000/api/employees/${employeeForm.value.id}`, employeeForm.value)
    } else {
      await axios.post('http://localhost:8000/api/employees', employeeForm.value)
    }
    await loadData()
    closeAddEditModal()
  } catch (err) {
    error.value = 'Ошибка сохранения: ' + (err.response?.data?.message || err.message)
    console.error('Ошибка сохранения:', err)
  }
}

async function deleteEmployee(id) {
  if (!confirm('Удалить сотрудника?')) return
  try {
    await axios.delete(`http://localhost:8000/api/employees?id=${id}`)
    await loadData()
  } catch (err) {
    error.value = 'Ошибка удаления: ' + (err.response?.data?.message || err.message)
    console.error('Ошибка удаления:', err)
  }
}

onMounted(loadData)
</script>

<style scoped>
.employees-view {
  min-height: 100vh;
  background: #f9f9f9;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px 0;
}
.employees-container {
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
.employee-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 24px;
}
.employee-card {
  background: #f8f8f8;
  border-radius: 16px;
  padding: 24px;
  border: 2px solid #e1e4e8;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
}
.employee-card:hover {
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
.employee-actions {
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
.modal-content {
  margin-top: 16px;
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
.save-button, .cancel-button {
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
