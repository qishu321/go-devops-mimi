<template>
  <div class="app-container">
    <el-form :inline="true" :model="params" class="demo-form-inline" size="mini">
      <el-form-item label="名称">
        <el-input v-model.trim="params.name" clearable placeholder="请输入名称" @clear="search" />
      </el-form-item>
      <el-form-item label="类型">
        <el-select v-model="params.cmd_type" clearable placeholder="请选择类型" @clear="search">
          <el-option label="命令" value="command" />
          <el-option label="脚本" value="script" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button :loading="loading" icon="el-icon-search" type="primary" @click="search">搜索</el-button>
      </el-form-item>
    </el-form>

    <el-table
      v-loading="loading"
      :data="tableData"
      border
      stripe
      height="500"
      highlight-current-row
      style="width: 95%"
      :cell-style="{padding: '5px 0', fontSize: '12px'}"
      :header-cell-style="{padding: '5px 0', fontSize: '12px'}"
    >
      <el-table-column prop="name" label="名称" width="180" align="center" />
      <el-table-column prop="status" label="状态" width="100" align="center">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'danger'">
            {{ row.status === 1 ? '成功' : '失败' }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="cmd_type" label="类型" width="100" align="center">
        <template #default="{ row }">
          {{ row.cmd_type === 'command' ? '命令' : '脚本' }}
        </template>
      </el-table-column>
      <el-table-column prop="startTime" label="开始时间" width="180" align="center">
        <template #default="{ row }">
          {{ formatTime(row.startTime) }}
        </template>
      </el-table-column>
      <el-table-column prop="endTime" label="结束时间" width="180" align="center">
        <template #default="{ row }">
          {{ formatTime(row.endTime) }}
        </template>
      </el-table-column>
      <el-table-column prop="timeCost" label="耗时(ms)" width="120" align="center" />
      <el-table-column prop="desc" label="描述" align="center" />
      <el-table-column prop="creator" label="执行用户" width="120" align="center" />
      <el-table-column label="操作" width="120" align="center">
        <template #default="{ row }">
          <el-tooltip content="查看详情" effect="dark" placement="top">
            <el-button icon="el-icon-setting" circle type="info" @click="showDetail(row)" />
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>
    <pagination
      v-show="total > 0"
      :total="total"
      :page.sync="params.pageNum"
      :limit.sync="params.pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      background  
      @pagination="getTableData"
    />

    <el-drawer
      title="执行日志详情"
      :visible.sync="drawerVisible"
      direction="rtl"
      size="75%"
    >
      <div v-if="currentLog" style="padding: 20px">
        <el-table
          :data="currentLog.t_script_log_s"
          border
          fit
          highlight-current-row
          style="width: 100%"
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

          <el-table-column prop="run_log" label="响应结果"  width="400" >
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
import { logList } from '@/api/exec/script'
import Pagination from '@/components/Pagination'
import dayjs from 'dayjs'

export default {
  name: 'ScriptLog',
  components: { Pagination },
  data() {
    return {
      // 查询参数
      params: {
        name: '',
        cmd_type: '',
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
    // 获取表格数据
    async getTableData() {
      this.loading = true
      try {
        const { data } = await logList(this.params)
        this.tableData = data.script_s
        this.total = data.total
      } finally {
        this.loading = false
      }
    },
    // 搜索
    search() {
      this.params.pageNum = 1
      this.getTableData()
    },
    formatTime(time) {
      return dayjs(time).format('YYYY-MM-DD HH:mm:ss')
    },
    showDetail(row) {
      this.currentLog = row
      this.drawerVisible = true
    }
  }
}
</script>

<style scoped>
.app-container {
  padding: 20px;
}
</style>