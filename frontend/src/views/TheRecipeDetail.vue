<template>
  <router-link :to="'/'">Overview</router-link>
  <div class="recipe" v-if="recipe">
    <h1>{{ recipe.name }}</h1>
    <p>{{ recipe.description }}</p>
  </div>
  <div v-else>Invalid recipe id</div>
</template>

<script setup lang="ts">
import { Ref, ref, onMounted } from 'vue';
import { IRecipe } from '../types/viewModels';
import * as repository from '../services/recipes-repository'
import { useRoute } from 'vue-router';
const route = useRoute();

if (typeof route.params.recipeId !== "string") {
  throw new Error(
    `Invalid route param ${JSON.stringify(route.params.recipeId)}`
  );
}
const recipeId = parseInt(route.params.recipeId);
onMounted(async() => {
 recipe.value = await repository.recipe(recipeId);
});
const recipe: Ref<IRecipe | null> = ref(null);

</script>

<style scoped lang="scss">
.recipe {
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>