<template>
  <div class="modal-overlay" @click.self="close">
    <div class="modal">
      <h2>{{ title }}</h2>
      <form @submit.prevent="submitForm">
        <label>
          Название:
          <input v-model="form.name" type="text" required />
        </label>

        <label>
          Местоположение:
          <input v-model="form.location" type="text" required />
        </label>

        <label>
          Характеристика:
          <input v-model="form.specification" type="text" required />
        </label>

        <label>
          Цена за ночь (₽):
          <input v-model.number="form.price" type="number" min="0" required />
        </label>

        <label>
          URL изображения:
          <input v-model="form.image" type="url" placeholder="Оставьте пустым для стандартного" />
        </label>

        <div class="buttons">
          <button type="submit" class="submit-button">Сохранить</button>
          <button type="button" @click="close" class="cancel-button">Отмена</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HotelForm',
  props: {
    initialData: {
      type: Object,
      default: null,
    },
    title: {
      type: String,
      default: 'Форма',
    },
  },
  data() {
    return {
      form: {
        id: this.initialData?.id || null,
        name: this.initialData?.name || '',
        location: this.initialData?.location || '',
        specification: this.initialData?.specification || '',
        price: this.initialData?.price || 0,
        image: this.initialData?.image || '',
      },
    };
  },
  methods: {
    submitForm() {
      if (!this.form.name || !this.form.location || !this.form.specification || this.form.price < 0) {
        alert('Пожалуйста, заполните все поля корректно.');
        return;
      }
      this.$emit('submit', { ...this.form });
    },
    close() {
      this.$emit('close');
    },
  },
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal {
  background: white;
  border-radius: 15px;
  padding: 25px 30px;
  width: 320px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.25);
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

h2 {
  margin-bottom: 20px;
  color: #2c3e50;
  font-weight: 700;
  font-size: 1.5rem;
  text-align: center;
}

form label {
  display: block;
  margin-bottom: 12px;
  font-weight: 600;
  color: #34495e;
}

input[type='text'],
input[type='number'],
input[type='url'] {
  width: 100%;
  padding: 7px 10px;
  margin-top: 4px;
  border: 1.5px solid #bdc3c7;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

input[type='text']:focus,
input[type='number']:focus,
input[type='url']:focus {
  border-color: #3498db;
  outline: none;
}

.buttons {
  display: flex;
  justify-content: space-between;
  margin-top: 25px;
}

.submit-button {
  background-color: #27ae60;
  border: none;
  border-radius: 10px;
  padding: 10px 25px;
  color: white;
  font-weight: 700;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.submit-button:hover {
  background-color: #1e8449;
}

.cancel-button {
  background-color: #e74c3c;
  border: none;
  border-radius: 10px;
  padding: 10px 25px;
  color: white;
  font-weight: 700;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.cancel-button:hover {
  background-color: #c0392b;
}
</style>
