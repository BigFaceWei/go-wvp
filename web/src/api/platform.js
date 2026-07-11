import service from '@/utils/request'

export function getPlatformList(params) {
  return service({
    url: '/platform/list',
    method: 'get',
    params
  })
}

export function getOnlinePlatforms() {
  return service({
    url: '/platform/online',
    method: 'get'
  })
}

export function getPlatform(id) {
  return service({
    url: `/platform/${id}`,
    method: 'get'
  })
}

export function createPlatform(data) {
  return service({
    url: '/platform',
    method: 'post',
    data
  })
}

export function updatePlatform(id, data) {
  return service({
    url: `/platform/${id}`,
    method: 'put',
    data
  })
}

export function deletePlatform(id) {
  return service({
    url: `/platform/${id}`,
    method: 'delete'
  })
}

export function getPlatformChannels(id) {
  return service({
    url: `/platform/${id}/channels`,
    method: 'get'
  })
}

export function syncPlatformCatalog(id) {
  return service({
    url: `/platform/${id}/sync`,
    method: 'post'
  })
}

export function startPlatformCascade(id) {
  return service({
    url: `/platform/${id}/start`,
    method: 'post'
  })
}

export function stopPlatformCascade(id) {
  return service({
    url: `/platform/${id}/stop`,
    method: 'post'
  })
}
