<script setup lang="ts">
import { useToastStore } from "~/stores/toast";

const toastStore = useToastStore();
</script>

<template>
  <div
    class="fixed bottom-4 right-4 z-[9999] flex flex-col gap-2 w-full max-w-sm pointer-events-none"
  >
    <TransitionGroup name="toast">
      <div
        v-for="toast in toastStore.toasts"
        :key="toast.id"
        class="pointer-events-auto flex items-center justify-between gap-4 p-4 rounded-xl border shadow-2xl backdrop-blur-md transition-all duration-300 transform translate-x-0"
        :class="{
          'bg-green-900/80 border-green-500/30 text-green-200':
            toast.type === 'success',
          'bg-red-900/80 border-red-500/30 text-red-200':
            toast.type === 'error',
          'bg-blue-900/80 border-blue-500/30 text-blue-200':
            toast.type === 'info',
        }"
      >
        <div class="flex items-center gap-3">
          <!-- Success Icon -->
          <svg
            v-if="toast.type === 'success'"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            class="w-5 h-5"
          >
            <path
              fill-rule="evenodd"
              d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
              clip-rule="evenodd"
            />
          </svg>
          <!-- Error Icon -->
          <svg
            v-else-if="toast.type === 'error'"
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

          <span class="font-medium text-sm">{{ toast.message }}</span>
        </div>

        <button @click="toastStore.remove(toast.id)" class="hover:opacity-70">
          ✕
        </button>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}
.toast-enter-from,
.toast-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>
