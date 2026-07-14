import service from '@/utils/request'

// 获取拉流代理列表
export const getStreamProxyList = (params) => {
  return service({
    url: '/stream/proxy/list',
    method: 'get',
    params
  })
}

// 获取拉流代理详情
export const getStreamProxy = (id) => {
  return service({
    url: `/stream/proxy/${id}`,
    method: 'get'
  })
}

// 创建拉流代理
export const createStreamProxy = (data) => {
  return service({
    url: '/stream/proxy',
    method: 'post',
    data
  })
}

// 更新拉流代理
export const updateStreamProxy = (id, data) => {
  return service({
    url: `/stream/proxy/${id}`,
    method: 'put',
    data
  })
}

// 删除拉流代理
export const deleteStreamProxy = (id) => {
  return service({
    url: `/stream/proxy/${id}`,
    method: 'delete'
  })
}

// 启动拉流代理
export const startStreamProxy = (id) => {
  return service({
    url: `/stream/proxy/${id}/start`,
    method: 'post'
  })
}

// 停止拉流代理
export const stopStreamProxy = (id) => {
  return service({
    url: `/stream/proxy/${id}/stop`,
    method: 'post'
  })
}
