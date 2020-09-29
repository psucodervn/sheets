import Vue from 'vue';
import '@/quasar';
import '@/styles/app.scss';

import App from './App.vue';
import './registerServiceWorker';
import router from '@/router';
import store from '@/store';

import Navigation from '@/plugins/navigation';
import NavigationBar from '@/components/NavigationBar.vue';
import axios from 'axios';
import VueApi from '@/plugins/api';
import GAuth from 'vue-google-oauth2';
import { gAuthOptions } from '@/modules/profile/auth';
import AddButton from '@/components/AddButton.vue';
import DatePicker from '@/components/DatePicker.vue';

Vue.config.productionTip = false;

Vue.use(Navigation);
Vue.component('navigation-bar', NavigationBar);
Vue.component('add-button', AddButton);
Vue.component('date-picker', DatePicker);

Vue.use(
  VueApi,
  axios.create({
    baseURL: '/api',
    validateStatus: () => true,
  })
);

Vue.use(GAuth, gAuthOptions);

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
