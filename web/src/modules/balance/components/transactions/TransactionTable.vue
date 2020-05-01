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
  >
    <template v-slot:body="props">
      <q-tr :props="props">
        <q-td
          v-for="col in props.cols"
          :key="col.name"
          :props="props"
          @click="props.expand = !props.expand"
        >
          {{ col.value }}
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
import { Component, Prop, Vue } from 'vue-property-decorator';
import { TableColumn, TablePagination } from '@/types/datatable';
import { ITransaction } from '@/modules/balance/types/transaction';
import formatter from '@/utils/formatter';

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
}
</script>

<style lang="scss">
td.date {
  width: 30px;
}

td.expand {
  text-align: left;
  white-space: normal;
}

.blur {
  opacity: 0.4;
}
</style>
