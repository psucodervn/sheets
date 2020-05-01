<template>
  <router-view />
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import { BalanceModule } from '@/store';
import NewTransactionBtn from '@/modules/balance/views/NewTransactionBtn.vue';

@Component({
  components: { NewTransactionBtn },
})
export default class Balance extends Vue {
  async created() {
    await Promise.all([
      BalanceModule.fetchUsers(),
      BalanceModule.fetchTransactions(),
    ]);
  }
}
</script>
