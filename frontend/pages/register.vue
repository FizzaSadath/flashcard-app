<script setup lang="ts">
const email = ref("");
const username = ref("");
const password = ref("");
const confirmPassword = ref("");
const authStore = useAuthStore();
const loading = ref(false);

const handleRegister = async () => {
  loading.value = true;
  await authStore.register(
    email.value,
    username.value,
    password.value,
    confirmPassword.value
  );
  loading.value = false;
};
</script>

<template>
  <div class="flex min-h-screen items-center justify-center">
    <div class="w-full max-w-md">
      <div class="text-center">
        <h2>Create New User</h2>
        <p>
          Or
          <NuxtLink to="/login"> Already registered? Login </NuxtLink>
        </p>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="handleRegister">
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
            <label for="username" class="sr-only">Username</label>
            <input
              v-model="username"
              id="username"
              name="username"
              type="text"
              required
              placeholder="User Name"
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
          <div>
            <label for="confirmPassword" class="sr-only"
              >Confirm Password</label
            >
            <input
              v-model="confirmPassword"
              id="confirmPassword"
              name="confirmPassword"
              type="password"
              required
              placeholder="Confirm Password"
            />
          </div>
        </div>

        <div>
          <button type="submit" :disabled="loading">
            {{ loading ? "Creating new user ..." : "Submit" }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
