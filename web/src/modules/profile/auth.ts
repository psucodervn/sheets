import { AuthenticateOptions } from '@/types/vue-authenticate';
import { AxiosRequestConfig, AxiosResponse } from 'axios';
import { ProfileModule } from '@/store';
import router from '@/router';

export const vueAuthenticateOptions: AuthenticateOptions = {
  tokenPath: 'data.accessToken',
  tokenPrefix: 'sheets',
  providers: {
    google: {
      clientId:
        '406743512267-59l14i8s6h7dk8o48he8i060gihfama2.apps.googleusercontent.com',
      redirectUri: window.location.origin + '/auth/callback',
    },
  },
  bindRequestInterceptor: function(VueAuth) {
    const tokenHeader = VueAuth.options.tokenHeader;
    VueAuth.$http.interceptors.request.use((config: AxiosRequestConfig) => {
      if (VueAuth.isAuthenticated()) {
        config.headers[tokenHeader] = [
          VueAuth.options.tokenType,
          VueAuth.getToken(),
        ].join(' ');
      } else {
        delete config.headers[tokenHeader];
      }
      return config;
    });

    VueAuth.$http.interceptors.response.use(
      async (response: AxiosResponse): Promise<AxiosResponse> => {
        if (response.status !== 401) {
          return response;
        }
        await ProfileModule.logout();

        if (router.currentRoute.meta.requiresNotAuth) {
          return response;
        }
        router
          .push('/auth/login?redirect=' + router.currentRoute.fullPath)
          .catch(console.log);
        return response;
      }
    );
  },
};
