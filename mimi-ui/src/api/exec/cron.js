import request from '@/utils/request'

// 创建数据
export function cronEnable(data) {
    return request({
      url: '/api/exec/cron/enable',
      method: 'post',
      data
    })
  }
  export function cronDelete(data) {
    return request({
      url: '/api/exec/cron/delete',
      method: 'post',
      data
    })
  }
  export function cronUpdate(data) {
    return request({
      url: '/api/exec/cron/update',
      method: 'post',
      data
    })
  }
  export function cronAdd(data) {
    return request({
      url: '/api/exec/cron/add',
      method: 'post',
      data
    })
  }

  export function cronInfo(params) {
    return request({
      url: '/api/exec/cron/info',
      method: 'get',
      params
    })
  }
  export function cronList(params) {
    return request({
      url: '/api/exec/cron/list',
      method: 'get',
      params
    })
  }
  
  export function cronLogList(params) {
    return request({
      url: '/api/exec/cron/log/list',
      method: 'get',
      params
    })
  }
