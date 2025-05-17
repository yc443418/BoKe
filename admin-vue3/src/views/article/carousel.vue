<template>
  <el-card>
    <!--条件搜索-->
    <el-form :inline="true" :model="queryParams" ref="queryRef" v-show="showSearch">
      <el-form-item label="是否发布" prop="isPublish">
        <el-select placeholder="请选择是否发布" v-model="queryParams.isPublish">
          <el-option v-for="item in isPublishList" :key="item.value" :label="item.label" :value="item.value"/>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button type="primary" icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button plain type="primary" icon="Plus" @click="handleAdd"
                   v-hasPermission="['article:carousel:add']">新增
        </el-button>
      </el-col>
      <!--搜索/刷新/隐藏-->
      <TipHover v-model:showSearch="showSearch" @queryTable="getCarouselList" :columns="columns"/>
    </el-row>
    <!--表格-->
    <el-table border stripe style="width: 100%;" :header-cell-style="{ background: '#eef1f6', color: '#606266' }"
              :data="carouselList">
      <el-table-column label="ID" prop="id" v-if="false"/>
      <el-table-column label="封面地址" prop="icon" v-if="columns[0].visible">
        <template #default="scope">
          <el-image style="width: 100px; height: 70px" :src="ip + scope.row.icon"
                    :preview-src-list="[ip + scope.row.icon]" :zoom-rate="1.2" :max-scale="7" :min-scale="0.2"
                    :initial-index="4" fit="cover" :preview-teleported="true" title="点击放大预览封面"/>
        </template>
      </el-table-column>
      <el-table-column label="是否发布" prop="isPublish" v-if="columns[1].visible">
        <template #default="scope">
          <el-tag v-if="scope.row.isPublish === 1" type="danger">否</el-tag>
          <el-tag v-else-if="scope.row.isPublish === 2" type="success">是</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" prop="createTime" v-if="columns[2].visible"/>
      <el-table-column label="更多操作">
        <template #default="scope">
          <el-button type="primary" link icon="Edit" @click="handleUpdate(scope.row.id)"
                     v-hasPermission="['article:carousel:edit']">编辑
          </el-button>
          <el-button type="danger" link icon="Delete" @click="handleDelete(scope.row.id)"
                     v-hasPermission="['article:carousel:delete']">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <Pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
                v-model:limit="queryParams.pageSize" @pagination="getCarouselList"/>
    <!--新增/编辑对话框-->
    <el-dialog :title="title" v-model="open" width="500px" @close="resetForm">
      <el-form ref="carouselRef" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="封面地址" prop="icon">
          <el-upload :headers=headers class="avatar-uploader" :action="uploadIconUrl" :show-file-list="false"
                     :on-success="handleAvatarSuccess">
            <img v-if="form.icon" :src="ip + form.icon" class="avatar" title="点击更换封面">
            <el-icon v-else class="avatar-uploader-icon">
              <Plus/>
            </el-icon>
          </el-upload>
        </el-form-item>
        <el-form-item label="是否发布" prop="isPublish">
          <el-radio-group v-model="form.isPublish">
            <el-radio :label="1">否</el-radio>
            <el-radio :label="2">是</el-radio>
          </el-radio-group>
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

const isPublishList = ref([
  {
    value: '1',
    label: '否'
  },
  {
    value: '2',
    label: '是'
  }
])

// 轮播图列表
const ip = import.meta.env.VITE_APP_ACCESS_IP
const queryParams = ref({})
const carouselList = ref([])
const total = ref(0)
const getCarouselList = async () => {
  const params = {
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
    isPublish: queryParams.value.isPublish,
  }
  const res = await proxy.$api.queryBCarouselList(params)
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    carouselList.value = res.data.list
    total.value = res.data.total
  }
}
getCarouselList()

// 搜索
const handleQuery = () => {
  getCarouselList()
}

// 重置
const queryRef = ref()
const resetQuery = () => {
  queryRef.value.resetFields()
  getCarouselList()
  proxy.$message.success('重置成功')
}

// 字段回显/隐藏
const showSearch = ref(true)
const columns = ref([
  {key: 0, label: `封面地址`, visible: true},
  {key: 1, label: `是否发布`, visible: true},
  {key: 2, label: `创建时间`, visible: true},
])

// 新增/编辑定义
const rules = {
  icon: [{required: true, message: "封面地址不能为空", trigger: "blur"}],
  isPublish: [{required: true, message: "是否发布不能为空", trigger: "blur"}],
}
const title = ref("")
const form = ref({})
const open = ref(false)
const resetForm = () => {
  form.value = {}
  proxy.$refs["carouselRef"].resetFields()
}

// 图片上传处理
const uploadIconUrl = import.meta.env.VITE_APP_UPLOAD_IP
const headers = ref({
  Authorization: 'Bearer ' + proxy.$store.state.token
})
const handleAvatarSuccess = (res, file) => {
  form.value.icon = res.data
}

// 弹出新增对话框
const handleAdd = () => {
  open.value = true
  title.value = '新增轮播图'
}

// 弹出修改对话框
const handleUpdate = (id) => {
  proxy.$api.bCarouselInfo(id).then(response => {
    form.value = response.data
    open.value = true
    title.value = '修改轮播图'
  })
}

// 取消
const cancel = () => {
  open.value = false
  proxy.$message.info('已取消')
}

// 新增/修改提交
const submitForm = () => {
  proxy.$refs["carouselRef"].validate(valid => {
    if (valid) {
      if (form.value.id != undefined) {
        proxy.$api.bCarouselUpdate(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('修改成功')
            open.value = false
            getCarouselList()
          }
        })
      } else {
        proxy.$api.addBCarousel(form.value).then(res => {
          if (res.code != 200) {
            proxy.$message.error(res.message)
          } else {
            proxy.$message.success('新增成功')
            open.value = false
            getCarouselList()
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
  const res = await proxy.$api.deleteBCarousel(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    proxy.$message.success('删除成功')
    await getCarouselList()
  }
}
</script>
<style lang="scss" scoped>
.avatar-uploader {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  width: 130px;
  height: 100px;
  position: relative;
  overflow: hidden;

  .avatar {
    width: 130px;
    height: 100px;
    display: block;
  }

  .el-icon.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 130px;
    height: 100px;
    text-align: center;
  }

}
</style>