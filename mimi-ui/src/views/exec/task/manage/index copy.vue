<template>
  <el-card class="container-card" shadow="always">
    <!-- 查询表单 -->
    <el-form size="mini" inline :model="listQuery" class="demo-form-inline">
      <el-form-item label="任务名称">
        <el-input
          v-model.trim="listQuery.name"
          clearable
          placeholder="任务名称"
          @input="debouncedSearch"
          size="small"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          :loading="loading"
          icon="el-icon-search"
          type="primary"
          size="small"
          @click="search"
        >查询</el-button>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          icon="el-icon-plus"
          size="small"
          @click="goToCreate"
        >创建任务</el-button>
      </el-form-item>
    </el-form>

    <!-- 任务列表 -->
    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      :row-key="row => row.ID"
      style="width: 100%; margin-top: 16px;"
    >
      <el-table-column label="ID" prop="ID" align="center" width="80" />
      <el-table-column label="任务名称" prop="name" min-width="150" />
      <el-table-column label="描述信息" prop="desc" min-width="150" />
      <el-table-column label="创建人" prop="creator" min-width="120" />
      <el-table-column label="操作" align="center" width="300">
        <template #default="{ row }">
          <el-button
            size="mini"
            type="success"
            icon="el-icon-play"
            @click="runUp(row)"
          >执行任务</el-button>
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-edit"
            @click="addUp(row)"
          >编辑</el-button>
          <el-button
            size="mini"
            type="danger"
            icon="el-icon-delete"
            @click="handleDelete(row)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <pagination
      v-if="total > 0"
      :total="total"
      :page.sync="listQuery.pageNum"
      :limit.sync="listQuery.pageSize"
      @pagination="getList"
    />

    <!-- 执行参数对话框 -->
    <el-dialog
      title="执行参数配置"
      :visible.sync="dialogVisible"
      width="50%"
      append-to-body
      destroy-on-close
    >
      <el-form
        v-if="formItems.length"
        ref="paramForm"
        :model="formParams"
        :rules="formRules"
        label-width="120px"
        size="small"
      >
        <el-form-item
          v-for="item in formItems"
          :key="item.name"
          :label="item.name"
          :prop="item.name"
        >
          <el-input
            v-if="item.type === 'string'"
            v-model="formParams[item.name]"
            :placeholder="'请输入' + item.name"
          />
          <el-select
            v-else-if="item.type === 'select'"
            v-model="formParams[item.name]"
            placeholder="请选择"
            style="width: 100%;"
          >
            <el-option
              v-for="(opt, idx) in item.options.split('\n')"
              :key="idx"
              :label="opt.split(':')[1] || opt"
              :value="opt.split(':')[0]"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <el-empty v-else description="当前无执行参数" />

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitParams">立即执行</el-button>
      </template>
    </el-dialog>
  </el-card>
</template>

<script>
import { listTaskManage, deleteTaskManage, addRun } from '@/api/exec/task'
import Pagination from '@/components/Pagination'
import debounce from 'lodash/debounce'

export default {
  name: 'TaskManage',
  components: { Pagination },
  data() {
    return {
      // 列表查询
      listQuery: { pageNum: 1, pageSize: 20, name: '' },
      listLoading: false,
      list: [],
      total: 0,
      loading: false,

      // 执行参数对话框
      dialogVisible: false,
      currentTask: null,
      formItems: [],
      formParams: {},
      formRules: {}
    }
  },
  created() {
    this.getList()
    this.debouncedSearch = debounce(this.search, 300)
  },
  methods: {
    // 拉列表
    async getList() {
      this.listLoading = true
      try {
        const { data = {} } = await listTaskManage(this.listQuery)
        this.list = data.task_manage_s || []
        this.total = data.total || 0
      } catch (err) {
        this.$message.error('列表加载失败')
      } finally {
        this.listLoading = false
      }
    },
    // 搜索
    search() {
      this.listQuery.pageNum = 1
      this.getList()
    },
    // 跳转创建
    goToCreate() {
      this.$router.push('/exec/task/manage/add')
    },
    // 跳转编辑
    addUp(row) {
      this.$router.push({
        path: '/exec/task/manage/update',
        query: { id: row.ID, name: row.name }
      })
    },
    // 删除
    async handleDelete(row) {
      try {
        await this.$confirm('确认删除该任务?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        await deleteTaskManage({ ids: [row.ID] })
        this.$message.success('删除成功')
        this.getList()
      } catch {}
    },
    // 点击“执行任务”
    runUp(row) {
      this.currentTask = row
      let args = []
      try {
        args = JSON.parse(row.args || '[]')
      } catch {
        this.$message.error('参数配置解析失败')
        return
      }
      if (!args.length) {
        // 无参数，直接触发
        this.doRun(row, {})
      } else {
        // 有参数，打开对话框
        this.formItems = args
        this.formParams = {}
        this.formRules = {}
        args.forEach(item => {
          this.$set(this.formParams, item.name, '')
          this.$set(this.formRules, item.name, [
            { required: item.required, message: `请输入${item.name}`, trigger: 'blur' }
          ])
        })
        this.dialogVisible = true
      }
    },
    // 参数对话框提交
    submitParams() {
      this.$refs.paramForm.validate(valid => {
        if (!valid) return
        this.doRun(this.currentTask, { ...this.formParams })
        this.dialogVisible = false
      })
    },
    // 真正发起任务
    async doRun(task, envParams) {
      this.loading = true
      try {
        const res = await addRun({
          id: task.ID,
          name: task.name,
          env_task_s: envParams
        })
        this.$message.success('任务开始执行')
        this.$router.push({
          path: '/exec/task/manage/info',
          query: { runID: res.data.runID }
        })
      } catch (err) {
        this.$message.error('任务执行失败: ' + (err.message || '未知错误'))
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
.container-card {
  margin: 10px;
}
</style>
