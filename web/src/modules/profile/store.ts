import {
  Action,
  Module,
  MutationAction,
  VuexModule,
} from 'vuex-module-decorators';
import { Vue } from 'vue-property-decorator';
import { User } from '@/modules/profile/models/user';
import { AuthToken } from '@/modules/profile/dtos/auth';

@Module({
  name: 'profile',
  namespaced: true,
})
export default class ProfileStore extends VuexModule {
  currentUser: User | null = null;
  token: string | null = null;

  @MutationAction({ mutate: ['token'], rawError: true })
  async authenticate(param: { provider: string }) {
    if (param.provider !== 'google') {
      throw new Error('Only Google sign-in is supported now.');
    }
    let authCode;
    try {
      authCode = await Vue.prototype.$gAuth.getAuthCode();
    } catch (e) {
      if (e.error === 'popup_closed_by_user') throw new Error('');
      throw new Error(`Get authorization code failed: ${e.error || e.message}`);
    }
    const res = await Vue.$api.post<AuthToken>('/auth/google', {
      code: authCode,
      redirectUri: 'postmessage',
    });
    if (!res.data || !res.success) {
      throw new Error(`Authenticate failed: ${res.message}`);
    }
    return { token: res.data.accessToken };
  }

  @MutationAction({ mutate: ['token', 'currentUser'], rawError: true })
  async logout() {
    try {
      await Vue.prototype.$gAuth.signOut();
    } catch (e) {
      console.log('$auth logout err:', e.message);
    }
    return { token: null, currentUser: null };
  }

  @Action({ rawError: true })
  async getToken() {
    return this.token;
  }

  @MutationAction({ mutate: ['currentUser'], rawError: true })
  async fetchMe() {
    const res = await Vue.$api.get<User>('/auth/me');
    if (!res.data || !res.success) {
      throw new Error(res.message);
    }
    return { currentUser: res.data };
  }
}
