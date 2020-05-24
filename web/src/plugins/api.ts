import {
  AxiosInstance,
  AxiosRequestConfig,
  AxiosResponse,
  AxiosStatic,
} from 'axios';
import { ApiResponse } from '@/types/api';
import { ProfileModule } from '@/store';

export class ApiWrapper {
  constructor(private axios: AxiosInstance) {
    this.axios.interceptors.request.use(async (config: AxiosRequestConfig) => {
      const token = await ProfileModule.getToken();
      if (token) {
        config.headers['Authorization'] = `Bearer ${token}`;
      } else {
        delete config.headers['Authorization'];
      }
      return config;
    });
  }

  async get<T = any>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<ApiResponse<T>> {
    return this.process<T>(this.axios.get<ApiResponse<T>>(url, config));
  }

  async post<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<ApiResponse<T>> {
    return this.process<T>(this.axios.post<ApiResponse<T>>(url, data, config));
  }

  async put<T = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<ApiResponse<T>> {
    return this.process<T>(this.axios.put<ApiResponse<T>>(url, data, config));
  }

  async delete<T = any>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<ApiResponse<T>> {
    return this.process<T>(this.axios.delete<ApiResponse<T>>(url, config));
  }

  async process<T = any>(
    promise: Promise<AxiosResponse<ApiResponse<T>>>
  ): Promise<ApiResponse<T>> {
    try {
      const res = await promise;
      if (!res.data) {
        return { success: false, message: res.statusText };
      }
      if (typeof res.data === 'string') {
        return { success: false, message: res.data };
      }
      if (!res.data.success) {
        return { success: false, message: res.data.message };
      }
      return { success: true, data: res.data.data };
    } catch (e) {
      return { success: false, message: e.message };
    }
  }
}

export default class VueApi {
  static installed = false;

  static install(Vue: any, axios: AxiosInstance): void {
    if (this.installed) return;
    this.installed = true;

    const apiWrapper = new ApiWrapper(axios);

    Vue.$http = axios;
    Vue.$api = apiWrapper;
    Vue.prototype.$http = axios;
    Vue.prototype.$api = apiWrapper;
  }
}

declare module 'vue/types/vue' {
  interface Vue {
    $http: AxiosStatic;
    $api: ApiWrapper;
  }

  interface VueConstructor {
    $http: AxiosStatic;
    $api: ApiWrapper;
  }
}
