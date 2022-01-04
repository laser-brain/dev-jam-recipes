<template>
  <router-link
    :to="props.featured ? `/recipe/${props.recipe.id}?source=external` : `/recipe/${props.recipe.id}`"
    custom
    v-slot="{ navigate }"
    :key="props.recipe.id"
  >
    <div :class="props.featured ? 'recipe-preview featured' : 'recipe-preview'" @click="navigate">
      <h1 v-if="props.featured">Featured recipe</h1>
      <img :src="recipe.thumbnail" alt="recipe" />
      <div class="recipe-data">
        <h2>{{ props.recipe.name }}</h2>
        <font-awesome-icon icon="exclamation-circle" />&nbsp;
        <span>Difficulty: {{ recipe.difficulty }}</span>
        <br />
        <font-awesome-icon icon="drumstick-bite" />&nbsp;
        <span>Type: {{ recipe.type }}</span>
        <br />
        <font-awesome-icon icon="hashtag" />&nbsp;
        <span>Servings: {{ recipe.servings }}</span>
      </div>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { PropType } from "vue"
import { IRecipe } from '../types/viewModels';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faDrumstickBite, faExclamationCircle, faHashtag } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faDrumstickBite, faExclamationCircle, faHashtag)

const props = defineProps({
  recipe: {
    type: Object as PropType<IRecipe>,
    required: true
  },
  featured: {
    type: Boolean,
    required: false
  }
});
</script>

<style scoped lang="scss">
.recipe-preview {
  display: flex;
  margin: 1em;
  padding: 0.5em;
  width: 30vw;
  border: 1px solid black;
  border-radius: 5px;
  background-color: #0677a1;

  h2 {
    font-family: "Parisienne", cursive;
    font-size: 2.5em;
  }

  > * {
    padding: 0.5rem;
  }

  img {
    height: 10rem;
    width: 10rem;
  }

  &:hover {
    transform: scale(1.05);
    transition: all 0.2s;
  }

  &.featured {
    height: 50vh;
    width: 30vw;
    flex-direction: column;
  }
}
</style>