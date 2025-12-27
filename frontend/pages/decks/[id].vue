<script setup lang="ts">
const route = useRoute()
const deckId = Number(route.params.id) 

const front = ref('')
const back = ref('')

const { data: decks } = await useAPI<any[]>('/decks')
const currentDeck = computed(() => decks.value?.find((d: any) => d.ID === deckId))

const { data: allCards, refresh: refreshCards } = await useAPI<any[]>('/cards')

const deckCards = computed(() => {
  if (!allCards.value) return []
  return allCards.value.filter((c: any) => c.DeckID === deckId)
})

async function createCard() {
  if (!front.value || !back.value) return

  await useAPI('/cards', {
    method: 'POST',
    body: {
      deck_id: deckId, 
      front: front.value,
      back: back.value
    }
  })

  front.value = ''
  back.value = ''
  refreshCards()
}


</script>

<template>
  <div class="p-8">
    <NuxtLink to="/dashboard" class="text-blue-500 underline mb-4 inline-block">
      ← Back to Dashboard
    </NuxtLink>

    <!-- Header -->
    <header class="mb-8 border-b pb-4">
      <h1 class="text-3xl font-bold">
        Deck: {{ currentDeck?.Name || 'Loading...' }}
      </h1>
      <p class="text-gray-500">Deck ID: {{ deckId }}</p>
    </header>

    <!-- Create Card Form -->
    <section class="mb-8 bg-gray-100 p-6 rounded-lg">
      <h2 class="text-xl font-bold mb-4">Add New Card</h2>
      <form @submit.prevent="createCard" class="flex flex-col gap-4 max-w-lg">
        <div>
          <label class="block font-bold text-sm">Front (Question)</label>
          <input 
            v-model="front" 
            type="text" 
            required 
            class="w-full border p-2 rounded"
            placeholder="e.g. What is the complexity of Binary Search?" 
          />
        </div>
        
        <div>
          <label class="block font-bold text-sm">Back (Answer)</label>
          <input 
            v-model="back" 
            type="text" 
            required 
            class="w-full border p-2 rounded"
            placeholder="e.g. O(log n)" 
          />
        </div>

        <button 
          type="submit" 
          class="bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 font-bold"
        >
          Add Card
        </button>
      </form>
    </section>

    <!-- Card List -->
    <section>
      <h2 class="text-xl font-bold mb-4">Cards in this Deck ({{ deckCards.length }})</h2>
      
      <div v-if="deckCards.length > 0" class="overflow-x-auto">
        <table class="w-full border-collapse border border-gray-300">
          <thead class="bg-gray-200">
            <tr>
              <th class="border p-2 text-left">Front</th>
              <th class="border p-2 text-left">Back</th>
              <th class="border p-2 text-left">Stats</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="card in deckCards" :key="card.ID" class="hover:bg-gray-50">
              <td class="border p-2">{{ card.Front }}</td>
              <td class="border p-2">{{ card.Back }}</td>
              <td class="border p-2 text-sm text-gray-600">
                Reps: {{ card.Stats.Repetitions }} <br/>
                Next: {{ card.Stats.Interval }} days
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-else class="text-gray-500 italic">
        This deck is empty. Add a card above!
      </div>
    </section>
  </div>
</template>