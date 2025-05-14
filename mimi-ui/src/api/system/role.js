import request from '@/utils/request'

// 获取角色列表（已完成）
export function getRoles(params) {
  return request({
    url: '/api/system/role/list',
    method: 'get',
    params
  })
}

// 创建角色（已完成）
export function createRole(data) {
  return request({
    url: '/api/system/role/add',
    method: 'post',
    data
  })
}

// 更新角色(已完成)
export function updateRoleById(data) {
  return request({
    url: '/api/system/role/update',
    method: 'post',
    data
  })
}

// 获取角色的权限菜单(已完成)
export function getRoleMenusById(params) {
  return request({
    url: '/api/system/role/getmenulist',
    method: 'get',
    params
  })
}

// 更新角色的权限菜单(已完成)
export function updateRoleMenusById(data) {
  return request({
    url: '/api/system/role/updatemenus',
    method: 'post',
    data
  })
}

// 获取角色的权限接口(已完成)
export function getRoleApisById(params) {
  return request({
    url: '/api/system/role/getapilist',
    method: 'get',
    params
  })
}

// 更新角色的权限接口(已完成)
export function updateRoleApisById(data) {
  return request({
    url: '/api/system/role/updateapis',
    method: 'post',
    data
  })
}

// 批量删除角色(已完成)
export function batchDeleteRoleByIds(data) {
  return request({
    url: '/api/system/role/delete',
    method: 'post',
    data
  })
}
