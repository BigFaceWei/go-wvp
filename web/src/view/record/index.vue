<template>
  <div class="record-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>录像管理</span>
          <el-button type="primary" @click="handleAdd">添加录像计划</el-button>
        </div>
      </template>
      <el-tabs v-model="activeTab">
        <el-tab-pane label="录像计划" name="plan">
          <el-table :data="planList" style="width: 100%" v-loading="loading">
            <el-table-column prop="name" label="计划名称" width="160" />
            <el-table-column prop="device_id" label="设备ID" width="180" />
            <el-table-column prop="channel_id" label="通道ID" width="180" />
            <el-table-column prop="record_type" label="录像类型" width="100">
              <template #default="scope">
                <span>{{ scope.row.record_type === 1 ? '连续录像' : scope.row.record_type === 2 ? '定时录像' : '告警录像' }}</span>
              </template>
            </el-table-column>
            <el-table-column label="状态" width="80">
              <template #default="scope">
                <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
                  {{ scope.row.status === 1 ? '录像中' : '停止' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="file_days" label="存储天数" width="90" />
            <el-table-column label="操作" width="320">
              <template #default="scope">
                <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
                <el-button size="small" :type="scope.row.status === 1 ? 'warning' : 'success'" @click="handleToggle(scope.row)">
                  {{ scope.row.status === 1 ? '停止' : '启动' }}
                </el-button>
                <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
            class="mt-4"
            v-model:current-page="planPage"
            v-model:page-size="planPageSize"
            :page-sizes="[10, 20, 50]"
            :total="planTotal"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="fetchPlans"
            @current-change="fetchPlans"
          />
        </el-tab-pane>
        <el-tab-pane label="录像文件" name="file">
          <el-form :inline="true" :model="fileSearchForm" class="demo-form-inline">
            <el-form-item label="设备ID">
              <el-input v-model="fileSearchForm.device_id" placeholder="请输入设备ID" clearable />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="handleFileSearch">搜索</el-button>
              <el-button @click="handleFileReset">重置</el-button>
            </el-form-item>
          </el-form>
          <el-table :data="fileList" style="width: 100%" v-loading="fileLoading">
            <el-table-column prop="device_id" label="设备ID" width="180" />
            <el-table-column prop="file_name" label="文件名" min-width="200" show-overflow-tooltip />
            <el-table-column prop="file_size" label="文件大小" width="100">
              <template #default="scope">
                {{ (scope.row.file_size / 1024 / 1024).toFixed(2) }} MB
              </template>
            </el-table-column>
            <el-table-column prop="start_time" label="开始时间" width="170" />
            <el-table-column prop="end_time" label="结束时间" width="170" />
            <el-table-column prop="duration" label="时长(秒)" width="90" />
            <el-table-column prop="storage_type" label="存储类型" width="80">
              <template #default="scope">
                <span>{{ scope.row.storage_type === 1 ? '本地' : '云存储' }}</span>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
            class="mt-4"
            v-model:current-page="filePage"
            v-model:page-size="filePageSize"
            :page-sizes="[10, 20, 50]"
            :total="fileTotal"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="fetchFiles"
            @current-change="fetchFiles"
          />
        </el-tab-pane>
      </el-tabs>
    </el-card>
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="planForm" label-width="100px">
        <el-form-item label="计划名称">
          <el-input v-model="planForm.name" />
        </el-form-item>
        <el-form-item label="设备ID">
          <el-input v-model="planForm.device_id" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="通道ID">
          <el-input v-model="planForm.channel_id" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="录像类型">
          <el-select v-model="planForm.record_type" placeholder="请选择">
            <el-option label="连续录像" :value="1" />
            <el-option label="定时录像" :value="2" />
            <el-option label="告警录像" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="存储天数">
          <el-input-number v-model="planForm.file_days" :min="1" :max="365" />
        </el-form-item>
        <el-form-item label="存储路径">
          <el-input v-model="planForm.storage_path" placeholder="默认: /data/record" />
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
import { getRecordPlanList, createRecordPlan, updateRecordPlan, deleteRecordPlan, startRecordPlan, stopRecordPlan, getRecordFileList } from '@/api/record'

const activeTab = ref('plan')

const loading = ref(false)
const planList = ref([])
const planTotal = ref(0)
const planPage = ref(1)
const planPageSize = ref(10)

const fileLoading = ref(false)
const fileList = ref([])
const fileTotal = ref(0)
const filePage = ref(1)
const filePageSize = ref(10)

const fileSearchForm = ref({ device_id: '' })

const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)

const planForm = ref({
  name: '',
  device_id: '',
  channel_id: '',
  record_type: 1,
  file_days: 30,
  storage_path: ''
})

const fetchPlans = async () => {
  loading.value = true
  try {
    const params = { page: planPage.value, page_size: planPageSize.value }
    const res = await getRecordPlanList(params)
    planList.value = res.data.list
    planTotal.value = res.data.total
  } catch (error) { console.error('获取录像计划失败:', error) }
  finally { loading.value = false }
}

const fetchFiles = async () => {
  fileLoading.value = true
  try {
    const params = { page: filePage.value, page_size: filePageSize.value }
    if (fileSearchForm.value.device_id) params.device_id = fileSearchForm.value.device_id
    const res = await getRecordFileList(params)
    fileList.value = res.data.list
    fileTotal.value = res.data.total
  } catch (error) { console.error('获取录像文件失败:', error) }
  finally { fileLoading.value = false }
}

const handleFileSearch = () => { filePage.value = 1; fetchFiles() }
const handleFileReset = () => { fileSearchForm.value = { device_id: '' }; handleFileSearch() }

const handleAdd = () => {
  dialogTitle.value = '添加录像计划'
  isEdit.value = false
  planForm.value = { name: '', device_id: '', channel_id: '', record_type: 1, file_days: 30, storage_path: '' }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑录像计划'
  isEdit.value = true
  planForm.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateRecordPlan(planForm.value.id, planForm.value)
      ElMessage.success('更新成功')
    } else {
      await createRecordPlan(planForm.value)
      ElMessage.success('添加成功')
    }
    dialogVisible.value = false
    fetchPlans()
  } catch (error) { console.error('操作失败:', error) }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该录像计划吗？', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    await deleteRecordPlan(row.id)
    ElMessage.success('删除成功')
    fetchPlans()
  } catch (error) { if (error !== 'cancel') console.error('删除失败:', error) }
}

const handleToggle = async (row) => {
  try {
    if (row.status === 1) {
      await stopRecordPlan(row.id)
      ElMessage.success('录像已停止')
    } else {
      await startRecordPlan(row.id)
      ElMessage.success('录像已启动')
    }
    fetchPlans()
  } catch (error) { console.error('操作失败:', error) }
}

onMounted(() => { fetchPlans(); fetchFiles() })
</script>

<style scoped>
.record-container { padding: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
.demo-form-inline { margin-bottom: 20px; }
</style>
