<script setup lang="ts">
import { useAuthStore } from "~/stores/auth";

const authStore = useAuthStore();
const newDeckName = ref("");

const {
  data: decks,
  refresh: refreshDecks,
  error,
} = await useAPI<any[]>("/decks");
const { data: deckStats, refresh: refreshStats } = await useAPI<any[]>(
  "/stats/decks"
);

// Helper to find due count for a specific deck
function getDueCount(deckId: number) {
  if (!deckStats.value) return 0;
  const stat = deckStats.value.find((s: any) => s.DeckID === deckId);
  return stat ? stat.Due : 0;
}

async function refreshAll() {
  await Promise.all([refreshDecks(), refreshStats()]);
}

async function createDeck() {
  if (!newDeckName.value) return;

  await useAPI("/decks", {
    method: "POST",
    body: { name: newDeckName.value },
  });

  newDeckName.value = "";
  refreshAll(); // Refresh both lists
}

async function deleteDeck(id: number) {
  if (
    !confirm(
      "Are you sure you want to delete this deck? All cards inside will be lost."
    )
  )
    return;

  await useAPI(`/decks/${id}`, { method: "DELETE" });
  refreshAll();
}
</script>

<template>
  <div class="min-h-[calc(100vh-6rem)] p-6 sm:p-10">
    <div class="mx-auto max-w-7xl">
      <div
        class="flex flex-col gap-6 md:flex-row md:items-center md:justify-between mb-10"
      >
        <div>
          <h1 class="text-3xl font-bold text-white tracking-tight">My Decks</h1>
          <p class="text-gray-400 mt-2 flex items-center gap-2">
            <span class="inline-block h-2 w-2 rounded-full bg-green-500"></span>
            Logged in as
            <span class="text-indigo-400 font-medium">{{
              authStore.user?.username || "Unknown"
            }}</span>
          </p>
        </div>
      </div>

      <!-- Create Deck -->
      <div class="mb-12">
        <form @submit.prevent="createDeck" class="relative max-w-xl">
          <div class="relative group">
            <div
              class="absolute -inset-0.5 rounded-xl bg-gradient-to-r from-indigo-500 to-purple-600 opacity-30 blur transition duration-1000 group-hover:opacity-70 group-hover:duration-200"
            ></div>
            <input
              v-model="newDeckName"
              type="text"
              placeholder="Create a new deck"
              class="relative block w-full rounded-xl border-none bg-gray-900 py-4 pl-6 pr-32 text-white placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/50"
              required
            />
            <button
              type="submit"
              class="absolute right-2 top-2 bottom-2 rounded-lg bg-indigo-600 px-6 text-sm font-bold text-white transition hover:bg-indigo-500 shadow-lg"
            >
              Add Deck
            </button>
          </div>
        </form>
      </div>

      <!-- Error message -->
      <div
        v-if="error"
        class="mb-8 flex items-center gap-3 rounded-lg border border-red-500/20 bg-red-500/10 p-4 text-red-200"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          class="w-5 h-5"
        >
          <path
            fill-rule="evenodd"
            d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-5a.75.75 0 01.75.75v4.5a.75.75 0 01-1.5 0v-4.5A.75.75 0 0110 5zm0 10a1 1 0 100-2 1 1 0 000 2z"
            clip-rule="evenodd"
          />
        </svg>
        Error loading decks: {{ error.message }}
      </div>

      <!-- Deck grid -->
      <section>
        <div
          v-if="decks && decks.length > 0"
          class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3"
        >
          <div
            v-for="deck in decks"
            :key="deck.ID"
            class="group relative flex flex-col justify-between overflow-hidden rounded-2xl border border-white/5 bg-gray-800/40 p-6 shadow-2xl backdrop-blur-sm transition-all hover:-translate-y-1 hover:border-white/10 hover:bg-gray-800/60"
          >
            <div>
              <div class="flex items-start justify-between">
                <div class="flex items-center gap-3">
                  <!-- Deck Icon -->
                  <div class="rounded-lg bg-gray-700/50 p-3 text-indigo-400">
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

                  <!-- NEW: Due Badge -->
                  <div
                    v-if="getDueCount(deck.ID) > 0"
                    class="flex items-center gap-1.5 rounded-full bg-rose-500/10 border border-rose-500/20 px-3 py-1 text-xs font-bold text-rose-400 shadow-[0_0_10px_rgba(244,63,94,0.1)]"
                  >
                    <span class="relative flex h-2 w-2">
                      <span
                        class="animate-ping absolute inline-flex h-full w-full rounded-full bg-rose-400 opacity-75"
                      ></span>
                      <span
                        class="relative inline-flex rounded-full h-2 w-2 bg-rose-500"
                      ></span>
                    </span>
                    {{ getDueCount(deck.ID) }} Due
                  </div>
                </div>

                <button
                  @click="deleteDeck(deck.ID)"
                  class="text-gray-600 transition hover:text-red-400 p-2"
                  title="Delete Deck"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                    class="w-5 h-5"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M8.75 1A2.75 2.75 0 006 3.75v.443c-.795.077-1.584.176-2.365.298a.75.75 0 10.23 1.482l.149-.022.841 10.518A2.75 2.75 0 007.596 19h4.807a2.75 2.75 0 002.742-2.53l.841-10.52.149.023a.75.75 0 00.23-1.482A41.03 41.03 0 0014 4.193V3.75A2.75 2.75 0 0011.25 1h-2.5zM10 4c.84 0 1.673.025 2.5.075V3.75c0-.69-.56-1.25-1.25-1.25h-2.5c-.69 0-1.25.56-1.25 1.25v.325C8.327 4.025 9.16 4 10 4zM8.58 7.72a.75.75 0 00-1.5.06l.3 7.5a.75.75 0 101.5-.06l-.3-7.5zm4.34.06a.75.75 0 10-1.5-.06l-.3 7.5a.75.75 0 001.5.06l.3-7.5z"
                      clip-rule="evenodd"
                    />
                  </svg>
                </button>
              </div>

              <h3 class="mt-4 text-xl font-bold text-white tracking-wide">
                {{ deck.Name }}
              </h3>
              <p
                class="mt-1 text-xs text-gray-500 font-medium uppercase tracking-wider"
              >
                Created: {{ new Date(deck.CreatedAt).toLocaleDateString() }}
              </p>
            </div>

            <!-- Card Actions -->
            <div class="mt-8 flex gap-3">
              <NuxtLink
                :to="`/study/${deck.ID}`"
                class="relative flex-1 flex items-center justify-center gap-2 rounded-lg bg-indigo-600 py-2.5 text-sm font-bold text-white shadow-lg shadow-indigo-500/20 transition hover:bg-indigo-500 hover:-translate-y-0.5"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                  fill="currentColor"
                  class="w-4 h-4"
                >
                  <path d="M10 12.5a2.5 2.5 0 100-5 2.5 2.5 0 000 5z" />
                  <path
                    fill-rule="evenodd"
                    d="M.664 10.59a1.651 1.651 0 010-1.186A10.004 10.004 0 0110 3c4.257 0 7.893 2.66 9.336 6.41.147.381.146.804 0 1.186A10.004 10.004 0 0110 17c-4.257 0-7.893-2.66-9.336-6.41zM14 10a4 4 0 11-8 0 4 4 0 018 0z"
                    clip-rule="evenodd"
                  />
                </svg>
                Study

                <!-- NEW: Mini alert dot on button if due -->
                <span
                  v-if="getDueCount(deck.ID) > 0"
                  class="absolute -top-1 -right-1 flex h-3 w-3"
                >
                  <span
                    class="animate-ping absolute inline-flex h-full w-full rounded-full bg-rose-400 opacity-75"
                  ></span>
                  <span
                    class="relative inline-flex rounded-full h-3 w-3 bg-rose-500"
                  ></span>
                </span>
              </NuxtLink>

              <NuxtLink
                :to="`/decks/${deck.ID}`"
                class="flex-1 flex items-center justify-center gap-2 rounded-lg border border-gray-600 bg-gray-700/50 py-2.5 text-sm font-semibold text-gray-200 transition hover:bg-gray-700 hover:text-white"
              >
                Cards
              </NuxtLink>
            </div>
          </div>
        </div>

        <div
          v-else
          class="mt-16 text-center py-12 rounded-2xl border border-white/5 bg-gray-900/50 backdrop-blur-sm"
        >
          <div
            class="mx-auto flex h-16 w-16 items-center justify-center rounded-full bg-gray-800 text-gray-500"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-8 h-8"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M12 9v6m3-3H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </div>
          <h3 class="mt-4 text-xl font-bold text-white">No decks found</h3>
          <p class="mt-2 text-gray-400 max-w-sm mx-auto">
            You haven't created any flashcard decks yet. Use the input field
            above to get started!
          </p>
        </div>
      </section>
    </div>
  </div>
</template>
