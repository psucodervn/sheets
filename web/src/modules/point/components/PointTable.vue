<template>
  <q-table
    :columns="columns"
    :data="items"
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
  import { TableColumn, TablePagination } from '@/types/datatable';
  import { IUserPoint } from '@/model/point';
  import { formatPoint } from '@/utils/formatter';

  @Component
  export default class PointTable extends Vue {
    @Prop({ type: Boolean, required: true }) loading!: boolean;
    @Prop({ type: Array, required: true }) users!: IUserPoint[];

    columns: Array<TableColumn> = [
      { name: 'name', label: 'Name', field: 'displayName', sortable: true, align: 'left' },
      {
        name: 'count', label: 'Issue Count', sortable: true,
        field: (row: IUserPoint) => row.issues.length,
      },
      {
        name: 'points', label: 'Points', field: 'pointTotal', sortable: true,
        format: formatPoint,
      },
    ];
    pagination: TablePagination = {
      sortBy: 'points', descending: true, rowsPerPage: -1,
    };

    get items(): IUserPoint[] {
      try {
        return this.users.sort(
          (a: IUserPoint, b: IUserPoint) => -(a.pointTotal - b.pointTotal),
        );
      } catch (e) {
        return [];
      }
    }
  };
</script>
