<script>
import PageLayout from './layout/PageLayout.vue'
import Home from './pages/Home.vue'
import {RouterView} from 'vue-router'
export default{
  name: "App",
  components: {
    PageLayout,
    RouterView,
    Home,
  }
}
</script>

<template>
  <div id="app">
    <PageLayout/>
    <RouterView />

  </div>
  <!-- <div id="app"> -->
    <!-- <h1>App</h1> -->
    <!-- <div v-if="isMobile"> -->
    <!-- <RouterView /> -->
    <!-- </div> -->
  <!-- </div> -->
</template>

<style>
  #app {
    /* 홈 */
    position: relative;
    width: 390px;
    height: 844px;

    /* BACKGROUND */
    background: #F5F5F5;
    margin: 0 auto;
  }
</style>