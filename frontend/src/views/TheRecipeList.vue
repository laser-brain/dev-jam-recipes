<template>
  <div class="recipes">
    <router-link
      v-for="recipe in recipes"
      :to="`/recipe/${recipe.id}`"
      custom
      v-slot="{ navigate }"
    >
      <RecipeThumbnail :recipe="recipe" @thumbnail-clicked="navigate" />
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { Ref, ref, onMounted } from "vue";
import RecipeThumbnail from "../components/RecipeThumbnail.vue"
import { IRecipe } from "../types/viewModels";
import * as repository from '../services/recipes-repository'
import * as themealdb from '../services/themealdb-repository'
import { useRoute } from 'vue-router';

const route = useRoute();
const recipes: Ref<IRecipe[] | null> = ref([])

onMounted(async () => {
  const q = route.query["q"]?.toString()
  const src = route.query["source"]?.toString();
  if (q) {
    switch (src) {
      case "themealdb":
        recipes.value = await themealdb.search(q)
        break;
      case "local":
      default:
        recipes.value = await repository.search(q)
        break;
    }
  } else {
    recipes.value = await repository.recipes()
  }
});
</script>

<style scoped lang="scss">
.recipes {
  display: flex;
  width: 100vw;
  flex-wrap: wrap;
  justify-content: center;
  padding: 0.5em;
}
</style>