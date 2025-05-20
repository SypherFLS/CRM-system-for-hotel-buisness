<template>
  <div class="body">
    <div id="home" class="container">
      <div class="left-side"></div>
      <div class="center-block">
        <h1>Сеть отелей «Брусника» — комфорт и доступность в центре города</h1>
        <p>
          «Брусника» — это сеть уютных отелей в Москве и других городах, предлагающая восхительные виды из окна, современные номера и качественный сервис по разумным ценам.
          Наши гостиницы идеально подходят как для деловых поездок, так и для туристов, ценящих комфорт и удобство.
        </p>
        <p>
          В каждом отеле «Брусника» вы найдете бесплатный Wi-Fi, круглосуточную стойку регистрации, удобные номера с необходимыми удобствами и возможность гибкого заселения, включая почасовую аренду.
          Мы заботимся о вашем комфорте и безопасности, обеспечивая чистоту и внимательное обслуживание.
        </p>
      </div>
      <div class="right-side"></div>
    </div>
    <section id="brusnichnye">
      <LHBalls />
    </section>
    <section id="hotel" class="info-block-section">
      <div class="info-blocks-inner">
        <ViewInfoBlock
          v-for="block in blocks"
          :key="block.id"
          :direction="block.direction"
          :title="block.title"
          :description="block.description"
          :imageSrc="block.imageSrc"
        >
          <template #text>
            <h2>{{ block.title }}</h2>
            <p>{{ block.description }}</p>
          </template>
          <template #image>
            <img :src="block.imageSrc" :alt="block.altText" />
          </template>
        </ViewInfoBlock>
      </div>
    </section>

    <footer id="contacts" class="footer">
      <div class="footer-container">
        <div class="footer-block contacts">
          <h4>Контакты</h4>
          <p>Москва, ул. Хорошевское шоссе, д. 26</p>
          <p><a href="tel:+79360006257">+7 936 000 62 57</a></p>
          <p><a href="mailto:info@brusnika24.ru">info@brusnika24.ru</a></p>
        </div>
        <div class="footer-block links">
          <h4>Меню</h4>
          <a href="#home">Главная</a><br />
          <a href="#hotel">Отели</a><br />
          <a href="#brusnichnye">Услуги</a><br />
          <a href="#contacts">Контакты</a>
        </div>
        <div class="footer-block social">
          <h4>Мы в соцсетях</h4>
          <a href="#" aria-label="Facebook" class="social-icon">📘</a>
          <a href="#" aria-label="Instagram" class="social-icon">📸</a>
          <a href="#" aria-label="VK" class="social-icon">🖖</a>
        </div>
      </div>
      <div class="footer-bottom">
        <p>© 2025 Сеть отелей «Брусника». Все права защищены.</p>
      </div>
    </footer>
  </div>
</template>

<script>
import hotel_image from '@/assets/visualB.png';
import gib from '@/assets/gibc.png';
import eco from '@/assets/price.png';
import service from '@/assets/service.png';
import comfortImage from '@/assets/comfort.png';
import ViewInfoBlock from '@/components/InfoBlock.vue';
import LHBalls from '@/components/LHBalls.vue';
export default {
  name: 'Home',
  components: {
    ViewInfoBlock,
    LHBalls,
  },
  data() {
    return {
      blocks: [
        {
          id: 1,
          title: 'Захватывающие красоты',
          description:
          'Отели «Брусника» расположены на локациях поражающий своей красотой',
          imageSrc: hotel_image,
          altText: 'Расположение отеля',
          direction: 'left',
        },
        {
          id: 2,
          title: 'Комфортные номера',
          description:
            'Современные номера с необходимыми удобствами: телевизор, бесплатный Wi-Fi, чайник, фен и туалетно-косметические принадлежности.',
          imageSrc: comfortImage,
          altText: 'Номера',
          direction: 'right',
        },
        {
          id: 3,
          title: 'Гибкие условия проживания',
          description:
          'Возможность почасовой аренды, поздний заезд и ранний выезд, а также продление проживания при наличии свободных номеров.',
          imageSrc: gib,
          altText: 'Гибкие условия',
          direction: 'left',
        },
        {
          id: 4,
          title: 'Сервис',
          description:
            'Общая кухня, прачечная, круглосуточная стойка регистрации и внимательный персонал, говорящий на 6 разных языках.',
          imageSrc: service,
          altText: 'Услуги',
          direction: 'right',
        },
        {
          id: 5,
          title: 'Доступные цены',
          description:
            'Оптимальное соотношение цены и качества, специальные предложения и акции для постоянных гостей.',
          imageSrc: eco,
          altText: 'Цены',
          direction: 'left',
        },
      ],
    };
  },
};
</script>

