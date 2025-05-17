import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({  // 解决@引入问题
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './src')
        }
    },
    plugins: [vue()],
    lintOnSave: false, //关闭校验
    productionSourceMap: false, //生产环境是否要生成 sourceMap
    publicPath: "/", // 部署应用包时的基本 URL(如果是'./'导致刷新页面出现cannot get/错误)
    outputDir: 'dist', // build 时输出的文件目录
    assetsDir: 'assets', // 放置静态文件夹目录
    server: {
        port: 5002, //运行时的端口
        host: '0.0.0.0', // 运行时域名，设置成'0.0.0.0',在同一个局域网下，如果你的项目在运行，同时可以通过你的http://ip:port/...访问你的项目
        https: false, // 是否启用 https
        open: false, // 是否直接打开浏览器
        proxy: { // 配置后端代理访问的地址
            "/api": {
                target: "http://localhost:5001",
                changeOrigin: true,
            }
        },
        client: {
            overlay: false
        }
    }
})
