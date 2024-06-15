<template>
   <v-card class="h-full">
        <v-layout>
          <v-navigation-drawer v-model="drawer" :rail="rail" permanent @click="rail = false">
            <v-list-item prepend-avatar="./public/vite.svg" title="John Leider" nav>
              <template v-slot:append>
                <v-btn icon="mdi-chevron-left" variant="text" @click.stop="rail = !rail"></v-btn>
              </template>
            </v-list-item>
            <v-divider></v-divider>
            <v-list density="compact" nav>
              <v-list-item prepend-icon="mdi-collage" title="最近文件" value="最近文件"></v-list-item>
              <v-list-item prepend-icon="mdi-folder" title="文件夹" value="文件夹"></v-list-item>
            </v-list>
          </v-navigation-drawer>
          <v-main class="h-full">
            <RouterView/>
          </v-main>
        </v-layout>
      </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const drawer = ref(true)
const rail = ref(true)
</script>