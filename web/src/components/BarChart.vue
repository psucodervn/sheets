<script lang="ts">
  import { Component, Mixins, Prop } from 'vue-property-decorator';
  import { Bar, mixins } from 'vue-chartjs';
  import { ChartOptions } from 'chart.js';

  const { reactiveProp } = mixins;

  @Component
  export default class BarChart extends Mixins(Bar, reactiveProp) {
    @Prop({ type: Object }) options!: ChartOptions;

    mounted() {
      this.renderChart(this.chartData, this.options);
    }
  };
</script>
