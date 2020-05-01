<template>
  <q-page-sticky position="bottom" :offset="[10, 10]">
    <q-btn fab icon="add" color="primary" @click="goToNewTransaction" />
  </q-page-sticky>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Routes } from '@/router/names';

@Component
export default class NewTransactionBtn extends Vue {
  goToNewTransaction() {
    this.$router.push({ name: Routes.BalanceTransactionsNew }).catch(() => {});
  }
}
</script>
