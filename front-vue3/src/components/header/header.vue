<template>
    <div class="header">
        <!--头部内容-->
        <div class="header-content">
            <!--标题-->
            <span class="title">CHEN</span>
            <!--多标签-->
            <div class="tags">
                <router-link :to="`/home`"
                    :class="['tags-item', path == '/home' ? 'active' : 'default']">首页</router-link>
                <router-link :to="`/category`"
                    :class="['tags-item', path == '/category' ? 'active' : 'default']">分类</router-link>
                <router-link :to="`/tag`" :class="['tags-item', path == '/tag' ? 'active' : 'default']">标签</router-link>
                <router-link :to="`/create/article`"
                    :class="['tags-item', path == '/create/article' ? 'active' : 'default']">创作</router-link>
                <router-link :to="`/search`"
                    :class="['tags-item', path == '/search' ? 'active' : 'default']">搜索</router-link>
            </div>
            <!--登录，注册，消息铃铛，头像-->
            <div class="login-register-icon">
                <!--登录，注册-->
                <div class="login-register" v-if="!userId">
                    <span class="login" @click="goToLogin">登录</span>
                    <el-divider direction="vertical"></el-divider>
                    <span class="register" @click="goToRegister">注册</span>
                </div>
                <!--消息铃铛-->
                <div class="message-info" v-if="userId">
                    <el-dropdown>
                        <el-badge :value="messageCountInfo.totalCount" class="item" :hidden="messageCountInfo.totalCount == 0 || messageCountInfo.totalCount == null || messageCountInfo.totalCount == NaN">
                            <div class="iconfont icon-xiaoxizhongxin"></div>
                        </el-badge>
                        <template #dropdown>
                            <el-dropdown-item class="message-item" @click="goToMessageList(1)">
                                <span class="text">回复我的</span>
                                <span class="count-tag" v-if="messageCountInfo.replyCount > 0" >{{ messageCountInfo.replyCount > 99 ? "99+" : messageCountInfo.replyCount }} }}</span>
                            </el-dropdown-item>
                            <el-dropdown-item class="message-item" @click="goToMessageList(2)">
                                <span class="text">赞了文章</span>
                                <span class="count-tag" v-if="messageCountInfo.starArticleCount > 0">{{ messageCountInfo.starArticleCount > 99 ? "99+" : messageCountInfo.starArticleCount }}</span>  
                            </el-dropdown-item>
                            <el-dropdown-item class="message-item" @click="goToMessageList(3)">
                                <span class="text">赞了评论</span>
                                <span class="count-tag" v-if="messageCountInfo.starCommentCount > 0">{{ messageCountInfo.starCommentCount > 99 ? "99+" : messageCountInfo.starCommentCount }}</span> 
                            </el-dropdown-item>
                            <el-dropdown-item class="message-item" @click="goToMessageList(4)">
                                <span class="text">系统消息</span>
                                <span class="count-tag" v-if="messageCountInfo.systemCount > 0">{{ messageCountInfo.systemCount > 99 ? "99+" : messageCountInfo.systemCount }}</span> 
                            </el-dropdown-item>
                        </template>
                    </el-dropdown>
                </div>
                <!--头像-->
                <div v-if="userId">
                    <el-dropdown>
                        <UserAvatar :userId="userId" :src="icon" :showIcon="false" />
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item @click="goToPersonal">修改信息</el-dropdown-item>
                                <el-dropdown-item @click="goToUser(userId)">个人中心</el-dropdown-item>
                                <el-dropdown-item @click="logout">退出</el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                </div>
            </div>
        </div>
    </div>
    <!--登录页面-->
    <Login ref="loginRef"></Login>
    <!--注册页面-->
    <Register ref="registerRef"></Register>
    <!--重置密码页面-->
    <ResetPassword ref="resetPwdRef"></ResetPassword>
</template>

