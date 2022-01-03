<template>
  <div class="recipes">
    <router-link
      v-for="recipe in recipes"
      :to="`/recipe/${recipe.id}`"
      custom
      v-slot="{ navigate }"
    >
      <div class="recipe-preview" @click="navigate">
        <img src="/src/assets/logo.png" alt="recipe" />
        <div class="recipe-data">
          <h2>{{ recipe.name }}</h2>
          <span>{{ recipe.difficulty }}</span>
          <br />
          <span>{{ recipe.type }}</span>
          <br />
          <span>Servings: {{ recipe.servings }}</span>
        </div>
      </div>
    </router-link>
  </div>
</template>

<script setup lang="ts">
import { IRecipe } from "../types/viewModels";
import { Ref, ref, onMounted } from "vue";
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