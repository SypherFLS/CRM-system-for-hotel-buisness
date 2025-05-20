<template>
  <div id="app" class="body">
    <div class="top-part">
      <div class="toper-container">
        <Header></Header>
        <img class="background-pic" src="@/assets/hotel_pic.png" alt="hotel"/>
      </div>
    </div>

    <section class="info-block-section">
      <div style="
           background-color: white;
           padding: 2rem;
           box-shadow: rgba(0, 0, 0, 0.1) 0px 10px 15px -3px, rgba(0, 0, 0, 0.05) 0px 4px 6px -2px;
           width: 100%;
           position: relative;
           top: -25%;
           border-radius: 0 0 5% 5%;
       ">
        
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
  </div>
</template>

<script>
import hotel_image from '@/assets/hotel_pic.png';
import Header from '@/components/Header.vue';
import RegistrationForm from '@/components/Registration-form.vue';
import ViewInfoBlock from '@/components/InfoBlock.vue';

export default {
  name: 'app',
  components: {
    Header,
    RegistrationForm,
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
  font-family: Roboto, Helvetica Neue, Arial, sans-serif;
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


.info-block-section {
  width: 100%;
  position: relative;
  top: -25%; 
  padding: 2rem;
  text-align: center;
  background-color: white;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1), 0 6px 6px rgba(0, 0, 0, 0.05);
  border-radius: 0 0 5% 5%; 
}


.info-blocks-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  gap: 2rem;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr)); 
}


h2 {
  font-size: 1.5em;
  color: var(--color-accent);
  margin-bottom: 0.5rem;
}

p {
  font-size: 1em;
  line-height: 1.6;
  color: #555;
}


img {
  width: 100%;
  height: auto;
  border-radius: 5%;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1), 0 6px 6px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s ease-in-out;

  &:hover {
    transform: scale(1.05); 
  }
}


button {
  display: inline-block;
  padding: 0.75em 1.5em;
  font-size: 1em;
  font-weight: bold;
  color: #fff;
  background-color: var(--color-accent);
  border: none;
  cursor: pointer;
  transition: background-color 0.3s ease-in-out;

  &:hover {
    background-color: darken(var(--color-accent), 10%); 
  }
}

hr {
  border: none;
  height: 1px;
  background-color: rgba(var(--color-accent), 0.5);
  margin: 1rem 0;
}

.block-title::before {
  content: '';
  display: inline-block;
  width: 30px;
  height: 3px;
  background-color: var(--color-accent);
  margin-right: 10px;
  vertical-align: middle;
}
</style>