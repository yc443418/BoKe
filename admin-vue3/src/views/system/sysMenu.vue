<template>
  <el-card>
    <!--条件搜索-->
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="菜单名称" prop="menuName">
        <el-input placeholder="请输入菜单名称" clearable v-model="queryParams.menuName"/>
      </el-form-item>
      <el-form-item label="菜单状态" prop="menuStatus">
        <el-select placeholder="请选择菜单状态" v-model="queryParams.menuStatus" style="width: 220px">
          <el-option v-for="item in menuStatusList" :key="item.value" :label="item.label" :value="item.value"/>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button type="primary" icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <!--表格按钮操作-->
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="Plus" @click="handleAdd" v-hasPermission="['system:sysMenu:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="info" plain icon="Sort" @click="toggleExpandAll">展开/折叠</el-button>
      </el-col>
      <!-- 搜索/刷新/隐藏组件-->
      <TipHover v-model:showSearch="showSearch" @queryTable="getMenuList" :columns="columns"></TipHover>
    </el-row>
    <!--表格-->
    <el-table border stripe :header-cell-style="{ background: '#eef1f6', color: '#606266' }" :data="menuList"
              row-key="id" :tree-props="{children: 'children'}" v-if="refreshTable" :default-expand-all="isExpandAll">
      <el-table-column label="菜单名称" prop="menuName" v-if="columns[0].visible"/>
      <el-table-column label="图标" prop="icon" v-if="columns[1].visible">
        <template #default="scope">
          <el-icon v-if="scope.row.icon">
            <component :is="scope.row.icon"></component>
          </el-icon>
        </template>
      </el-table-column>
      <el-table-column label="权限值" prop="value" v-if="columns[2].visible"/>
      <el-table-column label="菜单路径" prop="url" v-if="columns[3].visible"/>
      <el-table-column label="菜单类型" prop="menuType" v-if="columns[4].visible">
        <template #default="scope">
          <el-tag v-if="scope.row.menuType === 1">目录</el-tag>
          <el-tag v-else-if="scope.row.menuType === 2" type="success">菜单</el-tag>
          <el-tag v-else-if="scope.row.menuType === 3" type="danger">按钮</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="菜单状态" prop="menuStatus" v-if="columns[5].visible">
        <template #default="scope">
          <el-tag v-if="scope.row.menuStatus === 1" type="success">启用</el-tag>
          <el-tag v-else-if="scope.row.menuStatus === 2" type="danger">禁用</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="排序" prop="sort" v-if="columns[6].visible"/>
      <el-table-column label="创建时间" prop="createTime" v-if="columns[7].visible"/>
      <el-table-column label="更多操作" class-name="small-padding fixed-width">
        <template #default="scope">
          <el-button type="primary" link icon="Edit" @click="handleUpdate(scope.row.id)" v-hasPermission="['system:sysMenu:edit']">修改
          </el-button>
          <el-button type="danger" link icon="Delete" @click="handleDelete(scope.row.id)" v-hasPermission="['system:sysMenu:delete']">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--新增和修改对话框-->
    <el-dialog :totle="title" v-model="open" width="500px" @close="resetForm">
      <el-form ref="menuRef" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="菜单类型" prop="menuType">
          <el-radio-group v-model="form.menuType">
            <el-radio :label="1">目录</el-radio>
            <el-radio :label="2">菜单</el-radio>
            <el-radio :label="3">按钮</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="上级菜单">
          <el-tree-select v-model="form.parentId" :data="addList" :props="{value: 'id', label: 'menuName'}"
                          value-key="id" placeholder="请选择上级菜单" check-strictly/>
        </el-form-item>
        <el-form-item label="菜单图标" prop="icon" v-if="form.menuType != 3">
          <el-select v-model="form.icon">
            <el-option v-for="item in iconList" :key="item.value" :label="item.label" :value="item.value">
              <el-icon>
                <component :is="item.value"></component>
              </el-icon>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="菜单名称" prop="menuName">
          <el-input v-model="form.menuName" placeholder="请输入菜单名称"></el-input>
        </el-form-item>
        <el-form-item label="显示排序" prop="sort">
          <el-input-number v-model="form.sort" controls-position="right" :min="0"/>
        </el-form-item>
        <el-form-item label="菜单路径" prop="url" v-if="form.menuType == 2">
          <el-input v-model="form.url" placeholder="请输入菜单路径"></el-input>
        </el-form-item>
        <el-form-item label="权限标识" prop="value" v-if="form.menuType == 3">
          <el-input v-model="form.value" placeholder="请输入权限标识"/>
        </el-form-item>
        <el-form-item label="菜单状态" prop="menuStatus">
          <el-radio-group v-model="form.menuStatus">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="2">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm">确定</el-button>
          <el-button type="info" @click="cancel">取消</el-button>
        </div>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import {getCurrentInstance, nextTick, ref} from 'vue'

const {proxy} = getCurrentInstance()
const menuStatusList = ref([
  {
    value: '1',
    label: '启用',
  },
  {
    value: '2',
    label: '禁用',
  }
])
// 菜单列表
const queryParams = ref({})
const menuList = ref([])
const getMenuList = async () => {
  const params = {
    menuName: queryParams.value.menuName,
    menuStatus: queryParams.value.menuStatus,
  }
  const res = await proxy.$api.queryMenuList(params)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    menuList.value = proxy.$common.handleTree(res.data, "id")
    // console.log("菜单列表数据：", res.data)
  }
}
getMenuList()

