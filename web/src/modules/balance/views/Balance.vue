<template>
  <q-pull-to-refresh @refresh="fetchData">
    <balance-table :loading="loading" :users="users" />
    <q-space class="q-py-xs" />
    <transaction-table :transactions="transactions" />
  </q-pull-to-refresh>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { IUser } from "@/model/user";
import { BalanceModule, UserModule } from "@/store";
import TransactionTable from "@/modules/balance/components/TransactionTable.vue";
import { ITransaction } from "@/modules/balance/types/transaction";
import BalanceTable from "@/modules/balance/views/BalanceTable.vue";

@Component({
  components: { BalanceTable, TransactionTable }
})
export default class Balance extends Vue {
  get users(): IUser[] {
    return UserModule.users.sort(
      (a: IUser, b: IUser) => -(a.balance.value - b.balance.value)
    );
  }
  loading = false;
  transactions: ITransaction[] = [];

  async fetchData(done?: Function) {
    try {
      this.loading = true;
      await UserModule.fetchUsers();
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
    this.$navigation.title = "Overview";
    this.$navigation.to = null;
    await this.fetchData();
  }
}
</script>
