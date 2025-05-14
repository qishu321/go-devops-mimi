import request from '@/utils/request'

// 获取数据列表
export function listNode(params) {
  return request({
    url: '/api/cmdb/node/list',
    method: 'get',
    params
  })
}

// 创建数据
export function addNode(data) {
  return request({
    url: '/api/cmdb/node/add',
    method: 'post',
    data
  })
}

// 更新数据
export function updateNode(data) {
  return request({
    url: '/api/cmdb/node/update',
    method: 'post',
    data
  })
}
export function addNodesGroup (data) {
  return request({
    url: '/api/cmdb/node/add_nodes_group ',
    method: 'post',
    data
  })
}

// 批量删除数据
export function delNode(data) {
  return request({
    url: '/api/cmdb/node/delete',
    method: 'post',
    data
  })
}
