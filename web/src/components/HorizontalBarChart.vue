<script lang="ts">
  import { Component, Mixins, Prop } from 'vue-property-decorator';
  import { HorizontalBar, mixins } from 'vue-chartjs';
  import { ChartData, ChartOptions } from 'chart.js';

  const { reactiveProp } = mixins;

  @Component
  export default class HorizontalBarChart extends Mixins(HorizontalBar, reactiveProp) {
    @Prop({ type: Object }) options!: ChartOptions;

    mounted() {
      this.renderChart(this.chartData, this.options);
    }
  };
</script>
