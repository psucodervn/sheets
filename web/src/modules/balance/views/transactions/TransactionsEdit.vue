<template>
  <transaction-detail :transaction.sync="transaction" :is-edit="true" />
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { Routes } from '@/router/names';
import TransactionDetail from '@/modules/balance/components/transactions/TransactionDetail.vue';
import { ITransaction } from '@/modules/balance/types/transaction';
import { BalanceModule } from '@/store';

@Component({
  components: { TransactionDetail },
})
export default class TransactionsEdit extends Vue {
  transaction: Partial<ITransaction> = {
    payers: [],
    participants: [],
  };

  async created() {
    this.$navigation.title = 'Edit Transaction';
    this.$navigation.parent = { name: Routes.BalanceTransactions };
    try {
      this.transaction = await BalanceModule.fetchTransaction({
        id: this.$route.params.id,
      });
    } catch (e) {
      this.$q.notify({
        caption: 'Fetch transaction failed',
        message: e.message,
        type: 'negative',
      });
      await this.$router.push({ name: Routes.BalanceTransactions });
    }
  }
}
</script>

<style scoped lang="scss"></style>
