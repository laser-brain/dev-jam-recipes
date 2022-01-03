<template>
  <div class="recipes">
    <router-link
      v-for="recipe in recipes"
      :to="`/recipe/${recipe.id}`"
      custom
      v-slot="{ navigate }"
    >
      <RecipeThumbnail :recipe="recipe" @click="navigate" />
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { Ref, ref, onMounted } from "vue";
import RecipeThumbnail from "../components/RecipeThumbnail.vue"
import { IRecipe } from "../types/viewModels";
import * as repository from '../services/recipes-repository'

onMounted(async () => {
  recipes.value = await repository.recipes();
});

const recipes: Ref<IRecipe[] | null> = ref([])
</script>

<style lang="scss">
.recipes {
  display: flex;
  width: 100vw;
  flex-wrap: wrap;
  justify-content: center;
}

.recipe-preview {
  display: flex;
  margin: 1em;
  padding: 0.5em;
  width: 30vw;
  border: 1px solid black;
  border-radius: 5px;

  > * {
    padding: 1.5rem;
  }
  img {
    padding: 1rem;
    height: 10rem;
    width: 10rem;
  }
}

.recipe-data {
  flex-direction: column;
}
</style>