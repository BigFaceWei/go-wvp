<template>
  <div class="group-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>业务分组管理</span>
          <el-button type="primary" @click="handleAdd">添加分组</el-button>
        </div>
      </template>
      <el-table :data="groupList" style="width: 100%" v-loading="loading" row-key="ID">
        <el-table-column prop="ID" label="ID" width="70" />
        <el-table-column prop="name" label="分组名称" min-width="160" />
        <el-table-column prop="parent_id" label="父级ID" width="120" />
        <el-table-column prop="device_id" label="关联设备" width="120" />
        <el-table-column prop="region_id" label="关联区域" width="120" />
        <el-table-column prop="sort" label="排序" width="70" />
        <el-table-column prop="remark" label="备注" min-width="160" show-overflow-tooltip />
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="500px">
      <el-form :model="groupForm" label-width="100px">
        <el-form-item label="分组名称">
          <el-input v-model="groupForm.name" />
        </el-form-item>
        <el-form-item label="父级ID">
          <el-input v-model="groupForm.parent_id" placeholder="留空为顶级" />
        </el-form-item>
        <el-form-item label="关联设备">
          <el-input v-model="groupForm.device_id" />
        </el-form-item>
        <el-form-item label="关联区域">
          <el-input v-model="groupForm.region_id" />
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="groupForm.sort" :min="0" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="groupForm.remark" type="textarea" />
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
import { getGroupList, createGroup, updateGroup, deleteGroup } from '@/api/group'

const loading = ref(false)
const groupList = ref([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)

const groupForm = ref({
  name: '',
  parent_id: '',
  device_id: '',
  region_id: '',
  sort: 0,
  remark: ''
})

const fetchGroups = async () => {
  loading.value = true
  try {
    const res = await getGroupList()
    groupList.value = res.data.list || res.data
  } catch (error) {
    console.error('获取分组列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  dialogTitle.value = '添加分组'
  isEdit.value = false
  groupForm.value = { name: '', parent_id: '', device_id: '', region_id: '', sort: 0, remark: '' }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑分组'
  isEdit.value = true
  groupForm.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateGroup(groupForm.value.ID, groupForm.value)
      ElMessage.success('更新成功')
    } else {
      await createGroup(groupForm.value)
      ElMessage.success('添加成功')
    }
    dialogVisible.value = false
    fetchGroups()
  } catch (error) { console.error('操作失败:', error) }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该分组吗？', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    await deleteGroup(row.ID)
    ElMessage.success('删除成功')
    fetchGroups()
  } catch (error) { if (error !== 'cancel') console.error('删除失败:', error) }
}

onMounted(() => { fetchGroups() })
</script>

<style scoped>
.group-container { padding: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
</style>
