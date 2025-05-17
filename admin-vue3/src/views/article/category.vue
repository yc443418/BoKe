<template>
  <el-card>
    <!--条件搜索-->
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="分类名称" prop="categoryName">
        <el-input placeholder="请输入分类名称" clearable v-model="queryParams.categoryName"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button type="primary" icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button plain type="primary" icon="Plus" @click="handleAdd"
                   v-hasPermission="['article:category:add']">新增</el-button>
      </el-col>
      <!--搜索/刷新/隐藏-->
      <TipHover v-model:showSearch="showSearch" @queryTable="getCategoryList" :columns="columns" />
    </el-row>
    <!--表格-->
    <el-table border stripe style="width: 100%;" :header-cell-style="{ background: '#eef1f6', color: '#606266' }"
              :data="categoryList">
      <el-table-column label="ID" prop="id" v-if="false" />
      <el-table-column label="分类名称" prop="categoryName" v-if="columns[0].visible" />
      <el-table-column label="文章数量" prop="quantity" v-if="columns[1].visible" />
      <el-table-column label="排序" prop="sort" v-if="columns[2].visible" />
      <el-table-column label="创建时间" prop="createTime" v-if="columns[3].visible" />
      <el-table-column label="更多操作">
        <template #default="scope">
          <el-button type="primary" link icon="Edit" @click="handleUpdate(scope.row.id)"
                     v-hasPermission="['article:category:edit']">编辑</el-button>
          <el-button type="danger" link icon="Delete" @click="handleDelete(scope.row.id)"
                     v-hasPermission="['article:category:delete']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <Pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
                v-model:limit="queryParams.pageSize" @pagination="getCategoryList" />
    <!--新增/编辑对话框-->
    <el-dialog :title="title" v-model="open" width="500px" @close="resetForm">
      <el-form ref="categoryRef" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="分类名称" prop="categoryName">
          <el-input v-model="form.categoryName" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="form.sort" placeholder="请输入排序" />
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

const queryParams = ref({})

// 列表数据
const categoryList = ref([])
const total = ref(0)
const getCategoryList = async () => {
  const params = {
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
    categoryName: queryParams.value.categoryName,
  }
  const res = await proxy.$api.queryBCategoryList(params)
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    categoryList.value = res.data.list
    total.value = res.data.total
  }
}
getCategoryList()

// 搜索
const handleQuery = () => {
  getCategoryList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getCategoryList()
  proxy.$message.success('重置成功')
}

// 字段回显/隐藏
const showSearch = ref(true)
const columns = ref([
  { key: 0, label: `分类名称`, visible: true },
  { key: 1, label: `文章数量`, visible: true },
  { key: 2, label: `排序`, visible: true },
  { key: 3, label: `创建时间`, visible: true },
])

// 新增/编辑定义
const rules = {
  categoryName: [{ required: true, message: "分类名称不能为空", trigger: "blur" }],
  sort: [{ required: true, message: "排序不能为空", trigger: "blur" }],
}
const title = ref("")
const form = ref({})
const open = ref(false)
const resetForm = () => {
  form.value = {}
  proxy.$refs["categoryRef"].resetFields()
}

// 弹出新增对话框
const handleAdd = () => {
  open.value = true
  title.value = '新增标签'
}

// 弹出修改对话框
const handleUpdate = (id) => {
  proxy.$api.bCategoryInfo(id).then(response => {
    form.value = response.data
    open.value = true
    title.value = '修改标签'
  });
}

// 取消
const cancel = () => {
  open.value = false
  proxy.$message.info('已取消')
}

// 新增/修改提交
const submitForm = () => {
  proxy.$refs["categoryRef"].validate(valid => {
    if (valid) {
      if (form.value.id != undefined) {
        proxy.$api.bCategoryUpdate(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('修改成功')
            open.value = false
            getCategoryList()
          }
        })
      } else {
        proxy.$api.addBCategory(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('新增成功')
            open.value = false
            getCategoryList()
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
  const res = await proxy.$api.deleteBCategory(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('删除成功')
    await getCategoryList()
  }
}

</script>

<style lang="scss" scoped></style>