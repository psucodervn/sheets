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
  >
    <template v-slot:body-cell-name="props">
      <q-td :props="props" @click="goToIssues(props.row.name)">
        <span :to="`/point/issues/${props.row.name}`" class="name">
          {{ props.value }}
        </span>
      </q-td>
    </template>
  </q-table>
</template>

<script lang="ts">
  import { Component, Prop, Vue } from 'vue-property-decorator';
  import { TableColumn, TablePagination } from '@/types/datatable';
  import { IUserPoint } from '@/model/point';
  import { formatPoint } from '@/utils/formatter';
  import { Routes } from '@/router/names';

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
        return this.users.sort((a: IUserPoint, b: IUserPoint) => {
          if (a.pointTotal !== b.pointTotal) {
            return -(a.pointTotal - b.pointTotal);
          } else if (a.issues.length !== b.issues.length) {
            return -(a.issues.length - b.issues.length);
          } else {
            return a.name > b.name ? 1 : -1;
          }
        });
      } catch (e) {
        return [];
      }
    }

    goToIssues(name: string) {
      this.$router.push({ name: Routes.PointIssues, params: { name } });
    }
  };
</script>

<style lang="scss" scoped>
  span.name {
    cursor: pointer;
  }
</style>
