<template>
  <q-layout view="hHh lpR fFf">
    <q-page-container>
      <q-page class="container q-pa-sm">
        <navigation-bar/>
        <router-view/>
      </q-page>
    </q-page-container>
    <component :is="tabPosition" class="shadow-2 bg-dark">
      <q-tabs dense>
        <q-route-tab
          :label="t.label"
          :to="t.to"
          inline-label
          v-for="t in tabs"
        />
      </q-tabs>
    </component>
  </q-layout>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator';
  import { Routes } from '@/router/names';

  @Component({})
  export default class Layout extends Vue {
    tabs = [
      { label: 'Balance', to: Routes.BalanceDashboard },
      { label: 'Point', to: Routes.Point },
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
