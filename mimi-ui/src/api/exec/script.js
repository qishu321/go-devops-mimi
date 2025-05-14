import request from '@/utils/request'

// 获取数据列表
export function logList(params) {
  return request({
    url: '/api/exec/script/list',
    method: 'get',
    params
  })
}

// 创建数据
export function addRun(data) {
  return request({
    url: '/api/exec/script/add_run',
    method: 'post',
    data
  })
}
// 创建数据
export function addLibrary(data) {
    return request({
      url: '/api/exec/script_library/add',
      method: 'post',
      data
    })
  }
  //更新数据
  export function updateLibrary(data) {
    return request({
      url: '/api/exec/script_library/update',
      method: 'post',
      data
    })
  }
  // 批量删除数据
  export function deleteLibrary(data) {
    return request({
      url: '/api/exec/script_library/delete',
      method: 'post',
      data
    })
  }
  // 获取数据列表
  export function listLibrary(params) {
    return request({
      url: '/api/exec/script_library/list',
      method: 'get',
      params
    })
  }
  //获取指定数据
  export function infoLibrary(params) {
    return request({
      url: '/api/exec/script_library/info',
      method: 'get',
      params
    })
  }

  export function infoTransfer(params) {
    return request({
      url: '/api/exec/transfer/info',
      method: 'get',
      params
    })
  }
  export function listTransfer(params) {
    return request({
      url: '/api/exec/transfer/list',
      method: 'get',
      params
    })
  }
  export function addRunTransfer(data) {
    return request({
      url: '/api/exec/transfer/add_run',
      method: 'post',
      data
    })
  }

  export function uploadTransfer(data) {
    return request({
      url: '/api/exec/transfer/upload',
      method: 'post',
      data
    })
  }

