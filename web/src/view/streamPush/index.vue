<template>
  <div class="stream-push-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>推流管理</span>
          <el-button type="primary" @click="handleAdd">添加推流</el-button>
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
      <el-table :data="pushList" style="width: 100%" v-loading="loading">
        <el-table-column prop="name" label="名称" min-width="140" />
        <el-table-column prop="push_type" label="推流类型" width="110" />
        <el-table-column prop="push_url" label="推流URL" min-width="200" show-overflow-tooltip />
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
      <el-form :model="pushForm" label-width="120px">
        <el-form-item label="名称">
          <el-input v-model="pushForm.name" />
        </el-form-item>
        <el-form-item label="推流类型">
          <el-select v-model="pushForm.push_type" placeholder="请选择">
            <el-option label="RTMP" value="rtmp" />
            <el-option label="GB28181" value="gb28181" />
          </el-select>
        </el-form-item>
        <el-form-item label="推流URL">
          <el-input v-model="pushForm.push_url" />
        </el-form-item>
        <el-form-item label="应用名">
          <el-input v-model="pushForm.app" />
        </el-form-item>
        <el-form-item label="流ID">
          <el-input v-model="pushForm.stream" />
        </el-form-item>
        <el-form-item label="启用">
          <el-switch v-model="pushForm.enable" />
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
import { getStreamPushList, createStreamPush, updateStreamPush, deleteStreamPush, startStreamPush, stopStreamPush } from '@/api/streamPush'

const loading = ref(false)
const pushList = ref([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(10)
const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)

const searchForm = ref({ name: '' })

const pushForm = ref({
  name: '',
  push_type: 'rtmp',
  push_url: '',
  app: 'push',
  stream: '',
  enable: true
})

const fetchPushes = async () => {
  loading.value = true
  try {
    const params = { page: currentPage.value, page_size: pageSize.value, ...searchForm.value }
    const res = await getStreamPushList(params)
    pushList.value = res.data.list
    total.value = res.data.total
  } catch (error) {
    console.error('获取推流列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => { currentPage.value = 1; fetchPushes() }
const handleReset = () => { searchForm.value = { name: '' }; handleSearch() }
const handleSizeChange = (val) => { pageSize.value = val; fetchPushes() }
const handleCurrentChange = (val) => { currentPage.value = val; fetchPushes() }

const handleAdd = () => {
  dialogTitle.value = '添加推流'
  isEdit.value = false
  pushForm.value = { name: '', push_type: 'rtmp', push_url: '', app: 'push', stream: '', enable: true }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑推流'
  isEdit.value = true
  pushForm.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateStreamPush(pushForm.value.ID, pushForm.value)
      ElMessage.success('更新成功')
    } else {
      await createStreamPush(pushForm.value)
      ElMessage.success('添加成功')
    }
    dialogVisible.value = false
    fetchPushes()
  } catch (error) { console.error('操作失败:', error) }
}

const handleToggle = async (row) => {
  try {
    if (row.status) {
      await stopStreamPush(row.ID)
      ElMessage.success('推流已停止')
    } else {
      await startStreamPush(row.ID)
      ElMessage.success('推流已启动')
    }
    fetchPushes()
  } catch (error) { console.error('操作失败:', error) }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该推流吗？', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    await deleteStreamPush(row.ID)
    ElMessage.success('删除成功')
    fetchPushes()
  } catch (error) { if (error !== 'cancel') console.error('删除失败:', error) }
}

onMounted(() => { fetchPushes() })
</script>

<style scoped>
.stream-push-container { padding: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
.demo-form-inline { margin-bottom: 20px; }
</style>
