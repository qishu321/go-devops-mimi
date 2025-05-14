import request from '@/utils/request'

// 创建数据
export function addTaskManage(data) {
    return request({
      url: '/api/exec/task_manage/add',
      method: 'post',
      data
    })
  }
  //更新数据
  export function updateTaskManage(data) {
    return request({
      url: '/api/exec/task_manage/update',
      method: 'post',
      data
    })
  }
  // 批量删除数据
  export function deleteTaskManage(data) {
    return request({
      url: '/api/exec/task_manage/delete',
      method: 'post',
      data
    })
  }
  // 获取数据列表
  export function listTaskManage(params) {
    return request({
      url: '/api/exec/task_manage/list',
      method: 'get',
      params
    })
  }
  //获取指定数据
  export function infoTaskManage(params) {
    return request({
      url: '/api/exec/task_manage/info',
      method: 'get',
      params
    })
  }
  export function addRun(data) {
    return request({
      url: '/api/exec/run_task_manage/add_run',
      method: 'post',
      data
    })
  }
  export function runInfo(params) {
    return request({
      url: '/api/exec/run_task_manage/run_info',
      method: 'get',
      params
    })
  }
  export function runList(params) {
    return request({
      url: '/api/exec/run_task_manage/run_list',
      method: 'get',
      params
    })
  }

  export function runInfoWebSocket(params) {
    return request({
      url: '/api/exec/run_task_manage/run_info_webSocket',
      method: 'get',
      params
    })
  }


