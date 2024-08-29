<script>
import Home from './pages/Home.vue'
import {RouterView} from 'vue-router'
export default{
  name: "App",
  components: {
    RouterView,
    Home,
  }
}
</script>

<template>
  <div id="app">
    <h1>App</h1>
    <!-- <div v-if="isMobile"> -->
    <RouterView />
    <!-- </div> -->
  </div>
</template>
