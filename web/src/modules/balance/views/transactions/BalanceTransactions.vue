<template>
  <q-pull-to-refresh @refresh="fetchData" style="padding-bottom: 70px">
    <div class="row q-py-sm">
      <user-filter
        :users="users"
        :selected-users="selectedUserNames"
        @update:selectedUsers="onSelectUsers"
      />
    </div>
    <transaction-table
      :transactions="transactions"
      :expand-all="selectedUserNames.length > 0"
      @remove="remove"
    />
    <new-transaction-btn />
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
import TransactionTable from '@/modules/balance/components/transactions/TransactionTable.vue';
import UserFilter from '@/modules/balance/components/UserFilter.vue';
import NewTransactionBtn from '@/modules/balance/components/transactions/NewTransactionBtn.vue';

@Component({
  components: { NewTransactionBtn, UserFilter, TransactionTable },
})
export default class BalanceTransactions extends Vue {
  loading = false;
  selectedUserNames: string[] = [];

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
    if (!filter) this.selectedUserNames = [];
    else this.selectedUserNames = String(filter).split(',');
  }

  get transactions(): ITransaction[] {
    if (!this.selectedUserNames.length) {
      return BalanceModule.transactions;
    }

    const selectedUserIds = this.selectedUserNames.map(
      name => BalanceModule.userNames[name]!.id
    );
    return BalanceModule.transactions
      .filter(
        (tx: ITransaction) =>
          selectedUserIds.findIndex((id: string) => tx.changes[id]) >= 0
      )
      .map((tx: ITransaction) => {
        const changes: TTransactionChanges = {};
        for (const userID of Object.keys(tx.changes)) {
          changes[userID] = { ...tx.changes[userID] };
          changes[userID].filteredOut = !this.selectedUserNames.some(
            (n: string) => n === changes[userID].name
          );
        }
        return { ...tx, changes };
      });
  }

  async remove(txID: string) {
    const idx = BalanceModule.transactions.findIndex(tx => tx.id === txID);
    if (idx < 0) return;
    const tx = BalanceModule.transactions[idx];
    this.$q
      .dialog({
        message: `Remove transaction '${tx.summary}'?`,
        cancel: true,
      })
      .onOk(async () => {
        try {
          await BalanceModule.removeTransaction({ id: tx.id });
          this.$q.notify({
            caption: 'Remove succeed',
            message: `Transaction '${tx.summary}' was removed`,
            type: 'positive',
          });
          BalanceModule.transactions.splice(idx, 1);
        } catch (e) {
          this.$q.notify({
            caption: 'Remove failed',
            message: e.message,
            type: 'negative',
          });
        }
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
    this.$navigation.parent = { name: Routes.BalanceDashboard };
    await this.fetchData();
    if (this.$route.query.filter) {
      this.selectedUserNames = this.$route.query.filter.toString().split(',');
    }
  }
}
</script>
