<!-- RegistrationForm.vue -->
<template>
  <div class="form-back">
    <form @submit.prevent="registerUser">
      <label for="name">Имя:</label>
      <input type="text" v-model="user.name" required placeholder="Ваше имя"><br>
      
      <label for="email">Email:</label>
      <input type="email" v-model="user.email" required placeholder="example@example.com"><br>
      
      <label for="password">Пароль:</label>
      <input type="password" v-model="user.password" required placeholder="Введите пароль"><br>
      
      <label for="confirm-password">Повторите Пароль:</label>
      <input type="password" required placeholder="Введите пароль повторно"><br>
      
      <button type="submit">Зарегистрироваться</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        name: '',
        email: '',
        password: ''
      }
    };
  },
  methods: {
    registerUser() {
      alert(`Пользователь ${this.user.name} зарегистрирован.`);
      console.log(this.user);
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