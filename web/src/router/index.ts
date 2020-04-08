import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import { Routes } from '@/router/names';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    component: () => import('@/layouts/Layout.vue'),
    redirect: '/balance',
    children: [
      {
        path: '/balance',
        name: Routes.Balance,
        component: () => import('@/views/Balance.vue'),
      },
      {
        path: '/point',
        name: Routes.Point,
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "point" */ '@/modules/point/Point.vue'),
      },
    ],
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
