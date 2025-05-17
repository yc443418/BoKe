<template>
  <el-card>
    <!--搜索-->
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="用户账号" prop="username">
        <el-input placeholder="请输入用户账号" clearable v-model="queryParams.username" />
      </el-form-item>
      <el-form-item label="账号状态" prop="status">
        <el-select placeholder="请选择账号状态" v-model="queryParams.status">
          <el-option v-for="item in statusList" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item label="开始时间" prop="beginTime">
        <el-date-picker class="input-width" type="date" style="width: 199px;" value-format="YYYY-MM-DD" clearable
                        placeholder="请选择开始时间" v-model="queryParams.beginTime" />
      </el-form-item>
      <el-form-item label="结束时间" prop="endTime">
        <el-date-picker class="input-width" type="date" style="width: 199px;" value-format="YYYY-MM-DD" clearable
                        placeholder="请选择结束时间" v-model="queryParams.endTime" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button type="primary" icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <!--按钮操作-->
    <el-row :gutter="10" class="mb8">
      <!--搜索/刷新/隐藏-->
      <TipHover v-model:showSearch="showSearch" @queryTable="getUserList" :columns="columns" />
    </el-row>
    <!--表格-->
    <el-table border stripe style="width: 100%;" :header-cell-style="{ background: '#eef1f6', color: '#606266' }"
              :data="userList">
      <el-table-column label="ID" prop="id" v-if="false" />
      <el-table-column label="账号" prop="username" v-if="columns[0].visible" />
      <el-table-column label="状态" prop="status" width="150" v-if="columns[1].visible">
        <template #default="scope">
          <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="2" active-color="#13ce66"
                     inactive-color="#F5222D" active-text="启用" inactive-text="禁用" @change="userUpdateStatus(scope.row)">
          </el-switch>
        </template>
      </el-table-column>
      <el-table-column label="头像" prop="icon" v-if="columns[2].visible">
        <template #default="scope">
          <el-avatar shape="circle" :src="ip + scope.row.icon"></el-avatar>
        </template>
      </el-table-column>
      <el-table-column label="性别" prop="sex" v-if="columns[3].visible">
        <template #default="scope">
          <el-tag v-if="scope.row.sex == 1" type="success">男</el-tag>
          <el-tag v-else-if="scope.row.sex == 2" type="danger">女</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="文章数量" prop="quantity" v-if="columns[4].visible" />
      <el-table-column label="邮箱号" prop="email" v-if="columns[5].visible" />
      <el-table-column label="备注" prop="note" v-if="columns[6].visible" />
      <el-table-column label="登录ip" prop="loginIp" v-if="columns[7].visible" />
      <el-table-column label="登录地址" prop="loginAddress" v-if="columns[8].visible" />
      <el-table-column label="创建时间" prop="createTime" v-if="columns[9].visible" />
      <el-table-column label="更多操作" fixed="right" width="120">
        <template #default="scope">
          <el-button type="primary" link icon="ChatDotRound" @click="handleSendMessage(scope.row.id)"
                     v-hasPermission="['member:user:message']">发送消息</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--分页组件-->
    <Pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
                v-model:limit="queryParams.pageSize" @pagination="getUserList" />
    <!--发送消息-->
    <el-dialog :title="title" v-model="open" width="500px" @close="resetForm">
      <el-form ref="messageRef" :model="form" :rules="rules" label-width="80px">
        <el-form-item  prop="id" v-if="false">
          <el-input v-model="form.id"/>
        </el-form-item>
        <el-form-item label="消息内容" prop="messageContent">
          <el-input v-model="form.messageContent" type="textarea" placeholder="请输入内容" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div>
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button type="info" @click="cancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue'
const { proxy } = getCurrentInstance()
const statusList = ref([
  {
    value: '1',
    label: '启用'
  },
  {
    value: '2',
    label: '禁用'
  }
])
const queryParams = ref({})

// 数据列表
const ip = import.meta.env.VITE_APP_ACCESS_IP
const userList = ref([])
const total = ref(0)
const getUserList = async () => {
  const params = {
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
    username: queryParams.value.username,
    status: queryParams.value.status,
    beginTime: queryParams.value.beginTime,
    endTime: queryParams.value.endTime,
  }
  const res = await proxy.$api.queryUserList(params)
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    userList.value = res.data.list
    total.value = res.data.total
  }
}
getUserList()

// 搜索
const handleQuery = () => {
  getUserList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getUserList()
  proxy.$message.success('重置成功')
}

// 字段回显和隐藏
const showSearch = ref(true)
const columns = ref([
  { key: 0, label: `用户账号`, visible: true },
  { key: 1, label: `帐号状态`, visible: true },
  { key: 2, label: `头像`, visible: true },
  { key: 3, label: `性别`, visible: true },
  { key: 4, label: `发表文章数量`, visible: true },
  { key: 5, label: `邮箱`, visible: true },
  { key: 6, label: `备注`, visible: true },
  { key: 7, label: `登录ip`, visible: true },
  { key: 8, label: `登录地址`, visible: true },
  { key: 9, label: `创建时间`, visible: true },
])

// 状态修改
const userUpdateStatus = async (row) => {
  let text = row.status == 2 ? "禁用" : "启用"
  const confirmResult = await proxy.$confirm('确认要"' + text + '""' + row.username + '"用户吗?', "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).catch(err => err)
  if (confirmResult != 'confirm') {
    proxy.$message.info('已取消')
    await getUserList()
    return
  }
  const res = await proxy.$api.updateUserStatus(row.id, row.status)
  if (res.code != 200) {
    proxy.$message.error(res.message)
  } else {
    proxy.$message.success(text + "成功")
    await getUserList()
  }
}

// 发送消息相关定义
const rules = {
  messageContent: [{ required: true, message: "消息内容不能为空", trigger: "blur" }],
}
const title = ref("")
const form = ref({})
const open = ref(false)
const resetForm = () => {
  form.value = {}
  proxy.$refs["messageRef"].resetFields()
}

// 弹出发送消息对话框
const handleSendMessage = (id) => {
  open.value = true
  form.value.id = id
  console.log("用户id为：", form.value.id)
  title.value = '发送消息'
}

// 取消
const cancel = () => {
  open.value = false
  proxy.$message.info('已取消')
}

const submitForm = () => {
  proxy.$refs["messageRef"].validate(valid => {
    if (valid) {
      proxy.$api.sendMessage(form.value).then(res => {
        if (res.code != 200) {
          proxy.$message.error(res.message)
        } else {
          proxy.$message.success('发送消息成功')
          open.value = false
        }
      })
    }
  })
}


</script>

<style lang="scss" scoped></style>