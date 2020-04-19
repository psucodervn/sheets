import { AxiosRequestConfig, AxiosResponse } from "axios";
import { Vue } from "vue-property-decorator";
import { ApiResponse } from "@/types/api";

const process = async <T = any>(
  promise: Promise<AxiosResponse<ApiResponse<T>>>
): Promise<ApiResponse<T>> => {
  try {
    const res = await promise;
    if (!res.data) {
      return { success: false, message: res.statusText };
    }
    if (!res.data.success) {
      return { success: false, message: res.data.message };
    }
    return { success: true, data: res.data.data };
  } catch (e) {
    return { success: false, message: e.message };
  }
};

const get = async <T = any>(
  url: string,
  config?: AxiosRequestConfig
): Promise<ApiResponse<T>> => {
  return process<T>(Vue.axios.get<ApiResponse<T>>(url, config));
};

const post = async <T = any>(
  url: string,
  data?: any,
  config?: AxiosRequestConfig
): Promise<ApiResponse<T>> => {
  return process<T>(Vue.axios.post<ApiResponse<T>>(url, data, config));
};

const api = {
  get,
  post
};

export default api;
