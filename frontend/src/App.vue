<script setup lang="ts">
import TheHeader from "./components/TheHeader.vue"
</script>

<template>
  <the-header />
  <router-view :key="$route.fullPath" />
</template>

<style>
* {
  margin: 0;
  padding: 0;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background-color: #2d4159;
  min-height: 100vh;
}
</style>
