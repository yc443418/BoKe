/**
 * axios封装
 */


import axios from "axios"
import router from "../router/router"
import { ElMessage, ElLoading } from "element-plus"
import storage from "./storage"

// 创建axios实例
const service = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_API,
    timeout: 5000
});

let loading = null
// 请求拦截
service.interceptors.request.use((req)=>{
    const headers = req.headers
    const token = storage.getItem("token") || {}
    if (!headers.Authorization) {
        headers.Authorization = 'Bearer ' + token
    }
    // 加载loading
    if (req.showLoading) {
        loading = ElLoading.service({
            lock: true,
            text: 'loading',
            background: 'rgba(0, 0, 0, 0.8)',
        })
    }
    return req
}, (error) => {
    if (error.config.showLoading && loading) {
        loading.close()
        ElMessage.error('请求发送失败')
    }
})

// 响应拦截
service.interceptors.response.use((res)=>{
    const { showLoading } = res.config
    if (showLoading && loading) {
        loading.close()
    }
    const {code, date, message} = res.data
    if (code == 403) {
        ElMessage.error(message)
        storage.clearAll()
        setTimeout(()=>{
            router.push('/login')
        }, 1500)
    } else if (code == 200) {
        return res.data
    } else {
        ElMessage.error(message)
    }
}, (error)=>{
    if (error.config.showLoading && loading) {
        loading.close()
    }
})

// 核心请求函数
function request(options) {
    const { showLoading = true} = options
    options.method = options.method || 'get'

    if (options.method.toLowerCase() == 'get') {
        options.params = options.data
    }
    options.showLoading = showLoading
    service.defaults.baseURL = import.meta.env.VITE_APP_BASE_API
    return service(options)
}

export default request