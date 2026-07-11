import service from '@/utils/request'

export function getFileList(params) {
  return service({
    url: '/fileUploadAndDownload/getFileList',
    method: 'post',
    data: params
  })
}

export function editFileName(data) {
  return service({
    url: '/fileUploadAndDownload/editFileName',
    method: 'post',
    data
  })
}

export function deleteFile(data) {
  return service({
    url: '/fileUploadAndDownload/deleteFile',
    method: 'post',
    data
  })
}
