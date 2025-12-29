<script setup lang="ts">
const route = useRoute();
const deckId = Number(route.params.id);

const front = ref("");
const back = ref("");

const { data: decks } = await useAPI<any[]>("/decks");
const currentDeck = computed(() =>
  decks.value?.find((d: any) => d.ID === deckId)
);

const { data: allCards, refresh: refreshCards } = await useAPI<any[]>("/cards");

const deckCards = computed(() => {
  if (!allCards.value) return [];
  return allCards.value.filter((c: any) => c.DeckID === deckId);
});

async function createCard() {
  if (!front.value || !back.value) return;

  await useAPI("/cards", {
    method: "POST",
    body: {
      deck_id: deckId,
      front: front.value,
      back: back.value,
    },
  });

  front.value = "";
  back.value = "";
  refreshCards();
}
</script>

<template>
  <div class="min-h-[calc(100vh-6rem)] p-6 sm:p-10">
    <div class="mx-auto max-w-5xl">
      <!-- NAVIGATION -->
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

      <!-- HEADER -->
      <div
        class="flex items-end justify-between mb-10 border-b border-gray-800 pb-6"
      >
        <div>
          <h1 class="text-3xl font-bold text-white tracking-tight">
            {{ currentDeck?.Name || "Loading..." }}
          </h1>
        </div>
        <div class="text-right">
          <p class="text-2xl font-bold text-indigo-400">
            {{ deckCards.length }}
          </p>
          <p class="text-xs text-gray-500 uppercase tracking-wider">
            Total Cards
          </p>
        </div>
      </div>

      <!-- CREATE CARD FORM -->
      <section class="mb-12">
        <div
          class="relative overflow-hidden rounded-2xl border border-white/10 bg-gray-900/50 p-6 sm:p-8 shadow-2xl backdrop-blur-xl"
        >
          <!-- Glow Effect -->
          <div
            class="absolute -right-10 -top-10 h-32 w-32 rounded-full bg-indigo-500/10 blur-3xl"
          ></div>

          <h2
            class="relative text-xl font-bold text-white mb-6 flex items-center gap-2"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6 text-indigo-500"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 4.5v15m7.5-7.5h-15"
              />
            </svg>
            Add New Card
          </h2>

          <form @submit.prevent="createCard" class="relative space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <!-- Front Input -->
              <div class="space-y-2">
                <label class="block text-sm font-medium text-gray-300"
                  >Front (Question)</label
                >
                <input
                  v-model="front"
                  type="text"
                  required
                  class="block w-full rounded-xl border border-gray-700 bg-gray-800/50 px-4 py-3 text-white placeholder-gray-600 focus:border-indigo-500 focus:bg-gray-800 focus:ring-1 focus:ring-indigo-500 outline-none transition-all"
                  placeholder="e.g. What is the complexity of Binary Search?"
                />
              </div>

              <!-- Back Input -->
              <div class="space-y-2">
                <label class="block text-sm font-medium text-gray-300"
                  >Back (Answer)</label
                >
                <input
                  v-model="back"
                  type="text"
                  required
                  class="block w-full rounded-xl border border-gray-700 bg-gray-800/50 px-4 py-3 text-white placeholder-gray-600 focus:border-indigo-500 focus:bg-gray-800 focus:ring-1 focus:ring-indigo-500 outline-none transition-all"
                  placeholder="e.g. O(log n)"
                />
              </div>
            </div>

            <div class="flex justify-end">
              <button
                type="submit"
                class="rounded-lg bg-indigo-600 px-6 py-2.5 text-sm font-bold text-white shadow-lg shadow-indigo-500/20 transition hover:bg-indigo-500 hover:-translate-y-0.5 active:translate-y-0"
              >
                Add Card
              </button>
            </div>
          </form>
        </div>
      </section>

      <!-- CARD LIST -->
      <section>
        <h2 class="text-xl font-bold text-white mb-6">Cards in this Deck</h2>

        <div v-if="deckCards.length > 0" class="space-y-4">
          <!-- Card Row (Flexbox instead of Table for better styling) -->
          <div
            v-for="card in deckCards"
            :key="card.ID"
            class="group flex flex-col md:flex-row md:items-center gap-4 rounded-xl border border-white/5 bg-gray-800/30 p-4 hover:bg-gray-800/50 transition-colors"
          >
            <!-- Content -->
            <div class="flex-1 grid grid-cols-1 md:grid-cols-2 gap-4">
              <div class="flex flex-col gap-1">
                <span
                  class="text-xs uppercase tracking-wide text-gray-500 font-bold"
                  >Front</span
                >
                <p class="text-white font-medium">{{ card.Front }}</p>
              </div>

              <div class="flex flex-col gap-1">
                <span
                  class="text-xs uppercase tracking-wide text-gray-500 font-bold"
                  >Back</span
                >
                <p class="text-gray-300">{{ card.Back }}</p>
              </div>
            </div>

            <!-- Stats Badge -->
            <div
              class="flex items-center gap-4 md:border-l md:border-gray-700 md:pl-6 min-w-[140px]"
            >
              <div class="flex flex-col">
                <span class="text-xs text-gray-500">Streak</span>
                <span class="font-mono text-indigo-400 font-bold">{{
                  card.Stats.Repetitions
                }}</span>
              </div>
              <div class="flex flex-col">
                <span class="text-xs text-gray-500">Interval</span>
                <span class="font-mono text-emerald-400 font-bold"
                  >{{ card.Stats.Interval }}d</span
                >
              </div>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div
          v-else
          class="text-center py-12 rounded-xl border border-dashed border-gray-700"
        >
          <div
            class="inline-flex h-12 w-12 items-center justify-center rounded-full bg-gray-800 text-gray-600 mb-4"
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
                d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m0 12.75h7.5m-7.5 3H12M10.5 2.25H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z"
              />
            </svg>
          </div>
          <p class="text-gray-400 text-lg">This deck is empty.</p>
          <p class="text-gray-500 text-sm">
            Add your first flashcard using the form above.
          </p>
        </div>
      </section>
    </div>
  </div>
</template>
