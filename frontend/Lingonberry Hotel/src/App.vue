<template>
  <div id="app" :style="{ paddingTop: headerHeight + 'px' }">
    <div>
      <Header @header-height-change="updateHeaderHeight"></Header>
      <router-view />
    </div>
  </div>
</template>


<script>
import Header from '@/components/Header.vue';

export default {
  name: 'App',
  components: {
    Header,
  },
  data() {
    return {
      headerHeight: 120, 
    };
  },
  methods: {
    updateHeaderHeight(height) {
      this.headerHeight = height;
    },
  },
};
</script>

<style scoped>
nav a {
  margin: 0 10px;
  color: #6b4c3b;
  text-decoration: none;
  font-weight: 600;
}

nav a.router-link-exact-active {
  text-decoration: underline;
}
</style>
