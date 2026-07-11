import service from '@/utils/service'

export function getPlatformList(params) {
  return service({
    url: '/api/platform/list',
    method: 'get',
    params
  })
}

export function getOnlinePlatforms() {
  return service({
    url: '/api/platform/online',
    method: 'get'
  })
}

export function getPlatform(id) {
  return service({
    url: `/api/platform/${id}`,
    method: 'get'
  })
}

export function createPlatform(data) {
  return service({
    url: '/api/platform',
    method: 'post',
    data
  })
}

export function updatePlatform(id, data) {
  return service({
    url: `/api/platform/${id}`,
    method: 'put',
    data
  })
}

export function deletePlatform(id) {
  return service({
    url: `/api/platform/${id}`,
    method: 'delete'
  })
}

export function getPlatformChannels(id) {
  return service({
    url: `/api/platform/${id}/channels`,
    method: 'get'
  })
}

export function syncPlatformCatalog(id) {
  return service({
    url: `/api/platform/${id}/sync`,
    method: 'post'
  })
}

export function startPlatformCascade(id) {
  return service({
    url: `/api/platform/${id}/start`,
    method: 'post'
  })
}

export function stopPlatformCascade(id) {
  return service({
    url: `/api/platform/${id}/stop`,
    method: 'post'
  })
}
