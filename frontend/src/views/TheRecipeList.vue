<template>
  <div class="recipes">
    <router-link
      v-for="recipe in recipes"
      :to="`/recipe/${recipe.id}`"
      class="recipe-preview"
      custom
      v-slot="{ navigate }"
    >
      <div @click="navigate">
        <img :src="recipe.thumbnail" alt="recipe" />
        <h2>{{ recipe.name }}</h2>
      </div>
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { Ref, ref } from "vue";
import { IRecipe } from "../types/viewModels";
import * as repository from '../services/recipes-repository' 

const recipes = ref(repository.recipes());
</script>

<style lang="scss">
.recipes {
  display: flex;
  width: 100vw;
  flex-wrap: wrap;
  justify-content: center;
}

.recipe-preview {
  margin: 1em;
  border: 1px solid black;
  border-radius: 5px;

  img {
    height: 2.5rem;
  }
}
</style>