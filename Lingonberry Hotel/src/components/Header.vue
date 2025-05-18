<template>
  <header class="header">
    <div class="logo-container">
      <img src="@/assets/logo.png" alt="Logo" />
    </div>
    <nav class="nav-bar">
      <ul class="nav-list">
        <li><a href="#">Главная</a></li>
        <li><a href="#">Наши отели</a></li>
        <li><a href="#">Брусничные Баллы</a></li>
        <li><a href="#">Контакты</a></li>
        <li><a href="#">О нас</a></li>
      </ul>
    </nav>
    <div class="account-container">
        <img src="@/assets/login_pic.png" alt="account" />
        <img src="@/assets/lines_pic.png" alt="account" />
    </div>
  </header>
</template>

<script>
export default {
  name: 'Header'
}
</script>

<style scoped>
.header {
  margin: 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: transparent;
  position: sticky;
  top: 0;
  z-index: 1000;
}

.logo-container img {
  width: 100px;
  height: auto;
  padding-left: 20px;
  object-fit: cover;
}

.nav-bar {
  
  flex-grow: 1;
  text-align: center;
}

.nav-list {
  font-size: 30px;
  color: azure;
  font-family: sans-serif;
  display: inline-block;
}

.nav-bar ul {
  list-style-type: none;
  display: flex;
  gap: 1.5rem;
  justify-content: center; 
}

.account-container {
  justify-self: right;
}

.account-container img {
  filter: invert(100%);
  width: 40px;
  height: auto;
  padding-right: 20px;
}

.nav-bar a {
  text-decoration: none;
  color: inherit;
  transition: all .3s ease-in-out;
}

.nav-bar a:hover {
  transform: scale(1.05);
  color: #951c1c;
}

@media screen and (max-width: 768px) {
  .nav-bar ul {
    flex-direction: column;
    gap: 0.5rem;
  }

  .nav-bar li {
    margin-bottom: 0.5rem;
  }
}
</style>