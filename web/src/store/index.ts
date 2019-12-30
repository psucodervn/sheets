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
      // getState: (key, storage) => {
      //   let data;
      //   try {
      //     data = JSON.parse(storage.getItem(key));
      //     if (!data.user) {
      //       data.user = {
      //         users: [],
      //       };
      //     }
      //     return data;
      //   } catch (e) {
      //     data = {
      //       user: {
      //         users: [],
      //       },
      //     };
      //   }
      //   return data;
      // },
    }),
  ],
});
export default store;
