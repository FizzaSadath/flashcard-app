<script setup lang="ts">
const route = useRoute();
const router = useRouter();
const deckId = Number(route.params.deckId);

// Fetch all due cards of user
const { data: dueCards, pending } = await useAPI<any[]>("/cards/due");

// Return all cards belonging to the current deck
const queue = computed(() => {
  if (!dueCards.value) return [];
  return dueCards.value.filter((c: any) => c.DeckID === deckId);
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
  <div>
    <header>
      <button @click="goBack">← Exit Study</button>
      <div>
        <span v-if="currentCard"
          >Card {{ currentIndex + 1 }} of {{ queue.length }}</span
        >
      </div>
    </header>

    <hr />

    <div v-if="pending">Loading your cards...</div>

    <div v-else-if="sessionComplete">
      <h1>All Done!</h1>
      <p>You have reviewed all due cards for today.</p>
      <button @click="goBack">Back to Dashboard</button>
    </div>

    <div v-else-if="!currentCard">
      <h1>No cards due!</h1>
      <p>You are all caught up on this deck.</p>
      <button @click="goBack">Go Back</button>
    </div>

    <div v-else>
      <section
        style="
          border: 1px solid black;
          padding: 20px;
          margin: 20px 0;
          text-align: center;
        "
      >
        <div>
          <small>Question</small>
          <h2>{{ currentCard.Front }}</h2>
        </div>

        <div v-if="showAnswer">
          <hr />
          <small>Answer</small>
          <h2>{{ currentCard.Back }}</h2>
        </div>
      </section>

      <div>
        <button
          v-if="!showAnswer"
          @click="showAnswer = true"
          style="width: 100%; padding: 10px"
        >
          Show Answer
        </button>

        <div v-else style="display: flex; gap: 10px; justify-content: center">
          <!-- Failures (0-2) -->
          <button @click="submitGrade(0)">0 (Blackout)</button>
          <button @click="submitGrade(1)">1 (Incorrect)</button>
          <button @click="submitGrade(2)">2 (Close)</button>

          <!-- Successes (3-5) -->
          <button @click="submitGrade(3)">3 (Hard)</button>
          <button @click="submitGrade(4)">4 (Good)</button>
          <button @click="submitGrade(5)">5 (Easy)</button>
        </div>
      </div>
    </div>
  </div>
</template>
