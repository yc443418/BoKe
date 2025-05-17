<template>
    <div class="body-container">
        <div class="tag">
            <!--头部标签数量-->
            <div class="header">
                <span>共计{{ bTagsList.length - 1 }}个标签</span>
            </div>
            <!--左侧标签列表-->
            <div class="tag-panel">
                <el-tabs v-model="activeName" tab-position="left" @tab-click="handleTabClick">
                    <el-tab-pane v-for="(item, index) in bTagsList" :name="index + ''" :key="index">
                        <template #label>
                            <span>{{ item.tagName }}</span>
                        </template>
                    </el-tab-pane>
                </el-tabs>
            </div>
            <!--右侧文章列表-->
            <div class="tags-article" v-for="(item, index) in articleList" :key="item.id">
                <div class="article-item-inner">
                    <div class="article-body">
                        <!--用户信息-->
                        <div class="user-info">
                            <UserAvatar :width="30" :userId="item.userId" :src="item.icon" />
                            <router-link :to="'/user/' + item.userId" class="link-info">{{
                    item.username
                }}</router-link>
                            <el-divider direction="vertical"></el-divider>
                            <div class="create-time"><span class="iconfont icon-shijian"></span> {{
                        proxy.$date.formatDate(Date.parse(item.createTime)) }}</div>
                            <div>&nbsp;·&nbsp;{{ item.loginAddress }}</div>
                            <el-divider direction="vertical"></el-divider>
                            <div>{{ item.tagName }}</div>
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
            <div class="pagination" v-if="total > 0">
                <Pagination :total="total" v-model:page="queryParams.pageNum" v-model:limit="queryParams.pageSize"
                    @pagination="getArticleList" />
            </div>
            <!--无文章-->
            <div v-else class="noArticle">
                很抱歉，暂无文章
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue'
const { proxy } = getCurrentInstance()
const ip = import.meta.env.VITE_APP_ACCESS_IP
// 文章标签
const bTagsList = ref([
    {
        id: null,
        tagName: "全部"
    }
])
// 回显第一个标签页
const activeName = ref('0')
const getBTagsList = async () => {
    const res = await proxy.$api.bTagsList()
    // console.log("文章标签：", res)
    if (res.code !== 200) {
        proxy.$message.error(res.message)
    } else {
        bTagsList.value.push(...res.data)
    }
}
getBTagsList()

// 文章列表数据
const queryParams = ref({
    pageNum: 1,
    pageSize: 10,
    tagId: "",
})
const articleList = ref([])
const total = ref(0)
const getArticleList = async () => {
    const res = await proxy.$api.queryArticleList(queryParams.value)
    if (res.code != 200) {
        proxy.$message.error(res.message)
    } else {
        articleList.value = res.data.list
        total.value = res.data.total
    }
}
getArticleList()

// 点击分类获取文章标签数据
const handleTabClick = async (tab) => {
    let item = bTagsList.value[tab.index]
    queryParams.value.tagId = item.id
    const res = await proxy.$api.queryArticleList(queryParams.value)
    if (res.code != 200) {
        proxy.$message.error(res.message)
    } else {
        articleList.value = res.data.list
        total.value = res.data.total
    }
}
</script>

<style lang="scss" scoped>
.tag {

    // 头部分类数量
    .header {
        text-align: center;
        padding: 10px;
        color: #007fff;
        font-size: 1rem;
        font-weight: bold;
        border-bottom: 1px solid #61666d;
        margin-bottom: 10px;
    }

    // 左侧标签列表
    .tag-panel {
        margin-top: 150px;
        // 位置固定不变
        position: fixed;
    }

    // 右侧文章列表
    .tags-article {
        margin-left: 150px;
        background-color: #ffffff;
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
                    color: #4e5969;

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
        background-color: #ffffff;
        margin-left: 150px;
        padding: 0px 15px;
        margin-bottom: 10px;
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