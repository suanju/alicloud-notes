<template>
  <div class="w-screen h-screen flex flex-col">
    <div class="h-8 w-full flex flex-row-reverse z-36" style="-webkit-app-region: drag;">
      <div class="i-system-uicons-cross  h-6 w-6 bg-gray-500 m-2" @click="cross">
      </div>
      <div class="i-system-uicons-scale-extend  h-6 w-6 bg-gray-500 m-2" @click="maxus">
      </div>
      <div class="i-system-uicons-minus h-6 w-6 bg-gray-500 m-2" @click="minus">
      </div>
    </div>
    <div class="h-full">
      <RouterView />
    </div>
  </div>

</template>

<script setup lang="ts">
import { Quit, WindowMinimise, WindowMaximise, EventsOn, EventsEmit } from '@wails/runtime'
import { useUserStore } from '@/store/user';


const userStore = useUserStore()
const cross = () => {
  Quit()
}
const minus = () => {
  WindowMinimise()
}
const maxus = () => {
  WindowMaximise()
}

//登录过期
EventsOn("login_expired", () => {
  EventsEmit("logout-success")
  userStore.logout();
})

</script>
