<template>
  <div class="h-full">
    <a-layout class="layout-demo">
      <a-layout-sider class="h-full" :width="200" :collapsed="false">
        <div />
        <!-- 头部部分 -->
        <div class="w-full flex items-center"></div>
        <a-menu :style="{ width: '100%' }" @menu-item-click="onClickMenuItem">
          <a-sub-menu class="mt-2" key="0_1" title="我的文件夹">
            <template #icon>
              <IconCalendar></IconCalendar>
            </template>
            <a-tree
              class="ml-4"
              :data="[collectionDirectoryFolders]"
              :fieldNames="{
                key: 'value',
                title: 'name',
                children: 'children',
                icon: 'customIcon',
              }"
            >
              <template #icon>
                <IconFolder />
              </template>
            </a-tree>
          </a-sub-menu>
        </a-menu>
        <!-- trigger -->
        <template #trigger="{ collapsed }">
          <IconCaretRight v-if="collapsed"></IconCaretRight>
          <IconCaretLeft v-else></IconCaretLeft>
        </template>
      </a-layout-sider>
      <a-layout>
        <RouterView />
      </a-layout>
    </a-layout>
  </div>
</template>
<script setup lang="ts">
import { Message } from "@arco-design/web-vue";
import { onMounted, ref } from "vue";
import {
  IconCaretRight,
  IconCaretLeft,
  IconCalendar,
  IconFolder,
} from "@arco-design/web-vue/es/icon";
import { GetCollectionDirectoryFolders } from "@wails/go/backend/App";
import type { directory } from "@wails/go/models";

const onClickMenuItem = (key: String) => {
  Message.info({ content: `You select ${key}`, showIcon: true });
};

// 树状结构数据
const collectionDirectoryFolders = ref(<directory.DirectoryStructure>{});
const getFileTree = async () => {
  const GetCollectionDirectoryFoldersResult = await GetCollectionDirectoryFolders();
  collectionDirectoryFolders.value = GetCollectionDirectoryFoldersResult;
  console.log(collectionDirectoryFolders.value);
};

onMounted(async () => {
  await getFileTree();
});
</script>
<style scoped>
::v-deep(.arco-layout) {
  height: 100%;
}
</style>
