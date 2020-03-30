<template>
  <q-page class="home q-pa-md">
    <h6>Balance</h6>
    <horizontal-bar-chart
      :chart-data="chartData"
      :options="chartOptions"
      v-if="isMobile"
    />
    <bar-chart :chart-data="chartData" :options="chartOptions" v-else/>
  </q-page>
</template>

<script lang="ts">
  // @ is an alias to /src
  import { Component, Vue } from 'vue-property-decorator';
  import { IUser } from '@/model/user';
  import BarChart from '@/components/BarChart.vue';
  import { ChartData, ChartOptions, CommonAxe } from 'chart.js';
  import HorizontalBarChart from '@/components/HorizontalBarChart.vue';
  import { UserModule } from '@/store';

  @Component({
    components: { BarChart, HorizontalBarChart },
  })
  export default class Balance extends Vue {
    isMobile: boolean = false;
    max: number = 0;
    min: number = 0;

    get chartData(): ChartData {
      const labels = this.users.map((u: IUser) => (this.isMobile ? ' ' : u.name));
      const values = this.users.map((u: IUser) => ~~(u.balance.value / 1000));
      this.max = Math.max(1000, ...values) + 50;
      this.min = Math.min(-1000, ...values) - 50;
      const colors = values.map(val => {
        let percent;
        if (val >= 0) {
          percent = (val / this.max) * 50 + 50;
        } else {
          percent = 50 - (val / this.min) * 50;
        }
        return this.percentToColor(percent);
      });

      return {
        labels: labels,
        datasets: [
          {
            label: 'Balance',
            data: values,
            borderWidth: 1,
            backgroundColor: colors.map(c => c.fill),
            borderColor: colors.map(c => c.border),
          },
        ],
      };
    }

    get chartOptions(): ChartOptions {
      const axes: CommonAxe[] = [
        {
          display: true,
          scaleLabel: {
            display: true,
            labelString: 'Balance (x1000 vnđ)',
          },
          gridLines: {
            display: true,
          },
          ticks: {
            suggestedMin: this.min,
            suggestedMax: this.max,
            stepSize: 200,
          },
        },
      ];

      const opts: ChartOptions = {
        title: { display: false },
        legend: { display: false },
        responsive: true,
        scales: {},
        plugins: {
          datalabels: {
            formatter: value => {
              return Number(value).toLocaleString();
            },
            labels: {
              user: {
                align: this.isMobile ? 'start' : 'end',
                font: {
                  weight: 'bold',
                },
                formatter: (value, ctx) => {
                  return (
                    this.users[ctx.dataIndex].name +
                    (this.isMobile ? ': ' + value : '')
                  );
                },
              },
              balance: {
                display: !this.isMobile,
                font: {
                  weight: 'bold',
                },
                align: this.isMobile ? 'end' : 'start',
              },
            },
          },
        },
      };

      if (!this.isMobile) {
        opts.scales!.yAxes = axes;
      } else {
        opts.scales!.xAxes = axes;
      }

      return opts;
    }

    get users(): IUser[] {
      return UserModule.users.sort(
        (a: IUser, b: IUser) => -(a.balance.value - b.balance.value),
      );
    }

    async mounted() {
      this.handleResize();
      try {
        await UserModule.fetchUsers();
      } catch (e) {
        console.log(e.message);
      }
    }

    created() {
      window.addEventListener('resize', this.handleResize);
    }

    destroyed() {
      window.removeEventListener('resize', this.handleResize);
    }

    async handleResize() {
      this.isMobile = window.innerWidth < 800;
    }

    percentToColor(percent: number) {
      let r,
        g,
        b = 0;
      if (percent < 50) {
        r = 230;
        g = Math.round((5.1 * percent) / 2);
      } else {
        g = 190;
        r = Math.round(510 - 5.1 * percent);
      }
      let h = r * 0x10000 + g * 0x100 + b;
      return {
        fill: `rgba(${r}, ${g}, ${b}, 0.4)`,
        border: `rgba(${r}, ${g}, ${b}, 0.9)`,
      };
      // return '#' + ('000000' + h.toString(16)).slice(-6);
    }
  }
</script>

<style>
  canvas {
    margin: auto;
    max-width: 900px;
    max-height: 85vh;
  }
</style>
