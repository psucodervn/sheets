<template>
  <q-table
    :columns="columns"
    :data="items"
    :pagination.sync="pagination"
    binary-state-sort
    dense
    hide-bottom
    row-key="key"
  >
  </q-table>
</template>

<script lang="ts">
  import { Component, Prop, Vue } from 'vue-property-decorator';
  import { TableColumn, TablePagination } from '@/types/datatable';
  import { IIssue } from '@/model/point';

  @Component({})
  export default class IssueTable extends Vue {
    @Prop({ type: Array, required: true }) items!: IIssue[];

    columns: Array<TableColumn> = [
      { name: 'key', label: 'Key', field: 'key' },
      { name: 'summary', label: 'Summary', field: 'summary' },
    ];
    pagination: TablePagination = {
      sortBy: 'points', descending: true, rowsPerPage: -1,
    };
  }
</script>

<style lang="scss" scoped>

</style>
