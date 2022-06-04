import { createApp } from 'vue'
import App from './App.vue'
import router from './router/index.js'
import './assets/css/style.css'
import Toast from 'vue-toastification';
import "vue-toastification/dist/index.css";
import { store } from './plugins/store.js'

// Toast options
const options = {
    // You can set your default options here
};

createApp(App)
.use(router)
.use(store)
.use(Toast, options)
.mount('#app')
 