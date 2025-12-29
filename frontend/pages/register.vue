<script setup lang="ts">
import { useAuthStore } from "~/stores/auth";

const email = ref("");
const username = ref("");
const password = ref("");
const confirmPassword = ref("");
const authStore = useAuthStore();
const loading = ref(false);

useHead({ title: "Register - Flip" });

const handleRegister = async () => {
  loading.value = true;
  await authStore.register(
    email.value,
    username.value,
    password.value,
    confirmPassword.value
  );
  loading.value = false;
};
</script>

<template>
  <div class="flex flex-1 flex-col items-center justify-center px-4 py-12">
    <div class="w-full max-w-md">
      <div class="text-center mb-8">
        <h2 class="text-3xl font-bold tracking-tight text-white sm:text-4xl">
          Create an account
        </h2>
        <p class="mt-3 text-gray-400">
          Join Flip today to start mastering your decks.
        </p>
      </div>

      <div
        class="rounded-2xl border border-white/10 bg-gray-900/50 p-8 shadow-2xl backdrop-blur-xl"
      >
        <form class="space-y-5" @submit.prevent="handleRegister">
          <div class="space-y-1">
            <label for="email" class="block text-sm font-medium text-gray-300"
              >Email address</label
            >
            <div class="relative">
              <input
                v-model="email"
                id="email"
                name="email"
                type="email"
                required
                class="block w-full rounded-lg border border-gray-700 bg-gray-800/50 px-4 py-3 text-white placeholder-gray-500 transition-all focus:border-indigo-500 focus:bg-gray-800 focus:ring-2 focus:ring-indigo-500/20 outline-none"
                placeholder="you@example.com"
              />
            </div>
          </div>

          <div class="space-y-1">
            <label
              for="username"
              class="block text-sm font-medium text-gray-300"
              >Username</label
            >
            <div class="relative">
              <input
                v-model="username"
                id="username"
                name="username"
                type="text"
                required
                class="block w-full rounded-lg border border-gray-700 bg-gray-800/50 px-4 py-3 text-white placeholder-gray-500 transition-all focus:border-indigo-500 focus:bg-gray-800 focus:ring-2 focus:ring-indigo-500/20 outline-none"
                placeholder="username"
              />
            </div>
          </div>

          <div class="space-y-1">
            <label
              for="password"
              class="block text-sm font-medium text-gray-300"
              >Password</label
            >
            <div class="relative">
              <input
                v-model="password"
                id="password"
                name="password"
                type="password"
                required
                class="block w-full rounded-lg border border-gray-700 bg-gray-800/50 px-4 py-3 text-white placeholder-gray-500 transition-all focus:border-indigo-500 focus:bg-gray-800 focus:ring-2 focus:ring-indigo-500/20 outline-none"
                placeholder="••••••••"
              />
            </div>
          </div>

          <div class="space-y-1">
            <label
              for="confirmPassword"
              class="block text-sm font-medium text-gray-300"
              >Confirm Password</label
            >
            <div class="relative">
              <input
                v-model="confirmPassword"
                id="confirmPassword"
                name="confirmPassword"
                type="password"
                required
                class="block w-full rounded-lg border border-gray-700 bg-gray-800/50 px-4 py-3 text-white placeholder-gray-500 transition-all focus:border-indigo-500 focus:bg-gray-800 focus:ring-2 focus:ring-indigo-500/20 outline-none"
                placeholder="••••••••"
              />
            </div>
          </div>

          <div class="pt-2">
            <button
              type="submit"
              :disabled="loading"
              class="group relative flex w-full justify-center overflow-hidden rounded-lg bg-indigo-600 px-4 py-3 text-sm font-bold text-white shadow-lg shadow-indigo-500/30 transition-all hover:bg-indigo-500 hover:scale-[1.02] active:scale-[0.98] disabled:opacity-70 disabled:hover:scale-100"
            >
              <div
                class="absolute inset-0 bg-white/20 translate-y-full transition-transform group-hover:translate-y-0"
              ></div>
              <span class="relative flex items-center gap-2">
                <span
                  v-if="loading"
                  class="animate-spin h-4 w-4 border-2 border-white border-t-transparent rounded-full"
                ></span>
                {{ loading ? "Creating Account..." : "Create Account" }}
              </span>
            </button>
          </div>

          <div class="text-center text-sm">
            <p class="text-gray-500">
              Already registered?
              <NuxtLink
                to="/login"
                class="font-semibold text-indigo-400 hover:text-indigo-300 transition-colors"
              >
                Log in
              </NuxtLink>
            </p>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
