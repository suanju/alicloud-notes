<template>
  <div class="w-screen h-screen flex flex-col">
    <div class="h-full">
      <RouterView />
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  EventsOn,
  EventsEmit,
} from "@wails/runtime";
import { useUserStore } from "@/store/user";

const userStore = useUserStore();

//登录过期
EventsOn("login_expired", () => {
  EventsEmit("logout-success");
  userStore.logout();
});
</script>
