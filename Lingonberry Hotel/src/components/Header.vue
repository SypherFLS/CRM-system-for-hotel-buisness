<template>
  <header :class="{ stickyHeader: isSticky }" class="header">
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

      <h1 class="mobile-text">Lingonberry Hotel</h1>
    </nav>

    <div class="account-container">
      <div class="image-wrapper">
        <img src="@/assets/login_pic.png" alt="Login Icon" />
      </div>
      <div class="image-wrapper">
        <img src="@/assets/lines_pic.png" alt="Lines Icon" />
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'Header',
  data() {
    return {
      isSticky: false,
    };
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll);
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    handleScroll() {
      const scrollPosition = window.scrollY || document.documentElement.scrollTop;
      if (scrollPosition > 400) { 
        this.isSticky = true;
      } else {
        this.isSticky = false;
      }
    }
  }
};
</script>

<style scoped>
.header {
  margin: 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: transparent;
  position: fixed; 
  top: 0;
  left: 0;
  right: 0; 
  z-index: 1000;
  transition: background-color 0.3s ease, height 0.3s ease, box-shadow 0.3s ease; /* Плавность изменений */
  min-height: 70px; 
}

.stickyHeader {

  min-height: 50px; 
}

.logo-container img {
  width: 100px;
  height: auto;
  padding-left: 20px;
  object-fit: cover;
}
.logo-container img:hover {
  transform: scale(1.15);
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
  margin-left: auto;
  display: flex;
  align-items: center;
}

.image-wrapper {
  position: relative;
  display: inline-block;
  overflow: hidden;
  margin-right: 20px;
  transition: filter 0.3s ease; 
}

.image-wrapper img {
  width: 40px;
  height: auto;
  filter: invert(1); 
}

.image-wrapper:hover img {
  filter: grayscale(0%) sepia(100%) saturate(3000%) brightness(150%) hue-rotate(-50deg);
}

.nav-bar a {
  text-decoration: none;
  color: inherit;
  transition: all .3s ease-in-out;
  position: relative;
}

.nav-bar a:hover {
  transform: scale(1.05);
  color: #951C1C;
}

.mobile-text {
  display: none; 
  font-size: 24px;
  color: white; 
  font-weight: bold;
  text-transform: uppercase;
  letter-spacing: 2px;
  transition: color 0.3s ease; 
}

.mobile-text:hover {
  color: #B50D1E; 
  cursor: none; 
}

@media screen and (max-width: 968px) {
  .mobile-text {
    display: block;
  }

  .nav-bar ul.nav-list {
    display: none;
  }
}
</style>