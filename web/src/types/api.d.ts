interface ApiResponse<T> {
  success: boolean;
  message?: string;
  data?: T;
  status?: number;
}

export { ApiResponse };
