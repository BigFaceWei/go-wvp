import service from '@/utils/request'

export function getDeviceChannels(deviceId) {
  return service({
    url: `/device/${deviceId}/channels`,
    method: 'get'
  })
}

export function getChannelList(params) {
  return service({
    url: '/channel/list',
    method: 'get',
    params
  })
}

export function playVideo(data) {
  return service({
    url: '/video/play',
    method: 'post',
    data
  })
}

export function stopVideo(deviceId) {
  return service({
    url: `/video/${deviceId}/stop`,
    method: 'post'
  })
}
