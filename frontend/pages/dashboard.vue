<script setup lang="ts">
import { useAuthStore } from "~/stores/auth";

const authStore = useAuthStore();
const newDeckName = ref("");

const { data: decks, refresh, error } = await useAPI<any[]>("/decks");

async function createDeck() {
  if (!newDeckName.value) return;

  await useAPI("/decks", {
    method: "POST",
    body: { name: newDeckName.value },
  });

  newDeckName.value = "";
  refresh();
}

async function deleteDeck(id: number) {
  if (!confirm("Are you sure?")) return;

  await useAPI(`/decks/${id}`, { method: "DELETE" });
  refresh();
}
</script>

<template>
  <div>
    <header>
      <h1>My Decks</h1>
      <p>Logged in as user: {{ authStore.user?.username || "Unknown" }}</p>
      <button @click="authStore.logout()">Logout</button>
    </header>

    <hr />

    <section>
      <h2>Create New Deck</h2>
      <form @submit.prevent="createDeck">
        <input
          v-model="newDeckName"
          type="text"
          placeholder="Enter deck name"
          required
        />
        <button type="submit">Add Deck</button>
      </form>
    </section>

    <div v-if="error" style="color: red">
      Error loading decks: {{ error.message }}
    </div>

    <section>
      <div v-if="decks && decks.length > 0">
        <ul>
          <li v-for="deck in decks" :key="deck.ID">
            <h3>{{ deck.Name }}</h3>
            <small
              >Created:
              {{ new Date(deck.CreatedAt).toLocaleDateString() }}</small
            >

            <div>
              <button @click="deleteDeck(deck.ID)">Delete</button>
              <NuxtLink :to="`/decks/${deck.ID}`">Open Deck</NuxtLink>
              <NuxtLink :to="`/study/${deck.ID}`"> Study Now </NuxtLink>
            </div>
          </li>
        </ul>
      </div>

      <div v-else>
        <p>No decks found. Create one to get started.</p>
      </div>
    </section>
  </div>
</template>
