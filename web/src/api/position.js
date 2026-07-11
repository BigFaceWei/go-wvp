import service from '@/utils/request'

export function getPositionList(params) {
  return service({
    url: '/position/list',
    method: 'get',
    params
  })
}

export function getDevicePosition(id) {
  return service({
    url: `/position/${id}`,
    method: 'get'
  })
}

export function getPositionTrack(id) {
  return service({
    url: `/position/${id}/track`,
    method: 'get'
  })
}

export function deletePosition(id) {
  return service({
    url: `/position/${id}`,
    method: 'delete'
  })
}
