<template>
    <div class="font-sans">
        <div class="relative min-h-screen flex flex-col sm:justify-center items-center  ">
            <div class="relative sm:max-w-sm w-full">
                <div class="card bg-blue-400 shadow-lg  w-full h-full rounded-3xl absolute  transform -rotate-6"></div>
                <div class="card bg-red-400 shadow-lg  w-full h-full rounded-3xl absolute  transform rotate-6"></div>
                <div class="relative w-full rounded-3xl  px-6 py-4 bg-gray-100 shadow-md">
                    <div class="bg-white p-8 rounded-lg shadow-lg max-w-sm w-full text-center">
                        <h1 class="text-2xl font-semibold mb-4">Alicloud-Notes Login </h1>
                        <div class="mb-4">
                            <img :src="qrcodeResult.qrCode ? qrcodeResult.qrCode : Lodaing" alt="阿里云盘二维码"
                                class="mx-auto h-48 w-48">
                        </div>
                        <p class="mb-4">{{ tips }}</p>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script setup lang="ts">
import { CreatInstance, GenerateQrcode, QrcodeState } from "@wails/go/backend/App"
import type { types } from "@wails/go/models";
import { onMounted, ref } from "vue";
import { useUserStore } from "@/store/user";
import { useRouter } from "vue-router";
import Lodaing from "@/assets/loading.svg"
import { Message } from '@arco-design/web-vue';

const userStore = useUserStore()
const router = useRouter()
const qrcodeResult = ref(<types.GenerateQrcodeResp>{})
const tips = ref("请用阿里云盘 App 扫码")
let cl: NodeJS.Timeout
onMounted(async () => {
    qrcodeResult.value = await GenerateQrcode()
    cl = setInterval(async () => {
        try {
            //获取验证码状态
            const qrcodeStateResult = await QrcodeState(<types.QrcodeStateReq>{
                t: qrcodeResult.value.t,
                ck: qrcodeResult.value.ck
            })
            userStore.setRefreshToken(qrcodeStateResult.refreshToken)
            clearInterval(cl)
            //根据refreshToken获取accessToken并且得到实例
            const creatInstanceResult = await CreatInstance({ refreshToken: qrcodeStateResult.refreshToken })
            userStore.setUserInfo(creatInstanceResult)
            Message.success("登录成功")
            router.push({ name: "Home" })
            console.log("CreatInstance Response", creatInstanceResult)
        } catch (err: any) {
            console.log("tips", err)
            tips.value = err
        }
    }, 1000)
})
</script>

<style scoped></style>