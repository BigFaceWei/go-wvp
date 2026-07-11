import service from '@/utils/service'

export function getRecordPlanList(params) {
  return service({
    url: '/api/record/plan/list',
    method: 'get',
    params
  })
}

export function getRecordPlan(id) {
  return service({
    url: `/api/record/plan/${id}`,
    method: 'get'
  })
}

export function createRecordPlan(data) {
  return service({
    url: '/api/record/plan',
    method: 'post',
    data
  })
}

export function updateRecordPlan(id, data) {
  return service({
    url: `/api/record/plan/${id}`,
    method: 'put',
    data
  })
}

export function deleteRecordPlan(id) {
  return service({
    url: `/api/record/plan/${id}`,
    method: 'delete'
  })
}

export function startRecordPlan(id) {
  return service({
    url: `/api/record/plan/${id}/start`,
    method: 'post'
  })
}

export function stopRecordPlan(id) {
  return service({
    url: `/api/record/plan/${id}/stop`,
    method: 'post'
  })
}

export function getRecordFileList(params) {
  return service({
    url: '/api/record/file/list',
    method: 'get',
    params
  })
}
