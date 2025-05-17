<template>
    <div class="avatar" :style="{ width: width + 'px', height: width + 'px', 'border-radius': width / 2 + 'px' }">
        <el-image v-if="userId"
            :style="{ width: width + 'px', height: width + 'px', 'border-radius': width / 2 + 'px' }" :userId="userId"
            :src="ip + src" @click="goToUser" />
        <div v-else class="no-login">
            未登录
        </div>
    </div>
</template>

<script setup>
import { getCurrentInstance } from "vue"
import { useRouter } from "vue-router"
const { proxy } = getCurrentInstance()
const router = useRouter()

// 属性定义
const props = defineProps({
    userId: {
        type: Number
    },
    src: {
        type: String
    },
    width: {
        type: Number,
        default: 50,
    },
    // 点击头像是否跳转
    showIcon: {
        type: Boolean,
        default: true,
    },
})

const ip = import.meta.env.VITE_APP_ACCESS_IP

const goToUser = () => {
    if (props.showIcon) {
        router.push("/user/" + proxy.userId)
    }
}
</script>

<style lang="scss">
.avatar {
    display: flex;
    cursor: pointer;
    background: #f0f0f0;
    align-items: center;
    overflow: hidden;

    .no-login {
        width: 100%;
        text-align: center;
        font-size: 13px;
    }
}
</style>