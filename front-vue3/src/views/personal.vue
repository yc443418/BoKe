<template>
  <div class="body-container">
    <!--首页和修改信息-->
    <div class="user-banner">
      <router-link to="/home" class="link-info">首页</router-link>
      <span class="iconfont icon-xiangyoujiantou"></span>
      <span>修改信息</span>
    </div>
    <!--修改信息-->
    <div class="update-user">
      <el-tabs v-model="activeName">
        <el-tab-pane label="修改资料" name="1">
          <el-form
            :model="currentUserInfo"
            :rules="editFormRules"
            ref="editFormRefForm"
            label-width="80px"
          >
            <el-form-item label="用户头像" prop="icon">
              <el-upload
                :headers="headers"
                class="avatar-uploader"
                :action="uploadIconUrl"
                :show-file-list="false"
                :on-success="handleAvatarSuccess"
              >
                <img v-if="icon" :src="ip + icon" class="avatar" title="点击更换头像" />
                <img
                  v-else="!icon"
                  :src="ip + currentUserInfo.icon"
                  class="avatar"
                  title="点击更换头像"
                />
              </el-upload>
            </el-form-item>
            <el-form-item label="用户账号" prop="username">
              <el-input v-model="currentUserInfo.username" />
            </el-form-item>
            <el-form-item label="用户备注" prop="note">
              <el-input v-model="currentUserInfo.note" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="saveUser">保存</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="修改密码" name="2">
          <el-form
            :model="updateForm"
            :rules="updateFormRules"
            ref="updateFormRefForm"
            label-width="80px"
          >
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
              <el-button type="primary" @click="saveUserPwd">保存</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup>
import { ref, getCurrentInstance, watch } from "vue";
import { useStore } from "vuex";
const store = useStore();
import { useRouter } from "vue-router";
const { proxy } = getCurrentInstance();
const router = useRouter();

const activeName = ref("1");
// 获取当前登录的用户和图片访问ip
const currentUserInfo = ref({});
watch(
  () => proxy.$store.state.user,
  (newVal, oldVal) => {
    if (newVal == undefined) {
      router.push("/home");
    } else {
      currentUserInfo.value = newVal;
    }
  },
  { immediate: true, deep: true }
);

const ip = import.meta.env.VITE_APP_ACCESS_IP;

const uploadIconUrl = import.meta.env.VITE_APP_UPLOAD_IP;
const headers = ref({
  Authorization: "Bearer " + proxy.$store.state.token,
});
const icon = ref("");
const editFormRules = {
  username: [{ required: true, message: "请输入用户账号", trigger: "blur" }],
  note: [{ required: true, message: "请输入用户备注", trigger: "blur" }],
};
const handleAvatarSuccess = (res, file) => {
  icon.value = res.data;
};

// 修改个人信息提交
const saveUser = async () => {
  proxy.$refs.editFormRefForm.validate(async (valid) => {
    if (!valid) return;
    const res = await proxy.$api.updateUserInfo({
      icon: icon.value == "" ? currentUserInfo.value.icon : icon.value,
      username: currentUserInfo.value.username,
      note: currentUserInfo.value.note,
    });
    console.log("数据：", currentUserInfo)
    if (res.code != 200) {
      proxy.$message.error(res.message);
    } else {
      store.commit("saveUser", null);
      store.commit("saveUser", res.data);
      proxy.$message.success("修改资料成功");
    }
  });
};

// 修改密码
const updateForm = ref({});
const updateFormRules = {
  oldPassword: [{ required: true, message: "请输入旧密码", trigger: "blur" }],
  newPassword: [{ required: true, message: "请输入新密码", trigger: "blur" }],
  resetPassword: [{ required: true, message: "请输入重复密码", trigger: "blur" }],
};
const saveUserPwd = async () => {
  proxy.$refs.updateFormRefForm.validate(async (valid) => {
    if (!valid) return;
    const res = await proxy.$api.updateUserPassword({
      oldPassword: updateForm.value.oldPassword,
      newPassword: updateForm.value.newPassword,
      resetPassword: updateForm.value.resetPassword,
    });
    if (res.code != 200) {
      this.$message.error(res.message);
    } else {
      proxy.$storage.clearItem("user");
      store.commit("saveUser", res.data.user);
      proxy.$message.success("修改密码成功");
      updateForm.value = {};
    }
  });
};
</script>

<style lang="scss">
// 首页和修改信息
.user-banner {
  color: #9ba7b9;
  line-height: 35px;

  .link-info {
    text-decoration: none;
    color: #007fff;
  }

  .icon-xiangyoujiantou {
    padding: 0px 5px;
  }
}
// 修改信息
.update-user {
  background: #ffffff;
  padding: 5px 15px 0 15px;
  margin-bottom: 10px;

  // 修改资料头像
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
}
</style>
