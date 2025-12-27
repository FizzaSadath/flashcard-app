import { useAuthStore } from "~/stores/auth";

export const useAPI = <T>(url: string, options: any = {}) => {
  const config = useRuntimeConfig();
  const authStore = useAuthStore();

  const defaults = {
    baseURL: config.public.apiBase,
    headers: authStore.token
      ? { Authorization: `Bearer ${authStore.token}` }
      : {},
  };

  const params = {
    ...defaults,
    ...options,
    headers: { ...defaults.headers, ...options.headers },
  };

  return useFetch<T>(url, params);
};
