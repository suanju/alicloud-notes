<template>
  <div>
    <div class="h-8 w-full flex flex-row-reverse z-36" style="-webkit-app-region: drag">
      <div class="i-system-uicons-cross h-6 w-6 bg-gray-500 m-2" @click="cross"></div>
      <div
        class="i-system-uicons-scale-extend h-6 w-6 bg-gray-500 m-2"
        @click="maxus"
      ></div>
      <div class="i-system-uicons-minus h-6 w-6 bg-gray-500 m-2" @click="minus"></div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Quit, WindowMinimise, WindowMaximise } from "@wails/runtime";

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
