<template>
  <div class="hotels-page">
    <h1 class="title">Наши отели</h1>

    <div class="actions">
      <button @click="openAddForm" class="action-button">Добавить отель</button>
      <button @click="openEditForm" :disabled="!selectedHotel" class="action-button">Редактировать отель</button>
      <button @click="deleteHotel" :disabled="!selectedHotel" class="action-button delete">Удалить отель</button>
    </div>

    <div class="hotels-grid">
      <div
        v-for="hotel in hotels"
        :key="hotel.id"
        class="hotel-card"
        :class="{ selected: selectedHotel && selectedHotel.id === hotel.id }"
        @click="selectHotel(hotel)"
      >
        <img :src="hotel.image" :alt="hotel.name" class="hotel-image" />
        <div class="hotel-info">
          <h2 class="hotel-name">{{ hotel.name }}</h2>
          <p class="hotel-location">{{ hotel.location }}</p>
          <p class="hotel-spec">{{ hotel.specification }}</p>
          <p class="hotel-price">Цена за ночь: {{ hotel.price }} ₽</p>
        </div>
      </div>
    </div>

    
    <HotelForm
      v-if="showAddForm"
      @close="closeForms"
      @submit="addHotel"
      :initialData="null"
      title="Добавить новый отель"
    />

    <HotelForm
      v-if="showEditForm"
      @close="closeForms"
      @submit="updateHotel"
      :initialData="selectedHotel"
      title="Редактировать отель"
    />
  </div>
</template>

<script>
import hotelImage from '@/assets/hotel_futazh.png';
import HotelForm from '@/components/HotelForm.vue';

export default {
  name: 'HotelsPage',
  components: { HotelForm },
  data() {
    return {
      hotels: [
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
      ],
      selectedHotel: null,
      showAddForm: false,
      showEditForm: false,
      nextId: 13,
    };
  },
  methods: {
    selectHotel(hotel) {
      this.selectedHotel = this.selectedHotel && this.selectedHotel.id === hotel.id ? null : hotel;
    },
    openAddForm() {
      this.showAddForm = true;
      this.showEditForm = false;
      this.selectedHotel = null;
    },
    openEditForm() {
      if (!this.selectedHotel) return;
      this.showEditForm = true;
      this.showAddForm = false;
    },
    closeForms() {
      this.showAddForm = false;
      this.showEditForm = false;
    },
    addHotel(newHotel) {
      newHotel.id = this.nextId++;
      if (!newHotel.image) newHotel.image = hotelImage;
      this.hotels.push(newHotel);
      this.closeForms();
    },
    updateHotel(updatedHotel) {
      const index = this.hotels.findIndex(h => h.id === updatedHotel.id);
      if (index !== -1) {
        this.hotels.splice(index, 1, updatedHotel);
      }
      this.closeForms();
    },
    deleteHotel() {
      if (!this.selectedHotel) return;
      this.hotels = this.hotels.filter(h => h.id !== this.selectedHotel.id);
      this.selectedHotel = null;
    },
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

.actions {
  margin-bottom: 20px;
  display: flex;
  justify-content: center;
  gap: 15px;
}

.action-button {
  background-color: #3498db;
  border: none;
  border-radius: 10px;
  padding: 10px 18px;
  font-size: 1rem;
  color: white;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.action-button:hover:not(:disabled) {
  background-color: #2980b9;
}

.action-button:disabled {
  background-color: #95a5a6;
  cursor: not-allowed;
}

.action-button.delete {
  background-color: #e74c3c;
}

.action-button.delete:hover:not(:disabled) {
  background-color: #c0392b;
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
  cursor: pointer;
  transition: transform 0.2s ease;
  border: 2px solid transparent;
}

.hotel-card.selected {
  border-color: #3498db;
  box-shadow: 0 8px 18px rgba(52, 152, 219, 0.5);
  transform: translateY(-6px);
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
</style>
