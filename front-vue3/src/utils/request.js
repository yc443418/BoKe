/**
 * axios封装
 */
import axios from "axios"
import {
    ElMessage,
    ElLoading
} from 'element-plus'
import storage from './storage'
import store from "../store"

// 创建axios
const service = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_API,
    timeout: 5000
})

let loading = null
// 请求拦截
service.interceptors.request.use((req) => {
    const headers = req.headers
    const token = storage.getItem("token") || {}
    if (!headers.Authorization) {
        headers.Authorization = 'Bearer ' + token
    }
    // 加载loading
    if (req.showLoading) {
        loading = ElLoading.service({
            lock: true,
            text: "拼命加载中......",
            background: 'rgba(0,0,0,0.8)'
        })
    }
    return req
}, (error) => {
    if (error.config.showLoading && loading) {
        loading.close()
        ElMessage.error("请求发送失败")
    }
})

// 响应拦截（200-成功，403-登录过期，500-服务端失败）
service.interceptors.response.use((res) => {
    const {
        showLoading
    } = res.config
    if (showLoading && loading) {
        loading.close()
    }
    const {
        code,
        data,
        message
    } = res.data
    if (code == 403) {
        store.commit('saveShowLogin', true)
        storage.clearItem("token")
        storage.clearItem("user")
    } else if (code == 200) {
        return res.data
    } else {
        ElMessage.error(message)
    }
}, (error) => {
    if (error.config.showLoading && loading) {
        loading.close()
    }
})

// 核心函数
const request = (options) => {
    const { showLoading = true } = options
    options.method = options.method || 'get'
    if (options.method.toLowerCase() == 'get') {
        options.params = options.data
    }
    options.showLoading = showLoading
    service.defaults.baseURL =
        import.meta.env.VITE_APP_BASE_API
    return service(options)
}

export default request