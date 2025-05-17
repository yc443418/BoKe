<template>
    <el-image-viewer :initial-index="previewImgIndex" hide-on-click-modal :url-list="imageList" @close="closeImgViewer"
        v-if="previewImgIndex != null">
    </el-image-viewer>
</template>

<script setup>
import { ref } from 'vue'
const props = defineProps({
    imageList: {
        type: Array
    }
})
const previewImgIndex = ref(null)

// 开始滚动
const startScroll = () => {
    document.body.style.overflow = "auto"
}

// 禁止滚动
const stopScroll = () => {
    document.body.style.overflow = "hidden"
}

// 关闭
const closeImgViewer = () => {
    startScroll()
    previewImgIndex.value = null
}

// 开始展示
const show = (index) => {
    stopScroll()
    previewImgIndex.value = index
}

defineExpose({ show })
</script>

<style lang="scss"></style>