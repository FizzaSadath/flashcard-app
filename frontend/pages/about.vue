<script setup lang="ts">
useHead({ title: "About - Flip" });
import { useAuthStore } from "~/stores/auth";
const authStore = useAuthStore();
</script>

<template>
  <div class="min-h-[calc(100vh-6rem)] p-6 sm:p-10">
    <div class="mx-auto max-w-4xl space-y-12">
      <div class="text-center space-y-4">
        <h1
          class="text-4xl font-extrabold text-white tracking-tight sm:text-5xl"
        >
          Master your memory with <span class="text-indigo-400">Flip</span>.
        </h1>
        <p class="text-lg text-gray-400 max-w-2xl mx-auto leading-relaxed">
          Flip isn't just a flashcard app. It uses the
          <span class="text-gray-200 font-semibold"
            >SuperMemo-2 (SM-2) Algorithm</span
          >
          to calculate the perfect time for you to review a card, ensuring you
          learn efficiently without forgetting.
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div
          class="rounded-2xl border border-white/10 bg-gray-900/50 p-6 shadow-xl backdrop-blur-sm"
        >
          <div
            class="h-10 w-10 rounded-lg bg-indigo-500/20 text-indigo-400 flex items-center justify-center mb-4"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 6.042A8.967 8.967 0 006 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 016 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 016-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0018 18a8.967 8.967 0 00-6 2.292m0-14.25v14.25"
              />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-white mb-2">Create Decks</h3>
          <p class="text-gray-400 text-sm">
            Organize your learning into specific topics. Add cards with
            questions on the front and answers on the back.
          </p>
        </div>

        <div
          class="rounded-2xl border border-white/10 bg-gray-900/50 p-6 shadow-xl backdrop-blur-sm"
        >
          <div
            class="h-10 w-10 rounded-lg bg-emerald-500/20 text-emerald-400 flex items-center justify-center mb-4"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M3.75 13.5l10.5-11.25L12 10.5h8.25L9.75 21.75 12 13.5H3.75z"
              />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-white mb-2">Spaced Repetition</h3>
          <p class="text-gray-400 text-sm">
            The app schedules cards based on how well you know them. Hard cards
            appear often; easy cards appear rarely.
          </p>
        </div>

        <div
          class="rounded-2xl border border-white/10 bg-gray-900/50 p-6 shadow-xl backdrop-blur-sm"
        >
          <div
            class="h-10 w-10 rounded-lg bg-rose-500/20 text-rose-400 flex items-center justify-center mb-4"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 013 19.875v-6.75zM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V8.625zM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V4.125z"
              />
            </svg>
          </div>
          <h3 class="text-lg font-bold text-white mb-2">Track Progress</h3>
          <p class="text-gray-400 text-sm">
            Visualize your retention with detailed charts showing New, Learning,
            and Mature cards.
          </p>
        </div>
      </div>

      <div
        class="rounded-2xl border border-white/10 bg-gray-900/50 p-8 shadow-xl backdrop-blur-sm"
      >
        <h2
          class="text-2xl font-bold text-white mb-6 border-b border-gray-800 pb-4"
        >
          How to Grade Yourself
        </h2>
        <p class="text-gray-400 mb-8">
          When studying, be honest! Your grade determines when you will see the
          card next.
        </p>

        <div class="space-y-4">
          <div
            class="grid grid-cols-1 sm:grid-cols-[100px_1fr] gap-4 items-center p-3 rounded-lg hover:bg-white/5 transition"
          >
            <span
              class="text-center py-2 px-4 rounded-lg bg-rose-950/50 text-rose-300 border border-rose-900 font-bold text-sm"
            >
              0 - 2
            </span>
            <div>
              <h4 class="text-rose-400 font-bold">Fail (Reset)</h4>
              <p class="text-sm text-gray-400">
                Use this if you didn't know the answer, were incorrect, or
                struggled significantly. The card will reset to "New" and appear
                again very soon.
              </p>
            </div>
          </div>

          <div
            class="grid grid-cols-1 sm:grid-cols-[100px_1fr] gap-4 items-center p-3 rounded-lg hover:bg-white/5 transition"
          >
            <span
              class="text-center py-2 px-4 rounded-lg bg-amber-950/50 text-amber-300 border border-amber-900 font-bold text-sm"
            >
              3
            </span>
            <div>
              <h4 class="text-amber-400 font-bold">Hard</h4>
              <p class="text-sm text-gray-400">
                You remembered the answer, but it required intense mental effort
                or hesitation. The interval will grow slowly.
              </p>
            </div>
          </div>

          <div
            class="grid grid-cols-1 sm:grid-cols-[100px_1fr] gap-4 items-center p-3 rounded-lg hover:bg-white/5 transition"
          >
            <span
              class="text-center py-2 px-4 rounded-lg bg-emerald-950/50 text-emerald-300 border border-emerald-900 font-bold text-sm"
            >
              4
            </span>
            <div>
              <h4 class="text-emerald-400 font-bold">Good</h4>
              <p class="text-sm text-gray-400">
                You recalled the answer with a little hesitation. This is the
                standard "I know it" response.
              </p>
            </div>
          </div>

          <div
            class="grid grid-cols-1 sm:grid-cols-[100px_1fr] gap-4 items-center p-3 rounded-lg hover:bg-white/5 transition"
          >
            <span
              class="text-center py-2 px-4 rounded-lg bg-cyan-950/50 text-cyan-300 border border-cyan-900 font-bold text-sm"
            >
              5
            </span>
            <div>
              <h4 class="text-cyan-400 font-bold">Easy</h4>
              <p class="text-sm text-gray-400">
                You knew the answer instantly without thinking. The interval
                will grow aggressively to push the card far into the future.
              </p>
            </div>
          </div>
        </div>
      </div>
      <template v-if="!authStore.isAuthenticated">
        <div class="text-center pt-8">
          <NuxtLink
            to="/login"
            class="inline-flex items-center justify-center rounded-xl bg-indigo-600 px-8 py-3 text-base font-bold text-white shadow-lg transition hover:bg-indigo-500 hover:-translate-y-0.5"
          >
            Start Studying Now
          </NuxtLink>
        </div>
      </template>
      <template v-else>
        <div class="text-center pt-8">
          <NuxtLink
            to="/dashboard"
            class="inline-flex items-center justify-center rounded-xl bg-indigo-600 px-8 py-3 text-base font-bold text-white shadow-lg transition hover:bg-indigo-500 hover:-translate-y-0.5"
          >
            Start Studying Now
          </NuxtLink>
        </div>
      </template>
    </div>
  </div>
</template>
