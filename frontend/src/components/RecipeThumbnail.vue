<template>
  <router-link
    :to="props.featured ? `/recipe/${props.recipe.id}?source=external` : `/recipe/${props.recipe.id}`"
    custom
    v-slot="{ navigate }"
    :key="props.recipe.id"
  >
    <div :class="props.featured ? 'recipe-preview featured' : 'recipe-preview'" @click="navigate">
      <h1 v-if="props.featured">Featured recipe</h1>
      <h2 v-if="props.featured">{{ props.recipe.name }}</h2>
      <img :src="recipe.thumbnail" alt="recipe" />
      <div class="recipe-data">
        <h2 v-if="!props.featured">{{ props.recipe.name }}</h2>
        <font-awesome-icon icon="exclamation-circle" />&nbsp;
        <span>Difficulty: {{ recipe.difficulty }}</span>
        <br />
        <font-awesome-icon icon="drumstick-bite" />&nbsp;
        <span>Type: {{ recipe.type }}</span>
        <br />
        <font-awesome-icon icon="hashtag" />&nbsp;
        <span>Servings: {{ recipe.servings > 0 ? recipe.servings : 'unknown' }}</span>
      </div>
      <button v-if="props.addButtonEnabled" class="glued" @click.stop="addToLocalDatabase">+</button>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { PropType } from "vue"
import { IRecipe } from '../types/viewModels';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faDrumstickBite, faExclamationCircle, faHashtag } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import * as repository from '../services/recipes-repository'

library.add(faDrumstickBite, faExclamationCircle, faHashtag)

const props = defineProps({
  recipe: {
    type: Object as PropType<IRecipe>,
    required: true
  },
  featured: {
    type: Boolean,
    required: false
  },
  addButtonEnabled: {
    type: Boolean,
    required: true
  }
});

function addToLocalDatabase() {
  repository.addRecipe(props.recipe);
}

</script>

<style scoped lang="scss">
.glued {
  position: absolute;
  bottom: 10px;
  right: 10px;
  width: 50px;
  height: 50px;
  font-size: 2rem;
  font-weight: bold;
  background-color: transparent;
  z-index: 3;

  &:hover {
    transform: scale(1.2);
    transition: all 0.5s;
  }
}

.recipe-preview {
  position: relative;
  display: flex;
  margin: 1em;
  padding: 0.5em;
  width: 30vw;
  border: 1px solid black;
  border-radius: 5px;

  @media screen and (orientation: portrait) {
    width: 90vw;
    min-height: 30vh;
    flex-direction: column;
    align-items: center;
    margin: 1rem 0 1rem 0;
  }

  h2 {
    font-family: "Parisienne", cursive;
    font-size: 2.5em;
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
    min-height: 50vh;
    flex-direction: column;
    align-items: center;

    img {
      width: 25vw;
      height: 25vw;
      @media screen and (orientation: portrait) {
        width: 25vh;
        height: 25vh;
      }
    }

    @media screen and (orientation: landscape) {
      .recipe-data {
        text-align: left;
        width: 100%;
        > svg {
          margin-left: 3rem;
        }
      }
    }
  }
}
</style>