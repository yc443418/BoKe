<template>
  <div class="tags">
    <el-tag class="tag" size="large" v-for="item, index in tags" :key="item.path"
            :effect="item.title == route.meta.tTitle ? 'dark' : 'plain'" @click="goTo(item.path)" @close="close(index)"
            :disable-transitions="true" :closable="index > 0">
      <i class="circular" v-show="item.title == route.meta.tTitle"/>
      {{ item.title }}
    </el-tag>
  </div>
</template>

<script setup>
import {getCurrentInstance, ref, watch} from 'vue'
import {useRoute, useRouter, onBeforeRouteUpdate } from 'vue-router'

const router = useRouter()
const route = useRoute()
const { proxy } = getCurrentInstance()

// 定义数组
const tags = ref([{
  title: '首页',
  path: '/welcome'
}])

// 事件监听
watch(
    () => route,
    (newVal, oldVal) => {
      if (tags.value.length != undefined && tags.value.length != null && tags.value.length > 18) {
        return proxy.$message.error('z最大只能添加18个标签')
      }
      const boolean = tags.value.find(item => {
        return newVal.path == item.path
      })
      if (!boolean) {
        tags.value.push({
          title: newVal.meta.tTitle,
          path: newVal.path
        })
      }
    },
    { immediate: true, deep: true }
)

// 指定跳转
const goTo = (path)=> {
  router.push(path)
}

// 关闭
const close = (i)=> {
  if (tags.value[i].path == route.path && i != tags.value.length - 1) {
    router.push(tags.value[tags.value.length - 1].path)
  } else if (i == tags.value.length - 1) {
    router.push(tags.value[tags.value.length - 2].path)
  }
  tags.value.splice(i, 1)
}
</script>

<style lang="scss">
.tags {
  padding-left: 20px;
  padding-top: 2px;
  padding-bottom: 2px;

  .tag {
    cursor: pointer;
    margin-right: 3px;

    .circular {
      width: 8px;
      height: 8px;
      margin-right: 4px;
      background-color: #fff;
      border-radius: 50%;
      display: inline-block;
    }
  }
}
</style>