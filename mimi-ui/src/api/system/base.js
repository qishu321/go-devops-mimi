import request from '@/utils/request'
//用户登录（已完成）
export function login(data) {
  return request({
    url: '/api/system/base/login',
    method: 'post',
    data
  })
}

export function refreshToken() {
  return request({
    url: '/api/system/base/refreshToken',
    method: 'post'
  })
}
//用户退出接口（已完成）
export function logout() {
  return request({
    url: '/api/system/base/logout',
    method: 'post'
  })
}
