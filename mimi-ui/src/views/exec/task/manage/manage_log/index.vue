<template>
  <div class="app-container">
    <el-form :inline="true" :model="params" class="demo-form-inline" size="mini">
      <el-form-item label="名称">
        <el-input v-model.trim="params.name" clearable placeholder="请输入名称" @clear="search" />
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
      style="width: 97%"
      :cell-style="{padding: '10px 0', fontSize: '12px'}"
      :header-cell-style="{padding: '10px 0', fontSize: '12px'}"
    >
    <el-table-column prop="ID" label="ID" width="60" align="center" />
      <el-table-column prop="name" label="名称" width="180" align="center" />
      <el-table-column prop="args" label="环境变量" width="230" align="center" />
      <el-table-column prop="status" label="状态" width="100" align="center">
        <template #default="{ row }">
          <el-tag
            :type="getStatusTagType(row.status)"
            size="mini"
            effect="dark"
          >
            {{ row.status }}
          </el-tag>
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
      <el-table-column prop="desc" label="描述" width="240" align="center" />
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

  </div>
</template>

<script>
import { runList } from '@/api/exec/task'
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
        pageNum: 1,
        pageSize: 10
      },
      // 表格数据
      tableData: [],
      total: 0,
      loading: false,
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
        const { data } = await runList(this.params)
        this.tableData = data.manage_log_s
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
    getStatusTagType(status) {
      switch(status) {
        case 'running': return 'warning';
        case 'success': return 'success';
        case 'failed': return 'danger';
        default: return 'info';
      }
    },

    showDetail(row) {
      this.$router.push({
        path: '/exec/task/manage/info',
        query: { runID: row.ID }
      })
    }  }
}
</script>

<style scoped>
.app-container {
  padding: 20px;
}
</style>