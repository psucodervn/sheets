<template>
  <q-layout view="hHh Lpr lFf">
    <q-page-container class="page">
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
@Component
export default class SimpleLayout extends Vue {}
</script>

<style lang="scss" scoped></style>
