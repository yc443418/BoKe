<template>
  <div class="body-container article-detal-body" v-if="articleDetail.id > 0">
    <!--导航-->
    <div class="board-info">
      <router-link to="/home" class="link-info">首页</router-link>
      <span class="iconfont icon-xiangyoujiantou"></span>
      <span>{{ articleDetail.categoryName }}</span>
      <span class="iconfont icon-xiangyoujiantou"></span>
      <span>{{ articleDetail.title }}</span>
    </div>
    <!--详情-->
    <div class="detail-container" style="width: 1000px;">
      <div class="artilce-detail">
        <!--标题-->
        <div class="title">
          {{ articleDetail.title }}
        </div>
        <!--用户信息-->
        <div class="user-info">
          <UserAvatar :width="50" :userId="articleDetail.userId" :src="articleDetail.icon" />
          <div class="user-info-detail">
            <span class="username">{{ articleDetail.username }}</span>
            <div class="time-info">
              <span class="iconfont icon-shijian"></span>{{ proxy.$date.formatDate(Date.parse(articleDetail.createTime))
              }}
              <span class="iconfont icon-yanjing"> {{ articleDetail.quantity }}</span>
            </div>
          </div>
        </div>
        <!--文章内容-->
        <div class="detail" id="detail" v-html="articleDetail.content"></div>
      </div>
      <!--文章评论-->
      <div class="comment-panel" id="view-comment">
        <CommentList v-if="articleDetail.id" :articleId="articleDetail.id" :articleUserId="articleDetail.userId" @updateCommentTotal="updateCommentTotal"/>
      </div>
    </div>
    <!--目录-->
    <div class="menu-apnel">
      <div class="menu-container">
        <div class="menu-title">目录</div>
        <div class="toc-list">
          <template v-if="tocArrary.length == 0">
            <span class="no-toc">未解析到目录</span>
          </template>
          <template v-else>
            <div v-for="item in tocArrary">
              <span @click="goToToc(item.id)" class="toc-item" :style="{ 'padding-left': item.level * 5 + 'px' }">{{
    item.title }}</span>
            </div>
          </template>
        </div>
      </div>
    </div>
  </div>
  <!--点赞和评论-->
  <div class="quick-panel" :style="{ left: quickPanelLeft + 'px' }">
    <!--点赞-->
    <el-badge :value="articleDetail.goodCount" type="info">
      <div class="quick-item" @click="doLike">
        <span :class="['iconfont icon-dianzan1', haveLike ? 'have-like' : '']"></span>
      </div>
    </el-badge>
    <!--评论-->
    <el-badge :value="articleDetail.commentCount" type="info">
      <div class="quick-item" @click="goToPosition('view-comment')">
        <span class="iconfont icon-pinglun"></span>
      </div>
    </el-badge>
  </div>
  <!--图片预览-->
  <ImageViewer ref="imageViewerRef" :imageList="previewImgList"></ImageViewer>
</template>

<script setup>
import hljs from "highlight.js"
import "highlight.js/styles/atom-one-light.css" //样式
import ImageViewer from '../components/common/imageViewer.vue'
import CommentList from "../components/article/commentList.vue"
import { ref, getCurrentInstance, onMounted, nextTick } from 'vue'
const { proxy } = getCurrentInstance()
import { useRoute } from "vue-router"
import UserAvatar from '../components/common/userAvatar.vue';
const route = useRoute()

// 文章详情
const currentUserInfo = proxy.$store.state.user
const articleDetail = ref({})
const haveLike = ref(false)
const getArticleDetail = async () => {
  let params = {
    id: route.params.articleId,
    userId: currentUserInfo == undefined ? 0 : currentUserInfo.id
  }
  const res = await proxy.$api.articleDetail(params)
  if (res.code != 200) {
    proxy.$message.error(res.message)
  } else {
    articleDetail.value = res.data.bArticleDetailVo
    haveLike.value = res.data.bArticleDetailVo.haveLike
    // console.log("文章数据详情:", res.data.bArticleDetailVo)
    // 生成目录
    makeToc()
    // 图片预览
    imagePreview()
    // 代码高亮
    highlightCode()
  }
}
onMounted(() => {
  getArticleDetail()
})

// 计算快捷键位置
const quickPanelLeft = ref(
  (window.innerWidth - 1200) / 2 - 110
)

// 点赞
const doLike = async () => {
  let id = articleDetail.value.id
  const res = await proxy.$api.doLikeArticle(id)
  if (res.code != 200) {
    proxy.$message.error(res.message)
  } else {
    haveLike.value = !haveLike.value
    let goodCount = 1
    if (!haveLike.value) {
      goodCount = -1
    }
    articleDetail.value.goodCount = articleDetail.value.goodCount + goodCount
    proxy.$message.success(res.data)
  }
}

