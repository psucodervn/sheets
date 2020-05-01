<template>
  <q-pull-to-refresh @refresh="fetchData" style="padding-bottom: 70px">
    <balance-table :loading="loading" :users="users" />
    <new-transaction-btn />
  </q-pull-to-refresh>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { IUser } from '@/model/user';
import BalanceTable from '@/modules/balance/components/BalanceTable.vue';
import { Routes } from '@/router/names';
import { BalanceModule } from '@/store';
import NewTransactionBtn from '@/modules/balance/views/NewTransactionBtn.vue';

@Component({
  components: { NewTransactionBtn, BalanceTable },
})
export default class BalanceAccounts extends Vue {
  get users(): IUser[] {
    return BalanceModule.users.sort(
      (a: IUser, b: IUser) => -(a.balance - b.balance)
    );
  }

  loading = false;

  async fetchData(done?: Function) {
    try {
      this.loading = true;
      await BalanceModule.fetchUsers();
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
    this.$navigation.title = 'Accounts';
    this.$navigation.parent = { name: Routes.BalanceDashboard };
    await this.fetchData();
  }
}
</script>
