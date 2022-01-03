<template>
  <router-link :to="'/'">Overview</router-link>
  <div class="recipe" v-if="recipe">
    <h1>{{ recipe.name }}</h1>
    <p>{{ recipe.description }}</p>
  </div>
  <div v-else>Invalid recipe id</div>
</template>

<script setup lang="ts">import { PropType } from 'vue';
import { IRecipe } from '../types/viewModels';
import * as repository from '../services/recipes-repository'
import { useRoute } from 'vue-router';
const route = useRoute();

if (typeof route.params.recipeId !== "string") {
  throw new Error(
    `Invalid route param ${JSON.stringify(route.params.recipeId)}`
  );
}

const recipe = repository.recipe(parseInt(route.params.recipeId));
</script>

<style scoped lang="scss">
.recipe {
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>