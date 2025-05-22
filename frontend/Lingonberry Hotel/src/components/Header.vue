<template>
  <div class="header">
    <div class="logo">
      <img src="@/assets/logo.png" alt="Logo" />
    </div>
    <nav>
      <router-link to="/" class="nav-link">
        <img src="@/assets/an.png" alt="Page 1" />
      </router-link>
      <router-link to="/hotel" class="nav-link">
        <img src="@/assets/hotel.png" alt="Page 2" />
      </router-link>
      <router-link to="/brone" class="nav-link">
        <img src="@/assets/lines_pic.png" alt="Page 3" />
      </router-link>
      <router-link to="/rew" class="nav-link">
        <img src="@/assets/login_pic.png" alt="Page 3" />
      </router-link>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'Header',
};
</script>

<style scoped>
.header {
  position: fixed;
  top: 0;
  right: 0;
  height: 100vh;
  width: 45px; 
  background-color: rgba(255, 218, 185, 0.7);
  border-left: 1px solid black;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding-top: 12px;
  transition: background-color 0.3s ease;
}

.header:hover {
  background-color: rgba(255, 218, 185, 0.9);
}


.logo {
  width: 100%; 
  display: flex;
  justify-content: center; 
  margin-bottom: 20px;
}

.logo img {
  width: 55px; 
  height: 55px;
  animation: logoAppear 0.6s ease forwards;
  opacity: 0;
  transform: scale(0.8);
  transition: transform 0.3s ease;
  cursor: pointer;
}

.logo img:hover {
  transform: scale(1.15);
}

@keyframes logoAppear {
  to {
    opacity: 1;
    transform: scale(1);
  }
}

nav {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.nav-link img {
  width: 32px;
  height: 32px;
  cursor: pointer;
  border-radius: 6px;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  opacity: 0;
  animation: iconAppear 0.5s ease forwards;
}

.nav-link:nth-child(1) img {
  animation-delay: 0.7s;
}
.nav-link:nth-child(2) img {
  animation-delay: 0.8s;
}
.nav-link:nth-child(3) img {
  animation-delay: 0.9s;
}

.nav-link img:hover {
  transform: scale(1.3);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  filter: drop-shadow(0 0 2px rgba(0,0,0,0.3));
}

@keyframes iconAppear {
  to {
    opacity: 1;
  }
}
</style>
