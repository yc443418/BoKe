/**
 * 获取数据方法
*/
import { createStore } from "vuex"
import mutations from "./mutations"
import storage from "../utils/storage"

// 获取存储数据
const state = {
    token: "" || storage.getItem("token"),
    user: "" || storage.getItem("user"),
    showLogin: "" || storage.getItem("showLogin"),
    showRegister: "" || storage.getItem("showRegister"), 
    showResetPwd: "" || storage.getItem("showResetPwd"), 
    userMessage: "" || storage.getItem("userMessage"),
}

// 导出数据
export default createStore({
    state,
    mutations
})