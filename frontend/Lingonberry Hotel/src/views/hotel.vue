<template>
  <div class="hotels-page">
    <h1 class="title">Наши отели</h1>
    <div class="hotels-grid">
      <div v-for="hotel in hotels" :key="hotel.id" class="hotel-card">
        <img :src="hotel.image" :alt="hotel.name" class="hotel-image" />
        <div class="hotel-info">
          <h2 class="hotel-name">{{ hotel.name }}</h2>
          <p class="hotel-location">{{ hotel.location }}</p>
          <p class="hotel-spec">{{ hotel.specification }}</p>
          <p class="hotel-price">Цена за ночь: {{ hotel.price }} ₽</p>
        </div>
        <button class="book-button">Забронировать</button>
      </div>
    </div>
  </div>
</template>

<script>
import hotel from '@/assets/hotel_futazh.png';
export default {
  name: 'Hotel',
  data() {
    return {
      hotels: [
        {
          id: 1,
          name: 'Отель Солнечный',
          location: 'Сочи',
          specification: 'Для семейного отдыха',
          price: 4500,
          image: hotel
        },
        {
          id: 2,
          name: 'Ночной Рай',
          location: 'Москва',
          specification: 'Для тусовок',
          price: 5200,
          image: hotel,
        },
        {
          id: 3,
          name: 'Комфорт Плюс',
          location: 'Санкт-Петербург',
          specification: 'Для комфортного туризма',
          price: 4800,
          image: hotel
        },
        {
          id: 4,
          name: 'Летний Бриз',
          location: 'Крым',
          specification: 'Для семейного отдыха',
          price: 4300,
          image: hotel
        },
        {
          id: 5,
          name: 'Вечерний Клуб',
          location: 'Казань',
          specification: 'Для тусовок',
          price: 5100,
          image: hotel
        },
        {
          id: 6,
          name: 'Туристический Оазис',
          location: 'Екатеринбург',
          specification: 'Для комфортного туризма',
          price: 4700,
          image: hotel
        },
        {
          id: 7,
          name: 'Городской Уют',
          location: 'Новосибирск',
          specification: 'Для семейного отдыха',
          price: 4400,
          image: hotel
        },
        {
          id: 8,
          name: 'Ночная Жизнь',
          location: 'Ростов-на-Дону',
          specification: 'Для тусовок',
          price: 5300,
          image: hotel
        },
        {
          id: 9,
          name: 'Путешественник',
          location: 'Владивосток',
          specification: 'Для комфортного туризма',
          price: 4600,
          image: hotel
        },
        {
          id: 10,
          name: 'Семейный Очаг',
          location: 'Краснодар',
          specification: 'Для семейного отдыха',
          price: 4200,
          image: hotel
        },
        {
          id: 11,
          name: 'Клубный Дом',
          location: 'Нижний Новгород',
          specification: 'Для тусовок',
          price: 5000,
          image: hotel
        },
        {
          id: 12,
          name: 'Туристический Рай',
          location: 'Самара',
          specification: 'Для комфортного туризма',
          price: 4550,
          image: hotel
        },
      ],
    };
  },
};
</script>

<style scoped>
.hotels-page {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  text-align: center;
  background: #f9f9f9;
  color: #333;
}

.title {
  margin-bottom: 30px;
  font-size: 2.5rem;
  font-weight: 700;
  color: #2c3e50;
}

.hotels-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  gap: 25px;
  justify-items: center;
}

.hotel-card {
  background: white;
  border-radius: 15px;
  box-shadow: 0 6px 14px rgba(0, 0, 0, 0.1);
  width: 240px;
  padding: 15px;
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: transform 0.2s ease;
}

.hotel-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.15);
}

.hotel-image {
  width: 100%;
  height: 140px;
  object-fit: cover;
  border-radius: 15px 15px 10px 10px;
  margin-bottom: 12px;
}

.hotel-info {
  text-align: center;
  margin-bottom: 15px;
}

.hotel-name {
  font-size: 1.2rem;
  margin: 6px 0 4px;
  font-weight: 600;
  color: #34495e;
}

.hotel-location,
.hotel-spec,
.hotel-price {
  font-size: 0.9rem;
  margin: 2px 0;
  color: #7f8c8d;
}

.book-button {
  background-color: #3498db;
  border: none;
  border-radius: 10px;
  padding: 10px 18px;
  font-size: 1rem;
  color: white;
  cursor: pointer;
  width: 100%;
  transition: background-color 0.3s ease;
}

.book-button:hover {
  background-color: #2980b9;
}
</style>
