<script lang="ts">
  import { Component, Mixins, Prop } from 'vue-property-decorator';
  import { Bar, mixins } from 'vue-chartjs';
  import { ChartData, ChartOptions } from 'chart.js';

  const { reactiveData, reactiveProp } = mixins;

  @Component
  export default class BarChart extends Mixins(Bar) {
    @Prop({ type: Object }) chartData!: ChartData;
    @Prop({ type: Object }) options!: ChartOptions;

    mounted() {
      console.log(this.chartData);
      this.renderChart(this.chartData, this.options);
    }
  };
</script>
