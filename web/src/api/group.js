import service from '@/utils/request'

// 获取分组树
export const getGroupTree = () => {
  return service({
    url: '/group/tree',
    method: 'get'
  })
}

// 获取分组列表
export const getGroupList = (params) => {
  return service({
    url: '/group/list',
    method: 'get',
    params
  })
}

// 创建分组
export const createGroup = (data) => {
  return service({
    url: '/group',
    method: 'post',
    data
  })
}

// 更新分组
export const updateGroup = (id, data) => {
  return service({
    url: `/group/${id}`,
    method: 'put',
    data
  })
}

// 删除分组
export const deleteGroup = (id) => {
  return service({
    url: `/group/${id}`,
    method: 'delete'
  })
}
