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
            <v-treeview :items="collectionDirectoryFramework.children" class="custom-treeview">
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
import { onMounted, ref } from 'vue';
import { VTreeview } from 'vuetify/labs/VTreeview'
import { GetCollectionDirectoryFramework } from "@wails/go/backend/App"
import { directory } from "@wails/go/models"
components: {
  VTreeview
}

console.log()


const drawer = ref(true)
const rail = ref(false)
// 树状结构数据
const collectionDirectoryFramework = ref(<directory.DirectoryStructure>{})


onMounted(async () => {
  collectionDirectoryFramework.value = await GetCollectionDirectoryFramework()
})
</script>
