<template>
  <q-layout view="hHh lpR fFf">
    <navigation-bar title="Story Points"/>
    <q-page-container>
      <router-view/>
    </q-page-container>
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
  import NavigationBar from '@/components/NavigationBar.vue';
  import { Routes } from '@/router/names';

  @Component({
    components: {
      NavigationBar,
    },
  })
  export default class Layout extends Vue {
    tabs = [
      { label: 'Balance', to: Routes.Balance },
      { label: 'Point', to: Routes.Point },
    ];
  }
</script>

<style lang="scss">
</style>
