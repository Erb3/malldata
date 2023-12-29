<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
  <nav class="flex items-center space-x-6 text-sm font-medium">
    <router-link to="/" class="transition-colors hover:text-foreground/80 text-foreground/60">Home</router-link>
    <router-link to="/" class="transition-colors hover:text-foreground/80 text-foreground/60">Map</router-link>
  </nav>
</template>
