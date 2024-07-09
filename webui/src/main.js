import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import MessaggioErrore from './components/MessaggioErrore.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("MessaggioErrore", MessaggioErrore);

app.use(router)
app.mount('#app')
