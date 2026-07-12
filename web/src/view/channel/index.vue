<template>
  <div class="channel-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>通道管理 - {{ deviceName }}</span>
          <el-button @click="goBack">返回设备列表</el-button>
        </div>
      </template>

      <el-table :data="channelList" style="width: 100%" v-loading="loading">
        <el-table-column prop="channel_id" label="通道ID" width="180" />
        <el-table-column prop="name" label="通道名称" width="180" />
        <el-table-column prop="manufacturer" label="厂商" width="120" />
        <el-table-column prop="model" label="型号" width="120" />
        <el-table-column prop="civil_code" label="行政区域" width="120" />
        <el-table-column prop="address" label="地址" width="180" />
        <el-table-column label="音频" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.has_audio ? 'success' : 'info'" size="small">
              {{ scope.row.has_audio ? '有' : '无' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150">
          <template #default="scope">
            <el-button size="small" type="primary" @click="handlePlay(scope.row)">点播</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- 播放对话框 -->
    <el-dialog v-model="playDialogVisible" title="视频点播" width="800px" destroy-on-close>
      <div class="video-player">
        <video ref="videoPlayer" controls autoplay class="video-element">
          <source :src="playUrl" type="application/x-mpegURL">
          您的浏览器不支持视频播放
        </video>
      </div>
      <template #footer>
        <el-button @click="handleStop">停止播放</el-button>
        <el-button @click="playDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { getDeviceChannels, playVideo, stopVideo } from '@/api/channel'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const channelList = ref([])
const deviceName = ref('')
const deviceId = ref('')

const playDialogVisible = ref(false)
const playUrl = ref('')
const currentDeviceId = ref('')
const currentChannelId = ref('')
const videoPlayer = ref(null)

const fetchChannels = async () => {
  loading.value = true
  try {
    const res = await getDeviceChannels(deviceId.value)
    channelList.value = res.data || []
  } catch (error) {
    console.error('获取通道列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handlePlay = async (row) => {
  try {
    const res = await playVideo({
      device_id: deviceId.value,
      channel_id: row.channel_id
    })
    playUrl.value = res.data.play_url
    currentDeviceId.value = deviceId.value
    currentChannelId.value = row.channel_id
    playDialogVisible.value = true
    ElMessage.success('点播请求已发送')
  } catch (error) {
    console.error('点播失败:', error)
  }
}

const handleStop = async () => {
  try {
    await stopVideo(currentDeviceId.value)
    playDialogVisible.value = false
    playUrl.value = ''
    ElMessage.success('已停止播放')
  } catch (error) {
    console.error('停止播放失败:', error)
  }
}

const goBack = () => {
  router.back()
}

onMounted(() => {
  deviceId.value = route.query.device_id || ''
  deviceName.value = route.query.device_name || '未知设备'
  if (deviceId.value) {
    fetchChannels()
  }
})
</script>

<style scoped>
.channel-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.video-player {
  display: flex;
  justify-content: center;
  background: #000;
  border-radius: 4px;
  padding: 10px;
}

.video-element {
  width: 100%;
  max-height: 500px;
}
</style>
