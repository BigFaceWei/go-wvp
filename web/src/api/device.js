import service from '@/utils/request'

export function getDeviceList(params) {
  return service({
    url: '/device/list',
    method: 'get',
    params
  })
}

export function getDevice(id) {
  return service({
    url: `/device/${id}`,
    method: 'get'
  })
}

export function createDevice(data) {
  return service({
    url: '/device',
    method: 'post',
    data
  })
}

export function updateDevice(id, data) {
  return service({
    url: `/device/${id}`,
    method: 'put',
    data
  })
}

export function deleteDevice(id) {
  return service({
    url: `/device/${id}`,
    method: 'delete'
  })
}

export function getDeviceChannels(id) {
  return service({
    url: `/device/${id}/channels`,
    method: 'get'
  })
}

export function queryDeviceCatalog(id) {
  return service({
    url: `/device/${id}/catalog`,
    method: 'post'
  })
}

export function getDeviceStatus(id) {
  return service({
    url: `/device/${id}/status`,
    method: 'get'
  })
}