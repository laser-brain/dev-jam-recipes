<template>
  <div class="recipe-preview" @click="click">
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
  }
});

const emit = defineEmits(['thumbnailClicked']);
function click() {
  emit('thumbnailClicked');
}
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

  > * {
    padding: 1.5rem;
  }

  img {
    padding: 1rem;
    height: 10rem;
    width: 10rem;
  }

  &:hover {
    transform: scale(1.05);
    transition: all .2s
  }
}
</style>