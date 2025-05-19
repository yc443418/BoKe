<template>
  <el-dropdown trigger="click">
    <span>
      <img v-if="!icon" src="./../assets/image/logo.jpeg" class="admin-icon"/>
      <img v-else :src="ip + icon" class="admin-icon"/>
      <el-icon class="el-icon--right"><arrow-down /></el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item><span @click="goToPersonal">个人信息</span></el-dropdown-item>
        <el-dropdown-item><span @click="logout">退出登录</span></el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup>
import { ref, getCurrentInstance, watch } from "vue"
import { useRouter } from 'vue-router'
const router = useRouter()
const { proxy } = getCurrentInstance()

const ip = import.meta.env.VITE_APP_ACCESS_IP

// 头像监听
const icon = ref(null)
watch(
    () => proxy.$store.state.sysAdmin.icon,
    (newVal, oldVal) => {
      icon.value = newVal
      // console.log("当前的头像为：", ip + newVal)
    },
    { immediate: true, deep: true }
)

// 跳转个人信息页面
const goToPersonal = ()=> {
  router.push('/system/personal')
}

// 退出登录
const logout = async ()=> {
  const confirmResult = await proxy.$confirm('确定要退出登录吗？, 是否继续', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).catch(err => err)
  if (confirmResult != 'confirm') {
    proxy.$message.info('已取消')
    return
  }
  proxy.$storage.clearAll()
  proxy.$message.success('退出成功')
  router.push('/login')
}
</script>

<style lang="scss">
.el-dropdown {
  .admin-icon {
    width: 40px;
    height: 40px;
    border-radius: 20px;
  }

  .el-icon {
    cursor: pointer;
  }
}
</style>