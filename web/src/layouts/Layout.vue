<template>
  <q-layout view="hHh lpR fFf">
    <q-page class="container q-px-sm">
      <navigation-bar/>
      <router-view/>
    </q-page>
    <q-footer class="shadow-2 bg-dark">
      <q-tabs dense>
        <q-route-tab
          :label="t.label"
          :to="t.to" exact inline-label v-for="t in tabs"/>
      </q-tabs>
    </q-footer>
  </q-layout>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator';
  import { Routes } from '@/router/names';

  @Component({})
  export default class Layout extends Vue {
    tabs = [
      { label: 'Balance', to: Routes.Balance },
      { label: 'Point', to: Routes.Point },
    ];
  }
</script>

<style lang="scss" scoped>
  .container {
    border: transparent dashed 1px;
  }
</style>
