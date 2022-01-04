<template>
  <header>
    <a href="/">
      <img id="img-home" src="/src/assets/writing-pad.jpg" alt="recipe-writing-pad" />
    </a>
    <div id="search-container">
      <input
        id="search-query"
        placeholder="Search for recipes"
        @keyup="checkSubmit"
        v-model="query"
      />
      <select v-model="source" id="search-source">
        <option v-for="option in options" v-bind:value="option.id">{{ option.name }}</option>
      </select>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import router from '../router'

const query = ref('');
const source = ref('themealdb');

const options = [
  { id: "local", name: "Database" },
  { id: "themealdb", name: "themealdb.com" },
];

const checkSubmit = function (payload: KeyboardEvent) {
  if (payload.key === 'Enter') {
    router.push({ name: 'search', query: { q: query.value, source: source.value} });
  }
}
</script>

<style scoped lang="scss">
header {
  padding: 0.5em;
  padding-left: calc(3em - 3px);
  display: flex;
  align-items: center;

  > * {
    margin-right: 2em;
  }
}

#search-container {
  width: calc(100vw - 128px);
}

#search-query {
  width: calc(50vw - 128px);
  text-overflow: hidden;
  font-size: 2em;
  margin-left: calc(25vw - 250px - 2em - 1em);
  border-radius: 5px;
  padding: .1em;
  border: none;
  margin-right: .5em;
}
#search-source {
  padding: .1em;
  margin-left: .5em;
  width: 320px;
  font-size: 2em;
  background-color: #0677a1;
  border: none;
  border-radius: 5px;
}

#img-home {
  width: 128px;
}
</style>