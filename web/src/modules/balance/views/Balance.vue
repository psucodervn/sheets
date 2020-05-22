<template>
  <q-pull-to-refresh @refresh="fetchData">
    <router-view />
  </q-pull-to-refresh>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { BalanceModule } from '@/store';
import NewTransactionBtn from '@/modules/balance/components/transactions/NewTransactionBtn.vue';

@Component({
  components: { NewTransactionBtn },
})
export default class Balance extends Vue {
  async created() {
    await this.fetchData();
  }

  async fetchData(done?: Function) {
    try {
      await Promise.all([
        BalanceModule.fetchUsers(),
        BalanceModule.fetchTransactions(),
      ]);
    } catch (e) {
      this.$q.notify({
        message: `Fetch data failed: ${e.message}`,
        type: 'negative',
      });
    } finally {
      if (done) {
        done();
      }
    }
  }
}
</script>
