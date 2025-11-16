<template>
  <div class="common-layout">
    <el-container class="container">
      <el-header>
        <h1>HVC自动换盘系统</h1>
      </el-header>
      <el-main>
        <el-row class="main-row" :gutter="10">
          <el-col class="row-col" :xs="12" :sm="6" :md="6" :lg="6" :xl="6" v-for="d in devices">
            <el-card>
              <template #header>
                <div class="card-header">
                  <span>{{ d.name }}</span>
                  <el-switch v-model="d.status" @change="changeDeviceStatus(d)" />
                </div>
              </template>
              <el-descriptions :column="1" size="small" title="换盘成功">
                <el-descriptions-item label="换盘时间">{{ d.changeTime }}</el-descriptions-item>
                <el-descriptions-item label="粒子电压">{{ d.updateTime }}</el-descriptions-item>
                <template #extra>
                  <el-button size="small" type="text" @click="hanleDetail(d)">详情</el-button>
                </template>

              </el-descriptions>
            </el-card>
          </el-col>

        </el-row>

      </el-main>
    </el-container>
  </div>
</template>



<script setup>
import { ref, onMounted, reactive } from 'vue';

const device = {
  "id": "1",
  "name": "HVC-A01",
  "status": false,
  "checked": "1",
  "changeTime": "2021-09-01 10:00:00",
  "updateTime": "2021-09-01 10:00:00"
}


const devices = reactive([])


function mockData() {
  const devices = []
  for (let i = 1; i < 13; i++) {
    const device = {
      "id": i,
      "name": "HVC-A0" + i,
      "status": true,
      "changeTime": `2021-09-${i} 10:00:00`,
      "updateTime": `2021-09-${i} 10:10:00`
    }
    devices.push(device)
  }
  return devices
}

onMounted(() => {
  Object.assign(devices, mockData())
})


// 改变设备状态
function changeDeviceStatus(device) {
  console.log(device)
}

function hanleDetail(device) {
  console.log(device)
}


</script>



<style>
body {
  margin: 0 !important;
  padding: 0 !important;
}

.el-container {
  height: 100vh;

}

.el-header {
  display: flex;
  text-align: center;
  justify-content: center;
  align-items: center;
  background-color: #409EFF;
}

.main-row {
  height: 100%;
}


.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 10px;
}
</style>