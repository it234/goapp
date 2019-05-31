import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/menu/list',
    method: 'get',
    params: query
  })
}

export function requestAll() {
  return request({
    url: '/menu/allmenu',
    method: 'get'
  })
}

export function requestDetail(id) {
  return request({
    url: '/menu/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/menu/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/menu/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/menu/delete',
    method: 'post',
    data
  })
}

export function requestMenuButton(menucode) {
  return request({
    url: '/menu/menubuttonlist',
    method: 'get',
    params: { menucode }
  })
}

