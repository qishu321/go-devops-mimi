<template>
  <div>
    <el-card class="container-card" shadow="always">
      <el-form size="mini" :inline="true" :model="params" class="demo-form-inline">
        <el-form-item label="云账号">
          <el-input v-model.trim="params.cloudName" clearable placeholder="云账号" @change="search" @clear="search" />
        </el-form-item>
        <el-form-item label="云厂商">
          <el-select v-model.trim="params.cloudType" clearable placeholder="云厂商" @change="search" @clear="search">
            <el-option label="腾讯云" value="腾讯云" />
            <el-option label="阿里云" value="阿里云" />
          </el-select>
        </el-form-item>
        <el-form-item label="秘钥ID">
          <el-input v-model.trim="params.secretId" clearable placeholder="秘钥ID" @change="search" @clear="search" />
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" icon="el-icon-search" type="primary" @click="search">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" icon="el-icon-plus" type="warning" @click="create">新增</el-button>
        </el-form-item>
        <el-form-item>
          <el-button :disabled="multipleSelection.length === 0" :loading="loading" icon="el-icon-delete" type="danger" @click="batchDelete">批量删除</el-button>
        </el-form-item>
      </el-form>

      <el-table  :data="tableData" border  auto-fit-column   @selection-change="handleSelectionChange">
        <el-table-column type="selection"  align="center" />
        <el-table-column show-overflow-tooltip  prop="cloudName" label="云账户" />
        <el-table-column show-overflow-tooltip  prop="cloudType" label="类型" />
        <el-table-column show-overflow-tooltip  prop="secretId" label="秘钥ID" />
        <el-table-column show-overflow-tooltip  prop="secretKey" label="秘钥key" />
        <el-table-column show-overflow-tooltip auto-fit prop="remark" label="备注" />
        <el-table-column show-overflow-tooltip  prop="CreatedAt" label="创建时间" />
        <el-table-column show-overflow-tooltip  prop="UpdatedAt" label="更新时间" />
        <!-- <el-table-column show-overflow-tooltip sortable prop="category" label="实例ID" />
        <el-table-column show-overflow-tooltip sortable prop="method" label="请求方式" align="center">
          <template slot-scope="scope">
            <el-tag size="small" :type="scope.row.method | methodTagFilter" disable-transitions>{{ scope.row.method }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column show-overflow-tooltip sortable prop="creator" label="创建人" />
        <el-table-column show-overflow-tooltip sortable prop="remark" label="说明" /> -->
        <el-table-column fixed="right" label="操作" align="center" width="120">
          <template slot-scope="scope">
            <el-tooltip content="编辑" effect="dark" placement="top">
              <el-button size="mini" icon="el-icon-edit" circle type="primary" @click="update(scope.row)" />
            </el-tooltip>
            <el-tooltip class="delete-popover" content="删除" effect="dark" placement="top">
              <el-popconfirm title="确定删除吗？" @onConfirm="singleDelete(scope.row.ID)">
                <el-button slot="reference" size="mini" icon="el-icon-delete" circle type="danger" />
              </el-popconfirm>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        :current-page="params.pageNum"
        :page-size="params.pageSize"
        :total="total"
        :page-sizes="[20, 50, 100, 300]"
        layout="total, prev, pager, next, sizes"
        background
        style="margin-top: 10px;float:right;margin-bottom: 10px;"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />

      <el-dialog :title="dialogFormTitle" :visible.sync="dialogFormVisible">
        <el-form ref="dialogForm" size="small" :model="dialogFormData" :rules="dialogFormRules" label-width="120px">
          <el-form-item label="云账号名称" prop="cloudName">
            <el-input v-model.trim="dialogFormData.cloudName" placeholder="云账号名称" />
          </el-form-item>
          <el-form-item label="云厂商" prop="cloudType">
            <el-select v-model.trim="dialogFormData.cloudType" placeholder="云厂商">
              <el-option label="腾讯云" value="腾讯云" />
              <el-option label="阿里云" value="阿里云" />
            </el-select>
          </el-form-item>
          <el-form-item label="秘钥ID" prop="secretId">
            <el-input v-model.trim="dialogFormData.secretId" placeholder="秘钥ID" show-word-limit maxlength="64" />
          </el-form-item>
          <el-form-item label="秘钥KEY" prop="secretKey">
            <el-input v-model.trim="dialogFormData.secretKey" placeholder="秘钥KEY" show-word-limit maxlength="64" />
          </el-form-item>
          <el-form-item label="说明" prop="remark">
            <el-input v-model.trim="dialogFormData.remark" type="textarea" placeholder="说明" show-word-limit maxlength="128" />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="mini" @click="cancelForm()">取 消</el-button>
          <el-button size="mini" :loading="submitLoading" type="primary" @click="submitForm()">确 定</el-button>
        </div>
      </el-dialog>

    </el-card>
  </div>
</template>

<script>
import { getCloudAccounts,createCloudAccount, updateCloudAccountById, batchDeleteCloudAccountByIds } from '@/api/example/cloudAccount'
import { Message } from 'element-ui'


export default {
  name: 'CloudAccount',
  filters: {
  },
  data() {
    return {
      // 查询参数
      params: {
        cloudName: '',
        cloudType: '',
        secretId: '',
        CreatedAt: '',
        UpdatedAt: '',
        pageNum: 1,
        pageSize: 10
      },
      // 表格数据
      tableData: [],
      total: 0,
      loading: false,

      // dialog对话框
      submitLoading: false,
      dialogFormTitle: '',
      dialogType: '',
      dialogFormVisible: false,
      dialogFormData: {
        ID: '',
        cloudName: '',
        cloudType: '',
        secretId: '',
        secretKey: '',
        remark: ''
      },
      dialogFormRules: {
        path: [
          { required: true, message: '请输入云账号', trigger: 'blur' },
          { min: 1, max: 100, message: '长度在 1 到 100 个字符', trigger: 'blur' }
        ],
        category: [
          { required: true, message: '请输入实例ID', trigger: 'blur' },
          { min: 1, max: 50, message: '长度在 1 到 50 个字符', trigger: 'blur' }
        ],
        method: [
          { required: true, message: '请选择请求方式', trigger: 'change' }
        ],
        remark: [
          { required: false, message: '说明', trigger: 'blur' },
          { min: 0, max: 100, message: '长度在 0 到 100 个字符', trigger: 'blur' }
        ]
      },

      // 删除按钮弹出框
      popoverVisible: false,
      // 表格多选
      multipleSelection: []
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    // 查询
    search() {
      this.params.pageNum = 1
      this.getTableData()
    },

    // 获取表格数据
    async getTableData() {
      this.loading = true
      try {
        const { data } = await getCloudAccounts(this.params)
        this.tableData = data.cloudAccounts
        this.total = data.total
      } finally {
        this.loading = false
      }
    },

    // 新增
    create() {
      this.dialogFormTitle = '新增接口'
      this.dialogType = 'create'
      this.dialogFormVisible = true
    },

    // 修改
    update(row) {
      this.dialogFormData.ID = row.ID
      this.dialogFormData.cloudName = row.cloudName
      this.dialogFormData.cloudType = row.cloudType
      this.dialogFormData.secretId = row.secretId
      this.dialogFormData.secretKey = row.secretKey
      this.dialogFormData.remark = row.remark

      this.dialogFormTitle = '修改接口'
      this.dialogType = 'update'
      this.dialogFormVisible = true
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
    // 提交表单
    submitForm() {
      this.$refs['dialogForm'].validate(async valid => {
        if (valid) {
          this.submitLoading = true
          try {
            if (this.dialogType === 'create') {
              await createCloudAccount(this.dialogFormData).then(res =>{
                this.judgeResult(res)
              })
            } else {
              await updateCloudAccountById(this.dialogFormData).then(res =>{
                this.judgeResult(res)
              })
            }
          } finally {
            this.submitLoading = false
          }
          this.resetForm()
          this.getTableData()
        } else {
          Message({
            showClose: true,
            message: '表单校验失败',
            type: 'error'
          })
          return false
        }
      })
    },

    // 提交表单
    cancelForm() {
      this.resetForm()
    },

    resetForm() {
      this.dialogFormVisible = false
      this.$refs['dialogForm'].resetFields()
      this.dialogFormData = {
        cloudName: '',
        cloudType: '',
        secretId: '',
        secretKey: '',
        remark: ''
      }
    },

    // 批量删除
    batchDelete() {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async res => {
        this.loading = true
        const cloudAccountIds = []
        this.multipleSelection.forEach(x => {
          cloudAccountIds.push(x.ID)
        })
        try {
          await batchDeleteCloudAccountByIds({ cloudAccountIds: cloudAccountIds }).then(res =>{
            this.judgeResult(res)
          })
        } finally {
          this.loading = false
        }
        this.getTableData()
      }).catch(() => {
        Message({
          showClose: true,
          type: 'info',
          message: '已取消删除'
        })
      })
    },

    // 表格多选
    handleSelectionChange(val) {
      this.multipleSelection = val
    },

    // 单个删除
    async singleDelete(Id) {
      this.loading = true
      try {
        await batchDeleteCloudAccountByIds({ cloudAccountIds: [Id] }).then(res =>{
          this.judgeResult(res)
        })
      } finally {
        this.loading = false
      }

      this.getTableData()
    },

    // 分页
    handleSizeChange(val) {
      this.params.pageSize = val
      this.getTableData()
    },
    handleCurrentChange(val) {
      this.params.pageNum = val
      this.getTableData()
    }
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
