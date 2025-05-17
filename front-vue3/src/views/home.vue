<template>
    <div class="body-container">
        <!--轮播图-->
        <el-carousel>
            <el-carousel-item v-for="(item, index) in bCarouseList" :key="index">
                <img :src="ip + item.icon" style="width: 100%;" />
            </el-carousel-item>
        </el-carousel>
        <!--文章排行-->
        <div class="article-user">
            <div class="article-panel">
                <!--标题-->
                <div class="article">文章排行</div>
                <!--选项标签-->
                <div class="top-tab">
                    <div :class="['tab', orderType == 1 ? 'active' : '']" @click="changeOrderType(1)">热榜</div>
                    <el-divider direction="vertical"></el-divider>
                    <div :class="['tab', orderType == 2 ? 'active' : '']" @click="changeOrderType(2)">发布时间</div>
                    <el-divider direction="vertical"></el-divider>
                    <div :class="['tab', orderType == 3 ? 'active' : '']" @click="changeOrderType(3)">最新</div>
                </div>
                <!--列表-->
                <div class="article-list" v-for="(item, index) in articleList" :key="item.id">
                    <div class="article-item-inner">
                        <div class="article-body">
                            <!--用户信息-->
                            <div class="user-info">
                                <UserAvatar :width="30" :userId="item.userId" :src="item.icon" />
                                <router-link :to="'/user/' + item.userId" class="link-info">{{ item.username
                                    }}</router-link>
                                <el-divider direction="vertical"></el-divider>
                                <div class="create-time"><span class="iconfont icon-shijian"></span>{{
                proxy.$date.formatDate(Date.parse(item.createTime))
            }}</div>
                                <div>&nbsp;·&nbsp;{{ item.loginAddress }}</div>
                                <el-divider direction="vertical"></el-divider>
                                <div class="name" @click="goToCategory">{{ item.categoryName }}</div>
                                <el-divider direction="vertical"></el-divider>
                                <div class="name" @click="goToTag">{{ item.tagName }}</div>
                            </div>
                            <!--文章标题-->
                            <router-link :to="`/article/${item.id}`" class="title">
                                <span v-if="item.topType == 2" class="top">置顶</span>
                                <span>{{ item.title }}</span>
                            </router-link>
                            <div class="summary"> {{ item.summary }}</div>
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
                            <el-image :src="ip + item.image" style="width: 150px; height: 110px;" />
                        </router-link>
                    </div>
                </div>
                <!--分页-->
                <div class="pagination" v-if="total > 0">
                    <Pagination :total="total" v-model:page="queryParams.pageNum" v-model:limit="queryParams.pageSize"
                        @pagination="getArticleList" />
                </div>
            </div>
        </div>
        <!--用户排行-->
        <div class="user-rank">
            <div class="user-panel">
                <div class="user">作者排行</div>
                <div class="user-list" v-for="item in userArticleRankList" :key="item.id">
                    <div class="user-item-inner">
                        <div class="user-info">
                            <UserAvatar :width="50" :userId="item.id" :src="item.icon" />
                            <router-link :to="`/user/` + item.id" class="link-info">
                                {{ item.username }}
                            </router-link>
                        </div>
                        <div class="summary">
                            <span class="iconfont icon-wenzhang">{{ item.articleCount > 1000 ?
                "10000+" : item.articleCount }}</span>
                            <span class="iconfont icon-yuedu">{{ item.quantity > 1000 ? "1000+" : item.quantity
                                }}</span>
                            <span class="iconfont icon-dianzan1">{{ item.goodCount > 1000 ? "1000+" : item.goodCount
                                }}</span>
                            <span class="iconfont icon-pinglun">{{ item.commentCount > 1000 ? "10000+" :
                item.commentCount
                                }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, getCurrentInstance, onMounted } from 'vue'
const { proxy } = getCurrentInstance()
import { useRouter } from "vue-router"
import UserAvatar from '../components/common/userAvatar.vue';
const router = useRouter()

