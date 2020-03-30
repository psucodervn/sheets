import Chart from 'chart.js';
import ChartDataLabels from 'chartjs-plugin-datalabels';
import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';
import './quasar';

Chart.plugins.register(ChartDataLabels);
// @ts-ignore
Chart.helpers.merge(Chart.defaults.global.plugins.datalabels, {
  align: 'end',
  anchor: 'end',
});


Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
