import request from '@/utils/request'

// 获取数据列表
export function list_node_group(params) {
  return request({
    url: '/api/cmdb/node_group/list',
    method: 'get',
    params
  })
}

// 创建数据
export function add_node_group(data) {
  return request({
    url: '/api/cmdb/node_group/add',
    method: 'post',
    data
  })
}

// 更新数据
export function update_node_group(data) {
  return request({
    url: '/api/cmdb/node_group/update',
    method: 'post',
    data
  })
}

// 批量删除数据
export function del_node_group(data) {
  return request({
    url: '/api/cmdb/node_group/delete',
    method: 'post',
    data
  })
}
//服务器组添加服务器
export function add_node_to_group(data) {
  return request({
    url: '/api/cmdb/node_group/add_node_to_group',
    method: 'post',
    data
  })
}
//移除服务器
export function remonv_node_to_group(data) {
  return request({
    url: '/api/cmdb/node_group/remonv_node_to_group',
    method: 'post',
    data
  })
}
