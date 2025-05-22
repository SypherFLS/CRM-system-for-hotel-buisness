<template>
  <div class="home">
    <h1>Загруженность отелей</h1>
    <div class="content">
      <transition name="fade">
        <div v-if="selectedHotel" class="info-panel">
          <h2>Информация по отелю: {{ selectedHotel.name }}</h2>
          <p><em>{{ selectedHotel.location }} — {{ selectedHotel.specification }}</em></p>
          <p><strong>Цена за ночь:</strong> {{ selectedHotel.price }} ₽</p>

          <div class="room-status">
            <p><strong>Занятые номера:</strong> {{ roomData.occupied.join(', ') || 'нет' }}</p>
            <p><strong>Свободные номера:</strong> {{ roomData.free.join(', ') || 'нет' }}</p>
            <p><strong>Забронированные номера:</strong> {{ roomData.booked.join(', ') || 'нет' }}</p>
            <p><strong>Скоро освободятся:</strong> {{ roomData.soontojoin.join(', ') || 'нет' }}</p>
          </div>
        </div>
      </transition>

      <div class="hotels-list">
        <div
          v-for="hotel in hotels"
          :key="hotel.id"
          :class="['hotel-card', { selected: selectedHotel && selectedHotel.id === hotel.id }]"
          @click="selectHotel(hotel)"
        >
          <img :src="hotel.image" alt="Фото отеля" class="hotel-image" />
          <div class="hotel-info">
            <h3>{{ hotel.name }}</h3>
            <p class="location">{{ hotel.location }}</p>
            <p class="specification">{{ hotel.specification }}</p>
            <p class="price">{{ hotel.price }} ₽ / ночь</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import hotelImage from '@/assets/hotel_pic.png' // замените на свой путь или импорт

const hotels = [
  {
    id: 1,
    name: 'Отель Солнечный',
    location: 'Сочи',
    specification: 'Для семейного отдыха',
    price: 4500,
    image: hotelImage,
  },
  {
    id: 2,
    name: 'Ночной Рай',
    location: 'Москва',
    specification: 'Для тусовок',
    price: 5200,
    image: hotelImage,
  },
  {
    id: 3,
    name: 'Комфорт Плюс',
    location: 'Санкт-Петербург',
    specification: 'Для комфортного туризма',
    price: 4800,
    image: hotelImage,
  },
  {
    id: 4,
    name: 'Летний Бриз',
    location: 'Крым',
    specification: 'Для семейного отдыха',
    price: 4300,
    image: hotelImage,
  },
  {
    id: 5,
    name: 'Вечерний Клуб',
    location: 'Казань',
    specification: 'Для тусовок',
    price: 5100,
    image: hotelImage,
  },
  {
    id: 6,
    name: 'Туристический Оазис',
    location: 'Екатеринбург',
    specification: 'Для комфортного туризма',
    price: 4700,
    image: hotelImage,
  },
  {
    id: 7,
    name: 'Городской Уют',
    location: 'Новосибирск',
    specification: 'Для семейного отдыха',
    price: 4400,
    image: hotelImage,
  },
  {
    id: 8,
    name: 'Ночная Жизнь',
    location: 'Ростов-на-Дону',
    specification: 'Для тусовок',
    price: 5300,
    image: hotelImage,
  },
  {
    id: 9,
    name: 'Путешественник',
    location: 'Владивосток',
    specification: 'Для комфортного туризма',
    price: 4600,
    image: hotelImage,
  },
  {
    id: 10,
    name: 'Семейный Очаг',
    location: 'Краснодар',
    specification: 'Для семейного отдыха',
    price: 4200,
    image: hotelImage,
  },
  {
    id: 11,
    name: 'Клубный Дом',
    location: 'Нижний Новгород',
    specification: 'Для тусовок',
    price: 5000,
    image: hotelImage,
  },
  {
    id: 12,
    name: 'Туристический Рай',
    location: 'Самара',
    specification: 'Для комфортного туризма',
    price: 4550,
    image: hotelImage,
  },
]

