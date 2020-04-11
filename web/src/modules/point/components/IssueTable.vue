<template>
  <q-table
    :columns="columns"
    :data="issues"
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
    @Prop({ type: Array, required: true }) issues!: IIssue[];

    // TODO: implement custom sort for issue's key
    columns: Array<TableColumn> = [
      { name: 'key', label: 'Key', field: 'key', align: 'left', sortable: true },
      {
        name: 'summary', label: 'Summary', field: 'summary', align: 'left', sortable: true,
      },
    ];
    pagination: TablePagination = {
      sortBy: 'points', descending: true, rowsPerPage: -1,
    };
  }
</script>

<style lang="scss" scoped>

</style>
