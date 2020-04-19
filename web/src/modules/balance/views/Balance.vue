<template>
  <q-pull-to-refresh @refresh="fetchData">
    <q-table
      :columns="columns"
      :data="users"
      :loading="loading"
      :pagination.sync="pagination"
      binary-state-sort
      dense
      hide-bottom
      row-key="name"
    />
    <q-space class="q-py-xs" />
    <transaction-table :transactions="transactions" />
  </q-pull-to-refresh>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { IUser } from "@/model/user";
import { BalanceModule, UserModule } from "@/store";
import { formatCurrency } from "@/utils/formatter";
import { TableColumn, TablePagination } from "@/types/datatable";
import TransactionTable from "@/modules/balance/components/TransactionTable.vue";
import { ITransaction } from "@/modules/balance/types/transaction";

@Component({
  components: { TransactionTable }
})
export default class Balance extends Vue {
  get users(): IUser[] {
    return UserModule.users.sort(
      (a: IUser, b: IUser) => -(a.balance.value - b.balance.value)
    );
  }

  columns: Array<TableColumn> = [
    {
      name: "name",
      label: "Name",
      field: "name",
      sortable: true,
      align: "left"
    },
    {
      name: "balance",
      label: "Balance (vnđ)",
      field: (row: any) => row.balance.value,
      sortable: true,
      format: formatCurrency
    }
  ];
  pagination: TablePagination = {
    sortBy: "balance",
    rowsPerPage: -1,
    descending: true
  };
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

<style lang="scss" scoped></style>
