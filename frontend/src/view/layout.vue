<template>
  <v-card class="h-full">
    <v-layout>
      <v-navigation-drawer v-model="drawer" :rail="rail" permanent width="240">
        <v-list-item prepend-avatar="./public/vite.svg" title="Notes" nav>
        </v-list-item>
        <v-divider></v-divider>
        <v-list density="compact" nav>
          <!-- 文件夹树状结构 -->
          <v-list-group value="收藏夹">
            <template v-slot:activator="{ props }">
              <v-list-item v-bind="props" prepend-icon="mdi-folder" title="收藏夹" :ripple="false"></v-list-item>
            </template>
            <v-treeview :items="items" class="custom-treeview">
              <template v-slot:prepend="{ item, isActive }">
                <v-icon>
                  {{ item.children ? (isActive ? 'mdi-folder-open' : 'mdi-folder') : 'mdi-file' }}
                </v-icon>
              </template>
              <template v-slot:title="{ item }">
                <span class="custom-treeview-title text-xs">{{ item.name }}</span>
              </template>
            </v-treeview>
          </v-list-group>
        </v-list>
      </v-navigation-drawer>
      <v-main class="h-full">
        <RouterView />
      </v-main>
    </v-layout>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { VTreeview } from 'vuetify/labs/VTreeview'
components: {
  VTreeview
}

const drawer = ref(true)
const rail = ref(false)
  
// 树状结构数据
const items = ref([
  {
    id: 1,
    name: '文件夹1',
    children: [
      { id: 2, name: '子文件1' },
      { id: 3, name: '子文件2' }
    ]
  },
  {
    id: 4,
    name: '文件夹2',
    children: [
      { id: 5, name: '子文件3' },
      { id: 6, name: '子文件4' }
    ]
  }
])
</script>
