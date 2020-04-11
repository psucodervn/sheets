import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';
import '@/quasar';
import '@/styles/fonts.scss';
import Navigation from '@/plugins/navigation';

Vue.config.productionTip = false;
Vue.use(Navigation);

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
