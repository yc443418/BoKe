<template>
    <v-md-editor :model-value=modelValue :height="height + 'px'" :disabled-menus="[]"
        :include-level="[1, 2, 3, 4, 5, 6]" @change="change" @upload-image="uploadImageHandler">
    </v-md-editor>
</template>

<script setup>
import VMdEditor from "@kangc/v-md-editor"
import "@kangc/v-md-editor/lib/style/base-editor.css"
import githubTheme from "@kangc/v-md-editor/lib/theme/github.js"
import "@kangc/v-md-editor/lib/theme/style/github.css"
import hljs from "highlight.js"
import { getCurrentInstance } from "vue"
const { proxy } = getCurrentInstance()

// 属性定义
const props = defineProps({
    modelValue: {
        type: String,
        default: "",
    },
    height: {
        type: Number,
        default: 650,
    }
})

// 代码高亮风格
VMdEditor.use(githubTheme, {
    Hljs: hljs
})

// 对外方法
const emit = defineEmits()
const change = (markdownContent, htmlContent) => {
    emit("update:modelValue", markdownContent)
    emit("htmlContent", htmlContent)
}

// 图片上传
const ip = import.meta.env.VITE_APP_ACCESS_IP
const uploadImageHandler = async (event, insertImage, files) => {
    const formData = new FormData()
    formData.append("file", files[0])
    const res = await proxy.$api.upload(formData)
    if (res.code != 200) {
        proxy.$message.error(res.message)
    } else {
        insertImage({
            url: ip + res.data,
            desc: "图片"
        })
    }
}

</script>

<style lang="scss"></style>