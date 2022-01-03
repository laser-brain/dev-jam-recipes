<template>
  <div class="recipe-preview" @click="click">
    <img src="/src/assets/logo.png" alt="recipe" />
    <div class="recipe-data">
      <h2>{{ props.recipe.name }}</h2>
      <span>Difficulty: {{ recipe.difficulty }}</span>
      <br />
      <span>Type: {{ recipe.type }}</span>
      <br />
      <span>Servings: {{ recipe.servings }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { PropType } from "vue"
import { IRecipe } from '../types/viewModels';

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