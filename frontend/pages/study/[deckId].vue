<script setup lang="ts">
const route = useRoute();
const router = useRouter();
const deckId = Number(route.params.deckId);

useHead({ title: "Study - Flip" });

const { data: dueCards, pending } = await useAPI<any[]>(
  `/cards/due?deck_id=${deckId}`
);

const queue = computed(() => {
  return dueCards.value || [];
});

const currentIndex = ref(0);
const showAnswer = ref(false);
const sessionComplete = ref(false);

// Get current card
const currentCard = computed(() => {
  if (!queue.value || queue.value.length === 0) return null;
  return queue.value[currentIndex.value];
});

// Submit grade
async function submitGrade(grade: number) {
  if (!currentCard.value) return;

  await useAPI("/cards/review", {
    method: "POST",
    body: {
      card_id: currentCard.value.ID,
      grade: grade,
    },
  });

  if (currentIndex.value < queue.value.length - 1) {
    currentIndex.value++;
    showAnswer.value = false; // Reset flip
  } else {
    sessionComplete.value = true;
  }
}

function goBack() {
  router.push("/dashboard");
}
</script>

<template>
  <div class="min-h-[calc(100vh-6rem)] p-6 sm:p-10 flex flex-col items-center">
    <div class="w-full max-w-3xl">
      <div class="flex items-center justify-between mb-8">
        <button
          @click="goBack"
          class="flex items-center gap-2 text-sm font-medium text-gray-400 hover:text-white transition-colors group"
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
          Exit Study
        </button>

        <div
          class="text-sm font-mono text-indigo-400 bg-indigo-500/10 px-3 py-1 rounded-full border border-indigo-500/20"
        >
          <span v-if="currentCard">
            Card {{ currentIndex + 1 }} <span class="text-gray-500">/</span>
            {{ queue.length }}
          </span>
          <span v-else>Session Status</span>
        </div>
      </div>

      <div
        v-if="pending"
        class="flex flex-col items-center justify-center py-20"
      >
        <div
          class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-500"
        ></div>
        <p class="text-gray-500 mt-4 animate-pulse">Loading cards...</p>
      </div>

      <div
        v-else-if="sessionComplete"
        class="text-center py-16 rounded-3xl border border-white/5 bg-gray-900/50 backdrop-blur-xl shadow-2xl"
      >
        <div
          class="mx-auto flex h-20 w-20 items-center justify-center rounded-full bg-green-500/10 mb-6"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            class="w-10 h-10 text-green-500"
          >
            <path
              fill-rule="evenodd"
              d="M8.603 3.799A4.49 4.49 0 0112 2.25c1.357 0 2.573.6 3.397 1.549a4.49 4.49 0 013.498 1.307 4.491 4.491 0 011.307 3.497A4.49 4.49 0 0121.75 12a4.49 4.49 0 01-1.549 3.397 4.491 4.491 0 01-1.307 3.497 4.491 4.491 0 01-3.497 1.307A4.49 4.49 0 0112 21.75a4.49 4.49 0 01-3.397-1.549 4.49 4.49 0 01-3.498-1.306 4.491 4.491 0 01-1.307-3.498A4.49 4.49 0 012.25 12c0-1.357.6-2.573 1.549-3.397a4.49 4.49 0 011.307-3.497 4.49 4.49 0 013.497-1.307zm7.007 6.387a.75.75 0 10-1.22-.872l-3.236 4.53L9.53 12.22a.75.75 0 00-1.06 1.06l2.25 2.25a.75.75 0 001.14-.094l3.75-5.25z"
              clip-rule="evenodd"
            />
          </svg>
        </div>
        <h1 class="text-3xl font-bold text-white mb-2">All Done!</h1>
        <p class="text-gray-400 mb-8">
          You have reviewed all due cards for today.
        </p>
        <button
          @click="goBack"
          class="rounded-xl bg-indigo-600 px-8 py-3 text-sm font-bold text-white shadow-lg shadow-indigo-500/20 transition hover:bg-indigo-500 hover:-translate-y-0.5"
        >
          Back to Dashboard
        </button>
      </div>

      <div
        v-else-if="!currentCard"
        class="text-center py-16 rounded-3xl border border-white/5 bg-gray-900/50 backdrop-blur-xl shadow-2xl"
      >
        <div
          class="mx-auto flex h-20 w-20 items-center justify-center rounded-full bg-gray-800 text-gray-500 mb-6"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-10 h-10"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z"
            />
          </svg>
        </div>
        <h1 class="text-2xl font-bold text-white mb-2">No cards due!</h1>
        <p class="text-gray-400 mb-8">You are all caught up on this deck.</p>
        <button
          @click="goBack"
          class="rounded-xl bg-gray-700 px-6 py-2.5 text-sm font-bold text-white transition hover:bg-gray-600"
        >
          Go Back
        </button>
      </div>

      <div v-else>
        <div
          class="relative min-h-[400px] flex flex-col items-center justify-center rounded-3xl border border-white/10 bg-gradient-to-b from-gray-800/50 to-gray-900/50 p-8 sm:p-12 text-center shadow-2xl backdrop-blur-xl transition-all"
        >
          <div class="w-full">
            <span
              class="text-xs font-bold uppercase tracking-widest text-indigo-400 mb-4 block"
              >Question</span
            >
            <h2
              class="text-2xl sm:text-3xl font-medium text-white leading-relaxed"
            >
              {{ currentCard.Front }}
            </h2>
          </div>

          <div
            v-if="showAnswer"
            class="w-full mt-8 pt-8 border-t border-white/10 animate-[fadeIn_0.3s_ease-out]"
          >
            <span
              class="text-xs font-bold uppercase tracking-widest text-emerald-400 mb-4 block"
              >Answer</span
            >
            <h2
              class="text-xl sm:text-2xl font-medium text-gray-200 leading-relaxed"
            >
              {{ currentCard.Back }}
            </h2>
          </div>
        </div>

        <div class="mt-8">
          <button
            v-if="!showAnswer"
            @click="showAnswer = true"
            class="w-full rounded-xl bg-indigo-600 py-4 text-base font-bold text-white shadow-lg shadow-indigo-500/20 transition-all hover:bg-indigo-500 hover:shadow-indigo-500/40 active:scale-[0.98]"
          >
            Show Answer
          </button>

          <div
            v-else
            class="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-3"
          >
            <button
              @click="submitGrade(0)"
              class="grade-btn bg-rose-950/30 text-rose-200 border-rose-900 hover:bg-rose-900 hover:border-rose-700"
            >
              <span class="block text-lg font-bold">0</span>
              <span class="text-xs uppercase opacity-70">Blackout</span>
            </button>
            <button
              @click="submitGrade(1)"
              class="grade-btn bg-rose-950/30 text-rose-200 border-rose-900 hover:bg-rose-900 hover:border-rose-700"
            >
              <span class="block text-lg font-bold">1</span>
              <span class="text-xs uppercase opacity-70">Incorrect</span>
            </button>
            <button
              @click="submitGrade(2)"
              class="grade-btn bg-orange-950/30 text-orange-200 border-orange-900 hover:bg-orange-900 hover:border-orange-700"
            >
              <span class="block text-lg font-bold">2</span>
              <span class="text-xs uppercase opacity-70">Close</span>
            </button>

            <button
              @click="submitGrade(3)"
              class="grade-btn bg-amber-950/30 text-amber-200 border-amber-900 hover:bg-amber-900 hover:border-amber-700"
            >
              <span class="block text-lg font-bold">3</span>
              <span class="text-xs uppercase opacity-70">Hard</span>
            </button>
            <button
              @click="submitGrade(4)"
              class="grade-btn bg-emerald-950/30 text-emerald-200 border-emerald-900 hover:bg-emerald-900 hover:border-emerald-700"
            >
              <span class="block text-lg font-bold">4</span>
              <span class="text-xs uppercase opacity-70">Good</span>
            </button>
            <button
              @click="submitGrade(5)"
              class="grade-btn bg-cyan-950/30 text-cyan-200 border-cyan-900 hover:bg-cyan-900 hover:border-cyan-700"
            >
              <span class="block text-lg font-bold">5</span>
              <span class="text-xs uppercase opacity-70">Easy</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.grade-btn {
  @apply flex flex-col items-center justify-center py-3 rounded-xl border transition-all duration-200 active:scale-95;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
