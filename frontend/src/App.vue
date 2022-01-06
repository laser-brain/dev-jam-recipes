<script setup lang="ts">
import TheHeader from "./components/TheHeader.vue"
</script>

<template>
  <div class="paper">
    <the-header />
    <router-view :key="$route.fullPath" />
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
}

#app {
  font-family: "Noto Serif", serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background-color: #e6ddd2;
  min-height: 100vh;
}

.paper {
  background-image: url("/src/assets/beige-paper.png");
}
</style>
