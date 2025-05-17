<template>
  <el-card>
    <!--搜索-->
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="用户账号" prop="username">
        <el-input v-model="queryParams.username" placeholder="请输入用户账号"></el-input>
      </el-form-item>
      <el-form-item label="用户状态" prop="status">
        <el-select placeholder="请选择用户状态" v-model="queryParams.status" style="width: 220px">
          <el-option v-for="item in adminStatusList" :key="item.value" :label="item.label" :value="item.value"/>
        </el-select>
      </el-form-item>
      <el-form-item label="开始时间" prop="beginTime">
        <el-date-picker class="input-width" type="date" style="width: 199px" value-format="YYYY-MM-DD" clearable
                        placeholder="请选择开始时间" v-model="queryParams.beginTime"/>
      </el-form-item>
      <el-form-item label="结束时间" prop="endTime">
        <el-date-picker class="input-width" type="date" style="width: 199px" value-format="YYYY-MM-DD" clearable
                        placeholder="请选择结束时间" v-model="queryParams.endTime"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button type="primary" icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <!--表格按钮操作-->
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="Plus" @click="handleAdd" v-hasPermission="['system:sysAdmin:add']">新增</el-button>
      </el-col>
      <!-- 搜索/刷新/隐藏组件-->
      <TipHover v-model:showSearch="showSearch" @queryTable="getAdminList" :columns="columns"></TipHover>
    </el-row>
    <!-- 列表-->
    <el-table border stripe style="width: 100%" :header-cell-style="{ background: '#eef1f6', color: '#606266' }"
              :data="adminList">
      <el-table-column label="ID" prop="id" v-if="false"/>
      <el-table-column label="用户账号" prop="username" v-if="columns[0].visible"/>
      <el-table-column label="昵称" prop="nickname" v-if="columns[1].visible"/>
      <el-table-column label="手机号" prop="phone" v-if="columns[2].visible"/>
      <el-table-column label="邮箱号" prop="email" v-if="columns[3].visible"/>
      <el-table-column label="头像" prop="icon" v-if="columns[4].visible">
        <template #default="scope">
          <el-avatar shape="circle" :src="ip + scope.row.icon" />
        </template>
      </el-table-column>
      <el-table-column label="创建时间" prop="createTime" v-if="columns[5].visible" width="170"/>
      <el-table-column label="状态" prop="status" v-if="columns[6].visible" width="150">
        <template #default="scope">
          <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="2" active-text="启用"
                     inactive-text="禁用" active-color="#13ce66"
                     inactive-color="#F5222D" @change="adminUpdateStatus(scope.row)"/>
        </template>
      </el-table-column>
      <el-table-column label="角色名称" prop="roleName" v-if="columns[7].visible"/>
      <el-table-column label="性别" prop="sex" v-if="columns[8].visible"/>
      <el-table-column label="备注" prop="note" v-if="columns[9].visible"/>
      <el-table-column label="更多操作" fixed="right" width="300">
        <template #default="scope">
          <el-button type="primary" link icon="Edit" @click="handleUpdate(scope.row.id)" v-hasPermission="['system:sysAdmin:edit']">编辑</el-button>
          <el-button type="danger" link icon="Delete" @click="handleDelete(scope.row.id)" v-hasPermission="['system:sysAdmin:delete']">删除</el-button>
          <el-button type="info" link icon="Setting" @click="handleResetPwd(scope.row)" v-hasPermission="['system:sysAdmin:reset']">重置密码</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--分页-->
    <Pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
                v-model:limit="queryParams.pageSize" @pagination="getAdminList"/>
    <!--新增/编辑-->
    <el-dialog :title="title" v-model="open" width="600px" @close="resetForm">
      <el-form :model="form" :rules="rules" ref="userRef" label-width="80px">
        <el-row>
          <el-col :span="12">
            <el-form-item label="用户昵称" prop="nickname">
              <el-input v-model="form.nickname" placeholder="请输入用户昵称" maxlength="30" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="用户角色" prop="roleId">
              <el-select v-model="form.roleId" placeholder="请选择角色">
                <el-option v-for="item in roleList" :key="item.id" :label="item.roleName"
                           :value="item.id"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="手机号码" prop="phone">
              <el-input v-model="form.phone" placeholder="请输入手机号码" maxlength="11" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="用户邮箱" prop="email">
              <el-input v-model="form.email" placeholder="请输入邮箱" maxlength="50" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item v-if="form.id == undefined" label="用户名称" prop="username">
              <el-input v-model="form.username" placeholder="请输入用户名称" maxlength="30" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item v-if="form.id == undefined" label="用户密码" prop="password">
              <el-input v-model="form.password" placeholder="请输入用户密码" type="password" maxlength="20"
                        show-password />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="用户性别" prop="sex">
              <el-radio-group v-model="form.sex">
                <el-radio :label="1">男</el-radio>
                <el-radio :label="2">女</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="用户状态" prop="status">
              <el-radio-group v-model="form.status">
                <el-radio :label="1">启用</el-radio>
                <el-radio :label="2">禁用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <el-form-item label="用户备注" prop="note">
              <el-input v-model="form.note" type="textarea" placeholder="请输入内容"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
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
const adminStatusList = ref([
  {
    value: '1',
    label: '启用'
  },
  {
    value: '2',
    label: '禁用'
  }
])

