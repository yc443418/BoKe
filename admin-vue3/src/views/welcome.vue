<template>
  <!--数据类目统计-->
  <div class="part-panel">
    <el-card>
      <div class="card-title" slot="header">
        <span>近期数据类目统计</span>
      </div>
      <div class="data-list">
        <el-row :gutter="40">
          <el-col :span="6" v-for="item in dataStatisticsList">
            <div class="data-item">
              <div class="title">{{ item.statisticName }}</div>
              <div class="data-panel">
                <div class="data">{{ item.statisticCount }}</div>
                <div class="pre">近期新增：<span class="new">{{ item.yesterdayCount }}</span></div>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import * as echarts from 'echarts'
import { ref, getCurrentInstance, nextTick, shallowRef } from 'vue'
const { proxy } = getCurrentInstance()

// 数据类目统计
const dataStatisticsList = ref([])
const getDataStatisticsList = async () => {
  const res = await proxy.$api.queryDataStatisticsList()
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    dataStatisticsList.value = res.data.List
    // console.log("数据类目统计:", res.data)
  }
}
getDataStatisticsList()

</script>

<style lang="scss">
.part-panel {
  margin-top: 0px;
  margin-bottom: 10px;

  .card-title {
    font-weight: bold;
    font-size: 20px;
    margin-bottom: 10px;
  }

  .data-list {
    margin-right: 20px;

    .data-item {
      background: #f4f9fd;
      color: #9a9fa6;
      padding: 20px 10px 20px 10px;
      border-radius: 5px;
      width: 100%;

      .data-panel {
        display: flex;
        align-items: center;
        justify-content: space-between;

        .data {
          font-size: 25px;
          color: #007fff;
          font-weight: bold;
          margin-top: 10px;
        }

        .pre {
          margin-top: 5px;
          .new {
            color: #ff6873;
          }
        }
      }
    }
  }
}
</style>