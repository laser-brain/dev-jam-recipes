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
        <span v-if="!editFlags?.difficulty.enabled">{{ recipe.difficulty }}</span>
        <select v-if="editFlags?.difficulty.enabled" v-model="editFlags.difficulty.value">
          <option
            v-for="option in editFlags.difficulty.options"
            :value="option.value"
            :key="option.value"
          >{{ option.name }}</option>
        </select>
        &nbsp;
        <font-awesome-icon
          @click="edit('difficulty')"
          :icon="editFlags?.difficulty.enabled ? 'check' : 'edit'"
        />
        <br />
        <font-awesome-icon icon="drumstick-bite" />&nbsp;
        <label for="type">Type:&nbsp;</label>
        <span v-if="!editFlags?.type.enabled">{{ recipe.type }}</span>
        <select v-if="editFlags?.type.enabled" v-model="editFlags.type.value">
          <option
            v-for="option in editFlags.type.options"
            :value="option.value"
            :key="option.value"
          >{{ option.name }}</option>
        </select>
        &nbsp;
        <font-awesome-icon
          @click="edit('type')"
          :icon="editFlags?.type.enabled ? 'check' : 'edit'"
        />
        <br />
        <font-awesome-icon icon="hashtag" />&nbsp;
        <label for="servings">Servings:&nbsp;</label>
        <span v-if="!editFlags?.servings.enabled">{{ recipe.servings }}</span>
        <input v-if="editFlags?.servings.enabled" v-model="editFlags.servings.value" />
        &nbsp;
        <font-awesome-icon
          @click="edit('servings')"
          :icon="editFlags?.servings.enabled ? 'check' : 'edit'"
        />
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
.value
<script setup lang="ts">
import { Ref, ref, onMounted } from 'vue';
import { IRecipe } from '../types/viewModels';
import * as repository from '../services/recipes-repository'
import * as themealdb from '../services/themealdb-repository'
import { useRoute } from 'vue-router';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faDrumstickBite, faExclamationCircle, faHashtag, faEdit, faCheck } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

type EditFlags = {
  difficulty: EditElement,
  type: EditElement,
  servings: EditElement,
  [key: string]: EditElement;
}

type EditElement = {
  enabled: boolean,
  value: string,
  options?: { value: string, name: string }[]
}

library.add(faDrumstickBite, faExclamationCircle, faHashtag, faEdit, faCheck)

const recipe: Ref<IRecipe | null> = ref(null);
const editFlags: Ref<EditFlags | null> = ref(null);
const debug = ref('');

const edit = function (identifier: string) {
  if (!recipe || !recipe.value || !editFlags || !editFlags.value) {
    return;
  }

  var stateElement = editFlags.value[identifier];

  if (stateElement.enabled === true) {
    switch (identifier) {
      case "difficulty":
        recipe.value.difficulty = stateElement.value as string;
        break;
      case "type":
        recipe.value.type = stateElement.value as string;
        break;
      case "servings":
        recipe.value.servings = parseInt(stateElement.value as string);
        break;
    }
    repository.addRecipe(recipe.value);
  }

  stateElement.enabled = !stateElement.enabled;
}

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
  editFlags.value = {
    difficulty: {
      enabled: false,
      value: recipe.value?.difficulty || 'unknown',
      options: [
        { value: "unknown", name: "unknown" },
        { value: "beginner", name: "beginner" },
        { value: "intermediate", name: "intermediate" },
        { value: "advanced", name: "advanced" }
      ]
    },
    type: {
      enabled: false,
      value: recipe.value?.type || 'unknown',
      options: [
        { value: "unknown", name: "unknown" },
        { value: "breakfast", name: "breakfast" },
        { value: "lunch", name: "lunch" },
        { value: "supper", name: "supper" },
        { value: "snack", name: "snack" }
      ]
    },
    servings: { enabled: false, value: String(recipe.value?.servings || 0) }
  };

  mounted.value = true;
});

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