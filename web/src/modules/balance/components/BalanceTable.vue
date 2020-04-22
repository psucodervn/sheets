<template>
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
</template>
<script lang="ts">
  import { Component, Prop, Vue } from 'vue-property-decorator';
  import { IUser } from '@/model/user';
  import { TableColumn, TablePagination } from '@/types/datatable';
  import formatter from '@/utils/formatter';

  @Component
  export default class BalanceTable extends Vue {
    @Prop({ type: Array, required: true }) users!: IUser;
    @Prop({ type: Boolean, default: false }) loading!: boolean;

    columns: Array<TableColumn> = [
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
        format: formatter.currency,
      },
    ];
    pagination: TablePagination = {
      sortBy: 'balance',
      rowsPerPage: -1,
      descending: true,
    };
  }
</script>
