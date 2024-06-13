import { router } from '@/router/router';
import { createApp } from 'vue'
// Vuetify
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { aliases, mdi } from 'vuetify/iconsets/mdi'


import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css';
import '@/style.css'
import 'virtual:uno.css'
import App from './App.vue'

const vuetify = createVuetify({
    components,
    directives,
    icons: {
      defaultSet: 'mdi',
      aliases,
      sets: {
        mdi,
      },
    },
  })

createApp(App).use(router).use(vuetify).mount('#app')
