<template>
  <div class="comment-body">
    <!--评论，热榜，最新-->
    <div class="comment-title">
      <div class="title">
        评论<span class="count">{{ total }}</span>
      </div>
      <div class="tab">
        <span
          @click="orderChange(3)"
          :class="['tab-item', orderType == 3 ? 'active' : '']"
          >最新</span
        >
        <el-divider direction="vertical" />
        <span
          @click="orderChange(1)"
          :class="['tab-item', orderType == 1 ? 'active' : '']"
          >热榜</span
        >
      </div>
    </div>
    <!--发送评论-->
    <div class="comment-form-panel">
      <CommentArticle
        :width="50"
        :userId="currentUserInfo.id"
        :src="currentUserInfo.icon"
        :articleId="articleId"
        :pCommentId="0"
        @postCommentFinish="postCommentFinish"
      />
    </div>
    <!--评论列表-->
    <div v-for="data in commentListInfo">
      <CommentListItem
        :articleId="articleId"
        :articleUserId="articleUserId"
        :currentUserInfo="currentUserInfo"
        :commentData="data"
        @hiddenAllReply="hiddenAllReplyHandler"
        @reloadList="getCommentList"
      />
    </div>
    <!-- 分页按钮 -->
    <div class="pagination" v-if="total > 0">
      <Pagination
        :total="total"
        v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize"
        @pagination="getCommentList"
      />
    </div>
  </div>
</template>

<script setup>
import CommentArticle from "@/components/article/commentArticle.vue";
import CommentListItem from "@/components/article/commentListItem.vue";
import { ref, getCurrentInstance, watch } from "vue";
const { proxy } = getCurrentInstance();

const props = defineProps({
  articleId: {
    type: Number,
  },
  articleUserId: {
    type: Number,
  },
});

// 当前登录用户信息
const currentUserInfo = ref({});
watch(
  () => proxy.$store.state.user,
  (newVal, oldVal) => {
    currentUserInfo.value = newVal || {};
  },
  { immediate: true, deep: true }
);

// 最新，热榜排序
const orderType = ref(3);
const orderChange = (type) => {
  orderType.value = type;
  getCommentList();
};
// 评论列表
const queryParams = ref({});
const commentListInfo = ref([]);
const total = ref(0);
const getCommentList = async () => {
  let params = {
    pageSize: queryParams.value.pageSize,
    pageNum: queryParams.value.pageNum,
    orderType: orderType.value,
    articleId: props.articleId,
    userId: currentUserInfo.value.id == undefined ? 0 : currentUserInfo.value.id,
  };
  const res = await proxy.$api.commentList(params);
  if (res.code !== 200) {
    proxy.$message.error(res.message);
  } else {
    commentListInfo.value = res.data.list;
    total.value = res.data.total;
    // console.log("评论列表数据：", res.data);
  }
};
getCommentList();

const emit = defineEmits(["updateCommentTotal"]);
const postCommentFinish = (resData) => {
  if (commentListInfo.value == undefined) {
    getCommentList();
  } else {
    commentListInfo.value.unshift(resData);
    // 更新数量
    total.value = total.value + 1;
    commentListInfo.value.total = total;
    emit("updateCommentTotal", total);
  }
};

// 隐藏所有的回复框
const hiddenAllReplyHandler = () => {
  commentListInfo.value.forEach((element) => {
    element.showReply = false;
  });
};
</script>

<style lang="scss">
.comment-body {
  // 评论，热榜，最新
  .comment-title {
    display: flex;
    align-items: center;

    .title {
      font-size: 20px;

      .count {
        font-size: 14px;
        padding: 0px 15px;
      }
    }

    .tab {
      display: flex;
      font-size: 15px;

      .tab-item {
        cursor: pointer;
      }

      .active {
        color: #007fff;
      }
    }
  }

  // 发送评论
  .comment-form-panel {
    margin-top: 20px;
  }

  // 分页
  .pagination {
    margin-top: 10px;
  }
}
</style>
