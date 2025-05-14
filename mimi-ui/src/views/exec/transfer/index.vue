<template>
  <div class="app-container">
    <el-row :gutter="20">
      <el-col :span="10">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>文件分发</span>
          </div>
          <el-form ref="form" :model="form" :rules="rules" label-width="100px">
            <el-form-item label="任务名称" prop="name">
              <el-input v-model="form.name" placeholder="请输入分发任务名称" />
            </el-form-item>
            <el-form-item label="源文件" prop="source_path">
              <el-input v-model="form.source_path" placeholder="请输入文件路径或上传文件" style="width: calc(100% - 120px)" />
              <el-upload
                class="upload-demo"
                :http-request="uploadFile"
                :before-upload="beforeUpload"
                :show-file-list="false"
                action=""
                style="display: inline-block; margin-left: 10px;"
              >
                <el-button size="small" type="primary">上传文件</el-button>
              </el-upload>
            </el-form-item>
            <el-form-item label="目标路径" prop="target_path">
              <el-input v-model="form.target_path" placeholder="请输入目标主机路径" />
            </el-form-item>
            <el-form-item label="目标主机" prop="node_ids">
              <el-button 
                type="primary" 
                size="small" 
                @click="openHostSelectDialog">
                选择目标主机
              </el-button>
              <div v-if="form.node_ids.length > 0" style="margin-top: 10px;">
                <el-tag 
                  v-for="(nodeId, index) in form.node_ids" 
                  :key="index"
                  closable
                  @close="form.node_ids.splice(index, 1)"
                  style="margin-right: 5px; margin-bottom: 5px;">
                  {{ getNodeName(nodeId) }}
                </el-tag>
              </div>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm">立即分发</el-button>
              <el-button @click="resetForm">重置</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="14">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>分发记录</span>
            <el-button style="float: right; padding: 3px 0" type="text" @click="getList">刷新</el-button>
          </div>
          <el-table
            v-loading="listLoading"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%"
          >
            <el-table-column prop="name" label="任务名称" align="center" />
            <el-table-column prop="source_path" label="源文件" align="center" />
            <el-table-column prop="target_path" label="目标路径" align="center" />
            <el-table-column prop="status" label="状态" align="center">
              <template #default="{row}">
                <el-tag :type="row.status === 1 ? 'success' : 'danger'">
                  {{ row.status === 1 ? '成功' : '失败' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="creator" label="执行用户" align="center" />
            <el-table-column label="操作" width="100" align="center">
              <template #default="{row}">
                <el-button 
                  type="text" 
                  size="small" 
                  @click="showRunLog(row)">
                  查看执行结果
                </el-button>
              </template>
            </el-table-column>        
          </el-table>
          <pagination
            v-show="total>0"
            :total="total"
            :page.sync="listQuery.pageNum"
            :limit.sync="listQuery.pageSize"
            @pagination="getList"
          />
        </el-card>
      </el-col>
    </el-row>
    
    <!-- 主机选择弹窗 -->
    <el-dialog 
      title="选择目标主机" 
      :visible.sync="hostSelectDialogVisible" 
      width="70%"
      top="5vh">
      <div class="host-select-content">
        <el-form :inline="true" size="small">
          <el-form-item>
            <el-input
              v-model="hostSearchParams.nodeName"
              placeholder="输入主机名称搜索"
              clearable
              @clear="searchHosts"
              @keyup.enter.native="searchHosts">
              <el-button slot="append" icon="el-icon-search" @click="searchHosts"></el-button>
            </el-input>
          </el-form-item>
        </el-form>
        
        <el-table
          :data="hostTableData"
          border
          height="400"
          @selection-change="handleHostSelectionChange">
          <el-table-column type="selection" width="55"></el-table-column>
          <el-table-column prop="nodeName" label="主机名称"></el-table-column>
          <el-table-column prop="publicIP" label="IP地址"></el-table-column>
          <el-table-column prop="username" label="系统用户"></el-table-column>
        </el-table>
        
        <el-pagination
          :current-page="hostSearchParams.pageNum"
          :page-size="hostSearchParams.pageSize"
          :total="hostTotal"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next"
          background
          style="margin-top: 10px;"
          @size-change="handleHostSizeChange"
          @current-change="handleHostPageChange"
        />
      </div>
      
      <div slot="footer">
        <el-button @click="hostSelectDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmHostSelection">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listTransfer, addRunTransfer, uploadTransfer } from '@/api/exec/script'
import { listNode } from '@/api/cmdb/node'
import Pagination from '@/components/Pagination'

export default {
  name: 'Transfer',
  components: { Pagination },
  data() {
    return {
      form: {
        name: '',
        source_path: '',
        target_path: '',
        node_ids: []
      },
      rules: {
        name: [
          { required: true, message: '请输入任务名称', trigger: 'blur' },
          { min: 1, max: 64, message: '长度在 1 到 64 个字符', trigger: 'blur' }
        ],
        source_path: [
          { required: true, message: '请上传文件', trigger: 'blur' }
        ],
        target_path: [
          { required: true, message: '请输入目标路径', trigger: 'blur' }
        ],
        node_ids: [
          { required: true, message: '请选择目标主机', trigger: 'change' }
        ]
      },
      hostSelectDialogVisible: false,
      hostSearchParams: {
        nodeName: '',
        pageNum: 1,
        pageSize: 10
      },
      hostTableData: [],
      hostTotal: 0,
      selectedHosts: [],
      nodeOptions: [],
      list: [],
      total: 0,
      listLoading: false,
      listQuery: {
        pageNum: 1,
        pageSize: 10
      }
    }
  },
  created() {
    this.getList()
    this.getNodeList()
  },
  methods: {
    getList() {
      this.listLoading = true
      listTransfer(this.listQuery).then(response => {
        this.list = response.data.transfer_s
        this.total = response.data.total
        this.listLoading = false
      }).catch(() => {
        this.listLoading = false
      })
    },
    getNodeList() {
      listNode().then(response => {
        this.nodeOptions = response.data.nodeLists
      })
    },
    beforeUpload(file) {
      const isLt10M = file.size / 1024 / 1024 < 10
      if (!isLt10M) {
        this.$message.error('上传文件大小不能超过 10MB!')
        return false
      }
      return true
    },
    handleUploadSuccess(response) {
        // 后端成功时 code === 0
        if (response.code === 0) {
          this.form.source_path = response.data
          this.$message({
            message: '上传成功',
            type: 'success'
          })
        } else {
          // code !== 0 时才弹 error
          this.$message.error(response.msg)
        }
      },
      uploadFile(file) {
        const formData = new FormData()
        formData.append('file', file.file)
        uploadTransfer(formData, {
          headers: { 'Content-Type': 'multipart/form-data' }
        })
          .then(response => {
            // 先交给统一处理函数判断
            this.handleUploadSuccess(response)
          })
          .catch(err => {
            // 网络或其它异常
            this.$message.error(err.message)
          })
        },
    submitForm() {
      this.$refs.form.validate(valid => {
        if (valid) {
          addRunTransfer(this.form).then(response => {
            this.$message.success('分发任务已提交')
            this.resetForm()
            this.getList()
          })
        }
      })
    },
    resetForm() {
      this.$refs.form.resetFields()
    },
    getNodeName(nodeId) {
      const node = this.nodeOptions.find(item => item.ID === nodeId)
      return node ? node.nodeName : ''
    },
    openHostSelectDialog() {
      this.hostSelectDialogVisible = true
      this.hostSearchParams = {
        nodeName: '',
        pageNum: 1,
        pageSize: 10
      }
      this.hostTableData = []
      this.hostTotal = 0
      this.searchHosts()
    },
    searchHosts() {
      listNode(this.hostSearchParams).then(response => {
        this.hostTableData = response.data.nodeLists
        this.hostTotal = response.data.total
      })
    },
    handleHostSelectionChange(selection) {
      this.selectedHosts = selection
    },
    handleHostSizeChange(val) {
      this.hostSearchParams.pageSize = val
      this.searchHosts()
    },
    handleHostPageChange(val) {
      this.hostSearchParams.pageNum = val
      this.searchHosts()
    },
    confirmHostSelection() {
        if (this.selectedHosts.length) {
          // 注意：这里用的是大写 `ID`，和后台返回字段一一对应
          this.form.node_ids = this.selectedHosts.map(h => h.ID)
        } else {
          this.form.node_ids = []
        }
        this.hostSelectDialogVisible = false
      },
      showRunLog(row) {
        this.$alert(row.run_log || '暂无执行日志', '执行日志', {
          confirmButtonText: '确定',
          customClass: 'run-log-dialog',
          showClose: true
        })
      }
    }
}
</script>

<style scoped>
.upload-demo {
  margin-bottom: 20px;
}
.box-card {
  margin-bottom: 20px;
}
</style>