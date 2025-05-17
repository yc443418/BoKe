<template>
  <el-card>
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="参数名称" prop="name">
        <el-input placeholder="请输入参数名称" clearable v-model="queryParams.name"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button type="primary" icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button plain type="primary" icon="Plus" @click="handleAdd"
                   v-hasPermission="['system:sysConfig:add']">新增
        </el-button>
        <el-button plain type="primary" icon="Refresh" @click="refresh"
                   v-hasPermission="['system:sysConfig:refresh']">刷新同步
        </el-button>
      </el-col>
      <!--搜索/刷新/隐藏-->
      <TipHover v-model:showSearch="showSearch" @queryTable="getSysConfigList" :columns="columns"/>
    </el-row>
    <!--表格-->
    <el-table border stripe style="width: 100%;" :header-cell-style="{ background: '#eef1f6', color: '#606266' }"
              :data="sysConfigList">
      <el-table-column label="ID" prop="id" v-if="false"/>
      <el-table-column label="参数名称" prop="name" v-if="columns[0].visible"/>
      <el-table-column label="参数key" prop="configKey" v-if="columns[1].visible"/>
      <el-table-column label="参数value值" prop="configValue" v-if="columns[2].visible"/>
      <el-table-column label="排序" prop="sort" v-if="columns[3].visible"/>
      <el-table-column label="备注" prop="remark" v-if="columns[4].visible"/>
      <el-table-column label="创建时间" prop="createTime" v-if="columns[5].visible"/>
      <el-table-column label="更多操作">
        <template #default="scope">
          <el-button type="primary" link icon="Edit" @click="handleUpdate(scope.row.id)"
                     v-hasPermission="['system:sysConfig:edit']">编辑
          </el-button>
          <el-button type="danger" link icon="Delete" @click="handleDelete(scope.row.id)"
                     v-hasPermission="['system:sysConfig:delete']">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <Pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
                v-model:limit="queryParams.pageSize" @pagination="getSysConfigList"/>
    <!--新增/编辑对话框-->
    <el-dialog :title="title" v-model="open" width="500px" @close="resetForm">
      <el-form ref="sysConfigRef" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="参数名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入参数名称"/>
        </el-form-item>
        <el-form-item label="key值" prop="configKey">
          <el-input v-model="form.configKey" placeholder="请输入参数名称"/>
        </el-form-item>
        <el-form-item label="value值" prop="configValue">
          <el-input v-model="form.configValue" placeholder="请输入参数value"/>
        </el-form-item>
        <el-form-item label="参数排序" prop="sort">
          <el-input-number v-model="form.sort" placeholder="请输入排序"/>
        </el-form-item>
        <el-form-item label="参数备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入参数备注"></el-input>
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
import {getCurrentInstance, ref} from 'vue'

const {proxy} = getCurrentInstance()

const queryParams = ref({})

// 列表数据
const sysConfigList = ref([])
const total = ref(0)
const getSysConfigList = async () => {
  const params = {
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
    name: queryParams.value.name,
  }
  const res = await proxy.$api.queryConfigList(params)
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    sysConfigList.value = res.data.list
    total.value = res.data.total
  }
}
getSysConfigList()

// 搜索
const handleQuery = () => {
  getSysConfigList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getSysConfigList()
  proxy.$message.success('重置成功')
}

// 字段回显/隐藏
const showSearch = ref(true)
const columns = ref([
  {key: 0, label: `参数名称`, visible: true},
  {key: 1, label: `参数key`, visible: true},
  {key: 2, label: `参数value值`, visible: true},
  {key: 3, label: `排序`, visible: true},
  {key: 4, label: `备注`, visible: true},
  {key: 5, label: `创建时间`, visible: true},
])

// 新增/编辑定义
const rules = {
  name: [{required: true, message: "标签名称不能为空", trigger: "blur"}],
  configKey: [{required: true, message: "参数key不能为空", trigger: "blur"}],
  configValue: [{required: true, message: "参数value值不能为空", trigger: "blur"}],
  sort: [{required: true, message: "排序不能为空", trigger: "blur"}],
  remark: [{required: true, message: "备注不能为空", trigger: "blur"}],
}
const title = ref("")
const form = ref({})
const open = ref(false)
const resetForm = () => {
  form.value = {}
  proxy.$refs["sysConfigRef"].resetFields()
}

// 弹出新增对话框
const handleAdd = () => {
  open.value = true
  title.value = '新增参数'
}

// 弹出修改对话框
const handleUpdate = (id) => {
  proxy.$api.configInfo(id).then(response => {
    form.value = response.data
    open.value = true
    title.value = '修改参数'
  });
}

// 取消
const cancel = () => {
  open.value = false
  proxy.$message.info('已取消')
}

// 新增/修改提交
const submitForm = () => {
  proxy.$refs["sysConfigRef"].validate(valid => {
    if (valid) {
      if (form.value.id != undefined) {
        proxy.$api.configUpdate(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('修改成功')
            open.value = false
            getSysConfigList()
          }
        })
      } else {
        proxy.$api.addConfig(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('新增成功')
            open.value = false
            getSysConfigList()
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
  const res = await proxy.$api.deleteConfig(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('删除成功')
    await getSysConfigList()
  }
}

// 刷新
const refresh = async () => {
  const res = await proxy.$api.refresh()
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('刷新成功')
    await getSysConfigList()
  }
}

</script>

<style lang="scss" scoped></style>