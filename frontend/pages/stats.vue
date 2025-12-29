<script setup lang="ts">
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Doughnut } from "vue-chartjs";

ChartJS.register(ArcElement, Tooltip, Legend);

const { data: stats, pending } = await useAPI<any>("/stats");

// Colors for the chart matching Tailwind classes
const chartColors = {
  new: "#3B82F6", // Blue-500
  learning: "#F59E0B", // Amber-500
  mature: "#10B981", // Emerald-500
};

const chartData = computed(() => {
  if (!stats.value) return { labels: [], datasets: [] };

  return {
    labels: ["New", "Learning", "Mature"],
    datasets: [
      {
        backgroundColor: [
          chartColors.new,
          chartColors.learning,
          chartColors.mature,
        ],
        borderWidth: 0, // Cleaner look without borders
        hoverOffset: 4,
        data: [
          stats.value.NewCards,
          stats.value.LearningCards,
          stats.value.MatureCards,
        ],
      },
    ],
  };
});

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      display: false, // We build a custom legend HTML for better styling
    },
    tooltip: {
      backgroundColor: "rgba(17, 24, 39, 0.9)", // gray-900
      padding: 12,
      cornerRadius: 8,
      titleColor: "#fff",
      bodyColor: "#9CA3AF", // gray-400
    },
  },
};
</script>

<template>
  <div class="min-h-[calc(100vh-6rem)] p-6 sm:p-10">
    <div class="mx-auto max-w-5xl">
      <!-- Navigation -->
      <NuxtLink
        to="/dashboard"
        class="inline-flex items-center gap-2 text-sm font-medium text-gray-400 hover:text-white transition-colors mb-8 group"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="w-4 h-4 transition-transform group-hover:-translate-x-1"
        >
          <path
            fill-rule="evenodd"
            d="M17 10a.75.75 0 01-.75.75H5.612l4.158 3.96a.75.75 0 11-1.04 1.08l-5.5-5.25a.75.75 0 010-1.08l5.5-5.25a.75.75 0 111.04 1.08L5.612 9.25H16.25A.75.75 0 0117 10z"
            clip-rule="evenodd"
          />
        </svg>
        Back to Dashboard
      </NuxtLink>

      <!-- Header -->
      <h1 class="text-3xl font-bold text-white tracking-tight mb-10">
        Your Progress
      </h1>

      <!-- Loading State -->
      <div
        v-if="pending"
        class="flex flex-col items-center justify-center py-20"
      >
        <div
          class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"
        ></div>
        <p class="text-gray-500 mt-4 animate-pulse">Analyzing your brain...</p>
      </div>

      <div v-else class="space-y-8">
        <!-- Summary Cards Grid -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Total Cards Card -->
          <div
            class="relative overflow-hidden rounded-2xl border border-white/5 bg-gray-900/40 p-8 shadow-xl backdrop-blur-sm"
          >
            <div class="absolute top-0 right-0 p-4 opacity-10">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-24 h-24 text-indigo-500"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 6.042A8.967 8.967 0 006 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 016 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 016-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0018 18a8.967 8.967 0 00-6 2.292m0-14.25v14.25"
                />
              </svg>
            </div>
            <h2
              class="text-sm font-medium uppercase tracking-wider text-gray-400"
            >
              Total Cards
            </h2>
            <p class="mt-2 text-5xl font-bold text-white tracking-tight">
              {{ stats.TotalCards }}
            </p>
          </div>

          <!-- Due Today Card -->
          <div
            class="relative overflow-hidden rounded-2xl border border-white/5 bg-gray-900/40 p-8 shadow-xl backdrop-blur-sm"
          >
            <div class="absolute top-0 right-0 p-4 opacity-10">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-24 h-24 text-rose-500"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z"
                />
              </svg>
            </div>
            <h2
              class="text-sm font-medium uppercase tracking-wider text-gray-400"
            >
              Due Today
            </h2>
            <p class="mt-2 text-5xl font-bold text-rose-400 tracking-tight">
              {{ stats.CardsDue }}
            </p>
          </div>
        </div>

        <!-- Chart Section -->
        <div
          class="rounded-2xl border border-white/5 bg-gray-900/40 p-8 shadow-xl backdrop-blur-sm"
        >
          <h2
            class="text-xl font-bold text-white mb-8 border-b border-gray-800 pb-4"
          >
            Knowledge Distribution
          </h2>

          <div
            class="flex flex-col md:flex-row items-center justify-around gap-8"
          >
            <!-- Chart Container -->
            <ClientOnly>
              <div class="relative h-64 w-64 md:h-80 md:w-80">
                <Doughnut :data="chartData" :options="chartOptions" />
                <!-- Center Text Overlay -->
                <div
                  class="absolute inset-0 flex items-center justify-center pointer-events-none"
                >
                  <div class="text-center">
                    <span class="block text-3xl font-bold text-white">{{
                      stats.TotalCards
                    }}</span>
                    <span class="text-xs text-gray-500 uppercase">Cards</span>
                  </div>
                </div>
              </div>
            </ClientOnly>

            <!-- Custom Legend -->
            <div class="flex flex-col gap-4 w-full md:w-auto">
              <!-- New -->
              <div
                class="flex items-center justify-between gap-12 p-3 rounded-lg bg-gray-800/50 border border-gray-700/50"
              >
                <div class="flex items-center gap-3">
                  <span
                    class="h-3 w-3 rounded-full bg-blue-500 shadow-[0_0_8px_rgba(59,130,246,0.5)]"
                  ></span>
                  <span class="text-gray-300 font-medium">New</span>
                </div>
                <span class="text-xl font-bold text-white">{{
                  stats.NewCards
                }}</span>
              </div>

              <!-- Learning -->
              <div
                class="flex items-center justify-between gap-12 p-3 rounded-lg bg-gray-800/50 border border-gray-700/50"
              >
                <div class="flex items-center gap-3">
                  <span
                    class="h-3 w-3 rounded-full bg-amber-500 shadow-[0_0_8px_rgba(245,158,11,0.5)]"
                  ></span>
                  <span class="text-gray-300 font-medium">Learning</span>
                </div>
                <span class="text-xl font-bold text-white">{{
                  stats.LearningCards
                }}</span>
              </div>

              <!-- Mature -->
              <div
                class="flex items-center justify-between gap-12 p-3 rounded-lg bg-gray-800/50 border border-gray-700/50"
              >
                <div class="flex items-center gap-3">
                  <span
                    class="h-3 w-3 rounded-full bg-emerald-500 shadow-[0_0_8px_rgba(16,185,129,0.5)]"
                  ></span>
                  <span class="text-gray-300 font-medium">Mature</span>
                </div>
                <span class="text-xl font-bold text-white">{{
                  stats.MatureCards
                }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
