<template>
    <div class="post-comment-panel">
        <!--头像-->
        <UserAvatar :width="width" :userId="userId" :src="src" />
        <!--输入框-->
        <div class="comment-form">
            <el-form :model="formData" :rules="rules" ref="formDataRef">
                <el-form-item prop="content">
                    <el-input clearable :placeholder="placeholder" type="textarea" :rows="3" :maxlength="150"
                        resize="none" show-word-limit v-model="formData.content" />
                </el-form-item>
            </el-form>
        </div>
        <!--发表按钮-->
        <div class="send-btn" @click="createComment">发表</div>
    </div>
</template>

<script setup>
import { ref, getCurrentInstance } from "vue"
const { proxy } = getCurrentInstance()

// 对外定义的参数
const props = defineProps({
    width: {
        type: Number,
    },
    userId: {
        type: Number,
    },
    src: {
        type: String,
    },
    placeholder: {
        type: String,
        default: "请输入评论",
    },
    articleId: {
        type: Number
    },
    pCommentId: {
        type: Number
    },
    replyUserId: {
        type: Number
    }
})

// form信息
const formData = ref({})
const formDataRef = ref()
const rules = {
    content: [
        { required: true, message: "请输入评论内容" },
        { min: 5, message: "评论至少5个字" },
    ],
}

// 发表
const emit = defineEmits(["postCommentFinish"])
const createComment = () => {
    formDataRef.value.validate(async (valid) => {
        if (!valid) {
            return
        }
        let params = Object.assign({}, formData.value)
        params.articleId = props.articleId
        params.pCommentId = props.pCommentId
        params.replyUserId = props.replyUserId
        const res = await proxy.$api.addArticleComment(params)
        if (res.code != 200) {
            proxy.$message.error(res.message)
        } else {
            proxy.$message.success("评论发表成功")
            formDataRef.value.resetFields()
            // 回调（因为新增的评论添加到列表显示）
            emit("postCommentFinish", res.data)
        }
    })
}
</script>

<style lang="scss">
.post-comment-panel {
    display: flex;

    .comment-form {
        flex: 1;
        margin: 0px 10px;
    }

    .send-btn {
        width: 70px;
        height: 70px;
        background: #007fff;
        color: #fff;
        text-align: center;
        line-height: 70px;
        border-radius: 5px;
        cursor: pointer;
    }
}
</style>