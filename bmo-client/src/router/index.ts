import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import CreateSet from '../views/CreateSet.vue';
import BMO from '../views/BMO.vue';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'BMO',
    component: BMO,
  },
  {
    path: '/create',
    name: 'Create Set',
    component: CreateSet,
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
