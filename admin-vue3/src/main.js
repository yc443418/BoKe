import { createApp } from 'vue'
import App from './App.vue'
// router
import router from './router/router'
//引入element plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// 全局样式
import '@/styles/index.scss'
// 输出环境
console.log("环境地址为：",import.meta.env.VITE_APP_ENV)
console.log("baseUrl地址为：",import.meta.env.VITE_APP_BASE_API)
// api
import api from './api'
// storage
import storage from "./utils/storage"
// store
import store from "./store"
// 展示菜单树方法
import common from "@/utils/common"
// element-plus中文
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
// 搜索/刷新/隐藏组件
import tipHover from '@/components/tipHover.vue'
// 分页组件
import pagination from '@/components/pagination.vue'
// 权限
import { hasPermission } from "./permission/permission"

// app实列
const app = createApp(App)
app.use(router)
app.use(ElementPlus, {
    locale: zhCn
})
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.config.globalProperties.$api = api
app.config.globalProperties.$storage = storage
app.config.globalProperties.$common = common
app.component('TipHover', tipHover)
app.component('Pagination', pagination)
app.use(store)
app.use(hasPermission)
app.mount('#app')