// 轮播图列表
const ip = import.meta.env.VITE_APP_ACCESS_IP
const bCarouseList = ref([])
const getBCarouseList = async () => {
    const res = await proxy.$api.bCarouselList()
    if (res.code != 200) {
        proxy.$message.error(res.message)
    } else {
        bCarouseList.value = res.data
        // console.log("轮播图列表数据：", res.data)
    }
}
getBCarouseList()

// 文章列表
const pages = ref(0)
const orderType = ref(1)
const queryParams = ref({
    pageNum: 1,
    pageSize: 10,
    tagId: "",
    categoryId: "",
    orderType: ""
})
const articleList = ref([])
const total = ref(0)
const getArticleList = async () => {
    queryParams.value.orderType = orderType.value
    const res = await proxy.$api.queryArticleList(queryParams.value)
    if (res.code != 200) {
        proxy.$message.error(res.message)
    } else {
        // console.log("分页查询的文章列表数据为：", res)
        articleList.value = res.data.list
        total.value = res.data.total
        pages.value = res.data.pages
    }
}
onMounted(() => {
    getArticleList()
})

const changeOrderType = (type) => {
    orderType.value = type
    getArticleList()
}

// 跳转分类页面
const goToCategory = () => {
    router.push('/category')
}

// 跳转标签页面
const goToTag = () => {
    router.push('/tag')
}

// 文章用户排行
const userArticleRankList = ref([])
const getUserArticleRankList = async () => {
    const res = await proxy.$api.queryUserArticleRankList()
    if (res.code != 200) {
        proxy.$message.error(res.message)
    } else {
        // console.log("文章用户排行数据为：", res)
        userArticleRankList.value = res.data
    }
}
onMounted(() => {
    getUserArticleRankList()
})
</script>

<style lang="scss">
.article-user {
    margin-top: 10px;
    float: left;
    width: 850px;

    .article-panel {
        margin-bottom: 10px;
        background: #fff;

        .article {
            padding: 10px 15px;
            border-bottom: 1px solid #ddd;
            color: #007fff;
            font-size: 15px;
            font-weight: bold;
        }

        .top-tab {
            display: flex;
            align-items: center;
            padding: 10px 15px;
            font-size: 15px;
            border-bottom: 1px solid #ddd;

            .tab {
                cursor: pointer;
            }

            .active {
                color: #007fff;
            }
        }

        .article-list {
            padding: 5px 15px 0 15px;

            .article-item-inner {
                padding: 10px 0;
                border-bottom: 1px solid #ddd;
                display: flex;

                .article-body {
                    flex: 1;

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

                        .top {
                            font-size: 12px;
                            border-radius: 5px;
                            border: 1px solid #FF6699;
                            color: #FF6699;
                            padding: 0px 5px;
                            margin-right: 10px;
                        }
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
            margin-bottom: 10px;
        }
    }
}

// 作者排行
.user-rank {
    float: right;
    margin-top: 10px;
    width: 350px;

    .user-panel {
        margin-left: 20px;
        margin-bottom: 10px;
        background: #fff;

        .user {
            padding: 10px 15px;
            border-bottom: 1px solid #ddd;
            color: #007fff;
            font-size: 15px;
            font-weight: bold;
        }

        .user-list {
            padding: 5px 15px 0 15px;

            .user-item-inner {
                // 上下，左右
                padding: 10px 0px;
                border-bottom: 1px solid #ddd;

                .user-info {
                    display: flex;

                    .link-info {
                        color: #9a9a9a;
                        text-decoration: none;
                        margin-left: 10px;
                        margin-top: 15px;
                        font-size: 14px;
                    }

                    .link-info:hover {
                        color: #007fff;
                    }
                }

                .summary {

                    margin-top: 10px;
                    font-size: 14px;

                    .iconfont:before {
                        padding-right: 5px;
                    }

                    .iconfont {
                        color: #86909c;
                        margin-right: 35px;
                        font-size: 14px;
                    }
                }
            }
        }
    }
}
</style>