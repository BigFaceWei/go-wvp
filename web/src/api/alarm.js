import service from '@/utils/request'

export function getAlarmList(params) {
  return service({
    url: '/alarm/list',
    method: 'get',
    params
  })
}

export function getAlarmStatistics(params) {
  return service({
    url: '/alarm/statistics',
    method: 'get',
    params
  })
}

export function getAlarm(id) {
  return service({
    url: `/alarm/${id}`,
    method: 'get'
  })
}

export function handleAlarm(id, data) {
  return service({
    url: `/alarm/${id}/handle`,
    method: 'post',
    data
  })
}

export function deleteAlarm(id) {
  return service({
    url: `/alarm/${id}`,
    method: 'delete'
  })
}
