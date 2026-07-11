<template>
  <div class="alarm-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>报警管理</span>
          <el-button type="primary" @click="fetchAlarms">刷新</el-button>
        </div>
      </template>
      <el-row :gutter="20" class="stat-row">
        <el-col :span="6"><el-statistic title="总报警" :value="statistics.total" /></el-col>
        <el-col :span="6"><el-statistic title="未处理" :value="statistics.unhandled"><template #prefix><el-tag type="danger">!</el-tag></template></el-statistic></el-col>
        <el-col :span="6"><el-statistic title="已处理" :value="statistics.handled" /></el-col>
        <el-col :span="6"><el-statistic title="已忽略" :value="statistics.ignored" /></el-col>
      </el-row>
      <el-form :inline="true" :model="searchForm" class="demo-form-inline">
        <el-form-item label="设备ID">
          <el-input v-model="searchForm.device_id" placeholder="请输入设备ID" clearable />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.handle_status" placeholder="请选择" clearable>
            <el-option label="未处理" :value="0" />
            <el-option label="已处理" :value="1" />
            <el-option label="已忽略" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
      <el-table :data="alarmList" style="width: 100%" v-loading="loading">
        <el-table-column prop="device_id" label="设备ID" width="180" />
        <el-table-column prop="alarm_type" label="报警类型" width="120" />
        <el-table-column prop="alarm_priority" label="优先级" width="70">
          <template #default="scope">
            <el-tag :type="scope.row.alarm_priority > 3 ? 'danger' : scope.row.alarm_priority > 1 ? 'warning' : 'info'">
              {{ scope.row.alarm_priority }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="alarm_time" label="报警时间" width="170" />
        <el-table-column prop="description" label="描述" min-width="200" show-overflow-tooltip />
        <el-table-column label="处理状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.handle_status === 0 ? 'danger' : scope.row.handle_status === 1 ? 'success' : 'info'">
              {{ scope.row.handle_status === 0 ? '未处理' : scope.row.handle_status === 1 ? '已处理' : '已忽略' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="scope">
            <el-button size="small" type="primary" :disabled="scope.row.handle_status !== 0" @click="handleProcess(scope.row)">处理</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
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
    <el-dialog v-model="dialogVisible" title="处理报警" width="450px">
      <el-form :model="alarmForm" label-width="100px">
        <el-form-item label="处理结果">
          <el-select v-model="alarmForm.handle_status" placeholder="请选择">
            <el-option label="已处理" :value="1" />
            <el-option label="已忽略" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="处理说明">
          <el-input v-model="alarmForm.handle_result" type="textarea" :rows="3" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="handleSubmit">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getAlarmList, getAlarmStatistics, handleAlarm, deleteAlarm } from '@/api/alarm'

const loading = ref(false)
const alarmList = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const statistics = ref({ total: 0, unhandled: 0, handled: 0, ignored: 0 })

const searchForm = ref({ device_id: '', handle_status: null })
const dialogVisible = ref(false)
const alarmForm = ref({ handle_status: 1, handle_result: '' })
const currentAlarmId = ref('')

const fetchStatistics = async () => {
  try {
    const res = await getAlarmStatistics({})
    statistics.value = res.data
  } catch (error) { console.error('获取统计数据失败:', error) }
}

const fetchAlarms = async () => {
  loading.value = true
  try {
    const params = { page: currentPage.value, page_size: pageSize.value, ...searchForm.value }
    if (searchForm.value.handle_status === null || searchForm.value.handle_status === '') {
      delete params.handle_status
    }
    const res = await getAlarmList(params)
    alarmList.value = res.data.list
    total.value = res.data.total
  } catch (error) { console.error('获取报警列表失败:', error) }
  finally { loading.value = false }
}

const handleSearch = () => { currentPage.value = 1; fetchAlarms() }
const handleReset = () => { searchForm.value = { device_id: '', handle_status: null }; handleSearch() }
const handleSizeChange = (val) => { pageSize.value = val; fetchAlarms() }
const handleCurrentChange = (val) => { currentPage.value = val; fetchAlarms() }

const handleProcess = (row) => {
  currentAlarmId.value = row.id
  alarmForm.value = { handle_status: 1, handle_result: '' }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    await handleAlarm(currentAlarmId.value, alarmForm.value)
    ElMessage.success('处理成功')
    dialogVisible.value = false
    fetchAlarms()
    fetchStatistics()
  } catch (error) { console.error('处理失败:', error) }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该报警吗？', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    await deleteAlarm(row.id)
    ElMessage.success('删除成功')
    fetchAlarms()
    fetchStatistics()
  } catch (error) { if (error !== 'cancel') console.error('删除失败:', error) }
}

onMounted(() => { fetchAlarms(); fetchStatistics() })
</script>

<style scoped>
.alarm-container { padding: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
.demo-form-inline { margin-bottom: 20px; }
.stat-row { margin-bottom: 20px; }
</style>
