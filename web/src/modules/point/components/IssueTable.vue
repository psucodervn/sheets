<template>
  <q-table
    class="issue-table"
    :columns="columns"
    :data="rows"
    :pagination.sync="pagination"
    binary-state-sort
    dense
    row-key="key"
  >
    <template v-slot:body-cell-key="props">
      <q-td :props="props" class="key">
        <a :href="`https://pm.vzota.com.vn/browse/${props.row.key}`" class="link" target="_blank" rel="noreferrer">
          {{ props.value }}
        </a>
      </q-td>
    </template>
    <template v-slot:bottom="props">
      <div class="q-py-xs">
        <span class="q-mr-md">Total:</span>
        <q-badge class="q-px-sm q-py-xs" outline>{{ totalPoint }}</q-badge>
        <span class="q-mx-sm">point(s)</span>
        <q-badge class="q-px-sm q-py-xs" outline>{{ issueCount }}</q-badge>
        <span class="q-mx-sm">issue(s)</span>
      </div>
    </template>
  </q-table>
</template>

<script lang="ts">
  import { Component, Prop, Vue } from 'vue-property-decorator';
  import { TableColumn, TablePagination } from '@/types/datatable';
  import { IIssue } from '@/model/point';

  interface IItem extends IIssue {
    projectKey: string;
    index: number;
  }

  @Component({})
  export default class IssueTable extends Vue {
    @Prop({ type: Array, required: true }) issues!: IIssue[];

    // TODO: implement custom sort for issue's key
    columns: Array<TableColumn> = [
      {
        name: 'key', label: 'Key', field: 'key', align: 'left', sortable: true,
        sort: (keyA: string, keyB: string, a: IItem, b: IItem) => {
          if (a.projectKey !== b.projectKey) return a.index - b.index;
          return a.projectKey.localeCompare(b.projectKey);
        },
      },
      { name: 'point', label: 'Point', field: 'point', align: 'right', sortable: true },
      {
        name: 'summary', label: 'Summary', field: 'summary', align: 'left', sortable: true,
      },
    ];
    pagination: TablePagination = {
      sortBy: 'points', descending: true, rowsPerPage: -1,
    };

    get rows(): IItem[] {
      return this.issues.map((i: IIssue): IItem => {
        const ar = i.key.split('-');
        return { ...i, projectKey: ar[0], index: Number(ar[1]) };
      });
    }

    get issueCount() {
      return this.issues.length;
    }

    get totalPoint() {
      return this.issues.map(i => i.point).reduce((p, c) => p + c);
    }
  }
</script>

<style lang="scss">
  td.key {
    cursor: pointer;
  }

  a.link {
    color: white;
    text-decoration: none;
  }

  .issue-table {
    th:first-child, td.key:first-child {
      position: sticky;
      left: 0;
      z-index: 1;
      background-color: #161616;
    }

    td.sticky {
      position: sticky;
      left: 0;
      z-index: 1;
      background-color: #161616;
    }
  }
</style>
