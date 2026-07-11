import service from '@/utils/service'

export function getPositionList(params) {
  return service({
    url: '/api/position/list',
    method: 'get',
    params
  })
}

export function getDevicePosition(id) {
  return service({
    url: `/api/position/${id}`,
    method: 'get'
  })
}

export function getPositionTrack(id) {
  return service({
    url: `/api/position/${id}/track`,
    method: 'get'
  })
}

export function deletePosition(id) {
  return service({
    url: `/api/position/${id}`,
    method: 'delete'
  })
}
