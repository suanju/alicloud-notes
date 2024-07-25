<template>
  <a-layout>
    <a-layout-sider class="w-2xl">
      <div class="w-full flex items-center justify-center mt-4">
        <div class="w-4.4/5"><a-input-search class=" " placeholder="" /></div>
      </div>
    </a-layout-sider>
    <a-layout-content>
      <div class="flex flex-col">
        <div>
          <div
            class="h-8 w-full flex flex-row-reverse z-36"
            style="-webkit-app-region: drag"
          >
            <div
              class="i-system-uicons-cross h-6 w-6 bg-gray-500 m-2"
              @click="cross"
            ></div>
            <div
              class="i-system-uicons-scale-extend h-6 w-6 bg-gray-500 m-2"
              @click="maxus"
            ></div>
            <div
              class="i-system-uicons-minus h-6 w-6 bg-gray-500 m-2"
              @click="minus"
            ></div>
          </div>
        </div>
        <Editer></Editer>
      </div>
    </a-layout-content>
  </a-layout>
</template>
<script setup lang="ts">
import Editer from "@/components/editer/editer.vue";
import { Quit, WindowMinimise, WindowMaximise } from "@wails/runtime";

component: {
  Editer;
}

const cross = () => {
  Quit();
};
const minus = () => {
  WindowMinimise();
};
const maxus = () => {
  WindowMaximise();
};
</script>

<style scoped>
::v-deep(.arco-input-search) {
  border-radius: 16px;
  background-color: #fff;
  border-color: #999;
}
</style>
