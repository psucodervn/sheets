import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import { Routes } from '@/router/names';
import BalanceDashboard from '@/modules/balance/views/BalanceDashboard.vue';
import BalanceAccounts from '@/modules/balance/views/BalanceAccounts.vue';
import Layout from '@/layouts/Layout.vue';
import Point from '@/modules/point/views/Point.vue';
import Issues from '@/modules/point/views/Issues.vue';
import BalanceTransactions from '@/modules/balance/views/transactions/BalanceTransactions.vue';
import Report from '@/modules/point/views/Report.vue';
import Balance from '@/modules/balance/views/Balance.vue';
import TransactionsNew from '@/modules/balance/views/transactions/TransactionsNew.vue';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    component: Layout,
    redirect: Routes.BalanceDashboard,
    children: [
      {
        path: Routes.BalanceDashboard,
        name: Routes.BalanceDashboard,
        component: Balance,
        meta: {
          root: true,
        },
        children: [
          {
            name: Routes.BalanceDashboard,
            path: Routes.BalanceDashboard,
            component: BalanceDashboard,
          },
          {
            name: Routes.BalanceOverview,
            path: Routes.BalanceOverview,
            component: BalanceAccounts,
          },
          {
            name: Routes.BalanceTransactions,
            path: Routes.BalanceTransactions,
            component: BalanceTransactions,
          },
          {
            name: Routes.BalanceTransactionsNew,
            path: Routes.BalanceTransactionsNew,
            component: TransactionsNew,
          },
        ],
      },
      {
        path: Routes.Point,
        name: Routes.Point,
        component: Point,
        meta: {
          root: true,
        },
      },
      {
        path: Routes.PointIssues,
        name: Routes.PointIssues,
        component: Issues,
      },
      {
        path: Routes.Report,
        name: Routes.Report,
        component: Report,
        meta: {
          root: true,
        },
      },
    ],
  },
];

routes.push({
  path: '*',
  redirect: '/',
});

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

router.afterEach((to, from) => {
  // console.log('set from', from);
  // Vue.$navigation.from = from;
});

export default router;
