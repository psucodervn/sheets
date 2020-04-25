<template>
  <q-pull-to-refresh @refresh="fetchData">
    <report-time-filter :range="range" @input="onUpdateTimeRange"/>
    <report-table :points="points"></report-table>
  </q-pull-to-refresh>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator';
  import ReportTimeFilter from '@/modules/point/components/ReportTimeFilter.vue';
  import { TimeRange } from '@/types/datetime';
  import { date } from 'quasar';
  import ReportTable from '@/modules/point/components/ReportTable.vue';
  import { PointModule } from '@/store';
  import { IUserPoint } from '@/model/point';

  @Component({
    components: { ReportTable, ReportTimeFilter },
  })
  export default class Report extends Vue {
    range: TimeRange = this.defaultRange();
    points: IUserPoint[] = [];

    defaultRange(): TimeRange {
      const to = new Date();
      const from = date.addToDate(to, { days: -6 });
      return new TimeRange(from, to);
    }

    async mounted() {
      this.$navigation.title = 'Report';
      this.$navigation.to = null;
      await this.fetchData();
    }

    async fetchData(done?: Function) {
      try {
        this.points = await PointModule.fetchReport({ range: this.range });
      } catch (e) {
        this.$q.notify({
          message: 'Fetch report failed: ' + String(e.message),
          type: 'negative',
        });
      } finally {
        if (typeof done === 'function') {
          done();
        }
      }
    }

    async onUpdateTimeRange() {
      await this.fetchData();
    }
  }
</script>

<style lang="scss" scoped></style>
