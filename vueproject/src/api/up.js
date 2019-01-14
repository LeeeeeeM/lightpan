import request from '@/utils/requestUp'
import { getName } from '@/utils/auth'
export function postFile(path,formData) {
  return request({
    url: path,
    method: 'post',
    config:{
      headers:{'Content-Type':'multipart/form-data'}
    },
    data:formData
  })
}
export function getFile(path) {
  return request({
    url: getName()+path,
    method: 'get',
  })
}


