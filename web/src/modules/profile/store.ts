import {
  Action,
  Module,
  Mutation,
  MutationAction,
  VuexModule,
} from 'vuex-module-decorators';
import { Vue } from 'vue-property-decorator';
import { User } from '@/modules/profile/models/user';
import { AuthToken } from '@/modules/profile/dtos/auth';
import { AxiosResponse } from 'axios';

@Module({
  name: 'profile',
  namespaced: true,
})
export default class ProfileStore extends VuexModule {
  currentUser: User | null = null;
  token: string | null = null;
  refreshToken: string | null = null;

  @Action({ rawError: true })
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
    this.setTokens(res.data);
    return { token: res.data.accessToken, refreshToken: res.data.refreshToken };
  }

  @Mutation
  setTokens(tokens: AuthToken) {
    this.token = tokens.accessToken;
    this.refreshToken = tokens.refreshToken;
    console.log('setTokens done');
  }

  @MutationAction({
    mutate: ['token', 'refreshToken', 'currentUser'],
    rawError: true,
  })
  async logout() {
    try {
      await Vue.prototype.$gAuth.signOut();
    } catch (e) {
      console.log('$auth logout err:', e.message);
    }
    return { token: null, refreshToken: null, currentUser: null };
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

  @Action({ rawError: true })
  async generateTelegramLink() {
    const res = await Vue.$api.post<string>('/auth/telegram');
    if (!res.data || !res.success) {
      throw new Error(res.message);
    }
    return res.data;
  }

  @Action({ rawError: true })
  async postRefreshToken(response: AxiosResponse) {
    try {
      const res = await Vue.$api.post<AuthToken>(
        '/auth/refresh',
        {},
        {
          headers: { Authorization: `Bearer ${this.refreshToken}` },
        }
      );
      if (!res.data || !res.success) return response;
      this.setTokens(res.data);
      response.config.headers['Authorization'] = res.data.accessToken;
      return Vue.$api.call(response.config);
    } catch (e) {
      return response;
    }
  }
}
