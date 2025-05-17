<template>
  <el-card>
    <!--条件搜索-->
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="文章名称" prop="title">
        <el-input placeholder="请输入文章名称" clearable v-model="queryParams.title"></el-input>
      </el-form-item>
      <el-form-item label="分类名称" prop="categoryId">
        <el-select v-model="queryParams.categoryId" placeholder="请选择分类名称">
          <el-option v-for="item in bCategoryList" :key="item.id" :label="item.categoryName"
                     :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="标签名称" prop="tagId">
        <el-select v-model="queryParams.tagId" placeholder="请选择标签名称">
          <el-option v-for="item in bTagsList" :key="item.id" :label="item.tagName" :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="是否置顶" prop="topType">
        <el-select placeholder="请选择是否置顶" v-model="queryParams.topType">
          <el-option v-for="item in isTopTypeList" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
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
      <TipHover v-model:showSearch="showSearch" @queryTable="getArticleList" :columns="columns" />
    </el-row>
    <!--表格-->
    <el-table border stripe style="width: 100%;" :header-cell-style="{ background: '#eef1f6', color: '#606266' }"
              :data="articleList">
      <el-table-column label="ID" prop="id" v-if="false" />
      <el-table-column label="封面地址" prop="image" v-if="columns[0].visible">
        <template #default="scope">
          <el-image style="width: 100px; height: 70px" :src="ip + scope.row.image"
                    :preview-src-list="[ip + scope.row.image]" :zoom-rate="1.2" :max-scale="7" :min-scale="0.2"
                    :initial-index="4" fit="cover" :preview-teleported="true" title="点击放大预览封面" />
        </template>
      </el-table-column>
      <el-table-column label="文章名称" prop="title" v-if="columns[1].visible" />
      <el-table-column label="分类名称" prop="categoryName" v-if="columns[2].visible" />
      <el-table-column label="标签名称" prop="tagName" v-if="columns[3].visible" />
      <el-table-column label="发表用户" prop="username" v-if="columns[4].visible" />
      <el-table-column label="文章简介" prop="summary" v-if="columns[5].visible" />
      <el-table-column label="是否置顶" prop="topType" v-if="columns[6].visible" width="170">
        <template #default="scope">
          <el-switch v-model="scope.row.topType" :active-value="1" :inactive-value="2" active-color="#F5222D"
                     inactive-color="#13ce66" active-text="未置顶" inactive-text="已置顶" @click="updateTopType(scope.row)">
          </el-switch>
        </template>
      </el-table-column>
      <el-table-column label="阅读数" prop="quantity" v-if="columns[7].visible" />
      <el-table-column label="点赞数" prop="goodCount" v-if="columns[8].visible" />
      <el-table-column label="评论数" prop="commentCount" v-if="columns[9].visible" />
      <el-table-column label="创建时间" prop="createTime" v-if="columns[10].visible" />
      <el-table-column label="状态" prop="status" v-if="columns[11].visible">
        <template #default="scope">
          <el-tag v-if="scope.row.status == 2" type="success">未删除</el-tag>
          <el-tag v-else-if="scope.row.status == 1" type="danger">已删除</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="更多操作" fixed="right" width="180">
        <template #default="scope">
          <el-button type="danger" link icon="Delete" v-hasPermission="['article:article:delete']" @click="handleDelete(scope.row.id)">删除</el-button>
          <el-button type="primary" link icon="View" v-hasPermission="['article:article:query']" @click="queryComment(scope.row.id)">查看评论</el-button>
        </template>
      </el-table-column>
    </el-table>
    <Pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
                v-model:limit="queryParams.pageSize" @pagination="getArticleList" />
    <!--文章评论列表-->
    <el-dialog :title="title" v-model="open" width="600px" @close="resetForm">
      <div class="comment-item" v-for="item in commentList">
        <div class="p-comment-item">
          <el-image style="width: 40px; height: 40px; border-radius: 20px;" :src="ip + item.icon"/>
          <div class="content-info">
            <span>{{ item.username }}</span>
            <div v-html="item.content" class="comment-content"></div>
            <div class="time-address">
              时间：{{ item.createTime }}&nbsp;·&nbsp;地址：{{item.loginAddress}}
            </div>
          </div>
        </div>
      </div>
    </el-dialog>
  </el-card>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue'
