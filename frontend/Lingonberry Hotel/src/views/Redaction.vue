<template>
  <div class="page-container">
    <div class="booking-card">
      <h2 class="title">Календарь бронирования</h2>

      <div class="calendar">
        <div class="week-days">
          <div v-for="day in weekDays" :key="day" class="day-name">{{ day }}</div>
        </div>

        <div class="days-grid">
          <div
            v-for="day in daysInMonth"
            :key="day"
            class="day"
            :class="{ 
              disabled: day === null,
              selected: isSelected(day),
              booked: isBooked(day)
            }"
            @click="toggleDay(day)"
          >
            {{ day || '' }}
          </div>
        </div>
      </div>

      <div class="buttons">
        <button @click="addRange" class="btn btn-add">Добавить</button>
        <button @click="deleteRange" class="btn btn-delete">Удалить</button>
        <button @click="confirmBooking" class="btn btn-confirm">Подтвердить</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const today = new Date()
const year = today.getFullYear()
const month = today.getMonth()

const weekDays = ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс']

const bookedRanges = ref([])
const selectedRange = ref(null)

function getDaysInMonth(year, month) {
  return new Date(year, month + 1, 0).getDate()
}

const daysInMonth = computed(() => {
  const daysCount = getDaysInMonth(year, month)
  const firstDay = new Date(year, month, 1).getDay()
  const offset = firstDay === 0 ? 6 : firstDay - 1

  const days = Array(offset).fill(null)
  for (let i = 1; i <= daysCount; i++) {
    days.push(i)
  }
  return days
})

function isBooked(day) {
  if (!day) return false
  return bookedRanges.value.some(range => day >= range.start && day <= range.end)
}

function isSelected(day) {
  if (!day || !selectedRange.value) return false
  return day >= selectedRange.value.start && day <= selectedRange.value.end
}

function toggleDay(day) {
  if (!day) return

  if (!selectedRange.value) {
    selectedRange.value = { start: day, end: day }
  } else {
    if (day < selectedRange.value.start) {
      selectedRange.value = { start: day, end: selectedRange.value.end }
    } else if (day > selectedRange.value.start) {
      selectedRange.value.end = day
    } else {
      selectedRange.value = null
    }
  }
}

function parseRange(str) {
  const [start, end] = str.split(':').map(Number)
  return { start, end }
}

async function loadBookedRanges() {
  try {
    const res = await fetch('http://localhost:8642/bookings')
    if (!res.ok) throw new Error('Ошибка загрузки')
    const data = await res.json()
    bookedRanges.value = data.map(parseRange)
  } catch (e) {
    console.error(e)
  }
}

function addRange() {
  if (!selectedRange.value) {
    alert('Выберите диапазон дат на календаре')
    return
  }
  bookedRanges.value.push({ ...selectedRange.value })
  selectedRange.value = null
}

function deleteRange() {
  if (!selectedRange.value) {
    alert('Выберите диапазон дат для удаления')
    return
  }
  bookedRanges.value = bookedRanges.value.filter(
    r => !(r.start === selectedRange.value.start && r.end === selectedRange.value.end)
  )
  selectedRange.value = null
}

async function confirmBooking() {
  try {
    const payload = bookedRanges.value.map(r => `${r.start}:${r.end}`)
    const res = await fetch('http://localhost:8642/bookings', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    })
    if (!res.ok) throw new Error('Ошибка отправки')
    alert('Бронирование подтверждено')
  } catch (e) {
    alert('Ошибка при подтверждении бронирования')
    console.error(e)
  }
}

onMounted(() => {
  loadBookedRanges()
})
</script>

<style scoped>
.page-container {
  height: 100vh;
  background: linear-gradient(135deg, #f9f7f1, #d6e4e5);
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  box-sizing: border-box;
}

.booking-card {
  background: white;
  border-radius: 16px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.15);
  width: 360px;
  max-width: 100%;
  padding: 30px 25px;
  text-align: center;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.title {
  font-weight: 700;
  font-size: 24px;
  margin-bottom: 20px;
  color: #333;
}

.calendar {
  user-select: none;
}

.week-days {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  font-weight: 600;
  color: #666;
  margin-bottom: 8px;
}

.day-name {
  padding: 6px 0;
}

.days-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 6px;
}

.day {
  width: 40px;
  height: 40px;
  line-height: 40px;
  border-radius: 8px;
  cursor: pointer;
  background: transparent;
  transition: background-color 0.3s ease, color 0.3s ease;
  user-select: none;
  font-weight: 600;
  color: #444;
}

.day.disabled {
  cursor: default;
  background: transparent;
  color: transparent;
}

.day.booked {
  background-color: rgba(255, 0, 0, 0.3);
  color: #900;
  pointer-events: none;
}

.day.selected {
  background-color: rgba(0, 123, 255, 0.5);
  color: #003366;
}

.buttons {
  margin-top: 25px;
  display: flex;
  justify-content: space-between;
  gap: 10px;
}

.btn {
  flex: 1;
  padding: 10px 0;
  border-radius: 8px;
  font-weight: 700;
  cursor: pointer;
  border: none;
  transition: background-color 0.3s ease;
  color: white;
  user-select: none;
}

.btn-add {
  background-color: #28a745;
}

.btn-add:hover {
  background-color: #218838;
}

.btn-delete {
  background-color: #dc3545;
}

.btn-delete:hover {
  background-color: #c82333;
}

.btn-confirm {
  background-color: #007bff;
}

.btn-confirm:hover {
  background-color: #0056b3;
}
</style>
