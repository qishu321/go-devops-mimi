<template>
  <div>
    <el-card class="container-card" shadow="always">
      <el-form size="mini" :inline="true" :model="params" class="demo-form-inline">
        <el-form-item label="任务名称">
          <el-input v-model.trim="params.name" clearable placeholder="任务名称" @clear="search" />
        </el-form-item>
        <el-form-item label="任务类型">
          <el-radio-group v-model="params.cronType">
            <el-radio-button label="">全部</el-radio-button>
            <el-radio-button label="cron">Cron表达式</el-radio-button>
            <el-radio-button label="once">单次执行</el-radio-button>
            <el-radio-button label="interval">间隔执行</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" icon="el-icon-search" type="primary" @click="search">查询</el-button>
        </el-form-item>
      </el-form>

      <el-table v-loading="loading" :data="tableData" border stripe style="width: 100%">
        <el-table-column show-overflow-tooltip sortable prop="name" label="任务名称"  width="300" align="center" />
        <el-table-column prop="cronType" label="任务类型" width="100" align="center">
          <template slot-scope="scope">
            <el-tag :type="getCronTypeTag(scope.row.cronType)">
              {{ scope.row.cronType | cronTypeFilter }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column show-overflow-tooltip sortable prop="status" label="状态" width="100"  align="center">
          <template slot-scope="scope">
            <el-tag size="small" :type="scope.row.status | statusTagFilter" disable-transitions>{{ scope.row.status | statusTextFilter }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column show-overflow-tooltip sortable prop="desc" label="描述"align="center"  />

        <el-table-column show-overflow-tooltip sortable prop="startTime" label="开始时间"  width="160" align="center" >
          <template slot-scope="scope">
            {{ formatTime(scope.startTime) }}
          </template>
        </el-table-column>    
        <el-table-column show-overflow-tooltip sortable prop="endTime" label="结束时间" width="160" align="center" >
          <template slot-scope="scope">
            {{ formatTime(scope.endTime) }}
          </template>
        </el-table-column>    

        <el-table-column show-overflow-tooltip sortable prop="timeCost" label="耗时(ms)" align="center" width="100">
          <template slot-scope="scope">
            <el-tag size="small" :type="scope.row.timeCost | timeCostTagFilter" disable-transitions>{{ scope.row.timeCost }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="120" align="center">
          <template #default="{ row }">
            <el-button type="text" size="small"  @click="showDetail(row)">  查看日志详情 </el-button>
          </template>
        </el-table-column>
  
      </el-table>

      <el-pagination
        :current-page="params.pageNum"
        :page-size="params.pageSize"
        :total="total"
        :page-sizes="[1, 5, 10, 30]"
        layout="total, prev, pager, next, sizes"
        background
        style="margin-top: 10px;float:right;margin-bottom: 10px;"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </el-card>

    <el-drawer 
      title="定时任务执行详情" 
      :visible.sync="drawerVisible" 
      direction="rtl" 
      size="80%" 
    > 
      <div v-if="currentLog" style="padding: 20px"> 
        <el-descriptions :column="2" border>
          <el-descriptions-item label="任务名称">
            <el-tag type="primary" size="medium">任务名称：{{ currentLog.name }}</el-tag>
          </el-descriptions-item>
        </el-descriptions>
        <h3 style="margin: 20px 0 10px;">节点执行详情</h3>
        <el-table 
          :data="currentLog.t_cron_script_log_s" 
          border 
          fit 
          highlight-current-row 
          style="width: 100%; margin-top: 10px;" 
        > 
          <el-table-column prop="node_name" label="节点名称" width="100" align="center" /> 
          <el-table-column prop="type" label="类型" width="80" align="center" /> 
          <el-table-column prop="status" label="状态" width="80" align="center"> 
            <template #default="{ row }"> 
              <el-tag :type="row.status === 1 ? 'success' : 'danger'"> 
                {{ row.status === 1 ? '成功' : '失败' }} 
              </el-tag> 
            </template> 
          </el-table-column> 
          <el-table-column prop="timeout" label="命令超时时间(s)" width="100" align="center" /> 
          <el-table-column prop="content" label="执行命令" width="500" > 
            <template #default="{ row }"> 
              <pre style="background: #000; color: #fff; padding: 10px; border-radius: 4px; white-space: pre-wrap; font-family: monospace;">{{ row.content }}</pre> 
            </template> 
          </el-table-column> 
          <el-table-column prop="run_log" label="响应结果" width="400" > 
            <template #default="{ row }"> 
              <pre style="background: #000; color: #fff; padding: 10px; border-radius: 4px; white-space: pre-wrap; font-family: monospace;">{{ row.run_log }}</pre> 
            </template> 
          </el-table-column> 
          <el-table-column prop="startTime" label="开始时间" width="100" align="center"> 
            <template #default="{ row }"> 
              {{ formatTime(row.startTime) }} 
            </template> 
          </el-table-column> 
          <el-table-column prop="endTime" label="结束时间" width="100" align="center"> 
            <template #default="{ row }"> 
              {{ formatTime(row.endTime) }} 
            </template> 
          </el-table-column> 
          <el-table-column prop="timeCost" label="耗时(ms)" width="100" align="center" /> 
        </el-table> 
      </div> 
    </el-drawer>
</div>
</template>

<script>
import { Message } from 'element-ui'
import { cronLogList } from '@/api/exec/cron'
import dayjs from 'dayjs'  // 添加dayjs导入

export default {
  name: 'CronLog',
  filters: {
    statusTagFilter(val) {
      if (val === 1) {
        return 'success'
      } else {
        return 'danger'
      }
    },
    cronTypeFilter(type) {
      const map = { interval: '间隔执行', once: '单次执行', cron: 'Cron表达式' }
      return map[type] || type
    },
    statusTextFilter(val) {
      return val === 1 ? '成功' : '失败'
    },
    timeCostTagFilter(val) {
      if (val <= 200) {
        return 'success'
      } else if (val > 200 && val <= 1000) {
        return ''
      } else if (val > 1000 && val <= 2000) {
        return 'warning'
      } else {
        return 'danger'
      }
    }
  },
  data() {
    return {
      params: {
        name: '',
        cronType: '', // 新增任务类型参数
        pageNum: 1,
        pageSize: 10
      },
      // 表格数据
      tableData: [],
      total: 0,
      loading: false,
      drawerVisible: false,
      currentLog: null
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    // 移除未定义的parseGoTime引用
    // 查询
    search() {
      this.params.pageNum = 1
      this.getTableData()
    },
    formatTime(time) {
      return dayjs(time).format('YYYY-MM-DD HH:mm:ss')
    },

    // 获取表格数据
    async getTableData() {
      this.loading = true
      try {
        const { data } = await cronLogList(this.params)
        this.tableData = data.cron_log_s
        this.total = data.total
      } finally {
        this.loading = false
      }
    },

    // 判断结果
    judgeResult(res){
      if (res.code==0){
          Message({
            showClose: true,
            message: "操作成功",
            type: 'success'
          })
        }
    },
    getCronTypeTag(type) {
      return { interval: 'success', once: 'warning', cron: 'primary' }[type] || ''
    },

    // 分页
    handleSizeChange(val) {
      this.params.pageSize = val
      this.getTableData()
    },
    handleCurrentChange(val) {
      this.params.pageNum = val
      this.getTableData()
    },
    showDetail(row) {
      this.currentLog = row
      this.drawerVisible = true
    },
  }
}
</script>

<style scoped>
  .container-card{
    margin: 10px;
    margin-bottom: 100px;
  }

  .delete-popover{
    margin-left: 10px;
  }
</style>
