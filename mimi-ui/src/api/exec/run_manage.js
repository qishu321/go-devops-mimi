import request from '@/utils/request'

  export function runTaskManage(params) {
    return request({
      url: '/api/exec/run_task_manage/run',
      method: 'get',
      params
    })
  }


  

