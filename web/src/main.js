import {createApp} from "vue";
import App from "./App.vue";
import router from "./router/index.js"; // If using a router

const app = createApp(App);

app.use(router); // Use the router in the Vue app if needed

app.mount("#app");