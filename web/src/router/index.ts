import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import { Routes } from "@/router/names";
import BalanceDashboard from "@/modules/balance/views/Dashboard.vue";
import BalanceOverview from "@/modules/balance/views/BalanceOverview.vue";
import Layout from "@/layouts/Layout.vue";
import Point from "@/modules/point/views/Point.vue";
import Issues from "@/modules/point/views/Issues.vue";
import BalanceTransactions from "@/modules/balance/views/BalanceTransactions.vue";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: "/",
    component: Layout,
    redirect: Routes.BalanceDashboard,
    children: [
      {
        path: Routes.BalanceDashboard,
        name: Routes.BalanceDashboard,
        component: BalanceDashboard,
        meta: {
          root: true
        }
      },
      {
        name: Routes.BalanceOverview,
        path: Routes.BalanceOverview,
        component: BalanceOverview
      },
      {
        name: Routes.BalanceTransactions,
        path: Routes.BalanceTransactions,
        component: BalanceTransactions
      },
      {
        path: Routes.Point,
        name: Routes.Point,
        component: Point,
        meta: {
          root: true
        }
      },
      {
        path: Routes.PointIssues,
        name: Routes.PointIssues,
        component: Issues
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
