<template>
  <el-card>
    <!--搜索-->
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="角色名称" prop="roleName">
        <el-input v-model="queryParams.roleName" placeholder="请输入角色名称"></el-input>
      </el-form-item>
      <el-form-item label="角色状态" prop="status">
        <el-select placeholder="请选择角色状态" v-model="queryParams.status" style="width: 220px">
          <el-option v-for="item in roleStatusList" :key="item.value" :label="item.label" :value="item.value"/>
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
        <el-button type="primary" plain icon="Plus" @click="handleAdd" v-hasPermission="['system:sysRole:add']">新增</el-button>
      </el-col>
      <!-- 搜索/刷新/隐藏组件-->
      <TipHover v-model:showSearch="showSearch" @queryTable="getRoleList" :columns="columns"></TipHover>
    </el-row>
    <!-- 列表-->
    <el-table border stripe style="width: 100%" :header-cell-style="{ background: '#eef1f6', color: '#606266' }"
              :data="roleList">
      <el-table-column label="ID" prop="id" v-if="false"/>
      <el-table-column label="角色名称" prop="roleName" v-if="columns[0].visible"/>
      <el-table-column label="角色标识" prop="roleKey" v-if="columns[1].visible"/>
      <el-table-column label="创建时间" prop="createTime" v-if="columns[2].visible"/>
      <el-table-column label="角色状态" prop="status" v-if="columns[3].visible">
        <template #default="scope">
          <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="2" active-text="启用"
                     inactive-text="禁用" active-color="#13ce66"
                     inactive-color="#F5222D" @click="roleUpdateStatus(scope.row)"/>
        </template>
      </el-table-column>
      <el-table-column label="角色备注" prop="description" v-if="columns[4].visible"/>
      <el-table-column label="更多操作" fixed="right" width="300">
        <template #default="scope">
          <el-button type="primary" link icon="Edit" @click="handleUpdate(scope.row.id)" v-hasPermission="['system:sysRole:edit']">编辑</el-button>
          <el-button type="danger" link icon="Delete" @click="handleDelete(scope.row.id)" v-hasPermission="['system:sysRole:delete']">删除</el-button>
          <el-button type="info" link icon="Setting" @click="handlePermission(scope.row)" v-hasPermission="['system:sysRole:assign']">分配权限</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--分页-->
    <Pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
                v-model:limit="queryParams.pageSize" @pagination="getRoleList"/>
    <!--新增/编辑对话框-->
    <el-dialog :title="title" v-model="open" width="500px" @close="resetForm">
      <el-form ref="roleRef" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="角色名称" prop="roleName">
          <el-input v-model="form.roleName" placeholder="请输入角色名称"/>
        </el-form-item>
        <el-form-item label="角色标识" prop="roleKey">
          <el-input v-model="form.roleKey" placeholder="请输入角色标识"/>
        </el-form-item>
        <el-form-item label="角色状态" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="2">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="角色备注" prop="description">
          <el-input v-model="form.description" placeholder="请输入角色备注" type="textarea"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm">确定</el-button>
          <el-button type="info" @click="cancel">取消</el-button>
        </div>
      </template>
    </el-dialog>
    <!--分配权限-->
    <el-dialog :title="titlePermission" v-model="openPermission" width="300px" @close="resetFormPermission">
      <el-tree :data="permissionList" :props="{value: 'id', label: 'menuName'}" show-checkbox node-key="id"
        default-expand-all :default-checked-keys="defKeys" ref="treeRef" />
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="allotPermission">确定</el-button>
          <el-button type="info" @click="cancelPermission">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import {getCurrentInstance, ref} from "vue"

const {proxy} = getCurrentInstance()

const roleStatusList = ref([
  {
    value: '1',
    label: '启用',
  },
  {
    value: '2',
    label: '禁用',
  }
])

// 列表数据
const queryParams = ref({})
const roleList = ref([])
const total = ref(0)
const getRoleList = async () => {
  const params = {
    roleName: queryParams.value.roleName,
    status: queryParams.value.status,
    beginTime: queryParams.value.beginTime,
    endTime: queryParams.value.endTime,
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
  }
  const res = await proxy.$api.queryRoleList(params)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    roleList.value = res.data.list
    total.value = res.data.total
    // console.log("角色列表数据：", res.data.list)
  }
}
getRoleList()

