<template>
  <!--数据类目统计-->
  <div class="part-panel">
    <el-card>
      <div class="card-title" slot="header">
        <span>近周数据类目统计</span>
      </div>
      <div class="data-list">
          <el-row :gutter="40">
            <el-col :span="6" v-for="item in dataStatisticsList">
              <div class="data-item">
                <div class="title"> {{ item.statisticName }}</div>
                <div class="data-panel">
                  <div class="data"> {{ item.statisticCount }}</div>
                  <div class="pre"> 近周新增: <span class="new">{{ item.yesterdayCount }}</span></div>
                </div>
              </div>
            </el-col>
          </el-row>
      </div>
    </el-card>
  </div>
  <!--近周新增文章和用户统计-->
  <div class="article-user-panel">
    <el-card>
      <div class="article-user-title" slot="header">
        <span>近周新增文章和用户统计</span>
      </div>
      <div style="height: 388px; width: 50%; float: left" ref="chartsForArticleDataRef"></div>
      <div style="height: 388px; width: 50%; float: right" ref="chartsForUserDataRef"></div>
    </el-card>
  </div>
</template>

<script setup>
import * as echarts from 'echarts'
import { ref, getCurrentInstance, nextTick, shallowRef} from 'vue'
const { proxy } = getCurrentInstance()

// 数据类目统计
const dataStatisticsList = ref([])
const getDataStatisticsList = async ()=> {
  const res = await proxy.$api.queryDataStatisticsList()
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    dataStatisticsList.value = res.data.list
    // console.log("数据类目统计：", res.data)
  }
}
getDataStatisticsList()

const getOption = (title, xAxisData=[], seriesData=[]) => {
  return {
    title: {
      text: title
    },
    tooltip: {
      trigger: "axis",
      axisPointer: {
        type: "shadow",
        textStyle: {
          color: "red"
        }
      }
    },
    legend: {
      x: 200
    },
    grid: {
      left: 50,
      right: 0,
    },
    xAxis: {
      axisLine: {
        lineStyle: {
          color: "#90979c"
        }
      },
      data: xAxisData,
      axisLabel: {
        interval: 0,
        rotate: "45",
      }
    },
    yAxis: {
      type: "value"
    },
    series: seriesData
  }
}

// 近周新增文章统计
const chartsForArticleDataRef = ref()
const chartsForArticleData = shallowRef()
const initArticleData = ()=> {
  nextTick(()=> {
    chartsForArticleData.value = echarts.init(chartsForArticleDataRef.value)
    getDataArticleCreateList()
  })
}
const getDataArticleCreateList = async ()=> {
  const res = await proxy.$api.queryDataArticleCreateList()
  if (res.code != 200) {
    proxy.$message.error(res.message)
    return
  } else {
    const data = res.data
    const xAxisData = data.dateList
    const seriesData = []
    const colors = ["#1B9CFC"]
    data.articleCountList.forEach((item, index) => {
      seriesData.push({
        name: item.statisticName,
        type: "bar",
        smooth: true,
        data: item.dataList,
        itemStyle: {
          color: colors[index]
        }
      })
    })
    chartsForArticleData.value.setOption(getOption("文章", xAxisData, seriesData))
  }
}
initArticleData()

// 近一周用户新增统计
const chartsForUserDataRef = ref()
const chartsForUserData = shallowRef()
const initUserData = () => {
  nextTick(()=>{
    chartsForUserData.value = echarts.init(chartsForUserDataRef.value)
    getDataUserCreateList()
  })
}
const getDataUserCreateList = async ()=> {
  const res = await proxy.$api.queryDataUserCreateList()
  if (res.code !== 200) {
    proxy.$message.error(res.message)
  } else {
    const data = res.data
    const xAxisData = data.dateList
    const seriesData = []
    const colors = ["#1B9CFC"]
    data.userCountList.forEach((item, index) => {
      seriesData.push({
        name: item.statisticName,
        type: "bar",
        smooth: true,
        data: item.dataList,
        itemStyle: {
          color: colors[index]
        }
      })
    })
    chartsForUserData.value.setOption(getOption("用户", xAxisData, seriesData))
  }
}
initUserData()
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

.article-user-panel {
  margin-top: 0px;
  margin-bottom: 10px;

  .article-user-title {
    font-weight: bold;
    font-size: 20px;
    margin-bottom: 20px;


  }

}
</style>