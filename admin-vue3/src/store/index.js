/**
 * 获取的方法
 */
import { createStore } from 'vuex'
import mutations from "./mutations"
import storage from "../utils/storage"

const state = {
    token: "" || storage.getItem("token"),
    sysAdmin: "" || storage.getItem("sysAdmin"),
    leftMenuList: "" || storage.getItem("leftMenuList"),
    permissionList: "" || storage.getItem("permissionList"),
    saveActivePath: "" || storage.getItem("saveActivePath"),
};

export default createStore({
    state,
    mutations
})