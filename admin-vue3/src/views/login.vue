<template>
  <div class="login-container">
    <el-form class="login-form" :rules="rules" :model="loginForm" ref="loginFormRef">
      <!--标题-->
      <div class="title-container">博客运营后台</div>
      <!--用户名-->
      <el-form-item style="margin-bottom: 30px" prop="username">
        <el-input placeholder="请输入用户名" name="username" type="text" prefix-icon="User"
                  v-model="loginForm.username"></el-input>
      </el-form-item>
      <!--密码-->
      <el-form-item style="margin-bottom: 30px" prop="password">
        <el-input placeholder="请输入密码" name="password" :type="flagType" v-model="loginForm.password">
          <template #prefix>
            <el-icon>
              <lock/>
            </el-icon>
          </template>
          <template #suffix>
            <span @click="changeView">
              <el-icon v-if="flag == true" style="cursor: pointer">
                <Hide/>
              </el-icon>
              <el-icon v-else-if="flag == false" style="cursor: pointer">
                <View/>
              </el-icon>
            </span>
          </template>
        </el-input>
      </el-form-item>
      <!--登录按钮-->
      <el-button type="primary" style="width: 100%; margin-bottom: 30px" @click="handleLogin">登录</el-button>
    </el-form>
  </div>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
const router = useRouter()
const { proxy } = getCurrentInstance()
// 表单校验验证
const rules = {
  username: [{required: true, message: "请输入用户名", trigger: "blur"}],
  password: [{required: true, message: "请输入密码", trigger: "blur"}],
}

const loginForm = ref({})

// 切换小眼睛时间
const flagType = ref("password")
const flag = ref(true)
const changeView = () => {
  flag.value = !flag.value
  flagType.value = flag.value ? "password" : "text"
}

// 登录
const loginFormRef = ref()
const store = useStore()
const handleLogin = () => {
  loginFormRef.value.validate(valid => {
    if (!valid) return
    proxy.$api.login(loginForm.value).then(res => {
      if (res.code != 200) {
        proxy.$message.error(message)
      } else {
        // console.log("请求的响应数据：", res)
        proxy.$message.success("登录成功")
        store.commit('saveToken', res.data.token)
        store.commit('saveSysAdmin', res.data.sysAdmin)
        store.commit('saveLeftMenuList', res.data.leftMenuList)
        store.commit('savePermissionList', res.data.permissionList)
        router.push('/home')
      }
    }).catch(err => {
      console.log(err)
    })
  })
}
</script>

<style lang="scss">
.login-container {
  background-color: #2d3a4b;
  height: 100%;
  .login-form {
    width: 400px;
    border-radius: 1px;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    .title-container {
      font-size: 30px;
      line-height: 1.5;
      text-align: center;
      margin-bottom: 30px;
      font-weight: bold;
    }
  }
}
</style>