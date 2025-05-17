<template>
  <el-container class="home_container">
    <!--侧边栏-->
    <el-aside :width="isCollapse ? '64px' : '200px'">
      <!--logo和标题-->
      <div class="logo">
        <img src="./../assets/image/logo.jpg" class="aside-logo"/>
        <h4 v-show="!isCollapse">博客运营后台</h4>
      </div>
      <!--侧边菜单栏-->
      <el-menu background-color="#304156" text-color="#c1c1c1" unique-opened router :default-active="defaultActive" :collapse="isCollapse" :collapse-transition="false">
        <el-menu-item :index = "'/welcome'" @click="saveNavState('/welcome')">
          <el-icon>
            <component is="HomeFilled"></component>
          </el-icon>
          <template #title>
            <span>首页</span>
          </template>
        </el-menu-item>
        <el-sub-menu :index="item.id + ''" v-for="item in leftMenuList" :key="item.id">
          <template #title>
            <el-icon>
              <component :is="item.icon"></component>
            </el-icon>
            <span>{{ item.menuName }}</span>
          </template>
          <el-menu-item :index="subItem.url" v-for="subItem in item.menuSvoList" :key="subItem.id" @click="saveNavState(subItem.url)">
            <template #title>
              <el-icon>
                <component :is="subItem.icon"></component>
              </el-icon>
              <span>{{ subItem.menuName }}</span>
            </template>
          </el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>
    <el-container>
      <!--头部-->
      <el-header>
        <!--展开和折叠-->
        <div class="fold-btn">
          <el-icon @click="toggleCollapse">
            <component :is="collapseBtnClass"></component>
          </el-icon>
        </div>
        <!--面包屑-->
        <div class="bread-btn">
          <el-breadcrumb separator="/" v-if="router.currentRoute.value.path != '/welcome'">
            <el-breadcrumb-item :to="{ path: '/welcome'}">
              <span style="font-weight: bold">首页</span>
            </el-breadcrumb-item>
            <el-breadcrumb-item>{{ route.meta.sTitle}}</el-breadcrumb-item>
            <el-breadcrumb-item>{{ route.meta.tTitle}}</el-breadcrumb-item>
          </el-breadcrumb>
          <el-breadcrumb separator="/" v-else>
            <el-breadcrumb-item :to="{ path: '/welcome'}">
              <span style="font-weight: bold">首页</span>
            </el-breadcrumb-item>
          </el-breadcrumb>
        </div>
        <!--头部头像-->
        <HeadImage/>
      </el-header>
      <!--标签页-->
      <Tags/>
      <!--主体页面-->
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import HeadImage from '@/components/headImage.vue'
import Tags from '@/components/tags.vue'

import { ref, getCurrentInstance } from 'vue'
import { useStore } from 'vuex'
import {useRouter, useRoute, onBeforeRouteUpdate} from 'vue-router'
const {proxy} = getCurrentInstance()
const store = useStore()
const router = useRouter()
const route = useRoute()

// 菜单数据
const leftMenuList = proxy.$store.state.leftMenuList
// console.log("菜单数据：", leftMenuList)

// 保持路由激活
const defaultActive = ref(router.currentRoute.value.path)
const saveNavState = (activePath)=> {
  store.commit('saveActivePath', activePath)
}

// 侧边栏展开与折叠
const isCollapse = ref(false)
const collapseBtnClass = ref("Expand")
const toggleCollapse = ()=> {
  isCollapse.value = !isCollapse.value
  if (isCollapse.value) {
    collapseBtnClass.value = 'Fold'
  } else {
    collapseBtnClass.value = 'Expand'
  }
}

// 标签页与菜单联动
onBeforeRouteUpdate((to)=> {
  defaultActive.value = to.path
  store.commit('saveActivePath', to.path)
})
</script>

<style lang="scss">
.home_container {
  height: 100%;

  .el-aside {
    background-color: #304156;

    .logo {
      margin-top: 5px;
      display: flex;
      align-items: center;
      font-size: 13px;
      height: 50px;
      color: #fff;
      .aside-logo {
        width: 32px;
        height: 32px;
        margin: 0 16px;
        border-radius: 5px;
      }
    }
    .el-menu {
      border-right: none;
    }
  }

  .el-header {
    background-color: #f9fafc;
    align-items: center;
    justify-content: space-between;
    display: flex;

    .fold-btn {
      padding-top: 6px;
      font-size: 23px;
      cursor: pointer;
    }

    .bread-btn {
      position: fixed;
      margin-left: 40px;
    }
  }

  .el-main {
    background-color: #eaedf1;
  }
}

</style>