<style scoped>
:root {
  --color-primary: #4a7c59;
  --color-accent: #7dbf9e;
  --color-text-dark: #2e2e2e;
  --color-light: #f9f9f9;
  --color-button-bg: #4a7c59;
  --color-button-hover: #3a5c43;
}

.body {
  font-family: 'Open Sans', sans-serif;
  font-size: 18px;
  color: var(--color-text-dark);
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--color-light);
}

.container {
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-top: 50px;
}

.left-side,
.right-side {
  width: 10%;
  background-color: transparent;
}

.center-block {
  width: 80%;
  padding: 60px 40px;
  background-color: #fff;
  box-shadow: 0 12px 30px rgba(74, 124, 89, 0.15);
  border-radius: 15px;
  text-align: center;
  transition: box-shadow 0.3s ease;
}

.center-block:hover {
  box-shadow: 0 20px 40px rgba(74, 124, 89, 0.25);
}

h1 {
  font-size: 3rem;
  margin-bottom: 1.5rem;
  color: var(--color-primary);
  font-weight: 700;
}

p {
  font-size: 1.25rem;
  line-height: 1.8;
  color: var(--color-text-dark);
  max-width: 900px;
  margin: 0 auto 1.5rem auto;
}

.info-block-section {
  width: 100%;
  padding: 2rem 1rem;
  margin-top: 50px;
  background-color: white;
  border-radius: 15px;
  box-shadow: 0 10px 25px rgba(74, 124, 89, 0.1);
  text-align: center;
}

.loyalty-section {
  background-color: white;
  max-width: 900px;
  margin: 50px auto;
  padding: 3rem 2rem;
  border-radius: 15px;
  box-shadow: 0 10px 20px rgba(74, 124, 89, 0.1);
  text-align: center;
}

.loyalty-section h2 {
  color: var(--color-primary);
  margin-bottom: 1rem;
  font-size: 2.5rem;
}

.loyalty-section p {
  font-size: 1.2rem;
  line-height: 1.6;
  color: var(--color-text-dark);
}

.footer {
  background: linear-gradient(135deg, #2a2a2a 0%, #1e1e1e 100%);
  color: #dcdcdc;
  font-family: 'Open Sans', sans-serif;
  font-size: 16px;
  line-height: 1.7;
  padding: 50px 30px 30px;
  border-radius: 20px 20px 0 0;
  box-shadow:
    0 8px 24px rgba(0, 0, 0, 0.4),
    inset 0 0 50px rgba(255, 255, 255, 0.03);
  overflow: hidden;
}

.footer-container {
  max-width: 1100px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  gap: 2.5rem;
  flex-wrap: wrap;
}

.footer-block {
  flex: 1 1 250px;
  min-width: 220px;
}

.footer-block h4 {
  margin-bottom: 1.2rem;
  font-weight: 700;
  font-size: 1.5rem;
  color: #a3c4a1;
  letter-spacing: 1.5px;
  text-transform: uppercase;
}

.footer-block p,
.footer-block a {
  color: #dcdcdc;
  font-weight: 400;
  font-size: 1rem;
  text-decoration: none;
  transition: color 0.3s ease;
}

.footer-block a:hover {
  color: #a3c4a1;
  text-decoration: underline;
}

.social-icon {
  font-size: 26px;
  margin-right: 14px;
  transition: transform 0.3s ease, color 0.3s ease;
  display: inline-block;
  color: #dcdcdc;
}

.social-icon:hover {
  color: #a3c4a1;
  transform: scale(1.3);
}

.footer-bottom {
  margin-top: 40px;
  text-align: center;
  font-size: 14px;
  color: #a0a0a0;
  letter-spacing: 0.5px;
}

@media (max-width: 768px) {
  .container {
    flex-direction: column;
    align-items: center;
  }
  .center-block {
    width: 95%;
    padding: 30px 20px;
  }
  .footer-container {
    flex-direction: column;
    align-items: center;
  }
  .footer-block {
    text-align: center;
  }
  .left-side,
  .right-side {
    display: none;
  }
}
</style>
