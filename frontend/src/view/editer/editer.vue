<template>
    <div>
        <QuillEditor theme="snow" />
    </div>
</template>

<script setup lang="ts">
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';

components: {
    QuillEditor
}

</script>