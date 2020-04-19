interface ApiResponse<T> {
  success: boolean;
  message?: string;
  data?: T;
}

export { ApiResponse };
