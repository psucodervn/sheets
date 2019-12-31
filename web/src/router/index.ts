import Vue from 'vue';
import VueRouter from 'vue-router';
import { Routes } from '@/router/names';
import Overview from '@/views/Overview.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: Routes.Overview,
    component: Overview,
  },
  {
    path: '/transactions',
    name: Routes.Transactions,
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "transactions" */ '@/views/Transactions.vue'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
