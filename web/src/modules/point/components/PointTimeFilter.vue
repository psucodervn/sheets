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
  </div>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator';
  import { PointModule } from '@/store';
  import { Month } from '@/types/datetime';
  import formatter from '@/utils/formatter';

  interface MonthItem {
    label: string;
    value: Month;
  }

  @Component
  export default class PointTimeFilter extends Vue {
    months: MonthItem[] = [];

    get month(): MonthItem {
      return {
        value: PointModule.month,
        label: formatter.month(PointModule.month),
      };
    }

    set month(item: MonthItem) {
      PointModule.setMonth(item.value);
    }

    mounted() {
      let prev: Month = { month: 1, year: 2020 };
      const endYear = new Date().getUTCFullYear();
      const endMonth = new Date().getUTCMonth() + 1;
      this.months.push({ value: prev, label: formatter.month(prev) });
      while (true) {
        if (prev.year >= endYear && prev.month >= endMonth) {
          break;
        }
        prev = { ...prev };
        prev.month++;
        if (prev.month > 12) {
          prev.month = 1;
          prev.year++;
        }
        this.months.push({ value: prev, label: formatter.month(prev) });
      }
      this.months.reverse();
    }
  }
</script>
