import request from '@/utils/request'

// 获取数据列表
export function getCloudAccounts(params) {
  return request({
    url: '/api/example/cloudaccount/list',
    method: 'get',
    params
  })
}

// 创建数据
export function createCloudAccount(data) {
  return request({
    url: '/api/example/cloudaccount/add',
    method: 'post',
    data
  })
}

// 更新数据
export function updateCloudAccountById(data) {
  return request({
    url: '/api/example/cloudaccount/update',
    method: 'post',
    data
  })
}

// 批量删除数据
export function batchDeleteCloudAccountByIds(data) {
  return request({
    url: '/api/example/cloudaccount/delete',
    method: 'post',
    data
  })
}
