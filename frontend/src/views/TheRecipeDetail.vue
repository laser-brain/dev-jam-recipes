<template>
  <div class="flex-container">
    <img v-if="recipe" :src="recipe.thumbnail" />
    <div class="recipe" v-if="recipe">
      <h1>{{ recipe.name }}</h1>
      <h2>Ingredients</h2>
      <ul>
        <li v-for="ingredient in recipe.ingredients" :key="ingredient.name">
          <span>{{ ingredient.name }}: {{ ingredient.amount }}</span>
        </li>
      </ul>
      <h2>Steps</h2>
      <p v-html="recipe.description.replaceAll('\r\n', '<br />')"></p>
    </div>
    <div v-else>Invalid recipe id</div>
  </div>
</template>

<script setup lang="ts">
import { Ref, ref, onMounted } from 'vue';
import { IRecipe } from '../types/viewModels';
import * as repository from '../services/recipes-repository'
import * as themealdb from '../services/themealdb-repository'
import { useRoute } from 'vue-router';

const route = useRoute();

if (typeof route.params.recipeId !== "string") {
  throw new Error(
    `Invalid route param ${JSON.stringify(route.params.recipeId)}`
  );
}
const recipeId = parseInt(route.params.recipeId);
const recipeSource = route.query["source"];

onMounted(async () => {
  if (recipeSource && recipeSource === "external") {
    recipe.value = await themealdb.recipe(recipeId);
  }
  else {
    recipe.value = await repository.recipe(recipeId);
  }
});
const recipe: Ref<IRecipe | null> = ref(null);

</script>

<style scoped lang="scss">

.flex-container {
  display: flex;
}
img {
  width: 20vw;
  height: 20vw;
  margin: 2.5vw;
}
.recipe {
  margin-top: 2.5vw;
  margin-right: 25vw;
  display: flex;
  flex-direction: column;
  align-items: left;
  background-color: white;

  > h2 {
    margin-top: 1.5rem;
  }

  > ul {
    padding-left: 2rem;
  }
}
</style>