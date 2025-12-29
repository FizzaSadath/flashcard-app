<script setup lang="ts">
import { useAuthStore } from "~/stores/auth";
const authStore = useAuthStore();
function handleLogout() {
  authStore.logout();
}
</script>

<template>
  <nav
    class="sticky top-0 z-50 border-b border-white/5 bg-gray-900/80 backdrop-blur-xl transition-all duration-300"
  >
    <div class="flex h-24 w-full items-center justify-between px-8 sm:px-10">
      <!-- LOGO -->
      <NuxtLink to="/" class="flex items-center gap-5 group">
        <img
          src="/fliplogo.png"
          alt="Flip Logo"
          class="h-12 w-auto transition-transform duration-300 group-hover:scale-110"
        />
        <span
          class="text-3xl font-bold tracking-wide text-white group-hover:text-indigo-400 transition-colors"
          >Flip</span
        >
      </NuxtLink>

      <div class="flex items-center gap-10">
        <!-- GUEST VIEW -->
        <template v-if="!authStore.isAuthenticated">
          <!-- NEW: About Link -->
          <NuxtLink
            to="/about"
            class="text-lg font-medium text-gray-300 transition hover:text-white"
            active-class="text-indigo-400 font-bold"
          >
            About
          </NuxtLink>

          <NuxtLink
            to="/login"
            class="text-lg font-medium text-gray-300 transition hover:text-white"
          >
            Log in
          </NuxtLink>
          <NuxtLink
            to="/register"
            class="relative overflow-hidden rounded-xl bg-white px-8 py-3 text-lg font-bold text-gray-900 transition hover:bg-gray-100"
          >
            Sign up
          </NuxtLink>
        </template>

        <!-- LOGGED IN VIEW -->
        <template v-else>
          <NuxtLink
            to="/dashboard"
            class="text-lg font-medium text-gray-300 transition hover:text-indigo-400"
            active-class="text-indigo-400 font-bold"
          >
            Dashboard
          </NuxtLink>

          <NuxtLink
            to="/stats"
            class="text-lg font-medium text-gray-300 transition hover:text-indigo-400"
            active-class="text-indigo-400 font-bold"
          >
            Stats
          </NuxtLink>

          <!-- NEW: About Link -->
          <NuxtLink
            to="/about"
            class="text-lg font-medium text-gray-300 transition hover:text-indigo-400"
            active-class="text-indigo-400 font-bold"
          >
            About
          </NuxtLink>

          <div class="h-10 w-px bg-gray-700"></div>

          <button
            @click="handleLogout"
            class="text-lg font-medium text-red-400 transition hover:text-red-300"
          >
            Logout
          </button>
        </template>
      </div>
    </div>
  </nav>
</template>
