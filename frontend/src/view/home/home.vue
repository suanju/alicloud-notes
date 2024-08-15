<template>
  <a-layout>
    <a-layout-sider class="w-2xl">
      <SideBar></SideBar>
    </a-layout-sider>
    <a-layout-content>
      <div class="flex flex-col h-full overflow-hidden">  
        <WindowControlBar></WindowControlBar>
        <Editer></Editer>
      </div>
    </a-layout-content>
  </a-layout>
</template>
<script setup lang="ts">
import SideBar from "@/components/side_bar/side_bar.vue";
import WindowControlBar from "@/components/window_control_bar/window_control_bar.vue";
import Editer from "@/components/editer/editer.vue";
</script>

<style scoped>
::v-deep(.arco-input-search) {
  border-radius: 16px;
  background-color: #fff;
  border-color: #999;
}
</style>
