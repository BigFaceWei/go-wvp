import service from '@/utils/request'

export function getCategoryList(params) {
  return service({
    url: '/attachmentCategory/getCategoryList',
    method: 'get',
    params
  })
}

export function addCategory(data) {
  return service({
    url: '/attachmentCategory/addCategory',
    method: 'post',
    data
  })
}

export function deleteCategory(data) {
  return service({
    url: '/attachmentCategory/deleteCategory',
    method: 'post',
    data
  })
}
