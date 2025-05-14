import request from '@/utils/request'

// 获取分组树(已完成)
export function getGroupTree(params) {
  return request({
    url: '/api/system/group/tree',
    method: 'get',
    params
  })
}
// 分组列表(已完成)
export function groupList(params) {
  return request({
    url: '/api/system/group/list',
    method: 'get',
    params
  })
}
// 添加分组(已完成)
export function groupAdd(data) {
  return request({
    url: '/api/system/group/add',
    method: 'post',
    data
  })
}
// 更新分组(已完成)
export function groupUpdate(data) {
  return request({
    url: '/api/system/group/update',
    method: 'post',
    data
  })
}
// 删除分组（已完成）
export function groupDel(data) {
  return request({
    url: '/api/system/group/delete',
    method: 'post',
    data
  })
}
// 在分组内的用户（已完成）
export function useGroupList(params) {
  return request({
    url: '/api/system/group/useringroup',
    method: 'get',
    params
  })
}
// 不再分组内的用户（已完成）
export function useGroupRole(params) {
  return request({
    url: '/api/system/group/usernoingroup',
    method: 'get',
    params
  })
}
// 添加用户到分组（已完成）
export function groupInfo(data) {
  return request({
    url: '/api/system/group/adduser',
    method: 'post',
    data
  })
}
// 将用户从分组移出（已完成）
export function delGroup(data) {
  return request({
    url: '/api/system/group/removeuser',
    method: 'post',
    data
  })
}
