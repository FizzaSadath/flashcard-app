<script setup lang="ts">
const email = ref('')
const password = ref('')
const authStore = useAuthStore()
const loading = ref(false)

const handleLogin = async () => {
  loading.value = true
  await authStore.login(email.value, password.value)
  loading.value = false
}
</script>

<template>
  <div class="flex min-h-screen items-center justify-center">
    <div class="w-full max-w-md">
      
      <div class="text-center">
        <h2 >Sign in</h2>
        <p>
          Or
          <NuxtLink to="/register">
            create a new account
          </NuxtLink>
        </p>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
        <div class="space-y-4 rounded-md shadow-sm">
          <div>
            <label for="email" class="sr-only">Email address</label>
            <input
              v-model="email"
              id="email"
              name="email"
              type="email"
              required
              placeholder="Email address"
            />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input
              v-model="password"
              id="password"
              name="password"
              type="password"
              required
              placeholder="Password"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            :disabled="loading"
          >
            {{ loading ? 'Signing in...' : 'Sign in' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
