<template>
  <div class="comment-list">
    <!--左侧头像-->
    <UserAvatar
      :width="50"
      :userId="commentData.userId"
      :src="commentData.icon == undefined ? currentUserInfo.icon : commentData.icon"
    />
    <!--右侧信息-->
    <div class="comment-info">
      <!--用户名和作者-->
      <div class="username">
        <span @click="goToUserCenter(commentData.userId)" class="name">{{
          commentData.username == undefined
            ? currentUserInfo.username
            : commentData.username
        }}</span>
        <span v-if="commentData.userId == articleUserId" class="tag-author">作者</span>
      </div>
      <!--评论文章的内容-->
      <div class="comment-content">
        <span class="tag" v-if="commentData.topType == 2">置顶</span>
        <span v-html="commentData.content"></span>
      </div>
      <!--时间，地址，点赞，评论，置顶/取消置顶-->
      <div class="op-panel">
        <div class="time">
          <span>{{ commentData.createTime }}</span>
          <span class="address">&nbsp;·&nbsp;{{ commentData.loginAddress }}</span>
        </div>
        <div
          :class="['iconfont icon-dianzan1', commentData.likeType == 2 ? 'active' : '']"
          @click="commentLike(commentData)"
        >
          {{ commentData.goodCount > 0 ? commentData.goodCount : "点赞" }}
        </div>
        <div class="iconfont icon-pinglun" @click="showReplyPanel(commentData, 0)">
          回复
        </div>
        <div
          class="iconfont icon-zhiding"
          v-if="
            articleUserId == currentUserInfo.id && commentData.userId == articleUserId
          "
          @click="setTop(commentData)"
        >
          {{ commentData.topType == 1 ? "置顶" : "取消" }}
        </div>
      </div>
      <!--二级评论列表-->
      <div class="comment-sub-list" v-if="commentData.children">
        <div class="comment-sub-item" v-for="sub in commentData.children">
          <UserAvatar
            :width="30"
            :userId="sub.userId"
            :src="sub.icon == undefined ? currentUserInfo.icon : sub.icon"
          />
          <div class="comment-sub-info">
            <!--用户名和作者-->
            <div class="username">
              <span class="name" @click="goToUserCenter(sub.userId)">{{
                sub.username == undefined ? currentUserInfo.username : sub.username
              }}</span>
              <span v-if="sub.userId == articleUserId" class="tag-author">作者</span>
              <span class="reply">回复</span>
              <span @click="goToUserCenter(sub.replyUserId)" class="reply-name"
                >@{{ sub.replyUsername }}</span
              >
              <span>：</span>
              <span class="sub-content" v-html="sub.content"></span>
            </div>
            <div class="op-panel">
              <div class="time">
                <span>{{ sub.createTime }}</span>
                <span class="address">&nbsp;·&nbsp;{{ sub.loginAddress }}</span>
              </div>
              <div
                :class="['iconfont icon-dianzan1', sub.likeType == 2 ? 'active' : '']"
                @click="commentLike(sub)"
              >
                {{ sub.goodCount > 0 ? sub.goodCount : "点赞" }}
              </div>
              <div class="iconfont icon-pinglun" @click="showReplyPanel(sub, 1)">
                回复
              </div>
            </div>
          </div>
        </div>
      </div>
      <!--二级评论回复-->
      <div class="reply-info" v-if="commentData.showReply">
        <CommentArticle
          :width="30"
          :userId="currentUserInfo.id"
          :src="currentUserInfo.icon"
          :articleId="articleId"
          :pCommentId="pCommentId"
          :replyUserId="replyUserId"
          :placeholder="placeholder"
          @postCommentFinish="postCommentFinish"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import CommentArticle from "@/components/article/commentArticle.vue";

import { ref, getCurrentInstance } from "vue";
const { proxy } = getCurrentInstance();
import { useRouter } from "vue-router";
const router = useRouter();

