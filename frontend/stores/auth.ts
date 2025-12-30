import { defineStore } from "pinia";
import { useToastStore } from "./toast";
function parseJwt(token: string) {
  if (typeof window === "undefined") return null;

  try {
    const parts = token.split(".");

    if (parts.length < 2) {
      return null;
    }

    const base64Url = parts[1];

    if (!base64Url) {
      return null;
    }

    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    const jsonPayload = decodeURIComponent(
      window
        .atob(base64)
        .split("")
        .map(function (c) {
          return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2);
        })
        .join("")
    );
    return JSON.parse(jsonPayload);
  } catch (e) {
    return null;
  }
}

export const useAuthStore = defineStore("auth", () => {
  const token = useCookie<string | null>("token");
  const toast = useToastStore();

  const user = computed(() => {
    if (!token.value) return null;
    return parseJwt(token.value);
  });

  const isAuthenticated = computed(() => !!token.value);

  async function login(email: string, password: string) {
    const config = useRuntimeConfig();

    try {
      const { data, error } = await useFetch(
        `${config.public.apiBase}/auth/login`,
        {
          method: "POST",
          body: { email, password },
        }
      );

      if (error.value) {
        throw new Error(error.value.data?.error || "Login failed");
      }

      const response = data.value as any;
      token.value = response.token;

      toast.add("Welcome back!", "success");

      return navigateTo("/dashboard");
    } catch (err: any) {
      toast.add(err.message, "error");
    }
  }

  async function register(
    email: string,
    username: string,
    password: string,
    confirmPassword: string
  ) {
    const config = useRuntimeConfig();

    try {
      const { error } = await useFetch(
        `${config.public.apiBase}/auth/register`,
        {
          method: "POST",
          body: { email, username, password, confirmPassword },
        }
      );

      if (error.value) {
        throw new Error(error.value.data?.error || "Registration failed");
      }

      toast.add("Registration successful! Please login.", "success");
      return navigateTo("/login");
    } catch (err: any) {
      toast.add(err.message, "error");
    }
  }

  function logout() {
    token.value = null;
    toast.add("Logged out successfully", "info");
    return navigateTo("/login");
  }

  return { token, user, isAuthenticated, login, register, logout };
});
