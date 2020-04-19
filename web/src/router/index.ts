import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import { Routes } from "@/router/names";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    component: () =>
      import(/* webpackChunkName: "layout" */ "@/layouts/Layout.vue"),
    redirect: Routes.Balance,
    children: [
      {
        path: Routes.Balance,
        name: Routes.Balance,
        component: () =>
          import(
            /* webpackChunkName: "balance" */ "@/modules/balance/views/Balance.vue"
          ),
        meta: {
          root: true
        }
      },
      {
        path: Routes.Point,
        name: Routes.Point,
        component: () =>
          import(/* webpackChunkName: "point" */ "@/modules/point/Point.vue"),
        meta: {
          root: true
        }
      },
      {
        path: Routes.PointIssues,
        name: Routes.PointIssues,
        component: () =>
          import(/* webpackChunkName: "issues" */ "@/modules/point/Issues.vue")
      }
    ]
  }
];

routes.push({
  path: "*",
  redirect: "/"
});

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
