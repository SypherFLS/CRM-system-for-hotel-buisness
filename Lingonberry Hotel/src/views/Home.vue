<template>
  <div class="body">
    <div id="home" class="container">
      <div class="left-side"></div>
      <div class="center-block">
        <h1 style="font-size: 36px;">Добро пожаловать в мир комфорта и гостеприимства!</h1>
        <p style="line-height: 1.7;">
          Наша компания предлагает уникальные возможности для идеального отдыха и успешного ведения бизнеса. Отель расположен в живописном месте, идеально подходит как для деловых поездок, так и для романтического отпуска или семейного путешествия. Здесь вы найдете идеальное сочетание современного дизайна номеров, высоких стандартов обслуживания и неповторимой атмосферы уюта и спокойствия.

          Каждый номер оснащен современной мебелью, удобной кроватью, высококачественным постельным бельем, удобствами премиум-класса и всеми необходимыми элементами, такими как телевизор с плоским экраном, бесплатный Wi-Fi, мини-бар и сейф. Вы сможете наслаждаться великолепным видом из окна и окунуться в атмосферу гармонии и умиротворенности, оставаясь сосредоточенным на своих делах или наслаждаясь свободным временем.

          Для наших гостей предусмотрены разнообразные услуги, среди которых внимательное обслуживание персонала, вкусный завтрак каждое утро, возможность заказать дополнительные блюда и напитки прямо в номер, наличие конференц-залов для организации встреч и мероприятий различного формата. Специальная зона релаксации позволит отдохнуть душой и телом после насыщенного дня, воспользовавшись услугами СПА-салона или фитнес-центра отеля.

          Также наша команда специалистов готова организовать культурную программу, экскурсии и досуговые мероприятия, предоставив гостям уникальную возможность познакомиться с местной культурой и традициями региона, посетить знаменитые достопримечательности и интересные места поблизости.

          Выбирая наше заведение, вы получаете больше, чем просто проживание — вы открываете двери в мир удовольствия, удобства и безупречного сервиса. Позвольте нам стать частью ваших лучших воспоминаний и подарить впечатления, которыми захочется поделиться снова и снова.

          Присоединяйтесь к числу тех, кто уже оценил высокий уровень нашего гостиничного сервиса и остался доволен каждым моментом своего пребывания в нашем заведении. Мы уверены, что именно здесь вы почувствуете истинное значение понятия "гостеприимство" и обретете новый опыт путешествий и жизни.
        </p>
      </div>
      <div class="right-side"></div>
    </div>

    <section id="hotel" class="info-block-section">
      <div
        style="
          background-color: white;
          padding: 2rem;
          box-shadow: rgba(0, 0, 0, 0.1) 0px 10px 15px -3px,
            rgba(0, 0, 0, 0.05) 0px 4px 6px -2px;
          width: 100%;
          position: relative;
          top: -25%;
          border-radius: 0 0 5% 5%;
        "
      >
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

    <!-- Новый блок "Брусничные Баллы" -->
    <section id="brusnichnye" class="loyalty-section">
      <div class="loyalty-container">
        <h2>Брусничные Баллы — наша система лояльности</h2>
        <p>
          Накопительные брусничные баллы позволяют получать скидки, специальные предложения и подарки.
          Чем больше вы отдыхаете у нас, тем больше преимуществ получаете! Просто регистрируйтесь и начинайте накапливать баллы при каждом бронировании.
        </p>
      </div>
    </section>

    <!-- Футер с контактами -->
    <footer id="contacts" class="footer">
      <div class="footer-container">
        <h3>Контакты</h3>
        <p>Адрес: ул. Примерная, 123, Город</p>
        <p>Телефон: +7 (123) 456-78-90</p>
        <p>Email: info@lingonberryhotel.ru</p>
        <p>
          Мы в соцсетях:
          <a href="#" target="_blank" rel="noopener">Facebook</a> |
          <a href="#" target="_blank" rel="noopener">Instagram</a> |
          <a href="#" target="_blank" rel="noopener">VK</a>
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import hotel_image from '@/assets/hotel_pic.png';
import ViewInfoBlock from '@/components/InfoBlock.vue';

