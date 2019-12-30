import Vue from 'vue';
import Vuex from 'vuex';
import createPersistedState from 'vuex-persistedstate';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {},
  plugins: [
    createPersistedState({
      key: 'sheets',
    }),
  ],
});
export default store;
