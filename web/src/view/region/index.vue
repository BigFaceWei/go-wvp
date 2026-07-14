<template>
  <div class="region-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>行政区域管理</span>
          <el-button type="primary" @click="handleAdd">添加区域</el-button>
        </div>
      </template>
      <el-table :data="regionList" style="width: 100%" v-loading="loading" row-key="ID">
        <el-table-column prop="code" label="区域编码" width="120" />
        <el-table-column prop="name" label="区域名称" min-width="160" />
        <el-table-column prop="parent_id" label="父级编码" width="120" />
        <el-table-column prop="level" label="级别" width="80">
          <template #default="scope">
            <el-tag :type="['', 'info', 'primary', 'success'][scope.row.level]">
              {{ ['', '省/直辖市', '市', '区县', '街道'][scope.row.level] || '未知' }}
            </el-tag>
          </template>
        </el-table-column>
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
      <el-form :model="regionForm" label-width="100px">
        <el-form-item label="区域编码">
          <el-input v-model="regionForm.code" :disabled="isEdit" />
        </el-form-item>
        <el-form-item label="区域名称">
          <el-input v-model="regionForm.name" />
        </el-form-item>
        <el-form-item label="父级编码">
          <el-input v-model="regionForm.parent_id" placeholder="留空为顶级" />
        </el-form-item>
        <el-form-item label="级别">
          <el-select v-model="regionForm.level" placeholder="请选择">
            <el-option label="省/直辖市" :value="1" />
            <el-option label="市" :value="2" />
            <el-option label="区县" :value="3" />
            <el-option label="街道" :value="4" />
          </el-select>
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="regionForm.sort" :min="0" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="regionForm.remark" type="textarea" />
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
import { getRegionList, createRegion, updateRegion, deleteRegion } from '@/api/region'

const loading = ref(false)
const regionList = ref([])
const dialogVisible = ref(false)
const dialogTitle = ref('')
const isEdit = ref(false)

const regionForm = ref({
  code: '',
  name: '',
  parent_id: '',
  level: 1,
  sort: 0,
  remark: ''
})

const fetchRegions = async () => {
  loading.value = true
  try {
    const res = await getRegionList()
    regionList.value = res.data.list || res.data
  } catch (error) {
    console.error('获取区域列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleAdd = () => {
  dialogTitle.value = '添加区域'
  isEdit.value = false
  regionForm.value = { code: '', name: '', parent_id: '', level: 1, sort: 0, remark: '' }
  dialogVisible.value = true
}

const handleEdit = (row) => {
  dialogTitle.value = '编辑区域'
  isEdit.value = true
  regionForm.value = { ...row }
  dialogVisible.value = true
}

const handleSubmit = async () => {
  try {
    if (isEdit.value) {
      await updateRegion(regionForm.value.code, regionForm.value)
      ElMessage.success('更新成功')
    } else {
      await createRegion(regionForm.value)
      ElMessage.success('添加成功')
    }
    dialogVisible.value = false
    fetchRegions()
  } catch (error) { console.error('操作失败:', error) }
}

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除该区域吗？', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' })
    await deleteRegion(row.code)
    ElMessage.success('删除成功')
    fetchRegions()
  } catch (error) { if (error !== 'cancel') console.error('删除失败:', error) }
}

onMounted(() => { fetchRegions() })
</script>

<style scoped>
.region-container { padding: 20px; }
.card-header { display: flex; justify-content: space-between; align-items: center; }
</style>
