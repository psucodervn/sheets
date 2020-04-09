<template>
  <q-page class="home q-pa-md">
    <q-pull-to-refresh @refresh="fetchData">
      <p>Story Points</p>
      <PointTimeFilter/>
      <q-space class="q-pa-sm"/>
      <PointTable :loading="loading" :users="users"/>
    </q-pull-to-refresh>
  </q-page>
</template>

<script lang="ts">
  import { Component, Vue, Watch } from 'vue-property-decorator';
  import BarChart from '@/components/BarChart.vue';
  import HorizontalBarChart from '@/components/HorizontalBarChart.vue';
  import { PointModule } from '@/store';
  import PointTimeFilter from '@/modules/point/components/PointTimeFilter.vue';
  import PointTable from '@/modules/point/components/PointTable.vue';

  @Component({
    components: { PointTable, PointTimeFilter, BarChart, HorizontalBarChart },
  })
  export default class Points extends Vue {
    private loading = false;

    get year() {
      return PointModule.year;
    }

    get month() {
      return PointModule.month;
    }

    get users() {
      return PointModule.users;
    }

    @Watch('year')
    @Watch('month')
    async fetchData(done?: Function) {
      try {
        this.loading = true;
        await PointModule.fetchPoints({ year: this.year, month: this.month.value });
      } catch (e) {
        console.log(e.message);
      } finally {
        this.loading = false;
        if (typeof done === 'function') {
          done();
        }
      }
    }

    async mounted() {
      await this.fetchData();
    }
  }
</script>
