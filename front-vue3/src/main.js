import { createApp } from 'vue'
import App from './App.vue'
import router from './router/router'
//引入element plus
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
// 导入全局样式
import '@/styles/index.scss'
// 输出环境
console.log("环境地址为：",import.meta.env.VITE_APP_ENV)
console.log("baseUrl地址为：",import.meta.env.VITE_APP_BASE_API)
// 弹出组件
import dialog from '@/components/common/dialog.vue'
// 图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
// element-plus中文
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
// api
import api from './api'
// storage
import storage from './utils/storage'
// store
import store from './store'
// 图标
import '@/assets/icon/iconfont.css'
// 头像
import UserAvatar from '@/components/common/userAvatar.vue'
// 时间转换工具类
import * as date from '@/utils/date'
// 分页组件
import Pagination from '@/components/common/pagination.vue'

// app实例
const app = createApp(App)
app.use(router)
app.use(ElementPlus, {
    locale: zhCn
})
app.component("Dialog", dialog)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}
app.config.globalProperties.$api = api
app.config.globalProperties.$storage = storage
app.use(store)
app.component("UserAvatar", UserAvatar)
app.config.globalProperties.$date = date
app.component('Pagination', Pagination)
app.mount('#app')