// 数据列表
const queryParams = ref({})
const ip = import.meta.env.VITE_APP_ACCESS_IP
const adminList = ref([])
const total = ref(0)
const getAdminList = async ()=> {
  const params = {
    username: queryParams.value.username,
    status: queryParams.value.status,
    beginTime: queryParams.value.beginTime,
    endTime: queryParams.value.endTime,
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
  }
  const res = await proxy.$api.queryAdminList(params)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    adminList.value = res.data.list
    total.value = res.data.total
    // console.log("用户列表数据：", res.data.list)
  }
}
getAdminList()

// 搜索
const handleQuery = () => {
  getAdminList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getAdminList()
  proxy.$message.success('重置成功')
}

// 字段回显
const showSearch = ref(true)
const columns = ref([
  { key: 0, label: `用户账号`, visible: true},
  { key: 1, label: `昵称`, visible: true},
  { key: 2, label: `手机号`, visible: true},
  { key: 3, label: `邮箱号`, visible: true},
  { key: 4, label: `头像`, visible: true},
  { key: 5, label: `创建时间`, visible: true},
  { key: 6, label: `状态`, visible: true},
  { key: 7, label: `角色名称`, visible: true},
  { key: 8, label: `性别`, visible: true},
  { key: 9, label: `备注`, visible: true},
])

// 获取角色列表
const getRoleList = async () => {
  const res = await proxy.$api.querySysRoleVoList()
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    roleList.value = res.data
    // console.log("角色列表数据：", res.data)
  }
}
getRoleList()

// 新增和编辑定义
const rules = {
  deptId: [{ required: true, message: '请选择部门', trigger: 'blur' }],
  postId: [{ required: true, message: '请选择岗位', trigger: 'blur' }],
  roleId: [{ required: true, message: '请选择角色', trigger: 'blur' }],
  username: [{ required: true, message: '请输入用户账号', trigger: 'blur' }],
  nickname: [{ required: true, message: '请输入用户昵称', trigger: 'blur' }],
  password: [{ required: true, message: "用户密码不能为空", trigger: "blur" }, {
    min: 5,
    max: 20,
    message: "用户密码长度必须介于 5 和 20 之间",
    trigger: "blur"
  }],
  status: [{ required: true, message: '请选择状态', trigger: 'blur' }],
  email: [{ type: "email", message: "请输入正确的邮箱地址", trigger: ["blur", "change"] }],
  phone: [{ pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: "请输入正确的手机号码", trigger: "blur" }]
}
const title = ref("")
const form = ref({})
const open = ref(false)
const roleList = ref([])
const resetForm = () => {
  form.value = {}
  proxy.$refs["userRef"].resetFields()
}

// 弹出新增页面
const handleAdd = () => {
  open.value = true
  title.value = '新增用户'
}

// 弹出修改页面
const handleUpdate = (id) => {
  proxy.$api.adminInfo(id).then(response => {
    form.value = response.data;
    open.value = true;
    title.value = '修改用户'
  })
}

// 新增和编辑取消操作
const cancel = () => {
  open.value = false
  proxy.$message.info('已取消')
}

// 新增和编辑提交操作
const submitForm = () => {
  proxy.$refs["userRef"].validate(valid => {
    if (valid) {
      if (form.value.id != undefined) {
        proxy.$api.adminUpdate(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('修改成功')
            open.value = false
            getAdminList()
          }
        })
      } else {
        proxy.$api.addAdmin(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('新增成功')
            open.value = false
            getAdminList()
          }
        })
      }
    }
  })
}

// 删除
const handleDelete = async (id) => {
  const confirmResult = await proxy.$confirm('是否要删除编号为' + id + '的数据，' + '是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).catch(err => err)
  if (confirmResult != 'confirm') {
    proxy.$message.info('已取消')
    return
  }
  const res = await proxy.$api.deleteAdmin(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('删除成功')
    getAdminList()
  }
}

// 重置密码
const handleResetPwd = async (row) => {
  proxy.$prompt('请输入"' + row.username + '"的新密码', "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    closeOnClickModal: false,
    inputPattern: /^.{5,20}$/,
    inputErrorMessage: "用户密码长度必须介于 5 和 20 之间"
  }).then(({ value }) => {
    proxy.$api.resetPassword(row.id, value).then(response => {
      proxy.$message.success("修改成功，新密码是：" + value);
    })
  }).catch(() => {
  })
}

// 修改状态
const adminUpdateStatus = async (row) => {
  let text = row.status == 2 ? "禁用" : "启用"
  const confirmResult = await proxy.$confirm('确认要"' + text + '""' + row.username + '"用户吗?', "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).catch(err => err)
  if (confirmResult != 'confirm') {
    proxy.$message.info('已取消')
    await getAdminList()
    return
  }
  const res = await proxy.$api.updateAdminStatus(row.id, row.status)
  if (res.code != 200) {
    proxy.$message.error(res.message)
  } else {
    proxy.$message.success(text + "成功")
    await getAdminList()
  }
}
</script>

<style lang="scss">

</style>