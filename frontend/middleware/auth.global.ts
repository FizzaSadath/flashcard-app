import { useAuthStore } from "~/stores/auth";

export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStore();

  const publicRoutes = ["/login", "/register", "/", "/about"];

  if (authStore.isAuthenticated) {
    if (to.path === "/login" || to.path === "/register") {
      return navigateTo("/dashboard");
    }
  } else {
    if (!publicRoutes.includes(to.path)) {
      return navigateTo("/login");
    }
  }
});
