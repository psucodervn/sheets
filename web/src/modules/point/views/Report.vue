<template>
  <q-pull-to-refresh @refresh="fetchData">
    <report-time-filter :range="range" @input="onUpdateTimeRange" />
    <report-table :points="points" :loading="loading"></report-table>
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
  loading = false;

  defaultRange(): TimeRange {
    const to = new Date();
    const from = date.addToDate(to, { days: -6 });
    return new TimeRange(from, to);
  }

  async mounted() {
    this.$navigation.title = 'Report';
    this.$navigation.from = null;
    const fromQuery = this.$route.query.from;
    if (fromQuery && typeof fromQuery === 'string') {
      const d = date.extractDate(fromQuery, 'DD-MM-YYYY');
      if (d.getFullYear() >= 2019 && d.getFullYear() <= 3000) {
        this.range.from = d;
        this.range.to = date.addToDate(d, { days: 6 });
      } else {
        this.$router
          .replace({
            name: this.$route.name,
            query: {
              ...this.$route.query,
              from: date.formatDate(this.range.from, 'DD-MM-YYYY'),
            },
          })
          .catch(() => {});
      }
    }
    await this.fetchData();
  }

  async fetchData(done?: Function) {
    try {
      this.loading = true;
      this.points = await PointModule.fetchReport({ range: this.range });
    } catch (e) {
      this.$q.notify({
        message: 'Fetch report failed: ' + String(e.message),
        type: 'negative',
      });
    } finally {
      this.loading = false;
      if (typeof done === 'function') {
        done();
      }
    }
  }

  async onUpdateTimeRange() {
    try {
      await this.$router.replace({
        name: this.$route.name,
        query: {
          from: date.formatDate(this.range.from, 'DD-MM-YYYY'),
        },
      });
      await this.fetchData();
    } catch (e) {
      console.log(e.message);
    }
  }
}
</script>

<style lang="scss" scoped></style>
