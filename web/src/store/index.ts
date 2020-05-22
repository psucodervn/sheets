import Vue from 'vue';
import Vuex from 'vuex';
import createPersistedState from 'vuex-persistedstate';
import { getModule } from 'vuex-module-decorators';
import PointStore from '@/modules/point/store';
import BalanceStore from '@/modules/balance/store';
import ProfileStore from '@/modules/profile/store';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    point: PointStore,
    balance: BalanceStore,
    profile: ProfileStore,
  },
  plugins: [
    createPersistedState({
      key: 'sheets',
    }),
  ],
});

export default store;
export const PointModule = getModule(PointStore, store);
export const BalanceModule = getModule(BalanceStore, store);
export const ProfileModule = getModule(ProfileStore, store);
