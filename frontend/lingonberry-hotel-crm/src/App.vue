<template>
  <div class="main-content">
    <router-view class="page-card" />
  </div>
  <AppNav />
</template>

<script setup>
import AppNav from '@/components/AppNav.vue'
</script>

<style>
@import '@/assets/styles.css';
</style>
