import Vue from 'vue';
import '@/quasar';
import '@/styles/app.scss';

import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';

import Navigation from '@/plugins/navigation';
import NavigationBar from '@/components/NavigationBar.vue';
import VueAxios from 'vue-axios';
import axios from 'axios';

Vue.config.productionTip = false;

Vue.use(Navigation);
Vue.component('navigation-bar', NavigationBar);

Vue.use(
  VueAxios,
  axios.create({
    baseURL: '/api',
    validateStatus: () => true,
  })
);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
