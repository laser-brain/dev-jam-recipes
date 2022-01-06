<template>
  <router-link
    :to="props.addButtonEnabled ? `/recipe/${props.recipe.id}?source=external` : `/recipe/${props.recipe.id}`"
    custom
    v-slot="{ navigate }"
    :key="props.recipe.id"
  >
    <div
      :class="props.featured ? 'recipe-preview paper featured' : 'recipe-preview paper'"
      @click="navigate"
    >
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
      <button v-if="addButtonRef && !loading" class="glued" @click.stop="addToLocalDatabase">+</button>
      <font-awesome-icon :class="!successRef ? 'glued hidden' : 'glued'" icon="check-circle" />
      <font-awesome-icon v-if="loading" class="glued" icon="spinner" />
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { ref, PropType } from "vue"
import * as repository from '../services/recipes-repository'
import { IRecipe } from '../types/viewModels';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faDrumstickBite, faExclamationCircle, faHashtag, faCheckCircle, faSpinner } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faDrumstickBite, faExclamationCircle, faHashtag, faCheckCircle, faSpinner);

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

const addButtonRef = ref(props.addButtonEnabled);
const loading = ref(false);
const successRef = ref(false);

async function addToLocalDatabase() {
  loading.value = true;
  await repository.addRecipe(props.recipe);
  loading.value = false;
  addButtonRef.value = false;
  successRef.value = true;
  setTimeout(() => {
    successRef.value = false;
  }, 500)

  emit('addRecipe');
}

const emit = defineEmits(['addRecipe']);

</script>

<style scoped lang="scss">
.hidden {
  visibility: hidden;
  opacity: 0;
  transition: visibility 0s 2s, opacity 2s linear;
}

.fa-check-circle {
  color: green;
}

.fa-spinner {
  color: teal;
  animation: spin 1.2s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

button {
  border-radius: 30px;
  border: none;
  box-shadow: 4px -4px 6px;
}

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
  width: 15vw;
  min-height: 50vh;
  border: 1px solid black;
  border-radius: 5px;
  flex-direction: column;
  align-items: center;
  box-shadow: 3px -3px 5px darkgray;
  background-color: white;

  @media screen and (orientation: portrait) {
    width: 90vw;
    min-height: 30vh;
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
    width: 30vw;
    min-height: 50vh;
    flex-direction: column;
    align-items: center;

    @media screen and (orientation: portrait) {
      width: 100%;
    }

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