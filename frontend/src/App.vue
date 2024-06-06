<template>
   <div class="w-screen h-screen">
      <div class="h-8  w-full flex flex-row-reverse" style="-webkit-app-region: drag;">
         <div class="i-system-uicons-cross  h-6 w-6 bg-gray-500 m-2" @click="cross">
         </div>
         <div class="i-system-uicons-scale-extend   h-6 w-6 bg-gray-500 m-2" @click="maxus">
         </div>
         <div class="i-system-uicons-minus h-6 w-6 bg-gray-500 m-2" @click="minus">
         </div>
      </div>
      <div>
         <RouterView />
      </div>
   </div>

</template>

<script setup lang="ts">
import { Quit, WindowMinimise, WindowMaximise } from '@walis/runtime'
const cross = () => {
   Quit()
}
const minus = () => {
   WindowMinimise()
}
const maxus = () => {
   WindowMaximise()
}
</script>
