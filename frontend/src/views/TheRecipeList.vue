<template>
  <div class="recipes feature">
    <recipe-thumbnail v-if="featuredRecipe" :recipe="featuredRecipe" :featured="true" :add-button-enabled="true" />
  </div>
  <div class="recipes">
    <recipe-thumbnail v-for="recipe in recipes" :recipe="recipe" :key="recipe.id" :add-button-enabled="addButtonEnabled"  />
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
const recipes: Ref<IRecipe[] | null> = ref([]);
const featuredRecipe: Ref<IRecipe | null> = ref(null);
const addButtonEnabled = ref(false);
onMounted(async () => {
  const q = route.query["q"]?.toString();
  const src = route.query["source"]?.toString();

  if (q) {
    switch (src) {
      case "themealdb":
        recipes.value = await themealdb.search(q);
        addButtonEnabled.value = true;
        break;
      case "local":
      default:
        recipes.value = await repository.search(q);
        break;
    }
  } else {
    recipes.value = await repository.recipes();
    featuredRecipe.value = await themealdb.getRandomRecipe();
  }
});
</script>

<style scoped lang="scss">
.recipes {
  display: flex;
  width: 100vw;
  max-width: 100vw;
  flex-wrap: wrap;
  justify-content: center;
}
</style>