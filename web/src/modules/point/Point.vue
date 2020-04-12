<template>
  <q-pull-to-refresh @refresh="fetchData">
    <PointTimeFilter/>
    <q-space class="q-py-xs"/>
    <PointTable :loading="loading" :users="users"/>
  </q-pull-to-refresh>
</template>

<script lang="ts">
  import { Component, Vue, Watch } from 'vue-property-decorator';
  import { PointModule } from '@/store';
  import PointTimeFilter from '@/modules/point/components/PointTimeFilter.vue';
  import PointTable from '@/modules/point/components/PointTable.vue';
  import NavigationBar from '@/components/NavigationBar.vue';

  @Component({
    components: { NavigationBar, PointTable, PointTimeFilter },
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
      this.$navigation.title = 'Story Points';
      this.$navigation.to = null;
      await this.fetchData();
    }
  }
</script>
