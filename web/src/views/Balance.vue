<template>
  <q-page class="q-pa-md">
    <p>Balance</p>
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
  </q-page>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator';
  import { IUser } from '@/model/user';
  import BarChart from '@/components/BarChart.vue';
  import HorizontalBarChart from '@/components/HorizontalBarChart.vue';
  import { UserModule } from '@/store';
  import { formatCurrency } from '@/utils/formatter';
  import { TableColumn, TablePagination } from '@/types/datatable';

  @Component({
    components: { BarChart, HorizontalBarChart },
  })
  export default class Balance extends Vue {
    get users(): IUser[] {
      return UserModule.users.sort(
        (a: IUser, b: IUser) => -(a.balance.value - b.balance.value),
      );
    }

    private columns: Array<TableColumn> = [
      {
        name: 'name',
        label: 'Name',
        field: 'name',
        sortable: true,
        align: 'left',
      },
      {
        name: 'balance',
        label: 'Balance (vnđ)',
        field: (row: any) => row.balance.value,
        sortable: true,
        format: formatCurrency,
      },
    ];
    private pagination: TablePagination = { sortBy: 'balance', rowsPerPage: -1, descending: true };
    private loading = false;

    async fetchData() {
      try {
        this.loading = true;
        await UserModule.fetchUsers();
      } catch (e) {
        console.log(e.message);
      } finally {
        this.loading = false;
      }
    }

    async mounted() {
      await this.fetchData();
    }
  }
</script>

<style></style>
