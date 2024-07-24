import { defineStore } from "pinia";
import { ref } from "vue";
import type { aliyunpan } from "@wails/go/models";
import { router } from "@/router/router";
export interface UserInfo {
  refreshToken: string;
  user: aliyunpan.UserInfo;
}

export const useUserStore = defineStore(
  "user",
  () => {
    const user = ref(<UserInfo>{});
    const isLogin = () => {
      return !!user.value?.refreshToken;
    };
    const setRefreshToken = (token: string) => {
      user.value.refreshToken = token;
    };
    const setUserInfo = (info: aliyunpan.UserInfo) => {
      user.value.user = info;
    };

    const logout = () => {
      user.value = {} as UserInfo
      router.push({ name: "Login" })
    }
    return {
      user,
      isLogin,
      setRefreshToken,
      setUserInfo,
      logout,
    };
  },
  {
    persist: {
      storage: localStorage,
      paths: ["user"],
    },
  }
);
