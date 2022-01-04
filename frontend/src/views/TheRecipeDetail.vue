<template>
  <router-link :to="'/'">Overview</router-link>
  <div class="recipe" v-if="recipe">
    <h1>{{ recipe.name }}</h1>
    <h2>Ingredients</h2>
    <ul>
      <li v-for="ingredient in recipe.ingredients">
        <span>{{ ingredient.name }}: {{ ingredient.amount }}</span>
      </li>
    </ul>
    <h2>Steps</h2>
    <p>{{ recipe.description }}</p>
  </div>
  <div v-else>Invalid recipe id</div>
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
  if(recipeSource && recipeSource === "external") {
    recipe.value = await themealdb.recipe(recipeId);
  }
  else {
    recipe.value = await repository.recipe(recipeId);
  }
});
const recipe: Ref<IRecipe | null> = ref(null);

</script>

<style scoped lang="scss">
.recipe {
  margin-left: 25vw;
  margin-right: 25vw;
  display: flex;
  flex-direction: column;
  align-items: left;

  > h2 {
    margin-top: 1.5rem;
  }

  > ul {
    padding-left: 2rem;
  }
}
</style>