const props = defineProps({
  articleId: {
    type: Number,
  },
  commentData: {
    type: Object,
  },
  articleUserId: {
    type: Number,
  },
  currentUserInfo: {
    type: Object,
  },
});

// 跳转到指定个人信息
const goToUserCenter = (userId) => {
  router.push(`/user/${userId}`);
};
const emit = defineEmits(["hiddenAllReply", "reloadList"]);
//显示评论框
const pCommentId = ref(0);
const replyUserId = ref(null);
const placeholder = ref(null);
const showReplyPanel = (data, type) => {
  const haveShow =
    props.commentData.showReply == undefined ? false : props.commentData.showReply;
  emit("hiddenAllReply");
  if (type == 0) {
    props.commentData.showReply = !haveShow;
  } else {
    props.commentData.showReply = true;
  }
  pCommentId.value = props.commentData.id;
  replyUserId.value = data.userId;
  placeholder.value = "回复 @" + data.username;
};

const postCommentFinish = () => {
  emit("reloadList");
  placeholder.value = undefined;
};

// 点赞
const commentLike = async (data) => {
  const res = await proxy.$api.commentLike(data.id);
  if (res.code != 200) {
    proxy.$message.error(res.message);
  } else {
    data.goodCount = res.data.goodCount;
    data.likeType = res.data.likeType;
  }
};

// 置顶
const setTop = async (data) => {
  const res = await proxy.$api.commentTop({
    id: data.id,
    topType: data.topType == 2 ? 1 : 2,
  });
  if (res.code != 200) {
    proxy.$message.error(res.message);
  } else {
    if (data.topType == 2) {
      proxy.$message.success("已取消置顶");
    } else if (data.topType == 1) {
      proxy.$message.success("已置顶");
    }
    emit("reloadList")
  }
};
</script>

<style lang="scss">
// 评论列表
.comment-list {
  display: flex;
  padding-top: 15px;
  // 右侧信息
  .comment-info {
    flex: 1;
    margin-left: 10px;
    border-bottom: 1px solid #ddd;
    padding-bottom: 15px;

    // 用户名
    .username {
      .name {
        font-size: 14px;
        color: #61666d;
        margin-right: 10px;
        cursor: pointer;
      }

      .name:hover {
        color: #007fff;
      }

      .tag-author {
        background: #ff6699;
        color: #fff;
        font-size: 12px;
        border-radius: 2px;
        padding: 0px 3px;
      }
    }

    // 评论文章的内容
    .comment-content {
      margin-top: 5px;
      font-size: 15px;
      line-height: 22px;
      .tag {
        margin-right: 5px;
        font-size: 12px;
        border-radius: 3px;
        padding: 0px 5px;
        color: #ff6699;
        border: 1px solid #ff6699;
      }
    }

    // 时间，地址，点赞，评论，置顶/取消置顶
    .op-panel {
      display: flex;
      align-items: center;
      margin-top: 5px;
      font-size: 13px;
      color: #61666d;

      .time {
        margin-right: 20px;
      }

      .iconfont {
        margin-right: 15px;
        font-size: 14px;
        color: #9499a0;
        cursor: pointer;
      }

      .iconfont::before {
        margin-right: 3px;
      }

      .iconfont:hover {
        color: #007fff;
      }

      .active {
        color: #007fff;
      }

      .el-dropdown {
        :focus {
          outline: 0;
        }
      }
    }

    // 二级评论列表
    .comment-sub-list {
      margin-top: 10px;

      .comment-sub-item {
        display: flex;
        margin: 8px 0px;
        font-size: 14px;

        .comment-sub-info {
          margin-left: 5px;

          .username {
            .reply {
              margin: 0px 5px;
            }

            .reply-name {
              color: #007fff;
              cursor: pointer;
            }

            .sub-content {
              font-size: 15px;
            }
          }
        }
      }
    }
  }
}
</style>
