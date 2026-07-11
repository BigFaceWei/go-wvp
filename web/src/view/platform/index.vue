<template>
  <div class="platform-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>平台级联管理</span>
          <el-button type="primary" @click="handleAdd">添加平台</el-button>
        </div>
      </template>
      <el-form :inline="true" :model="searchForm" class="demo-form-inline">
        <el-form-item label="平台ID">
          <el-input v-model="searchForm.platform_id" placeholder="请输入平台ID" clearable />
        </el-form-item>
        <el-form-item label="平台名称">
          <el-input v-model="searchForm.name" placeholder="请输入平台名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
      <el-table :data="platformList" style="width: 100%" v-loading="loading">
        <el-table-column prop="platform_id" label="平台ID" width="180" />
        <el-table-column prop="name" label="平台名称" width="160" />
        <el-table-column prop="ip" label="IP地址" width="140" />
        <el-table-column prop="port" label="端口" width="70" />
        <el-table-column label="状态" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '在线' : '离线' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="channel_count" label="同步通道数" width="100" />
        <el-table-column label="操作" width="360">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="small" type="primary" @click="handleSync(scope.row)">同步目录</el-button>
            <el-button size="small" :type="scope.row.status === 1 ? 'warning' : 'success'" @click="handleCascade(scope.row)">
              {{ scope.row.status === 1 ? '停止级联' : '启动级联' }}
            </el-button>
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="platformForm" label-width="100px">
        <el-form-item label="平台ID">
          <el-input v-model="platformForm.platform_id" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="平台名称">
          <el-input v-model="platformForm.name" />
        </el-form-item>
        <el-form-item label="密钥">
          <el-input v-model="platformForm.secret" type="password" />
        </el-form-item>
        <el-form-item label="IP地址">
          <el-input v-model="platformForm.ip" />
        </el-form-item>
        <el-form-item label="端口">
          <el-input-number v-model="platformForm.port" :min="1" :max="65535" />
        </el-form-item>
        <el-form-item label="传输协议">
          <el-select v-model="platformForm.transport" placeholder="请选择">
            <el-option label="UDP" value="UDP" />
            <el-option label="TCP" value="TCP" />
          </el-select>
        </el-form-item>
        <el-form-item label="域">
          <el-input v-model="platformForm.domain" />
        </el-form-item>
        <el-form-item label="过期时间">
          <el-input-number v-model="platformForm.expires" :min="60" :max="86400" />
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
import { getPlatformList, createPlatform, updatePlatform, deletePlatform, syncPlatformCatalog, startPlatformCascade, stopPlatformCascade } from '@/api/platform'

const loading = ref(false)
const platformList = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)

const searchForm = ref({
  platform_id: '',
  name: ''
})

const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)

const platformForm = ref({
  platform_id: '',
  name: '',
  secret: '',
  ip: '',
  port: 5060,
  transport: 'UDP',
  domain: '',
  expires: 3600
})

const fetchPlatforms = async () => {
  loading.value = true
  try {
    const params = { page: currentPage.value, page_size: pageSize.value, ...searchForm.value }
    const res = await getPlatformList(params)
    platformList.value = res.data.list
    total.value = res.data.total
  } catch (error) {
    console.error('获取平台列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { currentPage.value = 1; fetchPlatforms() }
const handleReset = () => { searchForm.value = { platform_id: '', name: '' }; handleSearch() }
const handleSizeChange = (val) => { pageSize.value = val; fetchPlatforms() }
const handleCurrentChange = (val) => { currentPage.value = val; fetchPlatforms() }

const handleAdd = () => {
  dialogTitle.value = '添加平台'
  isEdit.value = false
  platformForm.value = { platform_id: '', name: '', secret: '', ip: '', port: 5060, transport: 'UDP', domain: '', expires: 3600 }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑平台'
  isEdit.value = true
  platformForm.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updatePlatform(platformForm.value.platform_id, platformForm.value)
      ElMessage.success('更新成功')
    } else {
      await createPlatform(platformForm.value)
      ElMessage.success('添加成功')
    }
    dialogVisible.value = false
    fetchPlatforms()
  } catch (error) { console.error('操作失败:', error) }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该平台吗？', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    await deletePlatform(row.platform_id)
    ElMessage.success('删除成功')
    fetchPlatforms()
  } catch (error) { if (error !== 'cancel') console.error('删除失败:', error) }
}

const handleSync = async (row) => {
  try {
    await syncPlatformCatalog(row.platform_id)
    ElMessage.success('目录同步请求已发送')
  } catch (error) { console.error('同步目录失败:', error) }
}

const handleCascade = async (row) => {
  try {
    if (row.status === 1) {
      await stopPlatformCascade(row.platform_id)
      ElMessage.success('级联已停止')
    } else {
      await startPlatformCascade(row.platform_id)
      ElMessage.success('级联已启动')
    }
    fetchPlatforms()
  } catch (error) { console.error('操作失败:', error) }
}

onMounted(() => { fetchPlatforms() })
</script>

<style scoped>
.platform-container { padding: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
.demo-form-inline { margin-bottom: 20px; }
</style>
