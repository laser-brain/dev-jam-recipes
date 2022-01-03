<template>
<div class="recipes">
  <div v-for="recipe in recipes" class="recipe-preview">
    <img :src="recipe.thumbnail" alt="recipe" />
    <h2>{{recipe.name}}</h2>
  </div>
</div>
</template>

<script setup lang="ts">
import { Ref, ref } from "vue";
interface Recipe {
  id: number,
  name: string,
  thumbnail: string
}
const recipes : Ref<Recipe[]> = ref([]);
for (let index = 1; index <= 20; index++) {
  recipes.value.push({
    id: index,
    name: `recipe ${index}`,
    thumbnail: "https://www.themealdb.com/images/logo-small.png"
  })  
}
</script>

<style lang="scss">
.recipes{
  display: flex;
  width: 100vw;
  flex-wrap: wrap;
  justify-content: center;
}

  .recipe-preview {
    margin: 1em;
    border: 1px solid black;
    border-radius: 5px;
  }
</style>