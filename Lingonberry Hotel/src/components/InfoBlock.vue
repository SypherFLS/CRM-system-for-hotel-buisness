<template>
  <div :class="['view-info-block', { animated: isVisible }, { 'right-to-left': direction === 'right' }]" ref="blockRef">
    <div v-if="direction !== 'right'" class="text-side">
      <slot name="text"></slot>
    </div>
    <div class="image-side">
      <slot name="image"></slot>
    </div>
    <div v-if="direction === 'right'" class="text-side">
      <slot name="text"></slot>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const props = defineProps({
  direction: {
    type: String,
    default: 'left', 
    validator: value => ['left', 'right'].includes(value)
  },
});

const isVisible = ref(false);
const blockRef = ref(null);

onMounted(() => {
  const observerOptions = {
    rootMargin: '0px',
    threshold: 0.5 
  };

  const handleIntersection = ([entry]) => {
    if (entry.isIntersecting) {
      isVisible.value = true;
    }
  };

  const observer = new IntersectionObserver(handleIntersection, observerOptions);
  observer.observe(blockRef.value);
});
</script>

<style scoped>
.view-info-block {
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-between;
  align-items: stretch;
  margin-bottom: 2em;
  position: relative;
  overflow: hidden;
  height: clamp(300px, 80vh, 600px); 
}

.text-side {
  flex-basis: 33%; 
  padding: 1em;
  display: flex;
  flex-direction: column;
  justify-content: center; 
  align-items: center; 
  order: 1;
}

.right-to-left .text-side {
  order: 2;
}

.image-side {
  align-items: center;
  flex-basis: 67%; 
  padding: 1em;
  order: 2;
  opacity: 0;
  transform: translateX(50%) scale(0); 
  transition: all 0.5s cubic-bezier(.2,.8,.2,1), transform 0.7s ease-out;
}

.right-to-left .image-side {
  transform: translateX(-50%) scale(0); 
}

.animated .image-side {
  opacity: 1;
  transform: translateX(0) scale(1); 
}

.image-side img {
  width: 100%; 
  height: auto; 
  object-fit: contain; 
  object-position: center center; 
}
</style>