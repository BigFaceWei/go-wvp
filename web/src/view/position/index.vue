<template>
  <div class="position-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>移动位置</span>
        </div>
      </template>
      <el-form :inline="true" :model="searchForm" class="demo-form-inline">
        <el-form-item label="设备ID">
          <el-input v-model="searchForm.device_id" placeholder="请输入设备ID" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
      <el-table :data="positionList" style="width: 100%" v-loading="loading">
        <el-table-column prop="device_id" label="设备ID" width="180" />
        <el-table-column prop="channel_id" label="通道ID" width="180" />
        <el-table-column prop="longitude" label="经度" width="120" />
        <el-table-column prop="latitude" label="纬度" width="120" />
        <el-table-column prop="altitude" label="海拔(m)" width="100" />
        <el-table-column prop="speed" label="速度(km/h)" width="110" />
        <el-table-column prop="direction" label="方向(°)" width="90" />
        <el-table-column prop="gps_status" label="GPS状态" width="90">
          <template #default="scope">
            <el-tag :type="scope.row.gps_status === 1 ? 'success' : 'danger'">
              {{ scope.row.gps_status === 1 ? '有效' : '无效' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="report_time" label="上报时间" width="170" />
        <el-table-column label="操作" width="120">
          <template #default="scope">
            <el-button size="small" @click="viewTrack(scope.row)">轨迹</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        class="mt-4"
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </el-card>
    <el-dialog v-model="trackVisible" title="位置轨迹" width="700px">
      <el-table :data="trackList" style="width: 100%" max-height="400">
        <el-table-column prop="longitude" label="经度" width="120" />
        <el-table-column prop="latitude" label="纬度" width="120" />
        <el-table-column prop="altitude" label="海拔(m)" width="100" />
        <el-table-column prop="speed" label="速度(km/h)" width="110" />
        <el-table-column prop="report_time" label="上报时间" width="180" />
      </el-table>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getPositionList, getPositionTrack } from '@/api/position'

const loading = ref(false)
const positionList = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)

const searchForm = ref({ device_id: '' })

const trackVisible = ref(false)
const trackList = ref([])

const fetchPositions = async () => {
  loading.value = true
  try {
    const params = { page: currentPage.value, page_size: pageSize.value }
    if (searchForm.value.device_id) params.device_id = searchForm.value.device_id
    const res = await getPositionList(params)
    positionList.value = res.data.list
    total.value = res.data.total
  } catch (error) { console.error('获取位置信息失败:', error) }
  finally { loading.value = false }
}

const handleSearch = () => { currentPage.value = 1; fetchPositions() }
const handleReset = () => { searchForm.value = { device_id: '' }; handleSearch() }
const handleSizeChange = (val) => { pageSize.value = val; fetchPositions() }
const handleCurrentChange = (val) => { currentPage.value = val; fetchPositions() }

const viewTrack = async (row) => {
  try {
    const res = await getPositionTrack(row.device_id)
    trackList.value = res.data
    trackVisible.value = true
  } catch (error) { console.error('获取轨迹失败:', error) }
}

onMounted(() => { fetchPositions() })
</script>

<style scoped>
.position-container { padding: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
.demo-form-inline { margin-bottom: 20px; }
</style>
