import service from '@/utils/request'

// 获取推流列表
export const getStreamPushList = (params) => {
  return service({
    url: '/stream/push/list',
    method: 'get',
    params
  })
}

// 获取推流详情
export const getStreamPush = (id) => {
  return service({
    url: `/stream/push/${id}`,
    method: 'get'
  })
}

// 创建推流
export const createStreamPush = (data) => {
  return service({
    url: '/stream/push',
    method: 'post',
    data
  })
}

// 更新推流
export const updateStreamPush = (id, data) => {
  return service({
    url: `/stream/push/${id}`,
    method: 'put',
    data
  })
}

// 删除推流
export const deleteStreamPush = (id) => {
  return service({
    url: `/stream/push/${id}`,
    method: 'delete'
  })
}

// 启动推流
export const startStreamPush = (id) => {
  return service({
    url: `/stream/push/${id}/start`,
    method: 'post'
  })
}

// 停止推流
export const stopStreamPush = (id) => {
  return service({
    url: `/stream/push/${id}/stop`,
    method: 'post'
  })
}
