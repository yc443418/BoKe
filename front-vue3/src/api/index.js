import request from "../utils/request"
export default {
    // 验证码
    captcha() {
        return request({
            url: "/captcha",
            method: 'get',
            showLoading: false
        })
    },
    // 登录
    login(params) {
        return request({
            url: "/login",
            method: 'post',
            data: params,
            showLoading: false
        })
    },
    // qq邮箱注册验证码
    qqMail(params) {
        return request({
            url: "/qqMail",
            method: 'post',
            data: params,
            showLoading: false
        })
    },
    // 注册
    register(params) {
        return request({
            url: '/register',
            method: 'post',
            data: params,
            showLoading: false
        })
    },
    // 重置密码
    resetPwd(params) {
        return request({
            url: '/user/reset/password',
            method: 'post',
            data: params,
            showLoading: false
        })
    },
    // 文章分类接口
    bCategorySelectList() {
        return request({
            url: "/bCategorySelect/list",
            method: 'get'
        })
    },
    // 文章标签接口
    bTagsList() {
        return request({
            url: "/bTagsSelect/list",
            method: 'get'
        })
    },
    // 图片接口
    upload(params) {
        return request({
            url: "/upload",
            method: 'post',
            data: params,
            showLoading: false
        })
    },
    // 新建文章
    createArticle(params) {
        return request({
            url: "/bArticle/add",
            method: 'post',
            data: params,
            showLoading: false
        })
    },
    // 轮播图接口
    bCarouselList() {
        return request({
            url: "/bCarousel/list",
            method: 'get'
        })
    },
    // 文章列表
    queryArticleList(params) {
        return request({
            url: "/bArticle/list",
            method: 'get',
            data: params
        })
    },
    // 文章用户排行
    queryUserArticleRankList() {
        return request({
            url: "/user/article/rank/list",
            method: 'get',
        })
    },
    // 根据关键字分页查询文章列表
    queryArticleKeywordList(params) {
        return request({
            url: "/bArticle/keyword/list",
            method: 'get',
            data: params
        })
    },
    // 文章详情
    articleDetail(params) {
        return request({
            url: "/bArticle/detail",
            method: 'get',
            data: params,
        })
    },
    // 文章点赞和取消点赞
    doLikeArticle(id) {
        const data = {
            id
        }
        return request({
            url: "/bArticle/like",
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 文章评论列表
    commentList(params) {
        return request({
            url: "/bArticle/commentVo/list",
            method: 'get',
            data: params,
        })
    },
    // 新增评论
    addArticleComment(params) {
        return request({
            url: "/bArticle/add/comment",
            method: 'post',
            data: params,
            showLoading: false
        })
    },
    // 评论点赞和取消点赞
    commentLike(id) {
        const data = {
            id
        }
        return request({
            url: "/bArticle/comment/like",
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 评论置顶
    commentTop(params) {
        return request({
            url: "/bArticle/comment/top",
            method: 'post',
            data: params,
            showLoading: false
        })
    },
    // 修改个人信息
    updateUserInfo(data) {
        return request({
            url: '/user/updateUserInfo',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 修改个人密码
    updateUserPassword(data) {
        return request({
            url: '/user/updateUserPassword',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 文章用户详情
    articleUserDetail(params) {
        return request({
            url: "/user/articleUserDetail",
            method: 'get',
            data: params,
        })
    },
    // 文章用户(文章，评论，点赞列表)
    queryArticleUserList(params) {
        return request({
            url: "/user/bArticle/list",
            method: 'get',
            data: params
        })
    },
    // 根据文章id查询文章
    queryBArticleById(id) {
        const data = {
            id
        }
        return request({
            url: "/bArticle/articleId",
            method: 'get',
            data: data,
        })
    },
    // 修改文章
    updateBArticle(params) {
        return request({
            url: "/bArticle/update",
            method: 'put',
            data: params,
            showLoading: false
        })
    },
    // 分页查询用户消息列表
    queryUserMessageList(params) {
        return request({
            url: "/user/message/list",
            method: 'get',
            data: params
        })
    },
    // 查询用户消息
    queryUserMessageCount(params) {
        return request({
            url: "/user/message/count",
            method: 'get',
            data: params
        })
    },
}