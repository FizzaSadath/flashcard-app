<script setup lang="ts">
const authStore = useAuthStore()
</script>

<template>
  <div class="min-h-screen">
    <div class="flex justify-between items-center mb-8">
      <h1 >Dashboard</h1>
      <button 
        @click="authStore.logout()" 
      >
        Logout
      </button>
    </div>
    
    <div>
      <p class="text-xl">Welcome! You are logged in.</p>
      <p class="text-gray-400 mt-2">Token: {{ authStore.token?.substring(0, 20) }}...</p>
    </div>
  </div>
</template>