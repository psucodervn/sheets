<template>
  <q-page-sticky position="bottom" :offset="[10, 10]">
    <q-btn fab icon="add" color="primary" @click="$emit('click')" />
  </q-page-sticky>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class AddButton extends Vue {}
</script>
