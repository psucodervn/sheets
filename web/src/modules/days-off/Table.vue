<template>
  <q-table
    :columns="columns"
    binary-state-sort
    :data="items"
    :pagination.sync="pagination"
    :rows-per-page-options="[20, 50, 0]"
    hide-bottom
  >
    <template #body-cell-date="{ row }">
      <q-td>
        <q-badge
          outline
          color="green"
          label="Today"
          v-if="isToday(row.date)"
        ></q-badge>
        <q-badge
          outline
          color="blue"
          label="Tomorrow"
          v-else-if="isTomorrow(row.date)"
        ></q-badge>
        <span v-else>{{ formatDate(row.date) }}</span>
      </q-td>
    </template>
    <template #body-cell-actions="{ row }">
      <q-td>
        <q-btn icon="delete_forever" dense rounded @click="remove(row)"></q-btn>
      </q-td>
    </template>
  </q-table>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';
import { TableColumn, TablePagination } from '@/types/datatable';
import { DayOff } from '@/types/logic';
import { BalanceModule } from '@/store';
import { date } from 'quasar';

@Component({})
export default class DayOffTable extends Vue {
  @Prop({ required: true }) items!: DayOff[];

  columns: TableColumn[] = [
    {
      name: 'date',
      field: (row: DayOff) => date.formatDate(row.date, 'DD/MM/YYYY'),
      label: 'Date',
      sortable: true,
      sort: (a: String, b: String, rowA: DayOff, rowB: DayOff) => {
        return new Date(rowA.date).getTime() - new Date(rowB.date).getTime();
      },
      align: 'left',
    },
    {
      name: 'user',
      field: (row: DayOff) => BalanceModule.userIds[row.userId].name,
      label: 'Name',
      align: 'left',
      sortable: true,
    },
    {
      name: 'part',
      field: 'part',
      label: 'When',
      sortable: true,
      align: 'left',
    },
    {
      name: 'note',
      field: 'note',
      label: 'Note',
      sortable: true,
      align: 'left',
    },
    {
      name: 'actions',
      field: 'actions',
      label: '',
    },
  ];

  pagination: TablePagination = {
    sortBy: 'date',
    descending: true,
    rowsPerPage: 50,
  };

  formatDate(d: Date) {
    return date.formatDate(d, 'ddd, DD/MM');
  }

  isToday(d: Date) {
    return this.formatDate(d) === this.formatDate(new Date());
  }

  isTomorrow(d: Date) {
    return (
      this.formatDate(date.addToDate(d, { days: -1 })) ===
      this.formatDate(new Date())
    );
  }

  @Emit('remove')
  remove(d: DayOff) {
    return d;
  }
}
</script>

<style scoped lang="scss"></style>
