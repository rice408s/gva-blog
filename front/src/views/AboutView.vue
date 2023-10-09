<script setup>
import HelloWorld from "../components/HelloWorld.vue"
</script>

<template>
  <div class="about">
    <HelloWorld class="hello" msg="这是一个博客" />
  </div>
</template>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}

.hello {
    margin: 0 auto;
    text-align: center;
    font-size: 3rem;
    /* color: #42b983; */
    font-weight: 300;
    margin-bottom: 2rem;

}

</style>
