import { defineStore } from "pinia";
import { ref } from "vue";

export interface UserInfo {
  accessToken: AccessToken;
  user: any;
}

export interface AccessToken {
  accessTokenType: string;
  accessToken: string;
  refreshToken: string;
  expiresIn: number;
  expireTime: string;
}

export const useUserStore = defineStore("user", () => {
  const user = ref(<UserInfo>{});
  const isLogin = () => {
    return !!user.value?.accessToken?.accessToken;
  };
  return {
    user,
    isLogin,
  };
});
