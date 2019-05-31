import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/role/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/role/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/role/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/role/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/role/delete',
    method: 'post',
    data
  })
}

export function requestRoleMenuIDList(roleid) {
  return request({
    url: '/role/rolemenuidlist',
    method: 'get',
    params: { roleid }
  })
}

export function requestSetRole(roleid, data) {
  return request({
    url: '/role/setrole',
    method: 'post',
    params: { roleid },
    data
  })
}

export function requestAll() {
  return request({
    url: '/role/allrole',
    method: 'get'
  })
}
