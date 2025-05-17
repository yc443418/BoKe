<template>
    <!--头部-->
    <Header></Header>
    <!--主页内容-->
    <div class="body-content">
        <router-view />
    </div>
</template>

<script setup>
import Header from '@/components/header/header.vue'
</script>

<style lang="scss">
.body-content {
    margin-top: 60px;
    position: relative;
}
</style>