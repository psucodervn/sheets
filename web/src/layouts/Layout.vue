<template>
  <q-layout view="hHh lpR fFf">
    <q-page-container>
      <router-view/>
    </q-page-container>
    <q-header class="shadow-2 bg-dark">
      <q-tabs dense>
        <q-route-tab exact inline-label label="Balance" to="/"/>
        <q-route-tab exact inline-label label="Points" to="/points"/>
      </q-tabs>
    </q-header>
  </q-layout>
</template>

<script lang="ts">
  import { Vue } from 'vue-property-decorator';

  export default class Layout extends Vue {
  }
</script>

<style lang="scss">
</style>
