import { createWebHistory, createRouter } from 'vue-router';
import TheRecipeDetail from '@/views/TheRecipeDetail.vue';
import TheRecipeList from '@/views/TheRecipeList.vue';

const routes = [
  {
    path: '/',
    name: 'recipe-list',
    component: TheRecipeList,
  },
  {
    path: '/recipe/:recipeId',
    name: 'recipe-detail',
    component: TheRecipeDetail,
  },
  {
    path: '/search',
    name: 'search',
    component: TheRecipeList,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
