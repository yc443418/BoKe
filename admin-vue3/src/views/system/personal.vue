<template>
  <el-row :gutter="20">
    <!--个人信息-->
    <el-col :xs="24" :sm="24" :md="8" :lg="6" :xl="5" style="margin-bottom: 10px">
      <el-card>
        <div style="text-align: center">
          <el-avatar :src="ip + sysAdminInfo.icon" />
        </div>
        <el-form label-width="100px">
          <el-row>
            <el-col :span="24">
              <el-form-item label="用户账号：">{{ sysAdminInfo.username }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="用户昵称：">{{ sysAdminInfo.nickname }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="用户邮箱：">{{ sysAdminInfo.email }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="用户电话：">{{ sysAdminInfo.phone }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="用户备注：">{{ sysAdminInfo.note }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="创建时间：">{{ sysAdminInfo.createTime }}</el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </el-card>
    </el-col>
    <!--基本资料-->
    <el-col :xs="24" :sm="24" :md="16" :lg="18" :xl="19">
      <el-card>
        <el-tabs v-model="activeName">
          <el-tab-pane label="基本资料" name="first">
            <el-form :model="sysAdminInfo" :rules="editFormRules" ref="editFormRefForm" label-width="80px">
              <el-form-item label="用户头像" prop="icon">
                <el-upload :headers=headers class="avatar-uploader" :action="uploadIconUrl"
                           :show-file-list="false" :on-success="handleAvatarSuccess">
                  <img v-if="icon" :src="ip + icon" class="avatar" title="点击更换头像">
                  <img v-else="!icon" :src="ip + sysAdminInfo.icon" class="avatar" title="点击更换头像">
                </el-upload>
              </el-form-item>
              <el-form-item label="用户账号" prop="username">
                <el-input v-model="sysAdminInfo.username" />
              </el-form-item>
              <el-form-item label="用户昵称" prop="nickname">
                <el-input v-model="sysAdminInfo.nickname" />
              </el-form-item>
              <el-form-item label="手机号码" prop="phone">
                <el-input v-model="sysAdminInfo.phone" maxlength="11" />
              </el-form-item>
              <el-form-item label="用户邮箱" prop="email">
                <el-input v-model="sysAdminInfo.email" maxlength="50" />
              </el-form-item>
              <el-form-item label="用户备注" prop="note">
                <el-input v-model="sysAdminInfo.note" />
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="submitFirst">保存</el-button>
                <el-button type="danger" @click="closeFirst">关闭</el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
          <el-tab-pane label="修改密码" name="second">
            <el-form :model="updateForm" :rules="updateFormRules" ref="updateFormRefForm" label-width="80px">
              <el-form-item label="旧密码" prop="oldPassword">
                <el-input v-model="updateForm.oldPassword" placeholder="请输入旧密码" />
              </el-form-item>
              <el-form-item label="新密码" prop="newPassword">
                <el-input v-model="updateForm.newPassword" placeholder="请输入新密码" />
              </el-form-item>
              <el-form-item label="确认密码" prop="resetPassword">
                <el-input v-model="updateForm.resetPassword" placeholder="请确认密码" />
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="submitSecond">保存</el-button>
                <el-button type="danger" @click="closeSecond">关闭</el-button>
              </el-form-item>
            </el-form>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue'
import { useRouter } from "vue-router"
import { useStore } from 'vuex'

const { proxy } = getCurrentInstance()
const router = useRouter()
const store = useStore()

// 修改个人信息提交
const sysAdminInfo = proxy.$store.state.sysAdmin
const ip = import.meta.env.VITE_APP_ACCESS_IP
const activeName = ref('first')
const editFormRules = {
  username: [{ required: true, message: '请输入用户账号', trigger: 'blur' }],
  email: [{ required: true, message: '请输入用户邮箱', trigger: 'blur' }],
  nickname: [{ required: true, message: '请输入用户昵称', trigger: 'blur' }],
  phone: [{ required: true, message: '请输入用户手机', trigger: 'blur' }],
  note: [{ required: true, message: '请输入用户备注', trigger: 'blur' }],
}
const uploadIconUrl = import.meta.env.VITE_APP_UPLOAD_IP
const headers = ref({
  Authorization: 'Bearer ' + proxy.$store.state.token
})
const icon = ref('')
const handleAvatarSuccess = (res, file) => {
  icon.value = res.data
}
const submitFirst = async () => {
  proxy.$refs.editFormRefForm.validate(async valid => {
    if (!valid) return
    const res = await proxy.$api.adminUpdatePersonal({
      icon: icon.value == '' ? sysAdminInfo.icon : icon.value,
      username: sysAdminInfo.username,
      email: sysAdminInfo.email,
      nickname: sysAdminInfo.nickname,
      phone: sysAdminInfo.phone,
      note: sysAdminInfo.note,
    })
    if (res.code != 200) {
      proxy.$message.error(res.message);
    } else {
      proxy.$storage.clearItem("sysAdmin")
      store.commit('saveSysAdmin', res.data)
      router.push('/home')
      proxy.$message.success('修改用户成功')
    }
  })
}
const closeFirst = () => {
  router.push('/home')
}

// 修改密码提交
const updateForm = ref({})
const updateFormRules = {
  oldPassword: [{ required: true, message: '请输入旧密码', trigger: 'blur' }],
  newPassword: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
  resetPassword: [{ required: true, message: '请输入重复密码', trigger: 'blur' }],
}
const submitSecond = async () => {
  proxy.$refs.updateFormRefForm.validate(async valid => {
    if (!valid) return
    const res = await proxy.$api.adminUpdatePersonalPassword({
      oldPassword: updateForm.value.oldPassword,
      newPassword: updateForm.value.newPassword,
      resetPassword: updateForm.value.resetPassword
    })
    if (res.code != 200) {
      this.$message.error(res.message);
    } else {
      proxy.$storage.clearAll()
      router.push("/login")
      proxy.$message.success('修改密码成功')
    }
  })
}
const closeSecond = () => {
  router.push('/home')
}
</script>

<style lang="scss">
.el-avatar {
  height: 130px;
  width: 130px;
}

.avatar-uploader {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  width: 100px;
  height: 100px;
  position: relative;
  overflow: hidden;

  .avatar {
    width: 100px;
    height: 100px;
    display: block;
  }
}
</style>