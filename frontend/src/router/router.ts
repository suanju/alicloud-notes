import { useUserStore } from "@/store/user";
import { createRouter, createWebHistory } from "vue-router";

import Login from "@/view/login/login.vue";
import Layout from "@/view/layout.vue";

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      component: Layout,
      children: [
        {
          path: "/",
          name: "Home",
          component: () => import("@/view/home/home.vue"),
        },
      ],
    },
    { path: "/login", name: "Login", component: Login },
  ],
});

router.beforeEach(async (to, _) => {
  const userStore = useUserStore();
  if (!userStore.isLogin() && to.name !== "Login") {
    return { name: "Login" };
  }
});
