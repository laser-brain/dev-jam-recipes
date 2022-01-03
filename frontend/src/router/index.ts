import { createWebHistory, createRouter } from 'vue-router';
import TheRecipeDetail from '@/views/TheRecipeDetail.vue';
import TheRecipeList from '@/views/TheRecipeList.vue';

const routes = [
  {
    path: '/',
    name: 'RecipeList',
    component: TheRecipeList,
  },
  {
    path: '/recipe/:recipeId',
    name: 'Recipe',
    component: TheRecipeDetail,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
