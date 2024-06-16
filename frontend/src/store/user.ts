import { defineStore } from "pinia";
import { ref } from "vue";
import type { aliyunpan } from "@walis/go/models";
export interface UserInfo {
  refreshToken: string;
  user: aliyunpan.UserInfo;
}

export const useUserStore = defineStore("user", () => {
  const user = ref(<UserInfo>{});
  const isLogin = () => {
    return !!user.value?.refreshToken;
  };
  const setUserInfo = (info :aliyunpan.UserInfo) =>{
    user.value.user = info
  }
  return {
    user,
    isLogin,
    setUserInfo
  };
});
