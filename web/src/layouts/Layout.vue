<template>
  <q-layout view="hHh lpR fFf">
    <q-page-container>
      <router-view/>
    </q-page-container>
    <q-footer class="shadow-2 bg-dark">
      <q-tabs dense>
        <q-route-tab exact inline-label label="Balance" to="/balance"/>
        <q-route-tab exact inline-label label="Point" to="/point"/>
      </q-tabs>
    </q-footer>
  </q-layout>
</template>

<script lang="ts">
  import { Vue } from 'vue-property-decorator';

  export default class Layout extends Vue {
  }
</script>

<style lang="scss">
</style>