// 搜索
const handleQuery = () => {
  getMenuList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getMenuList()
  proxy.$message.success('重置成功')
}

// 展开/折叠
const refreshTable = ref(true)
const isExpandAll = ref(true)
const toggleExpandAll = () => {
  refreshTable.value = false
  isExpandAll.value = !isExpandAll.value
  nextTick(() => {
    refreshTable.value = true
  })
}

const rules = {
  menuType: [{required: true, message: "菜单类型不能为空", trigger: "blur"}],
  menuName: [{required: true, message: "菜单名称不能为空", trigger: "blur"}],
}

const iconList = ref([
  {value: 'Aim', label: 'Aim'},
  {value: 'House', label: 'House'},
  {value: 'Service', label: 'Service'},
  {value: 'Pointer', label: 'Pointer'},
  {value: 'Star', label: 'Star'},
  {value: 'Notification', label: 'Notification'},
  {value: 'Connection', label: 'Connection'},
  {value: 'ChatDotRound', label: 'ChatDotRound'},
  {value: 'Setting', label: 'Setting'},
  {value: 'Clock', label: 'Clock'},
  {value: 'Position', label: 'Position'},
  {value: 'Discount', label: 'Discount'},
  {value: 'Odometer', label: 'Odometer'},
  {value: 'ChatSquare', label: 'ChatSquare'},
  {value: 'ChatRound', label: 'ChatRound'},
  {value: 'ChatLineRound', label: 'ChatLineRound'},
  {value: 'ChatLineSquare', label: 'ChatLineSquare'},
  {value: 'ChatDotSquare', label: 'ChatDotSquare'},
  {value: 'View', label: 'View'},
  {value: 'Hide', label: 'Hide'},
  {value: 'Unlock', label: 'Unlock'},
  {value: 'Lock', label: 'Lock'},
  {value: 'RefreshRight', label: 'RefreshRight'},
  {value: 'RefreshLeft', label: 'RefreshLeft'},
  {value: 'Bell', label: 'Bell'},
  {value: 'User', label: 'User'},
  {value: 'CircleCheck', label: 'CircleCheck'},
  {value: 'PieChart', label: 'PieChart'},
  {value: 'Compass', label: 'Compass'},
  {value: 'MessageBox', label: 'MessageBox'},
  {value: 'TurnOff', label: 'TurnOff'},
  {value: 'Crop', label: 'Crop'},
  {value: 'Open', label: 'Open'},
  {value: 'UserFilled', label: 'UserFilled'},
  {value: 'Tools', label: 'Tools'},
  {value: 'HomeFilled', label: 'HomeFilled'},
  {value: 'Avatar', label: 'Avatar'},
  {value: 'HelpFilled', label: 'HelpFilled'},
  {value: 'Promotion', label: 'Promotion'},
  {value: 'Grid', label: 'Grid'},
  {value: 'Printer', label: 'Printer'},
])

const title = ref("")
const form = ref({})
const open = ref(false)

// 重置表单
const resetForm = () => {
  form.value = {}
  proxy.$refs["menuRef"].resetFields()
}

// 新增上级菜单
const addList = ref([])
const getTreeList = async () => {
  addList.value = []
  const res = await proxy.$api.queryMenuList()
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    const menu = {id: 0, menuName: "顶级", children: []}
    menu.children = proxy.$common.handleTree(res.data, "id")
    addList.value.push(menu)
  }
}

// 弹出新增对话框
const handleAdd = async () => {
  getTreeList()
  open.value = true
  title.value = '新增菜单'
}

// 弹出修改对话框
const handleUpdate = async (id) => {
  await getTreeList()
  const res = await proxy.$api.menuInfo(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    form.value = res.data
    open.value = true
    title.value = "修改菜单"
  }
}

// 取消
const cancel = () => {
  open.value = false
  proxy.$message.info('已取消')
}

// 新增和修改提交操作
const submitForm = () => {
  proxy.$refs["menuRef"].validate(valid => {
    if (valid) {
      if (form.value.id != undefined) {
        proxy.$api.menuUpdate(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
            return
          } else {
            proxy.$message.success('修改成功')
            open.value = false
            getMenuList()
          }
        })
      } else {
        proxy.$api.addMenu(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('新增成功')
            open.value = false
            getMenuList()
          }
        })
      }
    }
  })
}

// 删除
const handleDelete = async (id)=> {
  const confirmResult = await proxy.$confirm('是否要删除编号为' + id + '的数据' + '是否继续？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).catch(err => err)
  if (confirmResult != 'confirm') {
    proxy.$message.info('已取消')
    return
  }
  const res = await proxy.$api.menuDelete(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('删除成功')
    getMenuList()
  }
}

// 字段回显和隐藏
const showSearch = ref(true)
const columns = ref([
  { key: 0, label: `菜单名称`, visible: true},
  { key: 1, label: `图标`, visible: true},
  { key: 2, label: `权限值`, visible: true},
  { key: 3, label: `菜单路径`, visible: true},
  { key: 4, label: `菜单类型`, visible: true},
  { key: 5, label: `菜单状态`, visible: true},
  { key: 6, label: `排序`, visible: true},
  { key: 7, label: `创建时间`, visible: true}
])
</script>

<style lang="scss">

</style>