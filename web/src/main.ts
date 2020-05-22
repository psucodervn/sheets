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
// eslint-disable-next-line
// @ts-ignore
import VueAuthenticate from 'vue-authenticate';
import { vueAuthenticateOptions } from '@/modules/profile/auth';

Vue.config.productionTip = false;

Vue.use(Navigation);
Vue.component('navigation-bar', NavigationBar);

Vue.use(
  VueApi,
  axios.create({
    baseURL: '/api',
    validateStatus: () => true,
  })
);

Vue.use(VueAuthenticate, vueAuthenticateOptions);
// it's magic
Vue.prototype.$auth;

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app');
