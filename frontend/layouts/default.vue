<template>
  <div
    class="min-h-screen flex flex-col font-sans text-gray-100 relative bg-gray-950"
  >
    <div
      class="absolute top-0 left-0 w-full h-full overflow-hidden -z-10 pointer-events-none"
    >
      <div
        class="absolute -top-[20%] -left-[10%] w-[50%] h-[50%] rounded-full bg-indigo-900/20 blur-[120px]"
      ></div>
      <div
        class="absolute top-[20%] -right-[10%] w-[40%] h-[40%] rounded-full bg-blue-900/10 blur-[100px]"
      ></div>
    </div>

    <navbar />

    <main
      class="flex-1 w-full max-w-[1920px] mx-auto flex flex-col relative z-10"
    >
      <slot />
    </main>
    <toast />
  </div>
</template>
