/**
 * store和本地存储
*/
import storage from '../utils/storage'

export default {
    // 存储token
    saveToken(state, token) {
        state.token = token
        storage.setItem('token', token)
    },
    // 存储用户信息
    saveUser(state, user) {
        state.user = user
        storage.setItem('user', user)
    },
    // 是否弹出登录框
    saveShowLogin(state, showLogin) {
        state.showLogin = showLogin
        storage.setItem('showLogin', false)
    },
    // 是否弹出注册框
    saveShowRegister(state, showRegister) {
        state.showRegister = showRegister
        storage.setItem('showRegister', false)
    },
    // 是否弹出重置密码框
    saveShowResetPwd(state, showResetPwd) {
        state.showResetPwd = showResetPwd
        storage.setItem('showResetPwd', false)
    },
    // 保存消息
    saveUserMessageCount(state, userMessage) {
        state.userMessage = userMessage
        storage.setItem("userMessage", userMessage)
    },
}