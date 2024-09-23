import {createApp} from "vue";
import App from "./App.vue";
import router from "./router/index";
import store from './store'; // 스토어 추가

const app = createApp(App);

app.use(router); // Use the router in the Vue app if needed
app.use(store); 
app.mount("#app");