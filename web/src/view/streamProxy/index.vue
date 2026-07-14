<template>
  <div class="stream-proxy-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>拉流代理管理</span>
          <el-button type="primary" @click="handleAdd">添加代理</el-button>
        </div>
      </template>
      <el-form :inline="true" :model="searchForm" class="demo-form-inline">
        <el-form-item label="名称">
          <el-input v-model="searchForm.name" placeholder="请输入名称" clearable />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
      <el-table :data="proxyList" style="width: 100%" v-loading="loading">
        <el-table-column prop="name" label="名称" min-width="140" />
        <el-table-column prop="stream_type" label="拉流类型" width="100" />
        <el-table-column prop="url" label="源URL" min-width="200" show-overflow-tooltip />
        <el-table-column prop="app" label="应用名" width="80" />
        <el-table-column prop="stream" label="流ID" min-width="140" />
        <el-table-column label="状态" width="80">
          <template #default="scope">
            <el-tag :type="scope.row.status ? 'success' : 'info'">
              {{ scope.row.status ? '运行中' : '停止' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="在线" width="70">
          <template #default="scope">
            <el-tag :type="scope.row.online ? 'success' : 'danger'">
              {{ scope.row.online ? '是' : '否' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="small" :type="scope.row.status ? 'warning' : 'success'" @click="handleToggle(scope.row)">
              {{ scope.row.status ? '停止' : '启动' }}
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
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form :model="proxyForm" label-width="120px">
        <el-form-item label="名称">
          <el-input v-model="proxyForm.name" />
        </el-form-item>
        <el-form-item label="拉流类型">
          <el-select v-model="proxyForm.stream_type" placeholder="请选择">
            <el-option label="RTSP" value="rtsp" />
            <el-option label="RTMP" value="rtmp" />
            <el-option label="HTTP-FLV" value="http-flv" />
          </el-select>
        </el-form-item>
        <el-form-item label="源URL">
          <el-input v-model="proxyForm.url" />
        </el-form-item>
        <el-form-item label="应用名">
          <el-input v-model="proxyForm.app" />
        </el-form-item>
        <el-form-item label="流ID">
          <el-input v-model="proxyForm.stream" />
        </el-form-item>
        <el-form-item label="超时(秒)">
          <el-input-number v-model="proxyForm.timeout_sec" :min="5" :max="300" />
        </el-form-item>
        <el-form-item label="启用">
          <el-switch v-model="proxyForm.enable" />
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
import { getStreamProxyList, createStreamProxy, updateStreamProxy, deleteStreamProxy, startStreamProxy, stopStreamProxy } from '@/api/streamProxy'

const loading = ref(false)
const proxyList = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)

const searchForm = ref({ name: '' })

const proxyForm = ref({
  name: '',
  stream_type: 'rtsp',
  url: '',
  app: 'proxy',
  stream: '',
  timeout_sec: 30,
  enable: true
})

const fetchProxies = async () => {
  loading.value = true
  try {
    const params = { page: currentPage.value, page_size: pageSize.value, ...searchForm.value }
    const res = await getStreamProxyList(params)
    proxyList.value = res.data.list
    total.value = res.data.total
  } catch (error) {
    console.error('获取代理列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { currentPage.value = 1; fetchProxies() }
const handleReset = () => { searchForm.value = { name: '' }; handleSearch() }
const handleSizeChange = (val) => { pageSize.value = val; fetchProxies() }
const handleCurrentChange = (val) => { currentPage.value = val; fetchProxies() }

const handleAdd = () => {
  dialogTitle.value = '添加代理'
  isEdit.value = false
  proxyForm.value = { name: '', stream_type: 'rtsp', url: '', app: 'proxy', stream: '', timeout_sec: 30, enable: true }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑代理'
  isEdit.value = true
  proxyForm.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateStreamProxy(proxyForm.value.ID, proxyForm.value)
      ElMessage.success('更新成功')
    } else {
      await createStreamProxy(proxyForm.value)
      ElMessage.success('添加成功')
    }
    dialogVisible.value = false
    fetchProxies()
  } catch (error) { console.error('操作失败:', error) }
}

const handleToggle = async (row) => {
  try {
    if (row.status) {
      await stopStreamProxy(row.ID)
      ElMessage.success('代理已停止')
    } else {
      await startStreamProxy(row.ID)
      ElMessage.success('代理已启动')
    }
    fetchProxies()
  } catch (error) { console.error('操作失败:', error) }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该代理吗？', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    await deleteStreamProxy(row.ID)
    ElMessage.success('删除成功')
    fetchProxies()
  } catch (error) { if (error !== 'cancel') console.error('删除失败:', error) }
}

onMounted(() => { fetchProxies() })
</script>

<style scoped>
.stream-proxy-container { padding: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
.demo-form-inline { margin-bottom: 20px; }
</style>
