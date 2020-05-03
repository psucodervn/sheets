<template>
  <q-table
    :columns="columns"
    :pagination.sync="pagination"
    :data="items"
    binary-state-sort
    :rows-per-page-options="[10, 20, 50, 0]"
    dense
    row-key="id"
    :expanded.sync="expanded"
    class="tx-table"
  >
    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td
          v-for="col in props.cols"
          :key="col.name"
          :props="props"
          @click="props.expand = !props.expand"
          :class="`td-${col.name}`"
        >
          <template v-if="col.name === 'actions'">
            <q-btn
              icon="edit"
              dense
              rounded
              flat
              @click.stop="edit(props.row.id)"
            />
            <q-btn
              icon="delete"
              dense
              rounded
              flat
              @click.stop="remove(props.row.id)"
            />
          </template>
          <template v-else> {{ col.value }}</template>
        </q-td>
      </q-tr>
      <q-tr v-show="props.expand || expandAll" :props="props">
        <td colspan="100%" class="expand">
          <div>
            <template v-for="(c, name) in props.row.changes">
              <q-badge
                v-if="c.value >= 0"
                outline
                color="green"
                class="q-ma-xs q-pa-xs"
                :class="{ blur: c.filteredOut }"
                :label="`${c.name}: +${formatValue(c.value)}`"
              />
              <q-badge
                v-else
                outline
                color="red"
                class="q-ma-xs q-pa-xs"
                :class="{ blur: c.filteredOut }"
                :label="`${c.name}: ${formatValue(c.value)}`"
              />
            </template>
          </div>
        </td>
      </q-tr>
    </template>
  </q-table>
</template>

<script lang="ts">
import { Component, Emit, Prop, Vue } from 'vue-property-decorator';
import { TableColumn, TablePagination } from '@/types/datatable';
import { ITransaction } from '@/modules/balance/types/transaction';
import formatter from '@/utils/formatter';
import { Routes } from '@/router/names';

@Component({})
export default class TransactionTable extends Vue {
  @Prop({ type: Array, required: true }) transactions!: ITransaction[];
  @Prop({ type: Boolean, default: false }) expandAll!: boolean;

  columns: TableColumn[] = [
    {
      name: 'time',
      field: 'time',
      sortable: true,
      label: 'Date',
      format: formatter.date,
      classes: 'date',
    },
    { name: 'summary', field: 'summary', label: 'Summary', align: 'left' },
    {
      name: 'value',
      field: 'value',
      label: 'Value (vnđ)',
      sortable: true,
      format: formatter.currency,
    },
    {
      name: 'actions',
      field: 'actions',
      label: '',
    },
  ];
  pagination: TablePagination = {
    sortBy: 'time',
    descending: true,
    rowsPerPage: 20,
  };
  expanded = [];

  get items(): ITransaction[] {
    return this.transactions;
  }

  formatValue(val: number) {
    return formatter.currency(val);
  }

  @Emit('remove')
  remove(id: string) {
    return id;
  }

  edit(id: string) {
    this.$router.push({ name: Routes.BalanceTransactionsEdit, params: { id } });
  }
}
</script>

<style lang="scss">
.tx-table td:first-child,
.tx-table th:first-child {
  padding-left: 8px !important;
}

.tx-table td:last-child,
.tx-table th:last-child {
  padding-right: 8px !important;
}

.tx-table {
  td.td-time {
    width: 30px;
  }

  td.td-actions {
    padding: 1px;
    width: 30px;
  }
}

td.expand {
  text-align: left;
  white-space: normal;
}

.blur {
  opacity: 0.4;
}
</style>
