import { router } from '@/router/router';
import { createApp } from 'vue'
import '@/style.css'
import 'virtual:uno.css'
import App from './App.vue'

createApp(App).use(router).mount('#app')
