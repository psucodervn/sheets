<template>
  <div class="row">
    <q-select
      :options="months"
      class="col"
      dense
      label="Month"
      outlined
      v-model="month"
    ></q-select>
    <q-space class="q-pa-xs"/>
    <q-select
      :options="years"
      class="col"
      dense
      label="Year"
      outlined
      v-model="year"
    ></q-select>
  </div>
</template>

<script lang="ts">
  import { Months } from '@/constants/datetime';
  import { Component, Vue } from 'vue-property-decorator';
  import { PointModule } from '@/store';
  import { Month } from '@/types/datetime';

  @Component
  export default class PointTimeFilter extends Vue {
    months = Months;
    years = [2019, 2020];

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
  };
</script>
