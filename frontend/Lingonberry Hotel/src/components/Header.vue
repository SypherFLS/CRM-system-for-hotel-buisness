<template>
  <header :class="{ stickyHeader: isSticky }" class="header">
    <div class="logo-container" v-show="!isSticky">
      <img src="@/assets/logo.png" alt="Logo" />
    </div>

    <nav class="nav-bar">
      <ul class="nav-list">
        <li>
          <router-link to="/" exact-active-class="router-link-exact-active">
            Главная
          </router-link>
        </li>
        <li>
          <router-link to="/hotel" exact-active-class="router-link-exact-active">
            Наши отели
          </router-link>
        </li>
        <li>
          <router-link to="/brone" exact-active-class="router-link-exact-active">
            Бронирование
          </router-link>
        </li>
        <li>
          <a
            href="#brusnichnye"
            :class="{ activeAnchor: activeSection === 'brusnichnye' }"
            @click.prevent="scrollToSection('brusnichnye')"
          >
            Брусничные Баллы
          </a>
        </li>
        <li>
          <a
            href="#hotel"
            :class="{ activeAnchor: activeSection === 'hotel' }"
            @click.prevent="scrollToSection('hotel')"
          >
            О нас
          </a>
        </li>
        <li>
          <a
            href="#contacts"
            :class="{ activeAnchor: activeSection === 'contacts' }"
            @click.prevent="scrollToSection('contacts')"
          >
            Контакты
          </a>
        </li>
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
      activeSection: null,
    };
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll);
    window.addEventListener('scroll', this.detectActiveSection);
    this.detectActiveSection();
  },
  beforeUnmount() {
    window.removeEventListener('scroll', this.handleScroll);
    window.removeEventListener('scroll', this.detectActiveSection);
  },
  methods: {
    handleScroll() {
      const scrollPosition = window.scrollY || document.documentElement.scrollTop;
      this.isSticky = scrollPosition > 400;
    },
    scrollToSection(id) {
      const el = document.getElementById(id);
      if (el) {
        el.scrollIntoView({ behavior: 'smooth' });
      }
    },
    detectActiveSection() {
      const sections = ['brusnichnye', 'hotel', 'contacts'];
      const scrollPos = window.scrollY + window.innerHeight / 3;

      for (let i = sections.length - 1; i >= 0; i--) {
        const el = document.getElementById(sections[i]);
        if (el && el.offsetTop <= scrollPos) {
          this.activeSection = sections[i];
          return;
        }
      }
      this.activeSection = null;
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
  filter: drop-shadow(0 2px 2px rgba(0, 0, 0, 0.15));
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

.nav-list a,
.nav-list router-link {
  text-decoration: none;
  color: inherit;
  transition: color 0.3s ease, transform 0.3s ease;
  position: relative;
  padding: 4px 0;
  cursor: pointer;
}

.nav-list a:hover,
.nav-list router-link:hover {
  color: #B50D1E;
  transform: scale(1.1);
}

.stickyHeader .nav-list {
  color: white;
}

.stickyHeader .nav-list a,
.stickyHeader .nav-list router-link {
  color: white;
}

.stickyHeader .nav-list a:hover,
.stickyHeader .nav-list router-link:hover {
  color: #B50D1E;
}

/* Подчеркивание для активного router-link */
.router-link-exact-active {
  text-decoration: underline;
  font-weight: 700;
}

/* Подчеркивание для активного якоря */
.activeAnchor {
  text-decoration: underline;
  font-weight: 700;
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

@media screen and (max-width: 968px) {
  .mobile-text {
    display: block;
  }

  .nav-bar ul.nav-list {
    display: none;
  }
}
</style>
