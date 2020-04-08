<template>
  <q-page class="home q-pa-md">
    <p>Story Points</p>
    <div class="row">
      <q-select
        :options="months" class="col" dense label="Month" outlined
        v-model="month"
      ></q-select>
      <q-space class="q-pa-sm"/>
      <q-select
        :options="years" class="col" dense label="Year" outlined
        v-model="year"
      ></q-select>
    </div>
    <q-space class="q-pa-sm"/>
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
  </q-page>
</template>

<script lang="ts">
  import { Component, Vue, Watch } from 'vue-property-decorator';
  import BarChart from '@/components/BarChart.vue';
  import HorizontalBarChart from '@/components/HorizontalBarChart.vue';
  import { PointModule } from '@/store';
  import { IUserPoint } from '@/model/point';
  import { Month, Months } from '@/constants/datetime';
  import { TableColumn, TablePagination } from '@/types/datatable';
  import { formatPoint } from '@/utils/formatter';

  @Component({
    components: { BarChart, HorizontalBarChart },
  })
  export default class Points extends Vue {
    private months = Months;
    private years = [2019, 2020];

    private columns: Array<TableColumn> = [
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
    private pagination: TablePagination = {
      sortBy: 'points', descending: true, rowsPerPage: -1,
    };
    private loading = false;

    get users(): IUserPoint[] {
      try {
        return PointModule.users.sort(
          (a: IUserPoint, b: IUserPoint) => -(a.pointTotal - b.pointTotal),
        );
      } catch (e) {
        return [];
      }
    }

    get year() {
      return PointModule.year;
    }

    set year(year: number) {
      PointModule.setYear(year);
    }

    get month() {
      return PointModule.month;
    }

    set month(month: Month) {
      PointModule.setMonth(month);
    }

    @Watch('year')
    @Watch('month')
    async fetchData() {
      try {
        this.loading = true;
        await PointModule.fetchPoints({ year: this.year, month: this.month.value });
      } catch (e) {
        console.log(e.message);
      } finally {
        this.loading = false;
      }
    }

    async mounted() {
      await this.fetchData();
    }
  }
</script>

<style>
</style>
