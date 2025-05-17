<template>
    <div class="body-container search-body">
        <div class="search-panel">
            <el-form :model="formData" :rules="rules" ref="formDataRef" @submit.prevent>
                <el-form-item prop="keyword">
                    <el-input size="large" clearable placeholder="输入文章标题或者文章用户名称" v-model="formData.keyword"
                        @keyup.enter="getArticleListSearch" @change="changeInput">
                        <template #suffix>
                            <span class="iconfont icon-sousuo" @click="getArticleListSearch"
                                @blur="formData.keyword = $event.target.value.trim()"></span>
                        </template>
                    </el-input>
                </el-form-item>
            </el-form>
        </div>
        <!--搜索列表-->
        <div class="article-list" v-for="(item, index) in articleListInfo" :key="item.id">
            <div class="article-item-inner">
                <div class="article-body">
                    <!--用户信息-->
                    <div class="user-info">
                        <UserAvatar :width="30" :userId="item.userId" :src="item.icon" />
                        <router-link :to="'/user/' + item.userId" class="link-info">{{ item.username }}</router-link>
                        <el-divider direction="vertical"></el-divider>
                        <div class="create-time">{{ item.createTime }}</div>
                        <div>&nbsp;·&nbsp;{{ item.loginAddress }}</div>
                        <el-divider direction="vertical"></el-divider>
                        <div class="name" @click="goToCategory">{{ item.categoryName }}</div>
                        <el-divider direction="vertical"></el-divider>
                        <div class="name" @click="goToTag">{{ item.tagName }}</div>
                    </div>
                    <!--文章标题-->
                    <router-link :to="`/article/${item.id}`" class="title">
                        <span>{{ item.title }}</span>
                    </router-link>
                    <div class="summary">{{ item.summary }}</div>
                    <div class="article-info">
                        <span class="iconfont icon-yuedu">
                            {{ item.quantity == 0 ? "阅读" : item.quantity }}
                        </span>
                        <span class="iconfont icon-dianzan1">
                            {{ item.goodCount == 0 ? "点赞" : item.goodCount }}
                        </span>
                        <span class="iconfont icon-pinglun">
                            {{ item.commentCount == 0 ? "评论" : item.commentCount }}
                        </span>
                    </div>
                </div>
                <router-link :to="`/article/${item.id}`">
                    <el-image style="width: 150px; height: 110px" :src="ip + item.image" />
                </router-link>
            </div>
        </div>
        <!-- 分页按钮 -->
        <div class="pagination" v-if="articleListInfo.length > 0">
            <Pagination :total="total" v-model:page="queryParams.pageNum" v-model:limit="queryParams.pageSize"
                @pagination="getArticleListSearch" />
        </div>
        <!--无文章-->
        <div v-if="articleListInfo.length == 0" class="noArticle">
            很抱歉，暂无文章
        </div>
    </div>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue'
const { proxy } = getCurrentInstance()
import { useRouter } from "vue-router"
const router = useRouter()

const ip = import.meta.env.VITE_APP_ACCESS_IP

const formData = ref({})
const formDataRef = ref()
const rules = {
    keyword: [
        { required: true, message: "输入文章标题或者文章用户名称" },
    ],
}
// 列表数据
const queryParams = ref({
    pageNum: 1,
    pageSize: 10,
    keywords: "",
})
const articleListInfo = ref([])
const total = ref(0)
const getArticleListSearch = async () => {
    formDataRef.value.validate(async (valid) => {
        if (!valid) {
            return
        }
        queryParams.value.keywords = formData.value.keyword
        const res = await proxy.$api.queryArticleKeywordList(queryParams.value)
        if (res.code != 200) {
            proxy.$message.error(res.message)
        } else {
            // console.log("搜索分页查询的文章列表数据为：", res)
            articleListInfo.value = res.data.list
            total.value = res.data.total
        }
    })
}

// 清空输入框并清空数据
const changeInput = ()=> {
    if (formData.value.keyword == "") {
        articleListInfo.value = []
        getArticleListSearch()
    }
}

// 跳转分类页面
const goToCategory = () => {
    router.push('/category')
}

// 跳转标签页面
const goToTag = () => {
    router.push('/tag')
}
</script>

<style lang="scss">
.search-body {
    background: #fff;
    margin-bottom: 10px;
    padding: 15px 0px 0px;
    min-height: calc(100vh - 210px);

    // 搜索框
    .search-panel {
        display: flex;
        justify-content: center;

        .el-input {
            width: 700px;
        }

        .icon-sousuo {
            cursor: pointer;
        }
    }

    // 搜索文章列表
    .article-list {
        padding: 5px 15px 0 15px;

        .article-item-inner {
            padding: 10px 0px;
            border-bottom: 1px solid #ddd;
            display: flex;

            .article-body {
                flex: 1;

                // 用户信息
                .user-info {
                    display: flex;
                    align-items: center;
                    font-size: 14px;

                    .link-info {
                        margin-left: 5px;
                        color: #9a9a9a;
                        text-decoration: none;
                    }

                    .link-info:hover {
                        color: #007fff;
                    }

                    .create-time {
                        font-size: 13px;
                        color: #9a9a9a;
                    }

                    .name {
                        color: #9a9a9a;
                        cursor: pointer;
                    }

                    .name:hover {
                        color: #007fff;
                    }
                }

                // 文章标题
                .title {
                    font-weight: bold;
                    text-decoration: none;
                    color: #4a4a4a;
                    font-size: 16px;
                    margin: 10px 0px;
                    display: inline-block;
                }

                // 文章简介
                .summary {
                    font-size: 14px;
                    color: #86909c;
                }

                .article-info {
                    margin-top: 15px;
                    display: flex;
                    align-items: center;
                    font-size: 13px;

                    .iconfont {
                        color: #86909c;
                        margin-right: 25px;
                        font-size: 14px;
                    }

                    .iconfont:before {
                        padding-right: 3px;
                    }

                    .edit-btn {
                        color: #007fff;
                        cursor: pointer;
                    }
                }
            }
        }
    }

    // 分页组件
    .pagination {
        padding: 0px 15px;
    }

    // 无文章
    .noArticle {
        display: flex;
        flex-direction: column;
        text-align: center;
        padding: 50px 0;
    }
}
</style>