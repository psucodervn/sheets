<template>
  <div class="row q-gutter-y-md q-py-sm" style="padding-bottom: 70px">
    <q-card
      bordered
      class="row full-width cursor-pointer"
      @click="$router.push('/balance/accounts')"
    >
      <div class="col flex items-center q-pa-md">
        <div class="full-width text-h4">{{ userCount }}</div>
        <div class="full-width text-italic">Active Accounts</div>
      </div>
      <div class="col-grow">
        <q-avatar size="100px" icon="person" style="top: calc(50% - 50px)" />
      </div>
    </q-card>
    <q-card
      bordered
      class="row full-width cursor-pointer"
      @click="$router.push('/balance/transactions')"
    >
      <div class="col-grow">
        <q-avatar
          size="100px"
          icon="shopping_cart"
          style="top: calc(50% - 50px)"
        />
      </div>
      <div class="col flex items-center q-pa-md text-right">
        <div class="full-width text-h4">{{ transactionCount }}</div>
        <div class="full-width text-italic">Transactions</div>
      </div>
    </q-card>
    <new-transaction-btn />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { BalanceModule } from '@/store';
import NewTransactionBtn from '@/modules/balance/components/transactions/NewTransactionBtn.vue';

@Component({
  components: { NewTransactionBtn },
})
export default class BalanceDashboard extends Vue {
  get userCount() {
    return BalanceModule.users.length;
  }

  get transactionCount() {
    return BalanceModule.transactions.length;
  }

  async mounted() {
    this.$navigation.title = 'Overview';
    this.$navigation.parent = null;
  }
}
</script>

<style scoped lang="scss"></style>
