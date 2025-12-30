<script setup lang="ts">
defineProps<{
  isOpen: boolean;
  title: string;
}>();

const emit = defineEmits(["close", "confirm"]);
</script>

<template>
  <Transition name="fade">
    <div
      v-if="isOpen"
      class="fixed inset-0 z-[100] flex items-center justify-center p-4"
    >
      <div
        class="absolute inset-0 bg-black/60 backdrop-blur-sm transition-opacity"
        @click="$emit('close')"
      ></div>

      <div
        class="relative w-full max-w-md transform overflow-hidden rounded-2xl border border-white/10 bg-gray-900 p-6 shadow-2xl transition-all"
      >
        <h3 class="text-xl font-bold text-white mb-2">
          {{ title }}
        </h3>

        <div class="mt-2 text-gray-400">
          <slot />
        </div>

        <div class="mt-6 flex justify-end gap-3">
          <button
            @click="$emit('close')"
            class="rounded-lg border border-gray-700 bg-transparent px-4 py-2 text-sm font-medium text-gray-300 hover:bg-gray-800 transition"
          >
            Cancel
          </button>
          <button
            @click="$emit('confirm')"
            class="rounded-lg bg-red-600 px-4 py-2 text-sm font-bold text-white hover:bg-red-500 shadow-lg shadow-red-500/20 transition"
          >
            Delete
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
