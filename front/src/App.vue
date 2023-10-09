<script setup>
import { RouterView } from 'vue-router'
import  NavBar  from './components/NavBar.vue'
</script>

<template>
  <NavBar />
  <RouterView />
</template>

<style>
.nav-bar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 999;
}

.router-view {
  width: 200%;
}
</style>