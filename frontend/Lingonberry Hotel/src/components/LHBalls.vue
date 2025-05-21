<template>
  <section class="loyalty-fullwidth-section">
    <div class="loyalty-content">
      <div class="text-panel">
        <h2>Брусничные Баллы — ваша выгода и забота о каждом госте</h2>
        <p>
          Наша уникальная программа лояльности позволяет накапливать баллы за каждое бронирование и использование услуг отеля. Обменивайте баллы на скидки, подарки и эксклюзивные предложения. Чем больше вы с нами, тем приятнее отдых!
        </p>
        <ul>
          <li>Начисление баллов за проживание и дополнительные услуги</li>
          <li>Подарки и скидки на будущие бронирования</li>
          <li>Персональные предложения и акции для участников программы</li>
          <li>Простая регистрация и удобное управление баллами в личном кабинете</li>
        </ul>
      </div>
    </div>
  </section>
</template>

<style scoped>
.loyalty-fullwidth-section {
  margin-top: 85px;
  width: 100vw;
  min-height: 300px;
  padding: 40px 20px;
  background-image: url('src/assets/ppp.png');
  background-size: cover;
  background-position: center center;
  background-repeat: no-repeat;
  color: #5a3e00;
  text-align: center;
  position: relative;
  overflow: hidden;
  font-family: 'Open Sans', sans-serif;
}

.loyalty-content {
  max-width: 900px;
  margin: 0 auto;
  padding: 0 10px;
}


.text-panel {
  background-color: rgba(255, 255, 255, 0.85); 
  border-radius: 12px;
  padding: 30px 40px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
  color: #5a3e00; 
  text-align: left;
}

.text-panel h2 {
  font-size: 2.4rem;
  font-weight: 700;
  margin-bottom: 1rem;
  color: #b36b00;
  text-shadow: none;
}

.text-panel p {
  font-size: 1.1rem;
  margin-bottom: 1.8rem;
  color: #7a4f00;
  text-shadow: none;
}

.text-panel ul {
  list-style: none;
  padding: 0;
  max-width: 600px;
  margin: 0;
}

.text-panel ul li {
  font-size: 1.05rem;
  margin-bottom: 0.7rem;
  padding-left: 1.5em;
  position: relative;
  color: #7a4f00;
}

.text-panel ul li::before {
  content: '✔';
  position: absolute;
  left: 0;
  color: #ffb97b;
  font-weight: bold;
  font-size: 1.1rem;
  top: 0;
  line-height: 1;
}

@media (max-width: 768px) {
  .loyalty-fullwidth-section {
    padding: 30px 15px;
    min-height: auto;
  }
  .text-panel {
    padding: 20px 25px;
  }
  .text-panel h2 {
    font-size: 1.8rem;
  }
  .text-panel p {
    font-size: 1rem;
  }
  .text-panel ul {
    text-align: center;
  }
  .text-panel ul li {
    padding-left: 0;
    display: inline-block;
    margin: 0 10px 10px 10px;
    position: static;
  }
  .text-panel ul li::before {
    content: '';
  }
}
</style>
