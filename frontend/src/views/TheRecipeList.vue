<template>
  <div>
    <div class="recipes feature">
      <recipe-thumbnail
        v-if="featuredRecipe"
        :recipe="featuredRecipe"
        :featured="true"
        :add-button-enabled="true"
        @add-recipe="updateRecipesList"
      />
    </div>
    <div class="recipes">
      <recipe-thumbnail
        v-for="recipe in recipesOrdered"
        :recipe="recipe"
        :key="recipe.id"
        :add-button-enabled="addButtonEnabled"
      />
    </div>
    <div class="no-content" v-if="mounted && recipes?.length === 0">
      <span>Sorry, we could not find any recipes for the search term "{{ q }}" :(</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Ref, ref, onMounted, computed } from "vue";
import RecipeThumbnail from "../components/RecipeThumbnail.vue"
import { IRecipe } from "../types/viewModels";
import * as repository from '../services/recipes-repository'
import * as themealdb from '../services/themealdb-repository'
import { useRoute } from 'vue-router';

const route = useRoute();
const recipes: Ref<IRecipe[] | null> = ref([]);
const featuredRecipe: Ref<IRecipe | null> = ref(null);
const addButtonEnabled = ref(false);
const mounted = ref(false);
const q = ref('');

onMounted(async () => {
  q.value = route.query["q"]?.toString() || '';
  const src = route.query["source"]?.toString();

  if (q.value || src) {
    switch (src) {
      case "themealdb":
        recipes.value = await themealdb.search(q.value);
        addButtonEnabled.value = true;
        break;
      case "local":
      default:
        recipes.value = await repository.search(q.value);
        break;
    }
  } else {
    recipes.value = await repository.recipes();
    featuredRecipe.value = await themealdb.getRandomRecipe();
  }
  mounted.value = true;
});

function updateRecipesList() {
  if (featuredRecipe && featuredRecipe.value) {
    if (recipes.value?.filter(r => r.id === featuredRecipe.value?.id).length === 0) {
      recipes.value.push(featuredRecipe.value);
    }
  }
}

const recipesOrdered = computed(() => {
  if (recipes && recipes.value) {
    return recipes.value.sort((a, b) => a.name > b.name ? 1 : 0);
  }
  return [];
})
</script>

<style scoped lang="scss">
.recipes {
  display: flex;
  width: 70vw;
  margin-left: 15vw;
  max-width: 70vw;
  flex-wrap: wrap;
  justify-content: center;
}

.no-content {
  margin: 1rem;
  width: 100vw;
  font-size: 3em;
  display: flex;
  justify-content: center;
}
</style>