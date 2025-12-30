import { defineStore } from "pinia";

export interface Toast {
  id: number;
  message: string;
  type: "success" | "error" | "info";
}

export const useToastStore = defineStore("toast", () => {
  const toasts = ref<Toast[]>([]);
  let nextId = 0;

  function add(message: string, type: "success" | "error" | "info" = "info") {
    const id = nextId++;

    toasts.value.push({ id, message, type });

    //timeout after 3 sec
    setTimeout(() => {
      remove(id);
    }, 3000);
  }

  function remove(id: number) {
    toasts.value = toasts.value.filter((t) => t.id !== id);
  }

  return { toasts, add, remove };
});
