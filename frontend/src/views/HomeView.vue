<script setup lang="ts">
import { Github } from 'lucide-vue-next';
import { Button } from '@/components/ui/button';
</script>

<template>
  <div>
    <main>
      <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl">MallData</h1>
      <a href="https://github.com/Pixium/malldata">
        <Button variant="link" size="icon">
          <Github />
        </Button>
      </a>
    </main>

    <footer>
      <p class="leading-7 [&:not(:first-child)]:mt-6 text-center">
        Made by <a class="green" href="https://github.com/Erb3">Erb3</a> and
        <a class="green" href="https://github.com/Pixium">Pixium</a> with ❤️
      </p>
    </footer>
  </div>
</template>

<style scoped>
div,
main {
  height: 100%;
  width: 100%;
  display: flex;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}

footer {
  justify-self: flex-end;
}
</style>
