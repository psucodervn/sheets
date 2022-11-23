<template>
  <q-table
    :columns="columns"
    :data="rows"
    :pagination.sync="pagination"
    :loading="loading"
    binary-state-sort
    grid
    grid-header
    hide-bottom
  >
    <template v-slot:item="{ row }">
      <div class="col-xs-12 q-my-xs">
        <q-card>
          <div class="row q-pa-sm">
            <div
              :id="row.name"
              @click="() => onClickHeader(row.name)"
              class="col text-yellow cursor-pointer"
            >
              {{ row.displayName }}
            </div>
            <div class="col text-center">
              <q-badge
                :label="`${row.pointTotal} pts`"
                class="q-pa-xs q-mr-xs text-bold"
                outline
              />
              <q-tooltip
                anchor="top middle"
                content-class="bg-blue-grey-9 text-white shadow-2"
                content-style="font-size: 1em; border: black solid 1px"
                self="bottom middle"
              >
                {{ row.pointTotal }} Story Points on JIRA
              </q-tooltip>
            </div>
            <div class="col text-center">
              <q-badge
                :label="`${row.wakatimeHuman || 'N/A'}`"
                class="q-pa-xs q-mr-xs text-bold"
                outline
              />
              <q-tooltip
                anchor="top middle"
                content-class="bg-blue-grey-9 text-white shadow-2"
                content-style="font-size: 1em; border: black solid 1px"
                self="bottom middle"
              >
                {{ row.wakatimeHuman || 'N/A' }} on Wakatime
              </q-tooltip>
            </div>
          </div>
          <template v-if="!collapsed">
            <q-separator />
            <div class="q-pa-sm">
              <div>
                <template v-for="is in row.issues">
                  <IssueRow :issue="is" />
                </template>
              </div>
            </div>
          </template>
        </q-card>
      </div>
    </template>
  </q-table>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { IIssue, IUserPoint } from '@/model/point';
import { TableColumn, TablePagination } from '@/types/datatable';
import IssueRow from '@/modules/point/components/IssueRow.vue';

@Component({
  components: { IssueRow },
})
export default class ReportTable extends Vue {
  @Prop({ type: Array, required: true }) points!: IUserPoint[];

  @Prop({ type: Boolean, default: false }) loading!: boolean;

  columns: TableColumn[] = [
    {
      name: 'name',
      field: 'name',
      label: 'Name',
      sortable: true,
      align: 'left',
    },
    {
      name: 'point',
      field: 'pointTotal',
      label: 'Point',
      sortable: true,
      align: 'right',
    },
    {
      name: 'wakatime',
      field: 'wakatimeSeconds',
      label: 'Wakatime',
      sortable: true,
      align: 'right',
    },
  ];

  pagination: TablePagination = {
    rowsPerPage: 0,
    sortBy: 'name',
  };

  collapsed = false;

  get rows(): IUserPoint[] {
    return this.points.map((up: IUserPoint) => ({
      ...up,
      issues: up.issues.sort((a: IIssue, b: IIssue) => a.status.localeCompare(b.status)),
    }));
  }

  onClickHeader(id: string) {
    this.collapsed = !this.collapsed;
    this.$nextTick().then(() => {
      const el = this.$el.querySelector(`#${id}`)!;
      el.scrollIntoView();
    });
  }
}
</script>
