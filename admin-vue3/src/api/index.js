/**
 * api接口的封装
 */
import request from "../utils/request"

export default {
    // 登录
    login(params) {
        return request({
            url: '/sysAdmin/login',
            method: 'post',
            data: params,
            showLoading: false,
        })
    },
    // 数据类目统计
    queryDataStatisticsList() {
        return request({
            url: "/index/statistics/list",
            method: 'get'
        })
    },
    // 近一周新增文章统计
    queryDataArticleCreateList() {
        return request({
            url: "/index/data/article/create/list",
            method: 'get'
        })
    },
    // 近一周用户新增统计
    queryDataUserCreateList() {
        return request({
            url: "/index/data/user/create/list",
            method: 'get'
        })
    },
    // 菜单列表
    queryMenuList(params) {
        return request({
            url: "/sysMenu/list",
            method: 'get',
            data: params
        })
    },
    // 新增菜单
    addMenu(data) {
        return request({
            url: '/sysMenu/add',
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 根据id查询菜单详情
    menuInfo(id) {
        const data = {
            id
        }
        return request({
            url: '/sysMenu/info',
            method: 'get',
            data: data
        })
    },
    // 修改菜单
    menuUpdate(data) {
        return request({
            url: '/sysMenu/update',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 根据id删除菜单
    menuDelete(id) {
        const data = {
            id
        }
        return request({
            url: '/sysMenu/delete',
            method: 'delete',
            data: data,
            showLoading: false
        })
    },
    // 角色列表
    queryRoleList(params) {
        return request({
            url: "/sysRole/list",
            method: 'get',
            data: params
        })
    },
    // 查询角色下拉列表
    querySysRoleVoList() {
        return request({
            url: "/sysRole/vo/list",
            method: 'get'
        })
    },
    // 新增角色
    addRole(data) {
        return request({
            url: '/sysRole/add',
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 根据id查询角色
    roleInfo(id) {
        const data = {
            id
        }
        return request({
            url: '/sysRole/info',
            method: 'get',
            data: data
        })
    },
    // 角色修改
    roleUpdate(data) {
        return request({
            url: '/sysRole/update',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 根据id删除角色
    deleteRole(id) {
        const data = {
            id
        }
        return request({
            url: '/sysRole/delete',
            method: 'delete',
            data: data,
            showLoading: false
        })
    },
    // 修改角色状态
    updateRoleStatus(id, status) {
        const data = {
            id,
            status
        }
        return request({
            url: "/sysRole/updateStatus",
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 根据id查询角色菜单列表
    queryRoleMenuIdList(id) {
        const data = {
            id
        }
        return request({
            url: "/sysRole/vo/idList",
            method: 'get',
            data: data
        })
    },
    // 为角色分配菜单
    assignPermissions(id, menuIds) {
        const data = {
            id,
            menuIds
        }
        return request({
            url: "/sysRole/assignPermissions",
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 后台用户列表
    queryAdminList(params) {
        return request({
            url: "/sysAdmin/list",
            method: 'get',
            data: params
        })
    },
    // 修改后台用户状态
    updateAdminStatus(id, status) {
        const data = {
            id,
            status
        }
        return request({
            url: "/sysAdmin/updateStatus",
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 新增后台用户
    addAdmin(data) {
        return request({
            url: '/sysAdmin/add',
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 根据id查询后台用户
    adminInfo(id) {
        const data = {
            id
        }
        return request({
            url: '/sysAdmin/info',
            method: 'get',
            data: data
        })
    },
    // 修改后台用户
    adminUpdate(data) {
        return request({
            url: '/sysAdmin/update',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 重置密码
    resetPassword(id, password) {
        const data = {
            id,
            password
        }
        return request({
            url: '/sysAdmin/updatePassword',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 根据id删除后台用户
    deleteAdmin(id) {
        const data = {
            id
        }
        return request({
            url: '/sysAdmin/delete',
            method: 'delete',
            data: data,
            showLoading: false
        })
    },
    // 修改后台用户个人信息
    adminUpdatePersonal(data) {
        return request({
            url: '/sysAdmin/updatePersonal',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 修改后台用户个人密码
    adminUpdatePersonalPassword(data) {
        return request({
            url: '/sysAdmin/updatePersonalPassword',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 系统参数信息列表
    queryConfigList(params) {
        return request({
            url: "/sysConfig/list",
            method: 'get',
            data: params
        })
    },
    // 新增参数信息
    addConfig(data) {
        return request({
            url: '/sysConfig/add',
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 根据id查询参数信息
    configInfo(id) {
        const data = {
            id
        }
        return request({
            url: '/sysConfig/info',
            method: 'get',
            data: data
        })
    },
    // 修改参数信息
    configUpdate(data) {
        return request({
            url: '/sysConfig/update',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 根据id删除参数信息
    deleteConfig(id) {
        const data = {
            id
        }
        return request({
            url: '/sysConfig/delete',
            method: 'delete',
            data: data,
            showLoading: false
        })
    },
    // 刷新缓存
    refresh() {
        return request({
            url: '/sysConfig/refresh',
            method: 'delete',
            showLoading: false
        })
    },
    // 文章轮播图列表
    queryBCarouselList(params) {
        return request({
            url: "/bCarousel/list",
            method: 'get',
            data: params
        })
    },
    // 新增文章轮播图
    addBCarousel(data) {
        return request({
            url: '/bCarousel/add',
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 根据id查询文章轮播图
    bCarouselInfo(id) {
        const data = {
            id
        }
        return request({
            url: '/bCarousel/info',
            method: 'get',
            data: data
        })
    },
    // 修改文章轮播图
    bCarouselUpdate(data) {
        return request({
            url: '/bCarousel/update',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 根据id删除文章轮播图
    deleteBCarousel(id) {
        const data = {
            id
        }
        return request({
            url: '/bCarousel/delete',
            method: 'delete',
            data: data,
            showLoading: false
        })
    },
    // 文章标签列表
    queryTagsList(params) {
        return request({
            url: "/bTags/list",
            method: 'get',
            data: params
        })
    },
    // 新增文章标签
    addTags(data) {
        return request({
            url: '/bTags/add',
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 根据id查询文章标签
    tagsInfo(id) {
        const data = {
            id
        }
        return request({
            url: '/bTags/info',
            method: 'get',
            data: data
        })
    },
    // 修改文章标签
    tagsUpdate(data) {
        return request({
            url: '/bTags/update',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 根据id删除文章标签
    deleteTags(id) {
        const data = {
            id
        }
        return request({
            url: '/bTags/delete',
            method: 'delete',
            data: data,
            showLoading: false
        })
    },
    // 文章标签下拉列表
    bTagsSelectList() {
        return request({
            url: "/bTagsSelect/list",
            method: 'get'
        })
    },
    // 文章分类列表
    queryBCategoryList(params) {
        return request({
            url: "/bCategory/list",
            method: 'get',
            data: params
        })
    },
    // 新增文章分类
    addBCategory(data) {
        return request({
            url: '/bCategory/add',
            method: 'post',
            data: data,
            showLoading: false
        })
    },
    // 根据id查询文章分类
    bCategoryInfo(id) {
        const data = {
            id
        }
        return request({
            url: '/bCategory/info',
            method: 'get',
            data: data
        })
    },
    // 修改文章分类
    bCategoryUpdate(data) {
        return request({
            url: '/bCategory/update',
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 根据id删除文章分类
    deleteBCategory(id) {
        const data = {
            id
        }
        return request({
            url: '/bCategory/delete',
            method: 'delete',
            data: data,
            showLoading: false
        })
    },
    // 文章分类下拉列表
    bCategorySelectList() {
        return request({
            url: "/bCategorySelect/list",
            method: 'get'
        })
    },
    // 文章信息列表
    queryArticleList(params) {
        return request({
            url: "/bArticle/list",
            method: 'get',
            data: params
        })
    },
    // 设置文章信息是否置顶
    updateTopType(id, topType) {
        const data = {
            id,
            topType
        }
        return request({
            url: "/bArticle/updateTopType",
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 设置文章信息是否状态删除
    updateStatus(id, status) {
        const data = {
            id,
            status
        }
        return request({
            url: "/bArticle/updateStatus",
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 查询文章评论列表
    queryArticleCommentListByArticleId(params) {
        return request({
            url: "/bArticle/comment",
            method: 'get',
            data: params
        })
    },
    // 文章评论信息
    queryArticleCommentList(params) {
        return request({
            url: "/bArticle/comment/list",
            method: 'get',
            data: params,
        })
    },
    // 设置是否状态删除文章评论
    updateCommentStatus(id, status) {
        const data = {
            id,
            status
        }
        return request({
            url: "/bArticle/comment/updateStatus",
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 前台用户列表
    queryUserList(params) {
        return request({
            url: "/user/list",
            method: 'get',
            data: params
        })
    },
    // 修改用户状态
    updateUserStatus(id, status) {
        const data = {
            id,
            status
        }
        return request({
            url: "/user/updateStatus",
            method: 'put',
            data: data,
            showLoading: false
        })
    },
    // 发送消息
    sendMessage(data) {
        return request({
            url: '/user/send/message',
            method: 'post',
            data: data,
            showLoading: false
        })
    },
}