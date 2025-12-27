import { defineStore } from "pinia";

export const useAuthStore = defineStore("auth", () => {
  // Loads cookies from browser
  const token = useCookie<string | null>("token");
  const user = useState("user", () => null);

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

      // Save token to cookie
      const response = data.value as any;
      token.value = response.token;

      return navigateTo("/dashboard");
    } catch (err: any) {
      alert(err.message);
    }
  }

  async function register(email: string, password: string) {
    const config = useRuntimeConfig();

    try {
      const { error } = await useFetch(
        `${config.public.apiBase}/auth/register`,
        {
          method: "POST",
          body: { email, password },
        }
      );

      if (error.value) {
        throw new Error(error.value.data?.error || "Registration failed");
      }

      alert("Registration successful! Please login.");
      return navigateTo("/login");
    } catch (err: any) {
      alert(err.message);
    }
  }

  function logout() {
    token.value = null;
    user.value = null;
    return navigateTo("/login");
  }

  return { token, user, isAuthenticated, login, register, logout };
});
