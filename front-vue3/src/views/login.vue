<template>
    <Dialog :show="dialogConfig.show" :title="dialogConfig.title" :buttons="dialogConfig.buttons" width="400px"
        :showCancel=false @close="closeDialog">
        <el-form :model="formData" :rules="rules" ref="formDataRef" class="login-btn">
            <el-form-item prop="email">
                <el-input prefix-icon="Message" size="large" clearable placeholder="请输入邮箱"
                    v-model="formData.email"></el-input>
            </el-form-item>
            <el-form-item prop="password">
                <el-input prefix-icon="Key" size="large" clearable placeholder="请输入密码"
                    v-model="formData.password"></el-input>
            </el-form-item>
            <el-form-item prop="checkCode">
                <div class="check-code-panel">
                    <el-input prefix-icon="View" size="large" clearable placeholder="请输入验证码"
                        v-model="formData.checkCode"></el-input>
                    <img class="check-code" @click="getCode" :src="imageCode"></img>
                </div>
            </el-form-item>
            <el-form-item>
                <div class="no-account">
                    <a href="javascript:void(0)" class="a-link" @click="goToRegister">没有账号?点击注册</a>
                    <a href="javascript:void(0)" class="a-link" @click="goToResetPwd">忘记密码?点击重置密码</a>
                </div>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" class="op-btn" @click="loginBtn">登录</el-button>
            </el-form-item>
        </el-form>
    </Dialog>
</template>

<script setup>
import { ref, reactive, getCurrentInstance, nextTick } from "vue"
const { proxy } = getCurrentInstance()
import { useStore } from 'vuex'

const store = useStore()

const dialogConfig = reactive({
    show: false,
    title: "标题"
})

const closeDialog = () => {
    dialogConfig.show = false
    store.commit('saveShowLogin', false)
}

const formData = ref({
    email: '',
    password: '',
    checkCode: '',
    idKey: '',
})
const formDataRef = ref()
const rules = {
    email: [{ required: true, message: "请输入邮箱" }],
    password: [{ required: true, message: "请输入密码" }],
    checkCode: [{ required: true, message: "请输入验证码" }]
}

const resetForm = () => {
    dialogConfig.show = true
    dialogConfig.title = '登录'
    nextTick(() => {
        getCode()
        formDataRef.value.resetFields()
    })
}
const showPanel = () => {
    resetForm()
}
defineExpose({ showPanel })

// 获取验证码
const imageCode = ref("")
const getCode = () => {
    proxy.$api.captcha().then(res => {
        imageCode.value = res.data.image
        formData.value.idKey = res.data.idKey
        // console.log("验证码：", res.data.image)
    }).catch(error => {
        console.log(error)
    })
}

// 登录
const loginBtn = () => {
    proxy.$refs.formDataRef.validate(valid => {
        if (valid) {
            proxy.$api.login(formData.value).then(res => {
                // console.log("登录的数据",res.data)
                if (res.code != 200) {
                    proxy.$message.error(res.message)
                } else {
                    dialogConfig.show = false
                    store.commit("saveToken", res.data.token)
                    store.commit("saveUser", res.data.user)
                    window.location.reload()
                }
            })
        }
    })
}

// 跳转注册页面
const goToRegister = () => {
    dialogConfig.show = false
    store.commit('saveShowRegister', true)
}

// 跳转到重置密码页面
const goToResetPwd = ()=> {
    dialogConfig.show = false
    store.commit("saveShowResetPwd", true)
}
</script>

<style lang="scss">
.login-btn {
    .check-code-panel {
        display: flex;

        .check-code {
            width: 142px;
            margin-left: 5px;
            cursor: pointer;
        }
    }

    .no-account {
        width: 100%;
        display: flex;
        justify-content: space-between;

        .a-link {
            text-decoration: none;
            color: #007fff;
        }
    }

    .op-btn {
        width: 100%;
    }

}
</style>