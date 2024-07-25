<template>
    <div class="flex flex-col mt-3.4">
        <div class="w-full flex">
            <div class="flex-1">
                <a-input placeholder="请输入标题" :max-length="30" allow-clear show-word-limit />
            </div>
            <div class="w-48 w-full flex flex-row-reverse">
                <a-button type="primary" class="mr-6 w-20" shape="round">
                    <template #default>保存</template>
                </a-button>
            </div>
        </div>
        <div class="mt-2">
            <QuillEditor theme="snow" class="h-full" />
        </div>

    </div>
</template>

<script setup lang="ts">
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';

components: {
    QuillEditor
}

</script>

<style scoped>
::v-deep(.arco-input-wrapper) {
    background-color: #fff;
}

::v-deep(.arco-input-size-medium) {
    background-color: #fff;
}
</style>