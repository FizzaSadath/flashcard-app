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
  <div>
    <NuxtLink to="/dashboard">
      ← Back to Dashboard
    </NuxtLink>

    <header>
      <h1>Deck: {{ currentDeck?.Name || 'Loading...' }}</h1>
      <p>Deck ID: {{ deckId }}</p>
    </header>

    <hr />

    <section>
      <h2>Add New Card</h2>
      <form @submit.prevent="createCard">
        <div>
          <label>Front (Question)</label>
          <br />
          <input 
            v-model="front" 
            type="text" 
            required 
            placeholder="e.g. What is the complexity of Binary Search?" 
          />
        </div>
        
        <div>
          <label>Back (Answer)</label>
          <br />
          <input 
            v-model="back" 
            type="text" 
            required 
            placeholder="e.g. O(log n)" 
          />
        </div>

        <button type="submit">Add Card</button>
      </form>
    </section>

    <hr />

    <!-- Card List -->
    <section>
      <h2>Cards in this Deck ({{ deckCards.length }})</h2>
      
      <div v-if="deckCards.length > 0">
        <table border="1">
          <thead>
            <tr>
              <th>Front</th>
              <th>Back</th>
              <th>Stats</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="card in deckCards" :key="card.ID">
              <td>{{ card.Front }}</td>
              <td>{{ card.Back }}</td>
              <td>
                Reps: {{ card.Stats.Repetitions }} <br/>
                Next: {{ card.Stats.Interval }} days
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-else>
        <p>This deck is empty. Add a card above!</p>
      </div>
    </section>
  </div>
</template>