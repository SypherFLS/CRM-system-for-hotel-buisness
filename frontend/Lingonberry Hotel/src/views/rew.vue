<template>
  <div class="chat-page">
    <h1>Отзывы клиентов</h1>

    <div class="reviews-widgets">
      <div
        v-for="review in reviews"
        :key="review.id"
        class="review-widget"
        :class="review.type"
        @click="selectReview(review)"
        :title="review.shortText"
      >
        {{ review.type === 'positive' ? '😊' : '⚠️' }}
      </div>
    </div>

    <transition name="fade">
      <div v-if="selectedReview" class="review-detail">
        <h2>Отзыв #{{ selectedReview.id }}</h2>
        <p class="full-text">{{ selectedReview.fullText }}</p>
        <p><strong>Решение по отзыву:</strong></p>
        <textarea v-model="selectedReview.solution" placeholder="Введите решение..." rows="4"></textarea>
        <div class="buttons">
          <button class="solve-btn" @click="solveReview">Решить</button>
          <button class="close-btn" @click="closeDetail">Закрыть</button>
        </div>
      </div>
    </transition>

    <p v-if="reviews.length === 0" class="empty-text">Все отзывы решены!</p>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const reviews = ref([
  {
    id: 1,
    type: 'positive',
    shortText: 'Отличный сервис и быстрое обслуживание',
    fullText:
      'Очень доволен обслуживанием, персонал внимательный и профессиональный. Все вопросы решаются быстро и качественно.',
    solution: '',
  },
  {
    id: 2,
    type: 'negative',
    shortText: 'Долгое ожидание ответа поддержки',
    fullText:
      'Обратился в поддержку и ждал ответа более суток, что неприемлемо для такого сервиса. Требуется улучшение времени реакции.',
    solution: '',
  },
  {
    id: 3,
    type: 'negative',
    shortText: 'Проблемы с бронированием номера',
    fullText:
      'При бронировании возникли ошибки, и номер не был подтверждён. Очень расстроен, надеюсь на скорое решение проблемы.',
    solution: '',
  },
])

const selectedReview = ref(null)

function selectReview(review) {
  selectedReview.value = review
}

function solveReview() {
  if (!selectedReview.value.solution.trim()) {
    alert('Пожалуйста, введите решение перед тем, как решить отзыв.')
    return
  }
  reviews.value = reviews.value.filter(r => r.id !== selectedReview.value.id)
  selectedReview.value = null
}

function closeDetail() {
  selectedReview.value = null
}
</script>

<style scoped>
.chat-page {
  max-width: 800px;
  margin: 40px auto;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background: #f7f9fc;
  padding: 20px 30px;
  border-radius: 12px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
  text-align: center;
}

h1 {
  color: #34495e;
  margin-bottom: 30px;
}


.reviews-widgets {
  display: flex;
  gap: 20px;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 40px;
}

.review-widget {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  font-size: 28px;
  line-height: 60px;
  cursor: pointer;
  user-select: none;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s, box-shadow 0.2s;
  display: flex;
  align-items: center;
  justify-content: center;
}

.review-widget:hover {
  transform: scale(1.15);
  box-shadow: 0 6px 14px rgba(0, 0, 0, 0.15);
}

.review-widget.positive {
  background-color: #27ae60;
  color: white;
}

.review-widget.negative {
  background-color: #e74c3c;
  color: white;
}


.review-detail {
  background: white;
  border-radius: 12px;
  padding: 25px 30px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.12);
  max-width: 600px;
  margin: 0 auto 30px auto;
  text-align: left;
}

.review-detail h2 {
  margin-top: 0;
  color: #2980b9;
}

.full-text {
  font-size: 16px;
  margin: 15px 0 25px 0;
  color: #2c3e50;
  white-space: pre-wrap;
}

textarea {
  width: 100%;
  resize: vertical;
  padding: 10px;
  font-size: 14px;
  border-radius: 8px;
  border: 1px solid #bdc3c7;
  margin-bottom: 20px;
  font-family: inherit;
  box-sizing: border-box;
}

.buttons {
  display: flex;
  gap: 10px;
}

.solve-btn {
  background-color: #2980b9;
  color: white;
  border: none;
  padding: 10px 20px;
  font-weight: 600;
  cursor: pointer;
  border-radius: 8px;
  transition: background-color 0.3s;
  flex-grow: 1;
}

.solve-btn:hover {
  background-color: #1c5980;
}

.close-btn {
  background-color: #7f8c8d;
  color: white;
  border: none;
  padding: 10px 18px;
  font-weight: 600;
  cursor: pointer;
  border-radius: 8px;
  transition: background-color 0.3s;
  flex-grow: 1;
}

.close-btn:hover {
  background-color: #626f73;
}

.empty-text {
  color: #7f8c8d;
  font-size: 18px;
  margin-top: 40px;
}


.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
