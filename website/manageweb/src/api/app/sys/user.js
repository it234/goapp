import request from '@/utils/request'

export function requestEditPwd(data) {
  return request({
    url: '/user/editpwd',
    method: 'post',
    data
  })
}
