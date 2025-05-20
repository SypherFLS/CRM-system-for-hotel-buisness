<template>
  <header :class="{ stickyHeader: isSticky }" class="header">
    <div class="logo-container" v-show="!isSticky">
      <img src="@/assets/logo.png" alt="Logo" />
    </div>

    <nav class="nav-bar">
      <ul class="nav-list">
        <li><router-link to="/">Главная</router-link></li>
        <li><router-link to="/hotel">Наши отели</router-link></li>
        <li><a href="#">Брусничные Баллы</a></li>
        <li><a href="#">О нас</a></li>
        <li><a href="#">Контакты</a></li>
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
      this.isSticky = scrollPosition > 400;
    },
  },
};
</script>

<style scoped>
.header {
  margin: 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #ffdabc; 
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  width: 100%;
  height: 120px;
  padding: 0 30px;
  box-sizing: border-box;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.05);
  border-bottom-left-radius: 15px;
  border-bottom-right-radius: 15px;
  transition:
    background-color 0.4s ease,
    height 0.4s ease,
    box-shadow 0.4s ease;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}


.stickyHeader {
  height: 70px;
  background-color: #000000; 
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.4);
  border-radius: 0;
}


.logo-container img {
  padding-left: 40px;
  width: 140px;
  height: auto;
  object-fit: contain;
  transition: transform 0.3s ease;
  filter: drop-shadow(0 2px 2px rgba(0,0,0,0.15));
}

.logo-container img:hover {
  transform: scale(1.15);
}


.nav-bar {
  flex-grow: 1;
  text-align: center;
}

.nav-list {
  font-size: 28px;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  display: inline-flex;
  list-style: none;
  padding: 0;
  margin: 0;
  gap: 2rem;
  align-items: center;
  color: azure;
  transition: color 0.3s ease;

}

.nav-list a {
  text-decoration: none;
  color: inherit;
  transition: color 0.3s ease, transform 0.3s ease;
  position: relative;
  padding: 4px 0;
}


.nav-list a:hover {
  color: #B50D1E;
  transform: scale(1.1);
}


.stickyHeader .nav-list {
  color: white;
}

.stickyHeader .nav-list a {
  color: white;
}

.stickyHeader .nav-list a:hover {
  color: #B50D1E;
}


.account-container {
  display: flex;
  align-items: center;
  gap: 25px;
  color: inherit;
}


.image-wrapper img {
  width: 40px;
  height: auto;
  filter: none;
  transition: filter 0.3s ease;
  cursor: pointer;
}


.stickyHeader .image-wrapper img {
  filter: brightness(0) invert(1);
}

.image-wrapper:hover img {
  filter: none;
}


.mobile-text {
  display: none;
  font-size: 24px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 2px;
  color: inherit;
  transition: color 0.3s ease;
}

.mobile-text:hover {
  color: #B50D1E;
  cursor: default;
}

.nav-list a.router-link-exact-active {
  text-decoration: underline;
  font-weight: 700;
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
