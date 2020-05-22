<template>
  <div>
    <balance-table :loading="false" :users="users" />
    <new-transaction-btn />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { IUser } from '@/model/user';
import BalanceTable from '@/modules/balance/components/BalanceTable.vue';
import { Routes } from '@/router/names';
import { BalanceModule } from '@/store';
import NewTransactionBtn from '@/modules/balance/components/transactions/NewTransactionBtn.vue';

@Component({
  components: { NewTransactionBtn, BalanceTable },
})
export default class BalanceAccounts extends Vue {
  get users(): IUser[] {
    return BalanceModule.users.sort(
      (a: IUser, b: IUser) => -(a.balance - b.balance)
    );
  }

  async mounted() {
    this.$navigation.title = 'Accounts';
    this.$navigation.parent = { name: Routes.BalanceDashboard };
  }
}
</script>
