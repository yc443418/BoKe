<template>
    <!-- <div class="body-container"> -->
    <div class="article">
        <el-form :model="formData" :rules="rules" ref="formDataRef" class="article-panel" label-width="60px">
            <!--内容-->
            <div class="article-post">
                <el-card :body-style="{ padding: '5px' }">
                    <template #header>
                        <span>正文</span>
                    </template>
                    <el-form-item prop="content" label-width="0">
                        <ArticleMarkdown v-model="formData.markdownContent" @htmlContent="setHtmlContent" />
                    </el-form-item>
                </el-card>
            </div>
            <!--设置-->
            <div class="article-setting">
                <el-card>
                    <template #header>
                        <span>设置</span>
                    </template>
                    <!--文章名称-->
                    <el-form-item label="标题" prop="title">
                        <el-input clearable placeholder="请输入标题" v-model="formData.title"></el-input>
                    </el-form-item>
                    <!--文章分类-->
                    <el-form-item label="分类" prop="categoryId">
                        <el-select v-model="formData.categoryId" placeholder="请选择分类" style="width: 100%;">
                            <el-option v-for="item in bCategoryList" :key="item.id" :label="item.categoryName"
                                :value="item.id"></el-option>
                        </el-select>
                    </el-form-item>
                    <!--文章标签-->
                    <el-form-item label="标签" prop="tagId">
                        <el-select v-model="formData.tagId" placeholder="请选择标签" style="width: 100%;">
                            <el-option v-for="item in bTagsList" :key="item.id" :label="item.tagName"
                                :value="item.id"></el-option>
                        </el-select>
                    </el-form-item>
                    <!--文章封面-->
                    <el-form-item label="封面" prop="image">
                        <el-upload :headers=headers class="avatar-uploader" :action="uploadIconUrl"
                            :show-file-list="false" :on-success="handleAvatarSuccess">
                            <img v-if="formData.image" :src="ip + formData.image" class="avatar" title="点击更换封面">
                            <el-icon v-else class="avatar-uploader-icon">
                                <Plus />
                            </el-icon>
                        </el-upload>
                    </el-form-item>
                    <!--文章摘要-->
                    <el-form-item label="摘要" prop="summary">
                        <el-input clearable type="textarea" :rows="5" :maxlength="150" resize="none" show-word-limit
                            placeholder="请输入摘要" v-model="formData.summary"></el-input>
                    </el-form-item>
                    <!--保存按钮-->
                    <el-form-item>
                        <el-button type="primary" style="width: 100%;" @click="saveArticle">保存</el-button>
                    </el-form-item>
                </el-card>
            </div>
        </el-form>
    </div>
    <!-- </div> -->
</template>

<script setup>
import ArticleMarkdown from '../components/article/articleMarkdown.vue'
import { ref, getCurrentInstance } from 'vue'
const { proxy } = getCurrentInstance()
import { useRouter } from "vue-router"
const router = useRouter()
import { useStore } from 'vuex'
const store = useStore()

// form相关定义
const formData = ref({})
const formDataRef = ref()
const rules = {
    markdownContent: [{ required: true, message: "正文内容不能为空", trigger: "blur" }],
    title: [{ required: true, message: "标题不能为空", trigger: "blur" }],
    categoryId: [{ required: true, message: "文章分类不能为空", trigger: "blur" }],
    tagId: [{ required: true, message: "文章标签不能为空", trigger: "blur" }],
    image: [{ required: true, message: "文章封面不能为空", trigger: "blur" }],
    summary: [{ required: true, message: "摘要不能为空", trigger: "blur" }],
}

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
    const res = await proxy.$api.bTagsList()
    if (res.code !== 200) {
        proxy.$message.error(res.message)
    } else {
        bTagsList.value = res.data
        // console.log(res.data)
    }
}
getBTagsList()

// 图片上传处理
const ip = import.meta.env.VITE_APP_ACCESS_IP
const uploadIconUrl = import.meta.env.VITE_APP_UPLOAD_IP
const headers = ref({
    Authorization: 'Bearer ' + proxy.$store.state.token
})
const handleAvatarSuccess = (res, file) => {
    // 未登录弹出登录页面
    if (!proxy.$store.state.token) {
        store.commit("saveShowLogin", true)
        return
    }
    formData.value.image = res.data
}

//设置markdown编辑器的富文本信息
const setHtmlContent = (htmlContent) => {
    formData.value.content = htmlContent
}

// 保存文章
const saveArticle = () => {
    // 未登录弹出登录页面
    if (!proxy.$store.state.token) {
        store.commit("saveShowLogin", true)
        return
    }
    proxy.$refs["formDataRef"].validate(valid => {
        if (valid) {
            proxy.$api.createArticle(formData.value).then(res => {
                if (res.code != 200) {
                    proxy.$message.error(res.message)
                } else {
                    proxy.$message.success('新增文章成功')
                    formData.value = {}
                    proxy.$refs["formDataRef"].resetFields()
                    router.push(`/article/${res.data.id}`)
                }
            })
        } else {
            return
        }
    })
}
</script>

<style lang="scss">
.article {

    // 表单
    .article-panel {
        display: flex;

        // 内容
        .article-post {
            flex: 1;
        }

        // 文章设置
        .article-setting {
            margin-left: 10px;
            width: 450px;

            // 封面样式
            .avatar-uploader {
                border: 1px dashed #8c939d;
                border-radius: 6px;
                cursor: pointer;
                width: 200px;
                height: 170px;
                position: relative;
                overflow: hidden;

                .avatar {
                    width: 200px;
                    height: 170px;
                    display: block;
                }

                .el-icon.avatar-uploader-icon {
                    font-size: 28px;
                    color: #8c939d;
                    width: 200px;
                    height: 170px;
                    text-align: center;
                }
            }
        }
    }
}
</style>