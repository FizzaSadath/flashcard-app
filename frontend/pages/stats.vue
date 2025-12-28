<script setup lang="ts">
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Doughnut } from "vue-chartjs";

ChartJS.register(ArcElement, Tooltip, Legend);

const { data: stats, pending } = await useAPI<any>("/stats");

const chartData = computed(() => {
  if (!stats.value) return { labels: [], datasets: [] };

  return {
    labels: ["New", "Learning", "Mature"],
    datasets: [
      {
        backgroundColor: ["#3B82F6", "#F59E0B", "#10B981"],
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
};
</script>

<template>
  <div>
    <NuxtLink to="/dashboard"> ← Back to Dashboard </NuxtLink>

    <h1>Your Progress</h1>

    <div v-if="pending">Loading stats...</div>

    <div v-else>
      <!-- Summary Grid -->
      <div>
        <!-- Total Cards Card -->
        <div>
          <h2>Total Cards</h2>
          <p>{{ stats.TotalCards }}</p>
        </div>

        <!-- Due Today Card -->
        <div>
          <h2>Due Today</h2>
          <p>{{ stats.CardsDue }}</p>
        </div>
      </div>

      <!-- Chart Section -->
      <div>
        <h2>Card Knowledge Distribution</h2>

        <!-- Chart Container -->
        <!-- Note: You might need to add specific height/width styles later for the Chart to render visibly -->
        <ClientOnly>
          <div style="height: 300px; position: relative">
            <Doughnut :data="chartData" :options="chartOptions" />
          </div>
        </ClientOnly>

        <!-- Legend / Explanation -->
        <div>
          <div><strong>New:</strong> {{ stats.NewCards }}</div>
          <div><strong>Learning:</strong> {{ stats.LearningCards }}</div>
          <div><strong>Mature:</strong> {{ stats.MatureCards }}</div>
        </div>
      </div>
    </div>
  </div>
</template>
