<template>
  <div class="top-right-btn">
    <el-row>
      <el-tooltip class="item" effect="dark" :content="showSearch ? '隐藏搜索' : '显示搜索'" placement="top" v-if="search">
        <el-button circle icon="Search" @click="toggleSearch()"></el-button>
      </el-tooltip>
      <el-tooltip class="item" effect="dark" content="刷新列表">
        <el-button circle icon="Refresh" @click="refresh()"></el-button>
      </el-tooltip>
      <el-tooltip class="item" effect="dark" content="显示或隐藏列表字段" v-if="columns">
        <el-button circle icon="Menu" @click="showColumn()"></el-button>
      </el-tooltip>
    </el-row>
    <el-dialog :title="title" v-model="open" append-to-body>
      <el-transfer :titles="['显示', '隐藏']" v-model="value" :data="columns" @change="dataChange"></el-transfer>
    </el-dialog>
  </div>
</template>

<script setup>
import {ref} from 'vue'
const props = defineProps({
  showSearch: {
    type: Boolean,
    default: true
  },
  columns: {
    type: Array,
  },
  search: {
    type: Boolean,
    default: true
  }
})
const emits = defineEmits(['update:showSearch', 'queryTable'])
const value = ref([])
const title = ref('显示或隐藏数据列表字段')
const open = ref(false)
const toggleSearch = ()=> {
  emits('update:showSearch', !props.showSearch)
}
const refresh = ()=> {
  emits('queryTable')
}
const showColumn = () => {
  open.value = true
}
const dataChange = (data) => {
  for (let item in props.columns) {
    const key = props.columns[item].key
    props.columns[item].visible = !data.includes(key)
  }
}
for (let item in props.columns) {
  if (props.columns[item].visible == false) {
    value.value.push(parseInt(item))
  }
}
</script>

<style lang="scss">
.top-right-btn {
  margin-left: auto;
}
</style>