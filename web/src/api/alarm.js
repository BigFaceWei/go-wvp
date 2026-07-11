import service from '@/utils/service'

export function getAlarmList(params) {
  return service({
    url: '/api/alarm/list',
    method: 'get',
    params
  })
}

export function getAlarmStatistics(params) {
  return service({
    url: '/api/alarm/statistics',
    method: 'get',
    params
  })
}

export function getAlarm(id) {
  return service({
    url: `/api/alarm/${id}`,
    method: 'get'
  })
}

export function handleAlarm(id, data) {
  return service({
    url: `/api/alarm/${id}/handle`,
    method: 'post',
    data
  })
}

export function deleteAlarm(id) {
  return service({
    url: `/api/alarm/${id}`,
    method: 'delete'
  })
}
