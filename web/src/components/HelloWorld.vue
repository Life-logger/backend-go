<script setup>
import { ref } from 'vue'

defineProps({
  msg: String,
})

const count = ref(0)
</script>

<template>
  <h1>{{ msg }}</h1>

  Let's Go TeamBuddy

  <button @click="helloWorldAction">Hello World</button>

</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>

<script>
import axios from '../store/axios'

export default {
  name: 'HelloWorld',
  methods: {
    async helloWorldAction() {
      try {

        // Call the helloWorld action to get hello from Backend Server
        axios.get('/hello')
                .then(response => {
                    console.log(response)
                })
                .catch(error => {
                    reject(error)
                })
      } catch (error) {
        console.error('Error receiving HelloWorld:', error);
      }
    },
  },
};
</script>
