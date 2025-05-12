/**
 * store和本地存储
 */
import storage from "../utils/storage"

export default {
    // 存token
    saveToken(state, token) {
        state.token = token
        storage.setItem('token', token)
    },
    // 存储用户信息
    saveSysAdmin(state, sysAdmin) {
        state.sysAdmin = sysAdmin
        storage.setItem('sysAdmin', sysAdmin)
    },
    // 左侧菜单数据列表
    saveLeftMenuList(state, leftMenuList) {
        state.leftMenuList = leftMenuList
        storage.setItem('leftMenuList', leftMenuList)
    },
    // 当前登录用户权限列表
    savePermissionList(state, permissionList) {
        state.permissionList = permissionList
        storage.setItem('permissionList', permissionList)
    },
    // 路由
    saveActivePath(state, activePath) {
        state.activePath = activePath
        storage.setItem('activePath', activePath)
    }
}