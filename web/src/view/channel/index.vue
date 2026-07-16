<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="通道编码">
          <el-input v-model="searchForm.channel_id" placeholder="请输入通道编码" clearable />
        </el-form-item>
        <el-form-item label="通道名称">
          <el-input v-model="searchForm.name" placeholder="请输入通道名称" clearable />
        </el-form-item>
        <el-form-item label="在线状态">
          <el-select v-model="searchForm.on_line" placeholder="请选择" clearable>
            <el-option label="在线" value="true" />
            <el-option label="离线" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="handleSearch">查询</el-button>
          <el-button icon="refresh" @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <span style="font-size:16px;font-weight:600">通道管理 - {{ deviceName }}</span>
      </div>
      <el-table :data="channelList" style="width: 100%" v-loading="loading">
        <el-table-column prop="channel_id" label="通道ID" width="180" />
        <el-table-column prop="name" label="通道名称" width="180" />
        <el-table-column prop="manufacturer" label="厂商" width="120" />
        <el-table-column prop="model" label="型号" width="120" />
        <el-table-column prop="civil_code" label="行政区域" width="120" />
        <el-table-column prop="address" label="地址" width="180" />
        <el-table-column label="在线状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 'ON' ? 'success' : 'danger'" size="small">
              {{ scope.row.status === 'ON' ? '在线' : '离线' }}
            </el-tag>
          </template>
        </el-table-column>
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

      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="fetchChannels"
          @current-change="fetchChannels"
        />
      </div>
    </div>

    <el-dialog v-model="playDialogVisible" title="视频点播" width="1000px" destroy-on-close @closed="handleDialogClosed">
      <JessibucaPlayer ref="playerRef" :urls="playUrls" />
      <template #footer>
        <el-button @click="handleStop">停止播放</el-button>
        <el-button @click="playDialogVisible = false">关闭</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { getChannelList, playVideo, stopVideo } from '@/api/channel'
import JessibucaPlayer from '@/components/JessibucaPlayer/index.vue'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const channelList = ref([])
const deviceName = ref('')
const deviceId = ref('')
const total = ref(0)
const page = ref(1)
const pageSize = ref(20)

const searchForm = reactive({
  channel_id: '',
  name: '',
  on_line: ''
})

const playDialogVisible = ref(false)
const playUrls = ref({})
const currentDeviceId = ref('')
const playerRef = ref(null)

const fetchChannels = async () => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      page_size: pageSize.value,
      device_id: deviceId.value
    }
    if (searchForm.channel_id) params.channel_id = searchForm.channel_id
    if (searchForm.name) params.name = searchForm.name
    if (searchForm.on_line) params.on_line = searchForm.on_line

    const res = await getChannelList(params)
    channelList.value = res.data.list || []
    total.value = res.data.total || 0
  } catch (error) {
    console.error('获取通道列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  page.value = 1
  fetchChannels()
}

const handleReset = () => {
  searchForm.channel_id = ''
  searchForm.name = ''
  searchForm.on_line = ''
  page.value = 1
  fetchChannels()
}

const handlePlay = async (row) => {
  try {
    const res = await playVideo({
      device_id: deviceId.value,
      channel_id: row.channel_id
    })
    playUrls.value = res.data || {}
    currentDeviceId.value = deviceId.value
    playDialogVisible.value = true
    ElMessage.success('点播请求已发送')
  } catch (error) {
    console.error('点播失败:', error)
  }
}

const handleStop = async () => {
  try {
    if (playerRef.value) {
      playerRef.value.stop()
    }
    await stopVideo(currentDeviceId.value)
    playDialogVisible.value = false
    playUrls.value = {}
    ElMessage.success('已停止播放')
  } catch (error) {
    console.error('停止播放失败:', error)
  }
}

const handleDialogClosed = () => {
  if (playerRef.value) {
    playerRef.value.stop()
  }
  playUrls.value = {}
}

onMounted(() => {
  deviceId.value = route.query.device_id || ''
  deviceName.value = route.query.device_name || ''
  if (deviceId.value) {
    fetchChannels()
  }
})
</script>
