<template>
  <div class="form-back">
    <form @submit.prevent="registerUser">
      <label for="name">Имя:</label>
      <input type="text" v-model.trim="user.name" required placeholder="Ваше имя" /><br />
      
      <label for="email">Email:</label>
      <input type="email" v-model.trim="user.email" required placeholder="example@example.com" /><br />
      
      <label for="password">Пароль:</label>
      <input type="password" v-model="user.password" required placeholder="Введите пароль" /><br />
      
      <label for="confirm-password">Повторите Пароль:</label>
      <input type="password" v-model="confirmPassword" required placeholder="Введите пароль повторно" /><br />

      <p v-if="errorMessage" style="color: red;">{{ errorMessage }}</p>
      
      <button type="submit">Зарегистрироваться</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios'; 

export default {
  data() {
    return {
      user: {
        name: '',
        email: '',
        password: ''
      },
      confirmPassword: '',
      errorMessage: ''
    };
  },
  methods: {
    async registerUser() {
      this.errorMessage = '';

      if (this.user.password !== this.confirmPassword) {
        this.errorMessage = 'Пароли не совпадают.';
        return;
      }

      
      if (this.user.password.length < 6) {
        this.errorMessage = 'Пароль должен содержать минимум 6 символов.';
        return;
      }

      try {
        
        const response = await axios.post('/api/register', this.user);

        if (response.data.success) {
          alert(`Пользователь ${this.user.name} успешно зарегистрирован.`);
          this.user.name = '';
          this.user.email = '';
          this.user.password = '';
          this.confirmPassword = '';
        } else {
          this.errorMessage = response.data.message || 'Ошибка регистрации.';
        }
      } catch (error) {
        this.errorMessage = error.response?.data?.message || 'Ошибка соединения с сервером.';
      }
    }
  }
};
</script>

<style scoped>
.form-back {
  background-color: brown;
  border-radius: 15px;
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  max-width: 400px;
  margin: 0 auto;
}

form {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  width: 100%;
}

input {
  margin-bottom: 10px;
  padding: 5px;
  width: 100%;
}

button {
  padding: 10px;
  cursor: pointer;
}
</style>
