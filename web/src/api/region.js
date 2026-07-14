import service from '@/utils/request'

// 获取区域树
export const getRegionTree = () => {
  return service({
    url: '/region/tree',
    method: 'get'
  })
}

// 获取区域列表
export const getRegionList = (params) => {
  return service({
    url: '/region/list',
    method: 'get',
    params
  })
}

// 创建区域
export const createRegion = (data) => {
  return service({
    url: '/region',
    method: 'post',
    data
  })
}

// 更新区域
export const updateRegion = (code, data) => {
  return service({
    url: `/region/${code}`,
    method: 'put',
    data
  })
}

// 删除区域
export const deleteRegion = (code) => {
  return service({
    url: `/region/${code}`,
    method: 'delete'
  })
}
