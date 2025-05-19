import request from '@/utils/request'

// 获取数据列表
export function listNav(params) {
  return request({
    url: '/api/nav/list',
    method: 'get',
    params
  })
}
export function infoNav(params) {
  return request({
    url: '/api/nav/info',
    method: 'get',
    params
  })
}

// 创建数据
export function addNav(data) {
  return request({
    url: '/api/nav/add',
    method: 'post',
    data
  })
}

// 更新数据
export function updateNav(data) {
  return request({
    url: '/api/nav/update',
    method: 'post',
    data
  })
}

// 批量删除数据
export function delAllNav(data) {
  return request({
    url: '/api/nav/delete_all',
    method: 'post',
    data
  })
}

export function infoLink(params) {
  return request({
    url: '/api/nav/link/info',
    method: 'get',
    params
  })
}
export function updateLink(data) {
  return request({
    url: '/api/nav/link/update',
    method: 'post',
    data
  })
}
export function addLink(data) {
  return request({
    url: '/api/nav/link/add',
    method: 'post',
    data
  })
}

export function delLink(data) {
  return request({
    url: '/api/nav/link/delete',
    method: 'post',
    data
  })
}
