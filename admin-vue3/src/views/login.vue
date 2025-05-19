<template>
  <div class="login-container">
    <div class="background-overlay"></div>

    <el-form class="login-form" :rules="rules" :model="loginForm" ref="loginFormRef">
      <!-- 头部区域 -->
      <div class="header-wrapper">
        <img src="./../assets/image/homeLogo.png" class="logo" alt="logo">
        <div class="title-wrapper">
          <h1 class="main-title">CHEN博客社区后台</h1>
          <p class="sub-title">欢迎您的使用</p>
        </div>
      </div>

      <!-- 用户名输入 -->
      <el-form-item prop="username">
        <el-input
            v-model="loginForm.username"
            placeholder="请输入用户名"
            size="large"
            class="modern-input"
        >
          <template #prefix>
            <el-icon class="input-icon"><User /></el-icon>
          </template>
        </el-input>
      </el-form-item>

      <!-- 密码输入 -->
      <el-form-item style="margin-bottom: 30px" prop="password">
        <el-input placeholder="请输入密码" name="password" :type="flagType" v-model="loginForm.password">
          <template #prefix>
            <el-icon>
              <lock/>
            </el-icon>
          </template>
          <template #suffix>
        <span @click="changeView" class="eye-clickable">
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

      <!-- 登录按钮 -->
      <el-button
          type="primary"
          size="large"
          class="login-btn"
          @click="handleLogin"
      >
        登 录
      </el-button>
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

<style lang="scss" scoped>
.login-container {
  position: relative;
  min-height: 100vh;
  background: url('./../assets/image/bg.jpg') no-repeat center/cover;
  display: flex;
  align-items: center;
  justify-content: center;

  .background-overlay {
    position: absolute;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.4);
    backdrop-filter: blur(2px);
  }

  .login-form {
    position: relative;
    z-index: 1;
    width: 520px;
    padding: 50px 40px;
    background: rgba(255, 255, 255, 0.95);
    border-radius: 20px;
    box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
    transform: translateY(-10%);

    .header-wrapper {
      text-align: center;
      margin-bottom: 40px;

      .logo {
        width: 100px;
        height: 100px;
        margin-bottom: 20px;
        filter: drop-shadow(0 4px 6px rgba(0, 0, 0, 0.1));
      }

      .title-wrapper {
        .main-title {
          font-size: 28px;
          color: #303133;
          margin-bottom: 8px;
          letter-spacing: 2px;
        }

        .sub-title {
          color: #909399;
          font-size: 14px;
          letter-spacing: 1px;
        }
      }
    }

    :deep(.modern-input) {
      .eye-wrapper {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 32px;
        height: 32px;
        border-radius: 50%;
        background: rgba(255, 255, 255, 0.9);
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
        margin-right: 8px;
        transition: all 0.3s ease;
        cursor: pointer;

        &:hover {
          background: #fff;
          box-shadow: 0 3px 12px rgba(64, 158, 255, 0.3);

          .eye-icon {
            color: #409EFF;
          }
        }
      }

      :deep(.el-input) {
        .eye-icon {
          color: rgba(0, 0, 0, 0.5) !important;
          margin-right: 8px;
          font-size: 18px;
          background: rgba(255, 255, 255, 0.9);
          padding: 4px;
          border-radius: 50%;
          box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
          transition: all 0.3s ease;

          &:hover {
            color: #409EFF !important;
            transform: scale(1.15);
            box-shadow: 0 2px 8px rgba(64, 158, 255, 0.3);
          }

          &.eye-active {
            color: #67C23A !important;
            background: rgba(103, 194, 58, 0.1);
          }
        }
      }

      .eye-icon {
        font-size: 20px;
        color: #606266;
        transition: all 0.3s ease;

        /* 当密码可见时显示渐变色 */
        &[data-visible="true"] {
          background: linear-gradient(45deg, #409EFF, #9254de);
          -webkit-background-clip: text;
          background-clip: text;
          color: transparent;
        }
      }
    }

    .login-btn {
      width: 100%;
      height: 56px;
      font-size: 18px;
      letter-spacing: 4px;
      border-radius: 12px;
      background: linear-gradient(135deg, #409EFF, #3375b9);
      border: none;
      transition: all 0.3s;

      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 8px 20px rgba(64, 158, 255, 0.3);
      }

      &:active {
        transform: translateY(0);
      }
    }
  }
  .eye-clickable {
    cursor: pointer;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    padding: 4px;

    .el-icon {
      transition: transform 0.2s;

      &:hover {
        transform: scale(1.1);
      }
    }
  }
}
</style>