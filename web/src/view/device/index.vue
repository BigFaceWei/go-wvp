<template>
  <div class="device-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>设备管理</span>
          <el-button type="primary" @click="handleAdd">添加设备</el-button>
        </div>
      </template>
      
      <!-- 搜索表单 -->
      <el-form :inline="true" :model="searchForm" class="demo-form-inline">
        <el-form-item label="设备ID">
          <el-input v-model="searchForm.device_id" placeholder="请输入设备ID" clearable />
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="searchForm.name" placeholder="请输入设备名称" clearable />
        </el-form-item>
        <el-form-item label="在线状态">
          <el-select v-model="searchForm.online" placeholder="请选择" clearable>
            <el-option label="在线" value="true" />
            <el-option label="离线" value="false" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
      
      <!-- 设备列表 -->
      <el-table :data="deviceList" style="width: 100%" v-loading="loading">
        <el-table-column prop="device_id" label="设备ID" width="180" />
        <el-table-column prop="name" label="设备名称" width="180" />
        <el-table-column prop="manufacturer" label="厂商" width="120" />
        <el-table-column prop="model" label="型号" width="120" />
        <el-table-column label="在线状态" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.on_line ? 'success' : 'danger'">
              {{ scope.row.on_line ? '在线' : '离线' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="ip" label="IP地址" width="140" />
        <el-table-column label="操作" width="350">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="small" type="success" @click="handleChannels(scope.row)">通道</el-button>
            <el-button size="small" type="primary" @click="handleCatalog(scope.row)">查询目录</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <!-- 分页 -->
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
    
    <!-- 添加/编辑对话框 -->
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="deviceForm" label-width="100px">
        <el-form-item label="设备ID">
          <el-input v-model="deviceForm.device_id" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="设备名称">
          <el-input v-model="deviceForm.name" />
        </el-form-item>
        <el-form-item label="厂商">
          <el-input v-model="deviceForm.manufacturer" />
        </el-form-item>
        <el-form-item label="型号">
          <el-input v-model="deviceForm.model" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="deviceForm.password" type="password" />
        </el-form-item>
        <el-form-item label="传输协议">
          <el-select v-model="deviceForm.transport" placeholder="请选择">
            <el-option label="UDP" value="UDP" />
            <el-option label="TCP" value="TCP" />
          </el-select>
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
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getDeviceList, createDevice, updateDevice, deleteDevice, queryDeviceCatalog } from '@/api/device'

const router = useRouter()

const loading = ref(false)
const deviceList = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)

const searchForm = ref({
  device_id: '',
  name: '',
  online: ''
})

const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)

const deviceForm = ref({
  device_id: '',
  name: '',
  manufacturer: '',
  model: '',
  password: '',
  transport: 'UDP'
})

const fetchDevices = async () => {
  loading.value = true
  try {
    const params = {
      page: currentPage.value,
      page_size: pageSize.value,
      ...searchForm.value
    }
    const res = await getDeviceList(params)
    deviceList.value = res.data.list
    total.value = res.data.total
  } catch (error) {
    console.error('获取设备列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  currentPage.value = 1
  fetchDevices()
}

const handleReset = () => {
  searchForm.value = {
    device_id: '',
    name: '',
    online: ''
  }
  handleSearch()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  fetchDevices()
}

const handleCurrentChange = (val) => {
  currentPage.value = val
  fetchDevices()
}

const handleAdd = () => {
  dialogTitle.value = '添加设备'
  isEdit.value = false
  deviceForm.value = {
    device_id: '',
    name: '',
    manufacturer: '',
    model: '',
    password: '',
    transport: 'UDP'
  }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑设备'
  isEdit.value = true
  deviceForm.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateDevice(deviceForm.value.device_id, deviceForm.value)
      ElMessage.success('更新成功')
    } else {
      await createDevice(deviceForm.value)
      ElMessage.success('添加成功')
    }
    dialogVisible.value = false
    fetchDevices()
  } catch (error) {
    console.error('操作失败:', error)
  }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该设备吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await deleteDevice(row.device_id)
    ElMessage.success('删除成功')
    fetchDevices()
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error)
    }
  }
}

const handleCatalog = async (row) => {
  try {
    await queryDeviceCatalog(row.device_id)
    ElMessage.success('目录查询请求已发送')
  } catch (error) {
    console.error('查询目录失败:', error)
  }
}

const handleChannels = (row) => {
  router.push({
    path: '/channel',
    query: {
      device_id: row.device_id,
      device_name: row.name
    }
  })
}

onMounted(() => {
  fetchDevices()
})
</script>

<style scoped>
.device-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.demo-form-inline {
  margin-bottom: 20px;
}
</style>