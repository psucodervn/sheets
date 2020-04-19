import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from "vuex-persistedstate";
import { UserStore } from "@/store/user";
import { getModule } from "vuex-module-decorators";
import { PointStore } from "@/store/point";
import BalanceStore from "@/modules/balance/store";

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    user: UserStore,
    point: PointStore,
    balance: BalanceStore
  },
  plugins: [
    createPersistedState({
      key: "sheets"
    })
  ]
});

export default store;
export const UserModule = getModule(UserStore, store);
export const PointModule = getModule(PointStore, store);
export const BalanceModule = getModule(BalanceStore, store);
