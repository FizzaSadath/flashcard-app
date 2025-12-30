<script setup lang="ts">
import { useAuthStore } from "~/stores/auth";

const authStore = useAuthStore();
const isMobileMenuOpen = ref(false);

function handleLogout() {
  isMobileMenuOpen.value = false;
  authStore.logout();
}

function toggleMobileMenu() {
  isMobileMenuOpen.value = !isMobileMenuOpen.value;
}

function closeMobileMenu() {
  isMobileMenuOpen.value = false;
}
</script>

<template>
  <nav
    class="sticky top-0 z-50 border-b border-white/5 bg-gray-900/80 backdrop-blur-xl transition-all duration-300"
  >
    <div
      class="mx-auto flex h-24 w-full max-w-[1920px] items-center justify-between px-6 lg:px-8"
    >
      <NuxtLink
        to="/"
        class="flex items-center gap-4 group"
        @click="closeMobileMenu"
      >
        <img
          src="/fliplogo.png"
          alt="Flip Logo"
          class="h-10 w-auto transition-transform duration-300 group-hover:scale-110 sm:h-12"
        />
        <span
          class="text-2xl font-bold tracking-wide text-white group-hover:text-indigo-400 transition-colors sm:text-3xl"
        >
          Flip
        </span>
      </NuxtLink>

      <div class="hidden md:flex md:items-center md:gap-8 lg:gap-10">
        <template v-if="!authStore.isAuthenticated">
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
            class="relative overflow-hidden rounded-xl bg-white px-6 py-2.5 text-lg font-bold text-gray-900 transition hover:bg-gray-100"
          >
            Sign up
          </NuxtLink>
        </template>

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

          <NuxtLink
            to="/about"
            class="text-lg font-medium text-gray-300 transition hover:text-indigo-400"
            active-class="text-indigo-400 font-bold"
          >
            About
          </NuxtLink>

          <div class="h-8 w-px bg-gray-700"></div>

          <button
            @click="handleLogout"
            class="text-lg font-medium text-red-400 transition hover:text-red-300"
          >
            Logout
          </button>
        </template>
      </div>

      <div class="flex md:hidden">
        <button
          @click="toggleMobileMenu"
          type="button"
          class="inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-800 hover:text-white focus:outline-none"
        >
          <span class="sr-only">Open main menu</span>
          <svg
            v-if="!isMobileMenuOpen"
            class="h-8 w-8"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
            />
          </svg>
          <svg
            v-else
            class="h-8 w-8"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>
    </div>

    <div
      v-if="isMobileMenuOpen"
      class="md:hidden border-t border-white/5 bg-gray-900/95 backdrop-blur-xl"
    >
      <div class="space-y-1 px-4 pb-6 pt-4">
        <template v-if="!authStore.isAuthenticated">
          <NuxtLink
            to="/about"
            @click="closeMobileMenu"
            class="block rounded-md px-3 py-3 text-base font-medium text-gray-300 hover:bg-gray-800 hover:text-white"
            active-class="bg-gray-800 text-indigo-400"
          >
            About
          </NuxtLink>
          <NuxtLink
            to="/login"
            @click="closeMobileMenu"
            class="block rounded-md px-3 py-3 text-base font-medium text-gray-300 hover:bg-gray-800 hover:text-white"
          >
            Log in
          </NuxtLink>
          <NuxtLink
            to="/register"
            @click="closeMobileMenu"
            class="block mt-4 text-center rounded-xl bg-indigo-600 px-3 py-3 text-base font-bold text-white hover:bg-indigo-500"
          >
            Sign up
          </NuxtLink>
        </template>

        <template v-else>
          <NuxtLink
            to="/dashboard"
            @click="closeMobileMenu"
            class="block rounded-md px-3 py-3 text-base font-medium text-gray-300 hover:bg-gray-800 hover:text-white"
            active-class="bg-gray-800 text-indigo-400"
          >
            Dashboard
          </NuxtLink>
          <NuxtLink
            to="/stats"
            @click="closeMobileMenu"
            class="block rounded-md px-3 py-3 text-base font-medium text-gray-300 hover:bg-gray-800 hover:text-white"
            active-class="bg-gray-800 text-indigo-400"
          >
            Stats
          </NuxtLink>
          <NuxtLink
            to="/about"
            @click="closeMobileMenu"
            class="block rounded-md px-3 py-3 text-base font-medium text-gray-300 hover:bg-gray-800 hover:text-white"
            active-class="bg-gray-800 text-indigo-400"
          >
            About
          </NuxtLink>
          <button
            @click="handleLogout"
            class="block w-full text-left rounded-md px-3 py-3 text-base font-medium text-red-400 hover:bg-gray-800 hover:text-red-300"
          >
            Logout
          </button>
        </template>
      </div>
    </div>
  </nav>
</template>
