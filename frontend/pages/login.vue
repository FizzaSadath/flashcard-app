<script setup lang="ts">
import { useAuthStore } from "~/stores/auth";

const email = ref("");
const password = ref("");
const authStore = useAuthStore();
const loading = ref(false);

useHead({ title: "Login - Flip" });

const handleLogin = async () => {
  loading.value = true;
  await authStore.login(email.value, password.value);
  loading.value = false;
};
</script>

<template>
  <div class="flex min-h-[calc(100vh-6rem)] bg-gray-900">
    <div
      class="flex flex-1 flex-col justify-center px-8 sm:px-10 lg:flex-none lg:w-[500px] xl:w-[600px] bg-gray-950 border-r border-white/5"
    >
      <div class="w-full max-w-sm lg:max-w-md">
        <div class="text-left">
          <h2 class="mt-6 text-3xl font-bold tracking-tight text-white">
            Welcome back
          </h2>
          <p class="mt-2 text-sm text-gray-400">
            Sign in to access your decks.
          </p>
        </div>

        <div class="mt-8">
          <form class="space-y-6" @submit.prevent="handleLogin">
            <div>
              <label
                for="email"
                class="block text-sm font-medium leading-6 text-gray-300 mb-2"
              >
                Email address
              </label>
              <div class="relative">
                <input
                  v-model="email"
                  id="email"
                  name="email"
                  type="email"
                  autocomplete="email"
                  required
                  class="block w-full rounded-lg border-0 bg-gray-800/50 px-4 py-3 text-white shadow-sm ring-1 ring-inset ring-gray-700 placeholder:text-gray-500 focus:ring-2 focus:ring-inset focus:ring-indigo-500 text-sm sm:leading-6 transition-all"
                  placeholder="you@example.com"
                />
              </div>
            </div>

            <div>
              <label
                for="password"
                class="block text-sm font-medium leading-6 text-gray-300 mb-2"
              >
                Password
              </label>
              <div class="relative">
                <input
                  v-model="password"
                  id="password"
                  name="password"
                  type="password"
                  autocomplete="current-password"
                  required
                  class="block w-full rounded-lg border-0 bg-gray-800/50 px-4 py-3 text-white shadow-sm ring-1 ring-inset ring-gray-700 placeholder:text-gray-500 focus:ring-2 focus:ring-inset focus:ring-indigo-500 text-sm sm:leading-6 transition-all"
                  placeholder="••••••••"
                />
              </div>
            </div>

            <div class="pt-2">
              <button
                type="submit"
                :disabled="loading"
                class="group relative flex w-full justify-center overflow-hidden rounded-lg bg-indigo-600 px-4 py-3 text-sm font-semibold text-white shadow-lg hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 disabled:opacity-70 disabled:cursor-not-allowed transition-all"
              >
                <div
                  class="absolute inset-0 bg-white/20 translate-y-full transition-transform group-hover:translate-y-0"
                ></div>
                <span class="relative flex items-center gap-2">
                  <span v-if="loading">
                    <svg
                      class="h-4 w-4 animate-spin text-white"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <circle
                        class="opacity-25"
                        cx="12"
                        cy="12"
                        r="10"
                        stroke="currentColor"
                        stroke-width="4"
                      ></circle>
                      <path
                        class="opacity-75"
                        fill="currentColor"
                        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                      ></path>
                    </svg>
                  </span>
                  {{ loading ? "Signing in..." : "Sign In" }}
                </span>
              </button>
            </div>

            <div class="text-sm text-center text-gray-400">
              <p>
                New to Flip?
                <NuxtLink
                  to="/register"
                  class="font-semibold text-indigo-400 hover:text-indigo-300 transition-colors ml-1"
                >
                  Create an account
                </NuxtLink>
              </p>
            </div>
          </form>
        </div>
      </div>
    </div>

    <div class="relative hidden w-0 flex-1 lg:block">
      <img
        class="absolute inset-0 h-full w-full object-cover"
        src="/auth_bg.jpg"
        alt="Abstract background pattern"
      />

      <div
        class="absolute inset-0 bg-gradient-to-t from-gray-950 via-gray-900/60 to-transparent"
      ></div>

      <div class="absolute bottom-0 left-0 right-0 p-16 text-white z-10">
        <div class="max-w-2xl">
          <h2 class="text-4xl font-bold tracking-tight mb-6 drop-shadow-lg">
            Master your memory
          </h2>
          <p class="text-lg text-gray-100 mb-8 leading-relaxed">
            Flip uses advanced algorithms to optimize your learning schedule.
            Connect the dots and retain information effortlessly.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>
