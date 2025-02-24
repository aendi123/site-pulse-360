<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <nav class="navbar bg-body-tertiary">
    <div class="container-fluid">
      <a class="navbar-brand" href="/">
        <img src="./assets/logo.png" alt="Logo" width="100" height="100" class="d-inline-block align-text-center">
        Site Pulse 360
      </a>
    </div>
  </nav>
  
  <main>
    <RouterView />
  </main>
</template>