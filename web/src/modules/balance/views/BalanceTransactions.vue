<template>
  <q-pull-to-refresh @refresh="fetchData">
    <div class="row q-py-sm">
      <user-filter
        :users="users"
        :selected-users="selectedUsers"
        @update:selectedUsers="onSelectUsers"
      />
    </div>
    <transaction-table
      :transactions="transactions"
      :expand-all="selectedUsers.length > 0"
    />
  </q-pull-to-refresh>
</template>

<script lang="ts">
import { BalanceModule } from '@/store';
import {
  ITransaction,
  TTransactionChanges,
} from '@/modules/balance/types/transaction';
import { Routes } from '@/router/names';
import { Component, Vue, Watch } from 'vue-property-decorator';
import TransactionTable from '@/modules/balance/components/TransactionTable.vue';
import UserFilter from '@/modules/balance/components/UserFilter.vue';

@Component({
  components: { UserFilter, TransactionTable },
})
export default class BalanceTransactions extends Vue {
  loading = false;
  selectedUsers: string[] = [];

  get users() {
    return BalanceModule.users;
  }

  onSelectUsers(val: string[] | null) {
    let qr = { ...this.$route.query };
    if (!val || !val.length) {
      delete qr.filter;
    } else {
      qr.filter = val.join(',');
    }
    this.$router.replace({
      name: this.$route.name,
      query: qr,
    });
  }

  @Watch('$route.query', { deep: true, immediate: true })
  setSelectedUsersBaseOnFilterQuery() {
    const filter = this.$route.query.filter;
    if (!filter) this.selectedUsers = [];
    else this.selectedUsers = String(filter).split(',');
  }

  get transactions(): ITransaction[] {
    if (!this.selectedUsers.length) {
      return BalanceModule.transactions;
    }
    return BalanceModule.transactions
      .filter(
        (tx: ITransaction) =>
          this.selectedUsers.findIndex((name: string) => tx.changes[name]) >= 0
      )
      .map((tx: ITransaction) => {
        const changes: TTransactionChanges = {};
        for (const name of this.selectedUsers) {
          if (tx.changes[name] != undefined) changes[name] = tx.changes[name];
        }
        return { ...tx, changes };
      });
  }

  async fetchData(done?: Function) {
    try {
      this.loading = true;
      await BalanceModule.fetchTransactions();
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
    this.$navigation.title = 'Transactions';
    this.$navigation.to = { name: Routes.BalanceDashboard };
    await this.fetchData();
    if (this.$route.query.filter) {
      this.selectedUsers = this.$route.query.filter.toString().split(',');
    }
  }
}
</script>
