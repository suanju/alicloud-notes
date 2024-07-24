<template>
  <div class="h-full"> <a-layout class="layout-demo">
      <a-layout-sider class="h-full" collapsible breakpoint="xl">
        <div class="logo" />
        <a-menu :default-open-keys="['1']" :default-selected-keys="['0_3']" :style="{ width: '100%' }"
          @menu-item-click="onClickMenuItem">
          <a-menu-item key="0_2">
            <IconCalendar></IconCalendar>
            Menu 2
          </a-menu-item>
          <a-menu-item key="0_3">
            <IconCalendar></IconCalendar>
            Menu 3
          </a-menu-item>
          <a-sub-menu key="0_4" title="我的文件夹">
            <template #icon>
              <IconCalendar></IconCalendar>
            </template>
            <a-tree class="ml-4" :data="[collectionDirectoryFramework]" :fieldNames="{
              key: 'value',
              title: 'name',
              children: 'children',
              icon: 'customIcon'
            }">
              <template #icon>
                <IconStar />
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
    </a-layout></div>
</template>
<script setup lang="ts">
import { Message } from '@arco-design/web-vue';
import { h, onMounted, ref } from 'vue';
import {
  IconCaretRight,
  IconCaretLeft,
  IconCalendar,
  IconStar,
  IconDriveFile
} from '@arco-design/web-vue/es/icon';
import { GetCollectionDirectoryFramework } from "@wails/go/backend/App"
import type { directory } from "@wails/go/models"

const onClickMenuItem = (key: String) => {
  Message.info({ content: `You select ${key}`, showIcon: true });
}

const treeData = [
  {
    title: 'Trunk',
    key: 'node1',
    children: [
      {
        title: 'Leaf',
        key: 'node2',
      },
    ],
  },
  {
    title: 'Trunk',
    key: 'node3',
    children: [
      {
        title: 'Leaf',
        key: 'node4',
        icon: () => h(IconDriveFile),
      },
      {
        title: 'Leaf',
        key: 'node5',
        icon: () => h(IconDriveFile),
      },
    ],
  },
];

console.log(treeData)
// 树状结构数据
const collectionDirectoryFramework = ref(<directory.DirectoryStructure>{})
const getFileTree = async () => {
  const GetCollectionDirectoryFrameworkResult = await GetCollectionDirectoryFramework()
  collectionDirectoryFramework.value = GetCollectionDirectoryFrameworkResult
  console.log(collectionDirectoryFramework.value)
}

onMounted(async () => {
  await getFileTree()
})
</script>
<style scoped>
::v-deep(.arco-layout) {
  height: 100%;
}
</style>
