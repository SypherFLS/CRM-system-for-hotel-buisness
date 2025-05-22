<template>
  <div>
    <div class="page-container">
      <div class="booking-card">
        <h2 class="title">Календарь бронирования</h2>

        
        <label class="label" for="hotel-select">Выберите отель:</label>
        <select id="hotel-select" v-model="selectedHotelId" class="select">
          <option disabled value="">-- Выберите отель --</option>
          <option v-for="hotel in hotels" :key="hotel.id" :value="hotel.id">
            {{ hotel.name }}
          </option>
        </select>

        
        <label class="label" for="month-select">Выберите месяц:</label>
        <select id="month-select" v-model="selectedMonth" class="select">
          <option v-for="(monthName, index) in monthNames" :key="index" :value="index">
            {{ monthName }}
          </option>
        </select>

        
        <label class="label" for="class-select">Выберите класс номера:</label>
        <select id="class-select" v-model="selectedClass" class="select" :disabled="!selectedHotelId">
          <option disabled value="">-- Выберите класс --</option>
          <option v-for="(className, index) in roomClasses" :key="index" :value="className">
            {{ className }}
          </option>
        </select>

        
        <div class="calendar" v-if="selectedHotelId !== ''">
          <div class="week-days">
            <div v-for="(day, index) in weekDays" :key="day + '-' + index" class="day-name">{{ day }}</div>
          </div>

          <div class="days-grid">
            <div
              v-for="(day, index) in daysInMonth"
              :key="day !== null ? day : 'empty-' + index"
              class="day"
              :class="{ 
                disabled: day === null,
                selected: isSelected(day)
              }"
              @click="toggleDay(day)"
            >
              {{ day || '' }}
            </div>
          </div>
        </div>

        
        <div class="price-info" v-if="selectedHotelId && selectedClass && selectedRange">
          <p>
            Цена за ночь: <strong>{{ pricePerNight }} ₽</strong><br />
            Всего за выбранный период ({{ selectedRange.end - selectedRange.start + 1 }} ночей): <strong>{{ totalPrice }} ₽</strong>
          </p>
        </div>

        <div class="buttons">
          <button class="btn btn-add" @click="onAddClick" :disabled="!canBook">Добавить</button>
          <button class="btn btn-delete" @click="onDeleteClick" :disabled="!selectedHotelId">Удалить</button>
          <button class="btn btn-confirm" @click="onConfirmClick" :disabled="!canBook">Подтвердить</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const hotels = [
  { id: 1, name: 'Отель Солнечный', basePrice: 4500 },
  { id: 2, name: 'Ночной Рай', basePrice: 5200 },
  { id: 3, name: 'Комфорт Плюс', basePrice: 4800 },
  { id: 4, name: 'Летний Бриз', basePrice: 4300 },
  { id: 5, name: 'Вечерний Клуб', basePrice: 5100 },
  { id: 6, name: 'Туристический Оазис', basePrice: 4700 },
  { id: 7, name: 'Городской Уют', basePrice: 4400 },
  { id: 8, name: 'Ночная Жизнь', basePrice: 5300 },
  { id: 9, name: 'Путешественник', basePrice: 4600 },
  { id: 10, name: 'Семейный Очаг', basePrice: 4200 },
  { id: 11, name: 'Клубный Дом', basePrice: 5000 },
  { id: 12, name: 'Туристический Рай', basePrice: 4550 },
]

const roomClasses = [
  'Стандарт',
  'Комфорт',
  'Люкс',
  'Бусничный магнат',
]

const weekDays = ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс']
const monthNames = [
  'Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь',
  'Июль', 'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь'
]

const today = new Date()
const selectedHotelId = ref('')
const selectedMonth = ref(today.getMonth())
const selectedYear = today.getFullYear()
const selectedClass = ref('')
const selectedRange = ref(null)

function getDaysInMonth(year, month) {
  return new Date(year, month + 1, 0).getDate()
}

const daysInMonth = computed(() => {
  const daysCount = getDaysInMonth(selectedYear, selectedMonth.value)
  const firstDay = new Date(selectedYear, selectedMonth.value, 1).getDay()
  const offset = firstDay === 0 ? 6 : firstDay - 1

  const days = Array(offset).fill(null)
  for (let i = 1; i <= daysCount; i++) {
    days.push(i)
  }
  return days
})

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


const classPriceMultipliers = {
  'Стандарт': 1,
  'Комфорт': 1.3,
  'Люкс': 1.7,
  'Бусничный магнат': 3,
}


const pricePerNight = computed(() => {
  if (!selectedHotelId.value || !selectedClass.value) return 0
  const hotel = hotels.find(h => h.id === selectedHotelId.value)
  if (!hotel) return 0
  const multiplier = classPriceMultipliers[selectedClass.value] || 1
  return Math.round(hotel.basePrice * multiplier)
})

// Общая цена за выбранный период
const totalPrice = computed(() => {
  if (!selectedRange.value) return 0
  const nights = selectedRange.value.end - selectedRange.value.start + 1
  return pricePerNight.value * nights
})

const canBook = computed(() => {
  return selectedHotelId.value && selectedClass.value && selectedRange.value
})

function onAddClick() {
  alert(`Добавлено бронирование:\nОтель: ${hotels.find(h => h.id === selectedHotelId.value)?.name}\nКласс: ${selectedClass.value}\nМесяц: ${monthNames[selectedMonth.value]}\nДаты: ${selectedRange.value.start} - ${selectedRange.value.end}\nЦена за ночь: ${pricePerNight.value} ₽\nИтого: ${totalPrice.value} ₽`)
}
function onDeleteClick() {
  alert(`Удалить бронирование: Отель ID ${selectedHotelId.value}, Месяц ${monthNames[selectedMonth.value]}`)
}
function onConfirmClick() {
  alert(`Подтверждено бронирование:\nОтель: ${hotels.find(h => h.id === selectedHotelId.value)?.name}\nКласс: ${selectedClass.value}\nМесяц: ${monthNames[selectedMonth.value]}\nДаты: ${selectedRange.value.start} - ${selectedRange.value.end}\nЦена за ночь: ${pricePerNight.value} ₽\nИтого: ${totalPrice.value} ₽`)
}

// Сброс выбранного диапазона и класса при смене месяца или отеля
watch([selectedMonth, selectedHotelId], () => {
  selectedRange.value = null
  selectedClass.value = ''
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
  width: 800px;
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

.label {
  display: block;
  margin-top: 10px;
  margin-bottom: 6px;
  font-weight: 600;
  color: #444;
  text-align: left;
}

.select {
  width: 100%;
  padding: 8px 12px;
  margin-bottom: 15px;
  border-radius: 8px;
  border: 1.5px solid #bdc3c7;
  font-size: 1rem;
  font-family: inherit;
  cursor: pointer;
  transition: border-color 0.3s ease;
}

.select:focus {
  outline: none;
  border-color: #3498db;
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

.day.selected {
  background-color: rgba(0, 123, 255, 0.5);
  color: #003366;
}

.price-info {
  margin-top: 15px;
  font-weight: 700;
  font-size: 1.1rem;
  color: #2c3e50;
  text-align: left;
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