const { proxy } = getCurrentInstance()

const queryParams = ref({})
const isTopTypeList = ref([
  {
    value: '1',
    label: '未置顶'
  },
  {
    value: '2',
    label: '已置顶'
  }
])

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

// 分类数据
const bCategoryList = ref([])
const getBCategoryList = async () => {
  const res = await proxy.$api.bCategorySelectList()
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    bCategoryList.value = res.data
    // console.log(res.data)
  }
}
getBCategoryList()

// 标签数据
const bTagsList = ref([])
const getBTagsList = async () => {
  const res = await proxy.$api.bTagsSelectList()
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    bTagsList.value = res.data
    // console.log(res.data)
  }
}
getBTagsList()

// 列表数据
const ip = import.meta.env.VITE_APP_ACCESS_IP
const articleList = ref([])
const total = ref(0)
const getArticleList = async () => {
  const params = {
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
    title: queryParams.value.title,
    categoryId: queryParams.value.categoryId,
    tagId: queryParams.value.tagId,
    isPublish: queryParams.value.isPublish,
    topType: queryParams.value.topType,
    status: queryParams.value.status,
  }
  const res = await proxy.$api.queryArticleList(params)
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    articleList.value = res.data.list
    total.value = res.data.total
  }
}
getArticleList()

// 搜索
const handleQuery = () => {
  getArticleList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getArticleList()
  proxy.$message.success('重置成功')
}

// 字段回显/隐藏
const showSearch = ref(true)
const columns = ref([
  { key: 0, label: `封面地址`, visible: true },
  { key: 1, label: `文章名称`, visible: true },
  { key: 2, label: `分类名称`, visible: true },
  { key: 3, label: `标签名称`, visible: true },
  { key: 4, label: `发表用户`, visible: true },
  { key: 5, label: `文章简介`, visible: true },
  { key: 6, label: `是否置顶`, visible: true },
  { key: 7, label: `阅读数`, visible: true },
  { key: 8, label: `点赞数`, visible: true },
  { key: 9, label: `评论数`, visible: true },
  { key: 10, label: `创建时间`, visible: true },
  { key: 11, label: `状态`, visible: true },
])

// 是否置顶
const updateTopType = async (row) => {
  let text = row.topType == 2 ? "置顶" : "取消置顶"
  const confirmResult = await proxy.$confirm('确认要"' + text + '""' + row.title + '"文章吗?', "警告", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).catch(err => err)
  if (confirmResult != 'confirm') {
    proxy.$message.info('已取消')
    await getArticleList()
    return
  }
  const res = await proxy.$api.updateTopType(row.id, row.topType)
  if (res.code != 200) {
    proxy.$message.error(res.message)
  } else {
    proxy.$message.success(text + "成功")
    await getArticleList()
  }
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
  const res = await proxy.$api.updateStatus(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('删除成功')
    await getArticleList()
  }
}

// 查询评论页面
const title = ref("")
const open = ref(false)
const resetForm = () => {
  open.value = false
}
const commentList = ref([])
const queryComment = async (id)=> {
  open.value = true
  title.value = '文章评论列表'
  const params = {
    id: id,
  }
  const res = await proxy.$api.queryArticleCommentListByArticleId(params)
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    commentList.value = res.data
    // console.log("查询的文章id:", id)
    // console.log("文章评论列表数据：", res.data)
  }
}
</script>

<style lang="scss" scoped>
.comment-item {
  border-bottom: 1px solid #ddd;
  padding: 10px 0px;
.p-comment-item {
  display: flex;
}
.content-info {
  margin-left: 5px;
}
.comment-content {
  margin: 5px 0px;
}
.time-address {
  font-size: 12px;
  color: rgb(135, 130, 130);
}
.s-comment-item {
  display: flex;
  margin-top: 10px;
}
}
</style>