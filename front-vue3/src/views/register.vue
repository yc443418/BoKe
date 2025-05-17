<template>
    <Dialog :show="dialogConfig.show" :title="dialogConfig.title" :buttons="dialogConfig.buttons" width="400px"
        :showCancel=false @close="closeDialog">
        <el-form :model="formData" :rules="rules" ref="formDataRef" class="register-btn">
            <!--邮箱-->
            <el-form-item prop="email">
                <el-input prefix-icon="Message" size="large" clearable placeholder="请输入邮箱" v-model="formData.email" />
            </el-form-item>
            <!--邮箱验证码-->
            <el-form-item prop="emailCode">
                <div class="send-email-panel">
                    <el-input prefix-icon="View" size="large" clearable placeholder="请输入邮箱验证码"
                        v-model="formData.emailCode" />
                    <el-button type="primary" size="large" class="send-mail-btn" @click="sendMail">{{ codeNum == 60 ?
        "获取验证码" : `(${codeNum}后发送验证码)` }}</el-button>
                </div>
            </el-form-item>
            <!-- 账号 -->
            <el-form-item prop="username">
                <el-input prefix-icon="User" size="large" clearable placeholder="请输入账号" v-model="formData.username" />
            </el-form-item>
            <!-- 密码 -->
            <el-form-item prop="password">
                <el-input prefix-icon="Key" size="large" clearable placeholder="请输入密码" v-model="formData.password" />
            </el-form-item>
            <!-- 重复密码 -->
            <el-form-item prop="resetPassword">
                <el-input prefix-icon="Key" size="large" clearable placeholder="请输入重复密码"
                    v-model="formData.resetPassword" />
            </el-form-item>
            <el-form-item>
                <a href="javascript:void(0)" class="a_link" @click="goToLogin">已有账号?点击登录</a>
            </el-form-item>
            <!-- 注册按钮 -->
            <el-form-item>
                <el-button type="primary" class="register" @click="registerBtn">注册</el-button>
            </el-form-item>
        </el-form>
    </Dialog>
</template>

<script setup>
import { ref, reactive, getCurrentInstance, nextTick } from "vue"
const { proxy } = getCurrentInstance()
import { useStore } from 'vuex'

const store = useStore()

const showPanel = () => {
    resetForm()
}
defineExpose({ showPanel })
const dialogConfig = reactive({
    show: false,
    title: "标题"
})
const formData = ref({})
const formDataRef = ref()
const rules = {
    email: [{ required: true, message: "请输入邮箱" }],
    emailCode: [{ required: true, message: "请输入验证码" }],
    username: [{ required: true, message: "请输入昵称" }],
    password: [{ required: true, message: "请输入密码" }],
    resetPassword: [{ required: true, message: "请输入重复密码" }]
}

// 重置表单
const resetForm = () => {
    dialogConfig.show = true
    dialogConfig.title = '注册'
    nextTick(() => {
        formDataRef.value.resetFields()
    })
}

const closeDialog = () => {
    dialogConfig.show = false
}

// 发送邮箱验证码
const codeNum = ref(60)
const clearId = ref()
const isCheckSend = ref(false)
const sendMail = async () => {
    if (formData.value.email == undefined) {
        return proxy.$message.error('邮箱为空。请输入邮箱')
    }
    if (isCheckSend.value || codeNum.value != 60) {
        return
    } else {
        isCheckSend.value = true
        const res = await proxy.$api.qqMail(formData.value)
        console.log("邮箱验证码：", res)
        clearId.value = setInterval(() => {
            codeNum.value--
            if (codeNum.value == 0) {
                clearInterval(clearId.value)
                codeNum.value = 60
                isCheckSend.value = false
            }
        }, 1000)
        if (res.code != 200) {
            proxy.$message.error(res.message)
        } else {
            proxy.$message.success('发送邮箱验证码成功')
        }
    }
}

// 注册
const registerBtn = async () => {
    proxy.$refs.formDataRef.validate(valid => {
        if (valid) {
            proxy.$api.register(formData.value).then(res => {
                if (res.code !== 200) {
                    return proxy.$message.error(res.message)
                } else {
                    dialogConfig.show = false
                    proxy.$message.success("注册成功！请登录。")
                }
            })
        }
    })
}

// 跳转登录页面
const goToLogin = () => {
    dialogConfig.show = false
    store.commit('saveShowLogin', true)
}
</script>

<style lang="scss">
.register-btn {
    .send-email-panel {
        display: flex;
        width: 100%;
        justify-content: space-between;

        .send-mail-btn {
            margin-left: 5px;
        }
    }

    .send-email-panel {
        display: flex;
    }

    .a_link {
        text-decoration: none;
        color: #007fff;
    }

    .register {
        width: 100%;
    }
}
</style>