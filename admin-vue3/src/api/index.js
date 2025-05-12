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
            url: "/index/data/statistics/list",
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
}