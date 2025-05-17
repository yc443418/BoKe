<template>
  <div class="body-container message-center">
    <!--首页和消息中心-->
    <div class="user-banner">
      <router-link to="/home" class="link-info">首页</router-link>
      <span class="iconfont icon-xiangyoujiantou"></span>
      <span>消息中心</span>
    </div>
    <!--消息列表-->
    <div class="message-panel">
      <!--多标签-->
      <div class="tab-list">
        <el-tabs :model-value="activeTabName" @tab-change="tabChange">
          <el-tab-pane name="1">
            <template #label>
              <div class="tab-item">
                <span>回复我的</span>
                <span class="count-tag" v-if="messageCountInfo.replyCount > 0">{{
                  messageCountInfo.replyCount > 99 ? "99+" : messageCountInfo.replyCount
                }}</span>
              </div>
            </template>
          </el-tab-pane>
          <el-tab-pane name="2">
            <template #label>
              <div class="tab-item">
                <span>赞了文章</span>
                <span class="count-tag" v-if="messageCountInfo.starArticleCount > 0">{{
                  messageCountInfo.starArticleCount > 99
                    ? "99+"
                    : messageCountInfo.starArticleCount
                }}</span>
              </div>
            </template>
          </el-tab-pane>
          <el-tab-pane name="3">
            <template #label>
              <div class="tab-item">
                <span>赞了评论</span>
                <span class="count-tag" v-if="messageCountInfo.starCommentCount > 0">{{
                  messageCountInfo.starCommentCount > 99
                    ? "99+"
                    : messageCountInfo.starCommentCount
                }}</span>
              </div>
            </template>
          </el-tab-pane>
          <el-tab-pane name="4">
            <template #label>
              <div class="tab-item">
                <span>系统消息</span>
                <span class="count-tag" v-if="messageCountInfo.systemCount > 0">{{
                  messageCountInfo.systemCount > 99 ? "99+" : messageCountInfo.systemCount
                }}</span>
              </div>
            </template>
          </el-tab-pane>
        </el-tabs>
      </div>
      <!--列表-->
      <div class="message-list" v-for="(item, index) in messageListInfo" :key="item.id">
        <!--回复我的-->
        <div class="message-item" v-if="item.messageType == 1">
          <UserAvatar
            :userId="item.sendUserId"
            :width="50"
            :src="item.sendUserIcon"
          ></UserAvatar>
          <div class="message-content">
            <router-link class="link" :to="`/user/${item.sendUserId}`"
              >@{{ item.sendUsername }}</router-link
            >
            对我的文章【<router-link class="link" :to="`/article/${item.articleId}`">{{
              item.articleTitle
            }}</router-link
            >】发表了评论
            <span class="create-time"
              ><span class="iconfont icon-shijian"></span>
              {{ proxy.$date.formatDate(Date.parse(item.createTime)) }}</span
            >
            <div class="reply-content" v-html="item.messageContent"></div>
          </div>
        </div>
        <!--赞了文章-->
        <div class="message-item" v-if="item.messageType == 2">
          <UserAvatar
            :userId="item.sendUserId"
            :width="50"
            :src="item.sendUserIcon"
          ></UserAvatar>
          <div class="message-content">
            <router-link class="link" :to="`/user/${item.sendUserId}`"
              >@{{ item.sendUsername }}</router-link
            >
            赞了我的文章【<router-link class="link" :to="`/article/${item.articleId}`">{{
              item.articleTitle
            }}</router-link
            >】
            <span class="create-time"
              ><span class="iconfont icon-shijian"></span>
              {{ proxy.$date.formatDate(Date.parse(item.createTime)) }}</span
            >
          </div>
        </div>
        <!--赞了评论-->
        <div class="message-item" v-if="item.messageType == 3">
          <UserAvatar
            :userId="item.sendUserId"
            :width="50"
            :src="item.sendUserIcon"
          ></UserAvatar>
          <div class="message-content">
            <router-link class="link" :to="`/user/${item.sendUserId}`"
              >@{{ item.sendUsername }}</router-link
            >
            在文章【<router-link class="link" :to="`/article/${item.articleId}`">{{
              item.articleTitle
            }}</router-link
            >】中赞了我的评论
            <span class="create-time"
              ><span class="iconfont icon-shijian"></span>
              {{ proxy.$date.formatDate(Date.parse(item.createTime)) }}</span
            >
            <div class="reply-content" v-html="item.messageContent"></div>
          </div>
        </div>
        <!--系统消息-->
        <div class="message-item" v-if="item.messageType == 4">
          <div class="message-content">
            <span v-html="item.messageContent"></span>
            <span class="create-time"
              ><span class="iconfont icon-shijian"></span>
              {{ proxy.$date.formatDate(Date.parse(item.createTime)) }}</span
            >
          </div>
        </div>
      </div>
      <!-- 分页 -->
      <div v-if="total > 0">
        <Pagination
          :total="total"
          v-model:page="queryParams.pageNum"
          v-model:limit="queryParams.pageSize"
          @pagination="getUserMessageList"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, getCurrentInstance, watch } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useStore } from "vuex";
