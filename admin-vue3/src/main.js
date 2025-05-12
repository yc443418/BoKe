import { createApp } from 'vue'
import App from './App.vue'
// router
import router from './router/router.js'
// 引入element plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// 全局样式
import '@/styles/index.scss'
// 输出环境
console.log("环境地址为：", import.meta.env.VITE_APP_ENV)
console.log("baseUrl地址为：", import.meta.env.VITE_APP_BASE_API)
// api
import api from './api'
// storage
import storage from "./utils/storage.js";
// store
import store from "./store";

// app实例
const app = createApp(App)
app.use(router)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.config.globalProperties.$api = api
app.config.globalProperties.$storage = storage
app.use(store)
app.mount('#app')
