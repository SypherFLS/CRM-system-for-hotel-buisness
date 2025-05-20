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
  align-items: center;
  margin-bottom: 3em;
  position: relative;
  overflow: hidden;
  height: clamp(320px, 70vh, 550px);
  border-radius: 15px;
  background-color: #fff;
  box-shadow: 0 8px 24px rgba(0,0,0,0.1);
  transition: box-shadow 0.3s ease;
}

.view-info-block:hover {
  box-shadow: 0 12px 36px rgba(0,0,0,0.15);
}

.text-side {
  flex-basis: 40%;
  padding: 2em 2.5em;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center; 
  color: #2c2c2c;
  font-size: 1.1rem;
  line-height: 1.6;
  order: 1;
  text-align: center; 
}

.right-to-left .text-side {
  order: 2;
  
  align-items: center !important;
  text-align: center !important;
}

.text-side h2 {
  width: 100%;
  margin-bottom: 1rem;
  font-weight: 700;
  color: var(--color-primary, #4a7c59);

  display: inline-block;
}

.image-side {
  flex-basis: 55%;
  padding: 1.5em 2em;
  order: 2;
  opacity: 0;
  transform: translateX(50%) scale(0.95);
  transition:
    opacity 0.6s cubic-bezier(.2,.8,.2,1),
    transform 0.7s cubic-bezier(.2,.8,.2,1);
  display: flex;
  justify-content: center;
  align-items: center;
}

.right-to-left .image-side {
  transform: translateX(-50%) scale(0.95);
}

.animated .image-side {
  opacity: 1;
  transform: translateX(0) scale(1);
}

.image-side img {
  width: 100%;
  max-height: 100%;
  object-fit: cover;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0,0,0,0.1);
  user-select: none;
  pointer-events: none;
  transition: transform 0.3s ease;
}


@media (max-width: 900px) {
  .view-info-block {
    flex-direction: column;
    height: auto;
  }
  .text-side, .right-to-left .text-side {
    flex-basis: 100%;
    padding: 1.5em 1.5em 0.5em;
    align-items: center !important;
    text-align: center !important;
  }
  .text-side h2 {
    display: block;
    width: 100%;
    text-align: center !important;
  }
  .image-side {
    flex-basis: 100%;
    padding: 0 1.5em 1.5em;
    order: 3 !important;
    transform: translateX(0) scale(1) !important;
    opacity: 1 !important;
  }
  .image-side img {
    width: 100%;
    height: auto; 
    max-height: none; 
    object-fit: contain; 
  }
}


</style>
