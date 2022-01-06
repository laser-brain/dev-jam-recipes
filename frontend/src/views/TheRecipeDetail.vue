<template>
  <div class="flex-container">
    <img v-if="recipe" :src="recipe.thumbnail" />
    <div class="recipe" v-if="recipe">
      <h1>{{ recipe.name }}</h1>
      <div class="meta-data">
        <div class="inline-block">
          <h2>Recipe info</h2>
        </div>
        <font-awesome-icon icon="exclamation-circle" />&nbsp;
        <label for="difficulty">Difficulty:&nbsp;</label>
        <span>{{ recipe.difficulty }}</span>
        <br />
        <font-awesome-icon icon="drumstick-bite" />&nbsp;
        <label for="type">Type:&nbsp;</label>
        <span>{{ recipe.type }}</span>
        <br />
        <font-awesome-icon icon="hashtag" />&nbsp;
        <label for="servings">Servings:&nbsp;</label>
        <span>{{ recipe.servings }}</span>
      </div>
      <h2>Ingredients</h2>
      <ul>
        <li v-for="ingredient in recipe.ingredients" :key="ingredient.name">
          <span>{{ ingredient.name }}: {{ ingredient.amount }}</span>
        </li>
      </ul>
      <h2>Steps</h2>
      <p v-html="recipe.description.replaceAll('\r\n', '<br />')"></p>
    </div>
    <div class="error" v-else-if="mounted">
      <span>Oops, something went wrong!</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Ref, ref, onMounted } from 'vue';
import { IRecipe } from '../types/viewModels';
import * as repository from '../services/recipes-repository'
import * as themealdb from '../services/themealdb-repository'
import { useRoute } from 'vue-router';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faDrumstickBite, faExclamationCircle, faHashtag, faEdit } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faDrumstickBite, faExclamationCircle, faHashtag, faEdit)

const route = useRoute();
const mounted = ref(false);
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

  mounted.value = true;
});
const recipe: Ref<IRecipe | null> = ref(null);

</script>

<style scoped lang="scss">
.inline-block {
  display: inline-block;
  width: 75vw;
}
.fa:before {
  display: inline-block;
}
.flex-container {
  display: flex;
}

.error {
  width: 100vw;
  font-size: 3em;
  display: flex;
  justify-content: center;
}

img {
  width: 20vw;
  height: 20vw;
  margin: 2.5vw;
}
h2 {
  margin-top: 1.5rem;
}

.recipe {
  margin-top: 2.5vw;
  margin-right: 25vw;
  display: flex;
  flex-direction: column;
  align-items: left;

  > ul {
    padding-left: 2rem;
  }
}
</style>