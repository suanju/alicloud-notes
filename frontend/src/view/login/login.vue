<template>
  <div class="flex flex-col">
    <div class="h-8 w-full flex flex-row-reverse z-36" style="-webkit-app-region: drag">
      <WindowControlBar></WindowControlBar>
    </div>

    <div class="font-sans">
      <div class="relative min-h-screen flex flex-col sm:justify-center items-center">
        <div class="relative sm:max-w-sm w-full">
          <div
            class="card bg-blue-400 shadow-lg w-full h-full rounded-3xl absolute transform -rotate-6"
          ></div>
          <div
            class="card bg-red-400 shadow-lg w-full h-full rounded-3xl absolute transform rotate-6"
          ></div>
          <div class="relative w-full rounded-3xl px-6 py-4 bg-gray-100 shadow-md">
            <div class="bg-white p-8 rounded-lg shadow-lg max-w-sm w-full text-center">
              <h1 class="text-2xl font-semibold mb-4">Alicloud-Notes Login</h1>
              <div class="mb-4 flex">
                <img
                  :src="qrcodeResult.qrCode ? qrcodeResult.qrCode : Lodaing"
                  alt="阿里云盘二维码"
                  class="mx-auto h-48 w-48"
                />
              </div>
              <div class="flex flex-col items-center">
                <p class="mb-4">{{ tips }}</p>
                <div v-show="isExpire" @click="reload" class="i-system-uicons-reset-alt size-4 bg-blue-400"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import WindowControlBar from "@/components/window_control_bar/window_control_bar.vue";
import { CreatInstance, GenerateQrcode, QrcodeState } from "@wails/go/backend/App";
import type { types } from "@wails/go/models";
import { onMounted, ref } from "vue";
import { useUserStore } from "@/store/user";
import { useRouter } from "vue-router";
import Lodaing from "@/assets/loading.svg";
import { Message } from "@arco-design/web-vue";

const userStore = useUserStore();
const router = useRouter();
const qrcodeResult = ref(<types.GenerateQrcodeResp>{});
const tips = ref("请用阿里云盘 App 扫码");
const isExpire = ref(false)
let cl: NodeJS.Timeout;

onMounted(() =>{
  getQrcode()  
});


const getQrcode = async () => {
  qrcodeResult.value = await GenerateQrcode();
  cl = setInterval(async () => {
    try {
      //获取验证码状态
      const result = await QrcodeState(<types.QrcodeStateReq>{
        t: qrcodeResult.value.t,
        ck: qrcodeResult.value.ck,
      });
      userStore.setRefreshToken(result.refreshToken);
      clearInterval(cl);
      //根据refreshToken获取accessToken并且得到实例
      const creatInstanceResult = await CreatInstance({
        refreshToken: result.refreshToken,
      });
      userStore.setUserInfo(creatInstanceResult);
      Message.success("登录成功");
      router.push({ name: "Home" });
      console.log("CreatInstance Response", creatInstanceResult);
    } catch (err: any) {
      console.log("tips", err);
      if(err === "二维码已过期"){
        isExpire.value = true
        clearInterval(cl);
      }
      tips.value = err;
    }
  }, 1000);
}
const reload = () => {
  isExpire.value = false
  tips.value = "请用阿里云盘 App 扫码"
  getQrcode()
}
</script>

<style scoped></style>
