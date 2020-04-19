<template>
  <q-pull-to-refresh @refresh="fetchData">
    <transaction-table :transactions="transactions" />
  </q-pull-to-refresh>
</template>

<script lang="ts">
import { BalanceModule } from "@/store";
import { ITransaction } from "@/modules/balance/types/transaction";
import { Routes } from "@/router/names";
import { Component, Vue } from "vue-property-decorator";
import TransactionTable from "@/modules/balance/components/TransactionTable.vue";

@Component({
  components: { TransactionTable }
})
export default class BalanceTransactions extends Vue {
  loading = false;
  transactions: ITransaction[] = [];

  async fetchData(done?: Function) {
    try {
      this.loading = true;
      this.transactions = (await BalanceModule.fetchTransactions())!;
    } catch (e) {
      console.log(e.message);
    } finally {
      this.loading = false;
      if (done) {
        done();
      }
    }
  }

  async mounted() {
    this.$navigation.title = "Transactions";
    this.$navigation.to = { name: Routes.BalanceDashboard };
    await this.fetchData();
  }
}
</script>