<script setup>
import Login from '../../views/login.vue'
import Register from '../../views/register.vue'
import ResetPassword from '../../views/resetPassword.vue'
import { ElMessageBox } from 'element-plus'
import { ref, getCurrentInstance, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
const { proxy } = getCurrentInstance()
const store = useStore()

const router = useRouter()
// 多标签页监听
const path = ref(null)
watch(
    () => router.currentRoute.value.path,
    (newVal, oldVal) => {
        path.value = newVal
        // console.log("当前的路由：", newVal)
    },
    { immediate: true, deep: true }
)

// 跳转登录页面
const loginRef = ref()
const goToLogin = () => {
    loginRef.value.showPanel()
}

// 跳转注册页面
const registerRef = ref()
const goToRegister = () => {
    registerRef.value.showPanel()
}

// 重置密码
const resetPwdRef = ref()
const handleResetPwd = () => {
    resetPwdRef.value.showPanel()
}

// 监听是否弹出登录框
watch(
    () => proxy.$store.state.showLogin,
    (newVal, oldVal) => {
        if (newVal) {
            goToLogin()
        }
    },
    { immediate: true, deep: true }
)

// 监听是否弹出注册框
watch(
    () => proxy.$store.state.showRegister,
    (newVal, oldVal) => {
        if (newVal) {
            goToRegister()
        }
    },
    { immediate: true, deep: true }
)

// 监听是否展示框
watch(
    () => proxy.$store.state.showResetPwd,
    (newVal, oldVal) => {
        if (newVal) {
            handleResetPwd()
        }
    },
    { immediate: true, deep: true }
)

// 查询消息
const messageCountInfo = ref({})
const queryUserMessageCount = async () => {
    let params = {
        userId: userId == undefined ? 0 : userId.value,
    }
    // console.log("消息数据的用户id：", userId.value)
    const res = await proxy.$api.queryUserMessageCount(params)
    if (res.code != 200) {
        proxy.$message.error(res.message)
    } else {
        messageCountInfo.value = res.data
        // console.log("消息数据：", res)
        // 保存消息数
        store.commit('saveUserMessageCount', res.data)
    }
}

// 监听消息数量
watch(
    () => proxy.$store.state.userMessage,
    (newVal, oldVal) => {
        messageCountInfo.value = newVal || {};
    },
    { immediate: true, deep: true }
)

// 跳转到消息中心
const goToMessageList = (type) => {
    router.push(`/user/message/${type}`)
}

// 用户信息监听
const icon = ref(null)
const userId = ref(null)
watch(
    () => proxy.$store.state.user,
    (newVal, oldVal) => {
        icon.value = newVal == undefined ? "" : newVal.icon
        userId.value = newVal == undefined ? "" : newVal.id
        // console.log("用户信息：", newVal)
        queryUserMessageCount()
    },
    { immediate: true, deep: true }
)

// 跳转到修改信息
const goToPersonal = () => {
    router.push('/personal')
}

// 跳转到个人中心
const goToUser = (userId) => {
    router.push(`/user/${userId}`)
}

// 退出
const logout = ()=> {
    ElMessageBox.confirm('确定要退出登录吗，是否继续？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(()=> {
        proxy.$storage.clearItem("token")
        proxy.$storage.clearItem("user")
        window.location.reload()
    }).catch((error)=> {
        console.log(error)
    })
}
</script>

<style lang="scss" scoped>
.header {
    top: 0px;
    height: 60px;
    box-shadow: 0 2px 6px 0 #ddd;
    position: fixed;
    width: 100%;
    z-index: 1000;
    background: #fff;

    // 头部内容
    .header-content {
        margin: 0px auto;
        align-items: center;
        height: 60px;
        width: 1300px;
        display: flex;

        // 标题
        .title {
            font-size: 20px;
            color: #007fff;
        }

        // 多标签
        .tags {
            flex: 1;
            margin-left: 25px;
            font-weight: bolder;
            font-size: 15px;

            .tags-item {
                margin-left: 25px;
            }

            .tags-item:hover {
                color: #007fff;
            }

            .active {
                color: #007fff;
            }

            .default {
                color: #71777c;
            }

            a {
                text-decoration: none;
            }
        }

        .login-register-icon {
            display: flex;
            align-items: center;

            // 登录，注册
            .login-register {
                color: #71777c;
                font-size: 15px;
                font-weight: bolder;

                .login {
                    cursor: pointer;
                }

                .login:hover {
                    color: #007fff;
                }

                .register {
                    cursor: pointer;
                }

                .register:hover {
                    color: #007fff;
                }
            }

            // 消息
            .message-info {
                margin-left: 5px;
                margin-right: 20px;
                cursor: pointer;

                .icon-xiaoxizhongxin {
                    font-size: 25px;
                    color: rgb(147, 147, 147);
                }
            }

            // 去掉下拉框黑色边框阴影
            .el-dropdown {
                :focus {
                    outline: 0;
                }
            }
        }
    }
}

.message-item {
    display: flex;
    justify-content: space-around;

    .text {
        flex: 1;
    }

    .count-tag {
        height: 15px;
        line-height: 15px;
        min-width: 20px;
        display: inline-block;
        background: #f56c6c;
        border-radius: 10px;
        font-size: 13px;
        text-align: center;
        color: #fff;
        margin-left: 2px;
    }
}
</style>