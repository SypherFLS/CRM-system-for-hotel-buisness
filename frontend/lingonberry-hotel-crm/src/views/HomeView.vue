<template>
  <div class="home">
    <div class="hero">
      <h1>Отель «Брусника»</h1>
      <p>Где каждый гость — как дома</p>
      <button class="glow-button" @click="startConfetti">Запустить праздник</button>
      <div v-if="confetti" class="confetti-container"></div>
    </div>
    <div class="features">
      <div class="feature-card" v-for="(feature, i) in features" :key="i" @mouseenter="hoverCard(i)" @mouseleave="unhoverCard(i)">
        <div class="icon">{{ feature.icon }}</div>
        <h3>{{ feature.title }}</h3>
        <p>{{ feature.text }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const confetti = ref(false)
const features = [
  { icon: '🛏️', title: 'Уютные номера', text: 'Комфорт и стиль в каждом номере' },
  { icon: '🧹', title: 'Чистота', text: 'Безупречная чистота и порядок' },
  { icon: '🍽️', title: 'Питание', text: 'Вкусные завтраки и ужины' }
]
const hoveredCard = ref(null)

function startConfetti() {
  confetti.value = true
  setTimeout(() => { confetti.value = false }, 3000)
}
function hoverCard(i) {
  hoveredCard.value = i
}
function unhoverCard(i) {
  if (hoveredCard.value === i) hoveredCard.value = null
}
</script>

<style scoped>
.home {
  padding: 40px;
  max-width: 1200px;
  margin: 0 auto;
  background: #fff;
  border-radius: 24px;
  box-shadow: 0 16px 40px rgba(0, 0, 0, 0.08);
}
.hero {
  text-align: center;
  margin-bottom: 60px;
  position: relative;
  overflow: hidden;
}
.hero h1 {
  font-size: 3rem;
  color: #a83b3b;
  margin-bottom: 16px;
  font-weight: 700;
}
.hero p {
  font-size: 1.5rem;
  color: #6d6d6d;
  margin-bottom: 32px;
}
.glow-button {
  background: #a83b3b;
  color: white;
  border: none;
  padding: 16px 32px;
  font-size: 1.2rem;
  border-radius: 50px;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 20px rgba(168, 59, 59, 0.3);
}
.glow-button:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(168, 59, 59, 0.4);
}
.glow-button:active {
  transform: translateY(0);
}
.features {
  display: flex;
  gap: 24px;
  justify-content: center;
  flex-wrap: wrap;
}
.feature-card {
  background: #f8f8f8;
  border-radius: 16px;
  padding: 24px;
  width: 280px;
  transition: all 0.3s ease;
  cursor: pointer;
  border: 2px solid transparent;
}
.feature-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 16px 32px rgba(168, 59, 59, 0.1);
  border-color: #a83b3b;
}
.feature-card .icon {
  font-size: 2.5rem;
  margin-bottom: 16px;
}
.feature-card h3 {
  color: #a83b3b;
  font-size: 1.5rem;
  margin-bottom: 12px;
}
.feature-card p {
  color: #6d6d6d;
  font-size: 1rem;
}
.confetti-container {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 100;
}

</style>
