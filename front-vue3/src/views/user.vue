<template>
  <div class="body-container user-center">
    <!--首页和个人中心-->
    <div class="user-banner">
      <router-link to="/home" class="link-info">首页</router-link>
      <span class="iconfont icon-xiangyoujiantou"></span>
      <span>个人中心</span>
    </div>
    <el-row :gutter="20">
      <!--个人信息-->
      <el-col :xs="24" :sm="24" :md="8" :lg="6" :xl="6" style="margin-bottom: 10px">
        <div class="user-info">
          <div class="user-icon">
            <UserAvatar :width="120" :userId="userInfo.id" :src="userInfo.icon" />
          </div>
          <el-form style="margin-top: 15px">
            <el-col :span="24">
              <el-form-item label="用户账号：">{{ userInfo.username }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="用户邮箱：">{{ userInfo.email }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="登录的IP：">{{ userInfo.loginIp }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="登录地址：">{{ userInfo.loginAddress }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="创建时间：">{{ userInfo.createTime }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="用户备注：">{{ userInfo.note }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="发布文章：">{{ userInfo.articleCount }}</el-form-item>
            </el-col>
            <el-col :span="24">
              <el-form-item label="文章点赞：">{{ userInfo.likeCount }}</el-form-item>
            </el-col>
          </el-form>
        </div>
      </el-col>
      <!--文章用户（列表，评论，点赞，修改资料，修改密码）-->
      <el-col :xs="24" :sm="24" :md="16" :lg="18" :xl="18">
        <div class="article-user-info">
          <el-tabs :model-value="activeName" @tab-change="changeTab">
            <el-tab-pane label="文章列表" :name="1"> </el-tab-pane>
            <el-tab-pane label="文章评论" :name="2"> </el-tab-pane>
            <el-tab-pane label="文章点赞" :name="3"> </el-tab-pane>
          </el-tabs>
          <!--列表-->
          <div
            class="tags-article"
            v-for="(item, index) in articleUserList"
            :key="item.id"
          >
            <div class="article-item-inner">
              <div class="article-body">
                <!--用户信息-->
                <div class="user-info">
                  <UserAvatar :width="30" :userId="item.userId" :src="item.icon" />
                  <router-link :to="'/user/' + item.userId" class="link-info">{{
                    item.username
                  }}</router-link>
                  <el-divider direction="vertical"></el-divider>
                  <div class="create-time">
                    <span class="iconfont icon-shijian"></span>
                    {{ proxy.$date.formatDate(Date.parse(item.createTime)) }}
                  </div>
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
                  <span
                    class="iconfont icon-bianji edit-btn"
                    v-if="currentUserInfo && activeName == 1"
                    @click="goToEditBArticle(item.id)"
                    >编辑</span
                  >
                </div>
              </div>
              <router-link :to="`/article/${item.id}`">
                <el-image style="width: 150px; height: 110px" :src="ip + item.image" />
              </router-link>
            </div>
          </div>
          <!-- 分页 -->
          <div v-if="total > 0">
            <Pagination
              :total="total"
              v-model:page="queryParams.pageNum"
              v-model:limit="queryParams.pageSize"
              @pagination="getArticleUserList"
            />
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, getCurrentInstance, watch } from "vue";
import { useRoute } from "vue-router";
const { proxy } = getCurrentInstance();
const route = useRoute();
import { useRouter } from "vue-router";
const router = useRouter();

// 图片访问ip
const ip = import.meta.env.VITE_APP_ACCESS_IP;

// 文章用户信息
const userId = ref(null);
const userInfo = ref({});
const getUserInfo = async () => {
  let params = {
    userId: userId.value,
  };
  const res = await proxy.$api.articleUserDetail(params);
  if (res.code !== 200) {
    proxy.$message.error(res.message);
  } else {
    userInfo.value = res.data;
    // console.log("文章用户数据：", res.data);
  }
};

// 右侧信息
const activeName = ref(1);
// 文章用户(文章，评论,点赞列表)
const changeTab = (type) => {
  activeName.value = type;
  if (type == 4 || type == 5) {
    return;
  }
  getArticleUserList();
};
const queryParams = ref({
  pageNum: 1,
  pageSize: 10,
  type: "",
  userId: "",
});
const articleUserList = ref([]);
const total = ref(0);
const getArticleUserList = async () => {
  let params = {
    pageNum: queryParams.value.pageNum,
    pageSize: queryParams.value.pageSize,
    type: activeName.value,
    userId: userId.value,
  };
  const res = await proxy.$api.queryArticleUserList(params);
  if (res.code != 200) {
    proxy.$message.error(res.message);
  } else {
    // console.log("分页查询的文章用户(文章，评论,点赞列表)数据为：", res);
    articleUserList.value = res.data.list;
    total.value = res.data.total;
  }
};

// 文章用户事件监听
watch(
  () => route.params.userId,
  (newVal, oldVal) => {
    if (newVal) {
      userId.value = newVal;
      getUserInfo();
      getArticleUserList();
    }
  },
  { immediate: true, deep: true }
);

// 重新设置当前的用户信息
const currentUserInfo = ref(false);
const resetCurrentUser = () => {
  const loginUser = proxy.$store.state.user;
  if (loginUser && loginUser.id == userId.value) {
    currentUserInfo.value = true;
  } else {
    currentUserInfo.value = false;
  }
};
watch(
  () => proxy.$store.state.user,
  (newVal, oldVal) => {
    resetCurrentUser();
  },
  { immediate: true, deep: true }
);

// 跳转分类页面
const goToCategory = () => {
    router.push('/category')
}

// 跳转标签页面
const goToTag = () => {
    router.push('/tag')
}

// 跳转到编辑页面
const goToEditBArticle = (articleId)=> {
    router.push(`/edit/article/${articleId}`)
}
</script>

<style lang="scss">
.user-center {
  // 左侧部分
  .user-banner {
    color: #9ba7b9;
    line-height: 35px;
    font-size: 15px;

    .link-info {
      text-decoration: none;
      color: #007fff;
    }

    .icon-xiangyoujiantou {
      padding: 0px 5px;
    }
  }

  // 左侧头像
  .user-info {
    background-color: #ffffff;

    .user-icon {
      padding: 20px 75px;
      border-bottom: 1px solid #ddd;
    }
  }

  // 右侧部分
  .article-user-info {
    background-color: #ffffff;
    padding: 5px 15px 0 15px;
    margin-bottom: 10px;

    // 修改资料头像
    .avatar-uploader {
      border: 1px dashed #d9d9d9;
      border-radius: 6px;
      cursor: pointer;
      width: 100px;
      height: 100px;
      position: relative;
      overflow: hidden;

      .avatar {
        width: 100px;
        height: 100px;
        display: block;
      }
    }

    // 列表
    .tags-article {
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
              color: #86909c;
              cursor: pointer;
            }

            .edit-btn:hover {
              color: #007fff;
            }
          }
        }
      }
    }
  }
}
</style>
