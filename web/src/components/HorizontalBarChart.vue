<script lang="ts">
  import { Component, Mixins, Prop } from 'vue-property-decorator';
  import { HorizontalBar } from 'vue-chartjs';
  import { ChartData, ChartOptions } from 'chart.js';

  @Component
  export default class HorizontalBarChart extends Mixins(HorizontalBar) {
    @Prop({ type: Object }) data!: ChartData;
    @Prop({ type: Object }) options!: ChartOptions;

    mounted() {
      this.renderChart(this.data, this.options);
    }
  };
</script>
