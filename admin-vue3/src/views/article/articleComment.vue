<template>
  <el-card>
    <!--条件搜索-->
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="文章名称" prop="title">
        <el-input placeholder="请输入文章名称" clearable v-model="queryParams.title"></el-input>
      </el-form-item>
      <el-form-item label="用户名称" prop="username">
        <el-input placeholder="请输入用户名称" clearable v-model="queryParams.username"></el-input>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select placeholder="请选择状态" v-model="queryParams.status">
          <el-option v-for="item in statusList" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button type="primary" icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <!--搜索/刷新/隐藏-->
      <TipHover v-model:showSearch="showSearch" @queryTable="getArticleCommentList" :columns="columns" />
    </el-row>
    <!--表格-->
    <el-table border stripe style="width: 100%;" :header-cell-style="{ background: '#eef1f6', color: '#606266' }"
              :data="articleCommentList">
      <el-table-column label="ID" prop="id" v-if="false" />
      <el-table-column label="文章名称" prop="title" v-if="columns[0].visible" />
      <el-table-column label="评论内容" prop="content" v-if="columns[1].visible" />
      <el-table-column label="评论人名称" prop="username" v-if="columns[2].visible" />
      <el-table-column label="评论人头像" prop="icon" v-if="columns[3].visible">
        <template #default="scope">
          <el-avatar shape="circle" :src="ip + scope.row.icon"></el-avatar>
        </template>
      </el-table-column>
      <el-table-column label="评论人地址" prop="loginAddress" v-if="columns[4].visible" />
      <el-table-column label="创建时间" prop="createTime" v-if="columns[5].visible" />
      <el-table-column label="状态" prop="status" v-if="columns[6].visible">
        <template #default="scope">
          <el-tag v-if="scope.row.status == 2" type="success">未删除</el-tag>
          <el-tag v-else-if="scope.row.status == 1" type="danger">已删除</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="更多操作" >
        <template #default="scope">
          <el-button v-if="scope.row.status == 2" type="danger" link icon="Delete" v-hasPermission="['article:comment:delete']" @click="handleDelete(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <Pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
                v-model:limit="queryParams.pageSize" @pagination="getArticleCommentList" />
  </el-card>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue'
const { proxy } = getCurrentInstance()

const queryParams = ref({})

const statusList = ref([
  {
    value: '2',
    label: '未删除'
  },
  {
    value: '1',
    label: '已删除'
  }
])

// 列表数据
const ip = import.meta.env.VITE_APP_ACCESS_IP
const articleCommentList = ref([])
const total = ref(0)
const getArticleCommentList = async () => {
  const params = {
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
    title: queryParams.value.title,
    username: queryParams.value.username,
    status: queryParams.value.status,
  }
  const res = await proxy.$api.queryArticleCommentList(params)
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    articleCommentList.value = res.data.list
    total.value = res.data.total
  }
}
getArticleCommentList()

// 搜索
const handleQuery = () => {
  getArticleCommentList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getArticleCommentList()
  proxy.$message.success('重置成功')
}

// 字段回显/隐藏
const showSearch = ref(true)
const columns = ref([
  { key: 0, label: `文章名称`, visible: true },
  { key: 1, label: `评论内容`, visible: true },
  { key: 2, label: `评论人名称`, visible: true },
  { key: 3, label: `评论人头像`, visible: true },
  { key: 4, label: `评论人地址`, visible: true },
  { key: 5, label: `创建时间`, visible: true },
  { key: 6, label: `状态`, visible: true },
])

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
  const res = await proxy.$api.updateCommentStatus(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('删除成功')
    await getArticleCommentList()
  }
}

</script>

<style lang="scss" scoped>

</style>