export default {
  name: 'Home',
  components: {
    ViewInfoBlock,
  },
  data() {
    return {
      blocks: [
        {
          id: 1,
          title: 'Первый блок',
          description: 'Описание первого блока.',
          imageSrc: hotel_image,
          altText: 'Image 1',
          direction: 'left',
        },
        {
          id: 2,
          title: 'Второй блок',
          description: 'Описание второго блока.',
          imageSrc: hotel_image,
          altText: 'Image 2',
          direction: 'right',
        },
        {
          id: 3,
          title: 'Третий блок',
          description: 'Описание третьего блока.',
          imageSrc: hotel_image,
          altText: 'Image 3',
          direction: 'left',
        },
        {
          id: 4,
          title: 'Четвёртый блок',
          description: 'Описание четвёртого блока.',
          imageSrc: hotel_image,
          altText: 'Image 4',
          direction: 'right',
        },
        {
          id: 5,
          title: 'Пятый блок',
          description: 'Описание пятого блока.',
          imageSrc: hotel_image,
          altText: 'Image 5',
          direction: 'left',
        },
      ],
    };
  },
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

:root {
  --color-primary: #FFDAB9;
  --color-accent: #CD853F;
  --color-text-dark: #333333;
  --color-light: #FDFCFB;
}

.body {
  font-family: 'Open Sans', sans-serif;
  font-size: 18px;
  color: var(--color-text-dark);
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: var(--color-light);
}

.top-part {
  width: 100%;
  height: fit-content;
  position: relative;
  overflow: hidden;
  padding-bottom: 10px;
}

.background-pic {
  display: block;
  width: 100%;
  height: auto;
  object-fit: cover;
  filter: brightness(0.9);
}

.container {
  display: flex;
  justify-content: space-between;
  width: 100%;
  height: 100%;
  margin-top: 50px;
}

.left-side,
.right-side {
  width: 10%;
  height: 100%;
  background-color: transparent;
}

.center-block {
  width: 80%;
  padding: 50px;
  padding-bottom: 50px;
  background-color: white;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1), 0 6px 6px rgba(0, 0, 0, 0.05);
  border-radius: 5%;
  text-align: center;
}

h1 {
  font-size: 50px;
  margin-bottom: 1rem;
  text-align: center;
  color: #333;
}

p {
  font-size: 30px;
  margin-bottom: 1.5rem;
  text-align: justify;
}

.info-block-section {
  width: 100%;
  position: relative;
  padding: 2rem;
  margin-top: 50px;
  text-align: center;
  background-color: white;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1), 0 6px 6px rgba(0, 0, 0, 0.05);
  border-radius: 0 0 5% 5%;
}

/* Новый блок Брусничные Баллы */
.loyalty-section {
  background-color: white;
  max-width: 900px;
  margin: 50px auto;
  padding: 3rem 2rem;
  border-radius: 5%;
  box-shadow: 0 10px 20px rgba(0,0,0,0.1), 0 6px 6px rgba(0,0,0,0.05);
  text-align: center;
}

.loyalty-section h2 {
  color: var(--color-accent);
  margin-bottom: 1rem;
  font-size: 36px;
}

.loyalty-section p {
  font-size: 22px;
  line-height: 1.6;
  color: var(--color-text-dark);
}

.footer {
  background: #181818;
  color: #fff;
  padding: 2rem 1rem 1rem 1rem;
  font-size: 16px;
}

.footer-container {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  max-width: 1200px;
  margin: 0 auto 1rem auto;
  gap: 2rem;
}

.footer-block {
  min-width: 180px;
  flex: 1 1 200px;
}

.footer-block h4 {
  margin-bottom: 1rem;
  font-size: 18px;
  color: #FFDAB9;
}

.footer-block a {
  color: #FFDAB9;
  text-decoration: none;
  margin-right: 0.5rem;
  transition: color 0.3s;
}

.footer-block a:hover {
  color: #fff;
}

.footer-block img {
  width: 28px;
  height: 28px;
  margin-right: 0.5rem;
  vertical-align: middle;
}

.footer-bottom {
  border-top: 1px solid #444;
  padding-top: 1rem;
  text-align: center;
  font-size: 14px;
  color: #aaa;
}

@media (max-width: 800px) {
  .footer-container {
    flex-direction: column;
    align-items: center;
    gap: 1.5rem;
  }
  .footer-block {
    min-width: unset;
    text-align: center;
  }
}

</style>
