<template>
  <q-pull-to-refresh @refresh="fetchData">
    <balance-table :loading="loading" :users="users"/>
  </q-pull-to-refresh>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator';
  import { IUser } from '@/model/user';
  import BalanceTable from '@/modules/balance/components/BalanceTable.vue';
  import { Routes } from '@/router/names';
  import { BalanceModule } from '@/store';

  @Component({
    components: { BalanceTable },
  })
  export default class BalanceOverview extends Vue {
    get users(): IUser[] {
      return BalanceModule.users.sort(
        (a: IUser, b: IUser) => -(a.balance.value - b.balance.value),
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
      this.$navigation.to = { name: Routes.BalanceDashboard };
      await this.fetchData();
    }
  }
</script>
