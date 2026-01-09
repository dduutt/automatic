<template>
  <div class="common-layout">
    <el-container class="container">
      <el-header>
        <h1>HVC自动换盘系统</h1>
      </el-header>
      <el-main>
        <el-row class="main-row" :gutter="5">
          <el-col class="row-col" :xs="24" :sm="12" :md="12" :lg="6" :xl="6" v-for="config in configs" :key="config.id">
            <el-card class="card" header-class="card-header">
              <template #header>

                <span>{{ config.name }}</span>
                <el-switch v-model="config.status" @change="changeDeviceStatus(config)" />
              </template>

              <div class="card-body">
                <el-table :data="infos[config.name]" border style="width: 100%;" height="200" max-height="200"
                  size="small">
                  <el-table-column prop="time" label="时间" align="center" />
                  <el-table-column prop="payload" label="消息" align="center" :show-overflow-tooltip="true" />
                </el-table>

              </div>

            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive, onUnmounted } from 'vue';
import api from './api.js'
import mqtt from 'mqtt';
import { Command } from '@tauri-apps/plugin-shell';

const configs = reactive([])
const infos = ref({})
const mqttClient = ref(null)
const sidecar = ref(null)
onMounted(async () => {
  try {
    await startServer()
    const interval = setInterval(async () => {
      try {
        const configsResp = await api.getConfigs()
        if (!configsResp.data) {
          ElMessage.error("获取配置信息失败")
          return
        }
        Object.assign(configs, configsResp.data)
        await initMQTT()
        clearInterval(interval)

      } catch (error) {
        console.log(error)
      }
    }, 1000)
  } catch (error) {
    console.log(error)
  }
})

onUnmounted(() => {
  if (mqttClient.value && mqttClient.value.connected) {
    mqttClient.value.end()
  }
  if (sidecar.value) {
    sidecar.value.kill()
  }
})

// 改变设备状态
async function changeDeviceStatus(config) {
  try {
    const resp = await api.updateConfigStatus(config.id, config.status)
    ElMessage.success(resp.message)
  } catch (error) {
    ElMessage.error(error.message)
  }
}

async function initMQTT() {
  if (mqttClient.value && mqttClient.value.connected) {
    mqttClient.value.end()
  }
  const cid = configs[0].name
  const r = Math.random()
  const options = {
    keepalive: 60,
    clientId: `hvc-${cid}-${r}`,
    clean: true,
  }
  mqttClient.value = mqtt.connect('ws://xs.jldg.com:8083/mqtt', options)
  mqttClient.value.on('connect', function () {
    ElMessage.info('消息服务连接成功')
    mqttClient.value.subscribe(`hvc/infos`, function (err) {
      if (err) {
        ElMessage.error('消息服务订阅失败')
        console.log(err)
      }
    })
    mqttClient.value.on('message', function (topic, message) {
      const data = JSON.parse(message.toString())
      // 最多100条
      if (infos.value[data.name] && infos.value[data.name].length > 100) {
        infos.value[data.name].shift()
      }
      infos.value[data.name] = [data,...(infos.value[data.name] || [])]
    })
    mqttClient.value.on('error', function (error) {
      ElMessage.error("消息服务发生错误", error)
    })
  })
}
async function startServer() {
  try {
    const resp = await api.ping()
    if (resp.code == 0) {
      ElMessage.success('服务器已启动');
    } else {
      ElMessage.error('服务器启动失败');
    }

  } catch (error) {
    const cmd = Command.sidecar("sidecar/hvc");
    sidecar.value = await cmd.spawn()
  }
}


</script>

<style>
body {
  margin: 0 !important;
  padding: 0 !important;
}

.el-header {
  display: flex;
  text-align: center;
  justify-content: center;
  align-items: center;
  background-color: #73b8fd;
}

.row-col {
  padding: 10px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 30px;
}
</style>