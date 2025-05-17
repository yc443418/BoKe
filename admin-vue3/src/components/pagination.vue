<template>
  <div class="pagination-container">
    <el-pagination :background="background" v-model:current-page="currentPage" v-model:page-size="pageSize"
                   :layout="layout" :page-sizes="pageSizes" :total="total" @size-change="handleSizeChange"
                   @current-change="handleCurrentChange"/>
  </div>
</template>

<script setup>
import {computed} from "vue";

const props = defineProps({
  total: {
    required: true,
    type: Number,
  },
  page: {
    type: Number,
    default: 1,
  },
  limit: {
    type: Number,
    default: 10,
  },
  pageSizes: {
    type: Array,
    default() {
      return [10, 50, 100, 500, 1000]
    },
  },
  layout: {
    type: String,
    default: 'total, sizes, prev, pager, next, jumper'
  },
  background: {
    type: Boolean,
    default: true
  }
})

const emits = defineEmits()
const currentPage = computed({
  get() {
    return props.page
  },
  set(val) {
    emits('update:page', val)
  }
})

const pageSize = computed({
  get() {
    return props.limit
  },
  set(val) {
    emits('update:limit', val)
  }
})

const handleSizeChange = (val) => {
  if (currentPage.value * val > props.total) {
    currentPage.value = 1
  }
  emits('pagination', {page: currentPage.value, limit: val})
}

const handleCurrentChange = (val) => {
  emits('pagination', {page: val, limit: pageSize.value})
}
</script>

<style lang="scss">
.pagination-container {
  padding: 5px 0px;
}
</style>