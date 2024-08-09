<template>
  <div class="w-full flex flex-col items-center justify-center mt-4">
    <div class="w-4.4/5"><a-input-search class=" " placeholder="" /></div>
    <div>
      <div class="flex w-full">
        <IconLeft />
        <div></div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useglobalStore } from "@/store/global";
import { GetDirectoryFrameworkByPath } from "@wails/go/backend/App";
import { directory } from "@wails/go/models";
import { onMounted, ref, watch } from "vue";
import { IconLeft } from "@arco-design/web-vue/es/icon";

const globalStore = useglobalStore();
const filesList = ref(<directory.DirectoryStructure[]>[]);

watch(
  () => globalStore.selectedCatalog,
  async (path, _) => {
    const result = await GetDirectoryFrameworkByPath(path);
    filesList.value = result.children as directory.DirectoryStructure[];
    console.log(result);
  }
);

onMounted(async () => {});
</script>

<style scoped></style>
