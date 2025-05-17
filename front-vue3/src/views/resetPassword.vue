<template>
    <Dialog :show="dialogConfig.show" :title="dialogConfig.title" :buttons="dialogConfig.buttons" width="400px"
        :showCancel=false @close="closeDialog">
        <el-form :model="formData" :rules="rules" ref="formDataRef" class="resetPwd-btn">
            <!-- 邮箱 -->
            <el-form-item prop="email">
                <el-input prefix-icon="Message" size="large" clearable placeholder="请输入邮箱" v-model="formData.email" />
            </el-form-item>
            <!-- 邮箱验证码 -->
            <el-form-item prop="emailCode">
                <div class="send-email-panel">
                    <el-input prefix-icon="View" size="large" clearable placeholder="请输入邮箱验证码"
                        v-model="formData.emailCode" />
                    <el-button type="primary" size="large" class="send-mail-btn" @click="sendMail">{{ codeNum == 60 ?
        "获取验证码" : `(${codeNum})后发送验证码` }}</el-button>
                </div>
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
            <!-- 重置密码按钮 -->
            <el-form-item>
                <el-button type="primary" class="resetPwd" @click="resetPwdBtn">重置密码</el-button>
            </el-form-item>
        </el-form>
    </Dialog>
</template>

<script setup>
import { ref, reactive, getCurrentInstance, nextTick } from "vue"
import { useStore } from 'vuex'

const store = useStore()
const { proxy } = getCurrentInstance()

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
    password: [{ required: true, message: "请输入密码" }],
    resetPassword: [{ required: true, message: "请输入重复密码" }]
}

// 重置表单
const resetForm = () => {
    dialogConfig.show = true
    dialogConfig.title = '重置密码'
    nextTick(() => {
        formDataRef.value.resetFields()
    })
}

// 发送qq邮箱验证码
const codeNum = ref(60)
const clearId = ref()
const isClickSend = ref(false)
const sendMail = async () => {
    if (formData.value.email === undefined) {
        return proxy.$message.error("邮箱为空，请输入邮箱!")
    }
    if (isClickSend.value || codeNum.value != 60) {
        return
    } else {
        isClickSend.value = true
        const res = await proxy.$api.qqMail(formData.value)
        // console.log("邮箱验证码：", res)
        clearId.value = setInterval(() => {
            codeNum.value--
            if (codeNum.value == 0) {
                clearInterval(clearId.value)
                codeNum.value = 60
                isClickSend.value = false
            }
        }, 1000)
        if (res.code !== 200) {
            proxy.$message.error(res.message)
        } else {
            proxy.$message.success("发送邮箱验证码成功!")
        }
    }
}

// 重置密码
function resetPwdBtn() {
    proxy.$refs.formDataRef.validate(valid => {
        if (valid) {
            proxy.$api.resetPwd(formData.value).then(res => {
                if (res.code !== 200) {
                    return proxy.$message.error(res.message)
                } else {
                    dialogConfig.show = false
                    proxy.$message.success("重置密码成功！请登录。")
                }
            })
        }
    })
}

// 保存是否弹出注册页面
const closeDialog = () => {
    dialogConfig.show = false
    store.commit("saveShowResetPwd", false)
}

// 跳转登录页面
const goToLogin = () => {
    dialogConfig.show = false
    store.commit("saveShowLogin", true)
}
</script>

<style lang="scss" scoped>
.resetPwd-btn {
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

    .resetPwd {
        width: 100%;
    }
}
</style>