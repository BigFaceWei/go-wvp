import service from '@/utils/service'

export function getDeviceList(params) {
  return service({
    url: '/api/device/list',
    method: 'get',
    params
  })
}

export function getDevice(id) {
  return service({
    url: `/api/device/${id}`,
    method: 'get'
  })
}

export function createDevice(data) {
  return service({
    url: '/api/device',
    method: 'post',
    data
  })
}

export function updateDevice(id, data) {
  return service({
    url: `/api/device/${id}`,
    method: 'put',
    data
  })
}

export function deleteDevice(id) {
  return service({
    url: `/api/device/${id}`,
    method: 'delete'
  })
}

export function getDeviceChannels(id) {
  return service({
    url: `/api/device/${id}/channels`,
    method: 'get'
  })
}

export function queryDeviceCatalog(id) {
  return service({
    url: `/api/device/${id}/catalog`,
    method: 'post'
  })
}

export function getDeviceStatus(id) {
  return service({
    url: `/api/device/${id}/status`,
    method: 'get'
  })
}