const { proxy } = getCurrentInstance();
const router = useRouter();
const route = useRoute();
const store = useStore();

// 监听用户信息
const userId = ref(null);
watch(
  () => proxy.$store.state.user,
  (newVal, oldVal) => {
    if (newVal) {
      userId.value = newVal.id;
    }
  },
  { immediate: true, deep: true }
);

// 标签跳转
const activeTabName = ref(1);
const tabChange = (type) => {
  router.push(`/user/message/${type}`);
};

// 列表数据
const queryParams = ref({
  pageNum: 1,
  pageSize: 10,
  type: "",
  userId: "",
});
const messageListInfo = ref([]);
const total = ref(0);
const getUserMessageList = async () => {
  let params = {
    pageNum: queryParams.value.pageNum,
    pageSize: queryParams.value.pageSize,
    type: activeTabName.value,
    userId: userId.value == undefined ? 0 : userId.value,
  };
  const res = await proxy.$api.queryUserMessageList(params);
  if (res.code != 200) {
    proxy.$message.error(res.message);
  } else {
    // console.log("分页查询的文章用户消息数据为：", res)
    messageListInfo.value = res.data.list;
    total.value = res.data.total;
    store.commit("saveUserMessageCount", res.data.userMessageVo);
  }
};

// 监听消息类型及拉取数据
watch(
  () => route.params.type,
  (newVal, oldVal) => {
    if (newVal) {
      activeTabName.value = newVal;
      getUserMessageList();
    }
  },
  { immediate: true, deep: true }
);

// 监听消息的变化
const messageCountInfo = ref({});
watch(
  () => proxy.$store.state.userMessage,
  (newVal, oldVal) => {
    messageCountInfo.value = newVal || {};
  },
  { immediate: true, deep: true }
);
</script>

<style lang="scss">
.message-center {
  margin-bottom: 10px;
  // 首页和消息中心
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
  // 消息列表
  .message-panel {
    background: #ffffff;
    padding: 5px 15px 0 15px;

    // 多标签
    .tab-list {
      position: relative;

      .tab-item {
        position: relative;
        padding: 0px 10px;

        .count-tag {
          position: absolute;
          right: -10px;
          height: 15px;
          bottom: 2px;
          line-height: 15px;
          min-width: 20px;
          display: inline-block;
          background: #f56c6c;
          border-radius: 10px;
          font-size: 13px;
          text-align: center;
          color: #fff;
          margin-left: 10px;
        }
      }
    }
    // 列表
    .message-list {
      .message-item {
        display: flex;
        justify-content: flex-start;
        align-items: center;
        font-size: 14px;
        border-bottom: 1px solid #ddd;
        padding: 10px;
        .message-content {
          margin-left: 5px;
          .link {
            text-decoration: none;
            color: #007fff;
          }
          .create-time {
            margin-left: 10px;
          }

          .reply-content {
            border-left: 2px solid #11a8f4;
            padding-left: 5px;
            margin-top: 5px;
          }
        }
      }
    }
  }
}
</style>
