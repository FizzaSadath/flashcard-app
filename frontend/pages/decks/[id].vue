<script setup lang="ts">
import { useToastStore } from "~/stores/toast";

useHead({ title: "Dashboard - Flip" });

const toast = useToastStore();

const route = useRoute();
const deckId = Number(route.params.id);

const front = ref("");
const back = ref("");

const showDeleteModal = ref(false);
const cardToDelete = ref<number | null>(null);

const { data: decks } = await useAPI<any[]>("/decks");
const currentDeck = computed(() =>
  decks.value?.find((d: any) => d.ID === deckId)
);

const { data: allCards, refresh: refreshCards } = await useAPI<any[]>("/cards");

const deckCards = computed(() => {
  if (!allCards.value) return [];
  return allCards.value.filter((c: any) => c.DeckID === deckId);
});

const fileInput = ref<HTMLInputElement | null>(null);
const isImporting = ref(false);

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

function openDeleteModal(id: number) {
  cardToDelete.value = id;
  showDeleteModal.value = true;
}

async function confirmDelete() {
  if (!cardToDelete.value) return;

  await useAPI(`/cards/${cardToDelete.value}`, { method: "DELETE" });

  refreshCards();
  showDeleteModal.value = false;
  cardToDelete.value = null;
}

async function handleFileUpload(event: Event) {
  const input = event.target as HTMLInputElement;

  if (!input.files || input.files.length === 0) return;
  const file = input.files[0];
  if (!file) return;

  const formData = new FormData();
  formData.append("file", file);

  isImporting.value = true;

  const { data, error } = await useAPI<any>(`/decks/${deckId}/import`, {
    method: "POST",
    body: formData,
  });

  isImporting.value = false;
  input.value = "";
  if (error.value) {
    const errorMsg =
      error.value.data?.error || "Import failed. Check file format.";
    toast.add(errorMsg, "error");
    return;
  }

  if (data.value && data.value.count === 0) {
    toast.add("No cards found in CSV. Check headers.", "error");
    return;
  }

  refreshCards();
  toast.add(`Imported ${data.value?.count} cards successfully!`, "success");
}

// Helper to click the hidden input
function triggerFileInput() {
  fileInput.value?.click();
}
</script>

<template>
  <div class="min-h-[calc(100vh-6rem)] p-6 sm:p-10">
    <div class="mx-auto max-w-5xl">
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

      <section class="mb-12">
        <div
          class="relative overflow-hidden rounded-2xl border border-white/10 bg-gray-900/50 p-6 sm:p-8 shadow-2xl backdrop-blur-xl"
        >
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
          <div class="flex justify-end mb-4">
            <input
              type="file"
              ref="fileInput"
              class="hidden"
              accept=".csv"
              @change="handleFileUpload"
            />

            <button
              @click="triggerFileInput"
              :disabled="isImporting"
              class="flex items-center gap-2 rounded-lg border border-gray-600 bg-gray-800 px-4 py-2 text-sm font-bold text-gray-300 hover:bg-gray-700 hover:text-white transition"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-5 h-5"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5"
                />
              </svg>
              {{ isImporting ? "Importing..." : "Import CSV" }}
            </button>
          </div>
          <form @submit.prevent="createCard" class="relative space-y-6">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
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

      <section>
        <h2 class="text-xl font-bold text-white mb-6">Cards in this Deck</h2>

        <div v-if="deckCards.length > 0" class="space-y-4">
          <div
            v-for="card in deckCards"
            :key="card.ID"
            class="group flex flex-col md:flex-row md:items-center gap-4 rounded-xl border border-white/5 bg-gray-800/30 p-4 hover:bg-gray-800/50 transition-colors relative"
          >
            <button
              @click="openDeleteModal(card.ID)"
              class="absolute top-2 right-2 p-2 text-gray-600 hover:text-red-400 opacity-100 sm:opacity-0 sm:group-hover:opacity-100 transition-all"
              title="Delete Card"
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

            <div class="flex-1 grid grid-cols-1 md:grid-cols-2 gap-4 pr-8">
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

    <modal
      :isOpen="showDeleteModal"
      title="Delete Card?"
      @close="showDeleteModal = false"
      @confirm="confirmDelete"
    >
      <p class="text-gray-300">
        Are you sure you want to delete this card? <br />This action can't be
        undone.
      </p>
    </modal>
  </div>
</template>
