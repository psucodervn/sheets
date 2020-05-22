<template>
  <q-layout view="hHh lpR fFf">
    <q-page-container class="page q-pa-sm">
      <navigation-bar />
      <router-view />
    </q-page-container>
    <component :is="tabPosition" class="shadow-2 bg-dark">
      <q-tabs dense>
        <q-route-tab :icon="t.icon" :to="t.to" v-for="t in tabs">
          <q-tooltip>{{ t.label }}</q-tooltip>
        </q-route-tab>
      </q-tabs>
    </component>
  </q-layout>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Routes } from '@/router/names';

@Component({})
export default class DefaultLayout extends Vue {
  tabs = [
    { icon: 'local_atm', label: 'Balance', to: Routes.BalanceDashboard },
    { icon: 'laptop_mac', label: 'Point', to: Routes.Point },
    { icon: 'bar_chart', label: 'Report', to: Routes.Report },
    { icon: 'person', label: 'Profile', to: Routes.Profile },
  ];

  get tabPosition() {
    if (this.$q.platform.is.desktop) return 'q-header';
    return 'q-footer';
  }
}
</script>

<style lang="scss" scoped>
.container {
  border: transparent dashed 1px;
}
</style>
