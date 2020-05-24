import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import { Routes } from '@/router/names';
import Point from '@/modules/point/views/Point.vue';
import Issues from '@/modules/point/views/Issues.vue';
import Report from '@/modules/point/views/Report.vue';
import Profile from '@/modules/profile/views/Profile.vue';
import { ProfileModule } from '@/store';
import Balance from '@/modules/balance/views/Balance.vue';
import BalanceDashboard from '@/modules/balance/views/BalanceDashboard.vue';
import BalanceAccounts from '@/modules/balance/views/BalanceAccounts.vue';
import BalanceTransactions from '@/modules/balance/views/transactions/BalanceTransactions.vue';
import TransactionsNew from '@/modules/balance/views/transactions/TransactionsNew.vue';
import TransactionsEdit from '@/modules/balance/views/transactions/TransactionsEdit.vue';
import Login from '@/modules/profile/views/Login.vue';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    redirect: Routes.BalanceDashboard,
  },
  {
    path: Routes.BalanceDashboard,
    name: Routes.BalanceDashboard,
    component: Balance,
    children: [
      {
        name: Routes.BalanceDashboard,
        path: Routes.BalanceDashboard,
        component: BalanceDashboard,
        meta: {
          root: true,
        },
      },
      {
        name: Routes.BalanceAccounts,
        path: Routes.BalanceAccounts,
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
        meta: {
          requiresAuth: true,
        },
      },
      {
        name: Routes.BalanceTransactionsEdit,
        path: Routes.BalanceTransactionsEdit,
        component: TransactionsEdit,
        meta: {
          requiresAuth: true,
        },
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
      {
        path: Routes.Profile,
        name: Routes.Profile,
        component: Profile,
        meta: {
          root: true,
          requiresAuth: true,
        },
      },
    ],
  },
  {
    path: Routes.Login,
    name: Routes.Login,
    component: Login,
    meta: {
      requiresNotAuth: true,
      layout: 'simple-layout',
    },
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

router.beforeEach(async (to, from, next) => {
  if (to.matched.some(value => value.meta.requiresNotAuth)) {
    if (ProfileModule.currentUser) {
      return next(false);
    }
    return next();
  }
  if (to.matched.some(value => value.meta.requiresAuth)) {
    if (!ProfileModule.currentUser) {
      return next({
        name: Routes.Login,
        query: {
          redirect: to.fullPath,
        },
      });
    }
  }
  return next();
});

export default router;
