import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/admins/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/admins/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/admins/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/admins/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/admins/delete',
    method: 'post',
    data
  })
}

export function requestAdminsRoleIDList(adminsid) {
  return request({
    url: '/admins/adminsroleidlist',
    method: 'get',
    params: { adminsid }
  })
}

export function requestSetRole(adminsid, data) {
  return request({
    url: '/admins/setrole',
    method: 'post',
    params: { adminsid },
    data
  })
}

