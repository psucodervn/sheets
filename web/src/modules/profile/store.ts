import {
  Action,
  Module,
  MutationAction,
  VuexModule,
} from 'vuex-module-decorators';
import { Vue } from 'vue-property-decorator';

@Module({
  name: 'profile',
  namespaced: true,
})
export default class ProfileStore extends VuexModule {
  isAuthenticated = false;

  @MutationAction({ mutate: ['isAuthenticated'], rawError: true })
  async login(payload: { email: string; password: string }) {
    const res = await Vue.prototype.$auth.login(payload, {
      validateStatus: () => true,
    });
    if (!res.data) {
      throw new Error(res.statusText);
    }
    if (!res.data.success) {
      throw new Error(res.data.message);
    }
    return { isAuthenticated: Vue.prototype.$auth.isAuthenticated() };
  }

  @MutationAction({ mutate: ['isAuthenticated'], rawError: true })
  async authenticate(param: { provider: string }) {
    const res = await Vue.prototype.$auth.authenticate(param.provider, {});
    return { isAuthenticated: Vue.prototype.$auth.isAuthenticated() };
  }

  @MutationAction({ mutate: ['isAuthenticated'], rawError: true })
  async logout() {
    try {
      await Vue.prototype.$auth.logout();
    } catch (e) {
      console.log('$auth logout err:', e.message);
    }
    return { isAuthenticated: Vue.prototype.$auth.isAuthenticated() };
  }

  @Action({ rawError: true })
  async getToken() {
    return Vue.prototype.$auth.getToken();
  }

  get token() {
    return Vue.prototype.$auth.getToken();
  }
}