// 搜索
const handleQuery = () => {
  getRoleList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getRoleList()
  proxy.$message.success('重置成功')
}
// 字段回显
const showSearch = ref(true)
const columns = ref([
  { key: 0, label: `角色名称`, visible: true},
  { key: 1, label: `角色标识`, visible: true},
  { key: 2, label: `创建时间`, visible: true},
  { key: 3, label: `角色状态`, visible: true},
  { key: 4, label: `角色备注`, visible: true},
])

// 新增/编辑定义
const rules = {
  roleName: [{required: true, message: "角色名称不能为空", trigger: "blur"}],
  roleKey: [{required: true, message: "角色标识不能为空", trigger: "blur"}],
  status: [{required: true, message: "角色状态不能为空", trigger: "blur"}],
}
const title = ref("")
const form = ref({})
const open = ref(false)
const resetForm = ()=> {
  form.value = {}
  proxy.$refs["roleRef"].resetFields()
}

// 弹出新增对话框
const handleAdd = ()=> {
  open.value = true
  title.value = '新增角色'
}

// 弹出修改角色对话框
const handleUpdate = (id)=> {
  proxy.$api.roleInfo(id).then(res=>{
    form.value = res.data
    open.value = true
    title.value = '修改角色'
  })
}

// 取消
const cancel = ()=> {
  open.value = false
  proxy.$message.info('已取消')
}

// 新增和修改提交操作
const submitForm = () => {
  proxy.$refs["roleRef"].validate(valid => {
    if (valid) {
      if (form.value.id != undefined) {
        proxy.$api.roleUpdate(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
            return
          } else {
            proxy.$message.success('修改成功')
            open.value = false
            getRoleList()
          }
        })
      } else {
        proxy.$api.addRole(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('新增成功')
            open.value = false
            getRoleList()
          }
        })
      }
    }
  })
}

// 状态修改
const roleUpdateStatus = async (row)=> {
  let text = row.status == 2 ? '禁用' : '启用'
  const confirmResult = await proxy.$confirm('确认要"' + text + '""' + row.roleName + '"角色吗?', "警告", {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).catch(err => err)
  if (confirmResult != 'confirm') {
    proxy.$message.info('已取消')
    await getRoleList()
    return
  }
  const res = await proxy.$api.updateRoleStatus(row.id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('修改成功')
    getRoleList()
  }
}

// 删除
const handleDelete = async (id) => {
  const confirmResult = await proxy.$confirm('是否要删除编号为' + id + '的数据' + '是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).catch(err => err)
  if (confirmResult != 'confirm') {
    proxy.$message.info('已取消')
    return
  }
  const res = await proxy.$api.deleteRole(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('删除成功')
    getRoleList()
  }
}

// 弹出分配权限对话框
const titlePermission = ref("")
const openPermission = ref(false)
const permissionList = ref([])
const defKeys = ref([])
const id = ref('')
const treeRef = ref(null)
const resetFormPermission = () =>{
  defKeys.value = []
}
const handlePermission = async (row)=>{
  id.value = row.id
  const resSysMenuVo = await proxy.$api.queryMenuList()
  if (resSysMenuVo.code != 200) {
    proxy.$message.error(resSysMenuVo.message)
    return
  } else {
    permissionList.value = proxy.$common.handleTree(resSysMenuVo.data, 'id')
    const resRoleMenuId = await proxy.$api.queryRoleMenuIdList(row.id)
    if (resRoleMenuId.code != 200) {
      proxy.$message.error(resSysMenuVo.message)
      return
    } else {
      defKeys.value = resRoleMenuId.data
    }
    openPermission.value = true
    titlePermission.value = '分配权限'
  }
}

// 取消
const cancelPermission = ()=> {
  openPermission.value = false
  proxy.$message.info('已取消')
}

// 分配权限
const allotPermission = async () => {
  const keys = [
      ...treeRef.value.getCheckedKeys(),
      ...treeRef.value.getHalfCheckedKeys(),
  ]
  const res = await proxy.$api.assignPermissions(id.value, keys)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    openPermission.value = false
    await getRoleList()
    proxy.$message.success('分配权限成功')
  }
}
</script>

<style scoped></style>