// 定位到评论
const goToPosition = (id) => {
  document.querySelector("#" + id).scrollIntoView()
}

// 获取目录
const tocArrary = ref([])
const makeToc = () => {
  nextTick(() => {
    const tocTags = ["H1", "H2", "H3", "H4", "H5", "H6"]
    const contentDom = document.querySelector("#detail")
    const childNodes = contentDom.childNodes
    let index = 0
    childNodes.forEach((item) => {
      let tagName = item.tagName
      if (tagName == undefined || !tocTags.includes(tagName.toUpperCase())) {
        return true
      }
      index++
      let id = "toc" + index
      item.setAttribute("id", id)
      tocArrary.value.push({
        id: id,
        title: item.innerText,
        level: Number.parseInt(tagName.substring(1)),
        offsetTop: item.offsetTop,
      })
    })
  })
}
const goToToc = (id) => {
  const dom = document.querySelector("#" + id)
  dom.scrollIntoView({
    behavior: "smooth",
    block: "start",
  })
}

// 图片预览
const imageViewerRef = ref(null)
const previewImgList = ref([])
const imagePreview = ()=> {
  nextTick(()=>{
    const imageNodeList = document.querySelector("#detail").querySelectorAll("img")
    const imageList = []
    imageNodeList.forEach((item, index) => {
      const src = item.getAttribute("src")
      imageList.push(src)
      item.addEventListener("click", ()=>{
        imageViewerRef.value.show(index)
      })
    })
    previewImgList.value = imageList
  })
}

// 代码高亮
const highlightCode = ()=>{
  nextTick(()=>{
    let blocks = document.querySelectorAll("pre code")
    blocks.forEach((item)=> {
      hljs.highlightBlock(item)
    })
  })
}

// 更新评论数量
const updateCommentTotal = (commentCount)=> {
  articleDetail.value.commentCount = commentCount
}
</script>

<style lang="scss">
.article-detal-body {
  position: relative;

  // 导航
  .board-info {
    font-size: 15px;
    line-height: 20px;

    .link-info {
      text-decoration: none;
      color: #007fff;
    }

    .icon-xiangyoujiantou {
      margin: 0px 5px;
    }
  }

  // 详情
  .detail-container {
    margin-top: 5px;

    .artilce-detail {
      background: #fff;
      padding: 10px;

      // 标题
      .title {
        font-weight: bolder;
        font-size: 15px;
      }

      // 用户信息
      .user-info {
        margin-top: 15px;
        display: flex;
        padding-bottom: 10px;
        border-bottom: 1px solid #ddd;

        .user-info-detail {
          .username {
            margin-left: 14px;
            text-decoration: none;
            color: #4e5969;
            font-size: 15px;
          }

          .time-info {
            margin-top: 5px;
            font-size: 13px;
            color: #61666d;

            .iconfont {
              margin-left: 14px;
            }

            .iconfont::before {
              padding-right: 3px;
            }
          }
        }
      }

      // 文章内容
      .detail {
        letter-spacing: 1px;
        line-height: 22px;

        img {
          max-width: 100%;
          cursor: pointer;
        }

        a {
          text-decoration: none;
          color: #007fff;
        }
      }
    }

    // 文章评论
    .comment-panel {
      margin-top: 15px;
      background: #fff;
      margin-bottom: 10px;
      padding: 20px;
    }
  }

  // 目录
  .menu-apnel {
    position: absolute;
    top: 26px;
    right: 0px;
    width: 190px;

    .menu-container {
      width: 235px;
      position: fixed;
      background: #fff;

      .menu-title {
        padding: 10px;
        border-bottom: 1px solid #ddd;
      }

      .toc-list {
        max-height: calc(100vh - 200px);
        overflow: auto;
        padding: 5px;

        .no-toc {
          margin-left: 70px;
          color: #5f5d5d;
          line-height: 40px;
          font-size: 15px;
        }

        .toc-item {
          cursor: pointer;
          display: block;
          line-height: 35px;
          overflow: hidden;
          white-space: nowrap;
          color: #555666;
          border-radius: 3px;
          font-size: 14px;
          border-left: 2px solid, #fff;
        }

        .toc-item:hover {
          background: #eeeded;
        }
      }
    }
  }
}

// 点赞和评论
.quick-panel {
  position: fixed;
  width: 50px;
  top: 250px;
  text-align: center;

  // 数字角标
  .el-badge__content.is-fixed {
    top: 5px;
    right: 15px;
  }

  .quick-item {
    width: 50px;
    height: 50px;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 50%;
    background: #fff;
    margin-bottom: 30px;
    cursor: pointer;

    .iconfont {
      font-size: 22px;
      color: #61666d;
    }

    // 点赞后颜色改变
    .have-like {
      color: red;
    }
  }
}
</style>