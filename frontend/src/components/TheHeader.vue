<template>
  <header>
    <router-link to="/" custom v-slot="{ navigate }">
      <button @click="navigate">
        <font-awesome-icon id="home" icon="utensils" />
      </button>
    </router-link>
    <div id="search-container">
      <input
        id="search-query"
        placeholder="Search for recipes"
        @keyup="checkSubmit"
        v-model="query"
      />
      <select v-model="source" id="search-source">
        <option
          v-for="option in options"
          :value="option.value"
          :key="option.value"
        >{{ option.name }}</option>
      </select>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import router from '../router'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faUtensils } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faUtensils)

const query = ref('');
const source = ref('themealdb');

const options = [
  { value: "themealdb", name: "themealdb.com" },
  { value: "local", name: "Database" },
];

const checkSubmit = function (payload: KeyboardEvent) {
  if (payload.key === 'Enter') {
    router.push({ name: 'search', query: { q: query.value, source: source.value || 'local' } });
  }
}
</script>

<style scoped lang="scss">
header {
  display: flex;
  align-items: center;
  * {
    font-size: 1.5rem;
  }
  @media screen and (orientation: landscape) {
    > * {
      margin-right: 2rem;
    }
  }

  @media screen and (orientation: portrait) {
    flex-direction: column;
    align-items: flex-start;
  }
}

#search-container {
  width: calc(100vw - 64px);
  @media screen and (orientation: portrait) {
    width: 100vw;
    > * {
      margin: 0.25rem;
    }
  }
}

#search-query {
  width: 30vw;
  text-overflow: hidden;
  margin-left: calc(25vw - 64px);
  border-radius: 5px;
  padding: 0.1em;
  border: 1px solid black;
  margin-right: 0.5em;
  @media screen and (orientation: portrait) {
    margin-left: 0.5em;
    width: 90vw;
  }
}

#search-source {
  padding: 0.1em;
  margin-left: 0.5em;
  width: 320px;
  background-color: white;
  border: 1px solid black;
  border-radius: 5px;
  @media screen and (orientation: portrait) {
    width: 90vw;
  }
}

button,
#home {
  height: 64px;
  width: 64px;

  @media screen and (orientation: portrait) {
    height: 32px;
    width: 32px;
  }
}

button {
  margin-top: 0.5em;
  margin-left: 0.5em;
  background-color: Transparent;
  border: none;
}
</style>