const selectedHotel = ref(null)

const roomsExample = {
  1: {
    occupied: [101, 102, 110],
    free: [103, 104, 105],
    booked: [106, 107],
    soontojoin: [108, 109],
  },
  2: {
    occupied: [201, 202],
    free: [203, 204, 205, 206],
    booked: [207],
    soontojoin: [208, 209],
  },
  3: {
    occupied: [301, 302, 303],
    free: [304, 305],
    booked: [306, 307],
    soontojoin: [],
  },
  // Добавьте данные для остальных отелей по необходимости
}

function selectHotel(hotel) {
  selectedHotel.value = hotel
}

const roomData = computed(() => {
  if (!selectedHotel.value) return { occupied: [], free: [], booked: [], soontojoin: [] }
  return roomsExample[selectedHotel.value.id] || { occupied: [], free: [], booked: [], soontojoin: [] }
})
</script>

<style scoped>
.home {
  max-width: 1000px;
  margin: 40px auto;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background: #f9fafb;
  padding: 20px 30px;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  text-align: center;
}

h1 {
  color: #2c3e50;
  margin-bottom: 30px;
}

/* Основной контейнер с двумя колонками */
.content {
  display: flex;
  gap: 30px;
  justify-content: center;
  align-items: flex-start;
}

/* Левая колонка - информация по отелю */
.info-panel {
  background: white;
  border-radius: 12px;
  padding: 30px 25px;
  box-shadow: 0 6px 18px rgba(44, 62, 80, 0.1);
  width: 45%;
  text-align: left;
  color: #34495e;
}

.info-panel h2 {
  color: #2980b9;
  margin-bottom: 12px;
  font-weight: 700;
}

.info-panel p em {
  color: #7f8c8d;
  font-style: normal;
  font-weight: 600;
  margin-bottom: 16px;
  display: block;
}

.info-panel p strong {
  color: #2c3e50;
}

.room-status p {
  font-size: 16px;
  margin: 10px 0;
  color: #2c3e50;
}

/* Правая колонка - список отелей */
.hotels-list {
  width: 45%;
  max-height: 600px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

/* Карточка отеля */
.hotel-card {
  background: white;
  border: 2px solid transparent;
  border-radius: 10px;
  cursor: pointer;
  box-shadow: 0 3px 8px rgba(44, 62, 80, 0.1);
  display: flex;
  padding: 12px;
  transition: all 0.3s ease;
  align-items: center;
  gap: 15px;
}

.hotel-card:hover {
  box-shadow: 0 6px 16px rgba(41, 128, 185, 0.3);
  border-color: #2980b9;
  transform: translateX(5px);
}

.hotel-card.selected {
  border-color: #2980b9;
  box-shadow: 0 8px 20px rgba(41, 128, 185, 0.5);
  transform: translateX(5px);
}

.hotel-image {
  width: 90px;
  height: 70px;
  object-fit: cover;
  border-radius: 8px;
  flex-shrink: 0;
}

.hotel-info {
  flex-grow: 1;
  text-align: left;
}

.hotel-info h3 {
  margin: 0 0 4px 0;
  color: #2980b9;
  font-size: 18px;
  font-weight: 700;
}

.hotel-info .location,
.hotel-info .specification {
  font-size: 13px;
  color: #7f8c8d;
  margin: 1px 0;
}

.hotel-info .price {
  font-weight: 600;
  color: #34495e;
  margin-top: 6px;
  font-size: 14px;
}

/* Плавное появление плашки */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.4s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}


.hotels-list::-webkit-scrollbar {
  width: 8px;
}

.hotels-list::-webkit-scrollbar-thumb {
  background-color: #bdc3c7;
  border-radius: 4px;
}

.hotels-list::-webkit-scrollbar-track {
  background-color: #ecf0f1;
}
</style>
