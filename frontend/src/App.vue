<script setup lang="ts">
import TheHeader from "./components/TheHeader.vue"
</script>

<template>
  <div>
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
  background-color: white;
  min-height: 100vh;
}
</style>
