<template>
    <div class="app-container">
      <el-card shadow="always">
        <!-- 查询 & 操作按钮区域 -->
        <el-form :inline="true" :model="params" size="mini">
          <el-form-item label="任务名称">
            <el-input
              v-model.trim="params.name"
              placeholder="请输入任务名称"
              clearable
              @clear="search"
            />
          </el-form-item>
          <el-form-item>
            <el-button type="primary" icon="el-icon-search" @click="search">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="warning" icon="el-icon-plus" @click="create">新增</el-button>
          </el-form-item>
          <el-form-item>
            <el-button
              type="danger"
              icon="el-icon-delete"
              :disabled="multipleSelection.length === 0"
              @click="batchDelete"
            >批量删除</el-button>
          </el-form-item>
        </el-form>
  
        <!-- 表格区域 -->
        <el-table
          :data="tableData"
          border
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column prop="name"    label="任务名称" min-width="120" />
          <el-table-column prop="desc"    label="描述"     min-width="150" show-overflow-tooltip />
          <el-table-column prop="cronType" label="任务类型" width="100" align="center">
            <template slot-scope="scope">
              <el-tag :type="getCronTypeTag(scope.row.cronType)">
                {{ scope.row.cronType | cronTypeFilter }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="cronession" label="执行规则" width="150" align="center" />
          <el-table-column prop="status" label="状态" width="100" align="center">
            <template slot-scope="scope">
              <el-switch
                v-model="scope.row.enable"
                :active-value="1"
                :inactive-value="0"
                @change="handleStatusChange(scope.row)"
              />
            </template>
          </el-table-column>
          <el-table-column prop="createdAt" label="创建时间" width="160" />
          <el-table-column label="操作" width="180" align="center" fixed="right">
            <template slot-scope="scope">
              <el-button
                size="mini"
                type="primary"
                icon="el-icon-edit"
                @click="update(scope.row)"
              />
              <el-popconfirm title="确认删除该任务？" @confirm="singleDelete(scope.row.id)">
                <el-button slot="reference" size="mini" type="danger" icon="el-icon-delete" />
              </el-popconfirm>
            </template>
          </el-table-column>
        </el-table>
  
        <!-- 分页 -->
        <el-pagination
          :current-page="params.pageNum"
          :page-size="params.pageSize"
          :total="total"
          :page-sizes="[20, 50, 100, 300]"
          layout="total, prev, pager, next, sizes"
          background
          style="margin-top: 10px; float: right; margin-bottom: 10px;"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </el-card>
  
      <!-- 新增/编辑抽屉 -->
      <el-drawer
        :title="dialogFormTitle"
        :visible.sync="dialogFormVisible"
        size="80%"
        direction="rtl"
        custom-class="drawer-form"
        :before-close="handleCloseDrawer"
      >  
      <div class="drawer-scroll">
        <el-form
          ref="dialogForm"
          :model="dialogFormData"
          :rules="dialogFormRules"
          label-width="100px"
          style="padding-right: 20px;"
        >
          <el-form-item label="任务名称" prop="name">
            <el-input v-model.trim="dialogFormData.name" placeholder="请输入任务名称" />
          </el-form-item>
  
          <el-form-item label="任务描述" prop="desc">
            <el-input
              v-model.trim="dialogFormData.desc"
              type="textarea"
              :rows="2"
              placeholder="请输入任务描述"
            />
          </el-form-item>
  
          <el-form-item label="任务类型" prop="cronType">
            <el-radio-group
              v-model="dialogFormData.cronType"
              @change="handleCronTypeChange"
            >
              <el-radio-button label="interval">间隔执行</el-radio-button>
              <el-radio-button label="once">单次执行</el-radio-button>
              <el-radio-button label="cron">Cron表达式</el-radio-button>
            </el-radio-group>
          </el-form-item>
  
          <el-form-item
            v-if="dialogFormData.cronType === 'interval'"
            label="间隔秒数"
            prop="interval"
          >
            <el-input-number v-model.number="dialogFormData.interval" :min="1" />
          </el-form-item>
  
          <el-form-item
            v-if="dialogFormData.cronType === 'once'"
            label="执行时间"
            prop="onceTime"
          >
            <el-date-picker
              v-model="dialogFormData.onceTime"
              type="datetime"
              placeholder="选择执行时间"
              value-format="yyyy-MM-dd HH:mm:ss"
            />
          </el-form-item>
  
          <el-form-item
            v-if="dialogFormData.cronType === 'cron'"
            label="Cron表达式"
            prop="cronession"
          >
            <el-input
              v-model.trim="dialogFormData.cronession"
              placeholder="请输入Cron表达式"
            />
          </el-form-item>
  
          <el-form-item label="命令类型" prop="cmdType">
            <el-radio-group v-model="dialogFormData.cmdType">
              <el-radio-button label="command">命令</el-radio-button>
              <el-radio-button label="script">脚本</el-radio-button>
            </el-radio-group>
          </el-form-item>
  
          <el-form-item label="脚本类型" prop="type">
            <el-radio-group v-model="dialogFormData.type">
              <el-radio-button label="shell">Shell</el-radio-button>
              <el-radio-button label="python">Python</el-radio-button>
            </el-radio-group>
          </el-form-item>
  
          <!-- 这里用 CodeMirror 替代 el-input -->
          <el-form-item label="内容" prop="content">
            <div ref="codeEditor" class="code-editor-container"></div>
          </el-form-item>
  
          <el-form-item label="超时时间(秒)" prop="timeout">
            <el-input-number v-model.number="dialogFormData.timeout" :min="1" />
          </el-form-item>
  
          <el-form-item label="执行主机" prop="nodeIds">
            <el-select
              v-model="dialogFormData.nodeIds"
              multiple
              placeholder="请选择执行主机"
              style="width: 100%;"
            >
              <el-option
                v-for="item in nodeOptions"
                :key="item.id"
                :label="item.name"
                :value="item.id"
              />
            </el-select>
          </el-form-item>
  
          <el-form-item>
            <el-button @click="dialogFormVisible = false">取 消</el-button>
            <el-button type="primary" @click="submitForm">确 定</el-button>
          </el-form-item>
        </el-form>
      </div>
      </el-drawer>
      <el-dialog 
      title="选择执行主机" 
      :visible.sync="hostSelectDialogVisible" 
      width="70%"
      top="5vh">
      <div class="host-select-tabs">
        <el-radio-group v-model="hostSelectMode" size="small">
          <el-radio-button label="host">选择主机</el-radio-button>
          <el-radio-button label="group">选择主机组</el-radio-button>
        </el-radio-group>
      </div>
      
      <!-- 主机选择模式 -->
      <div v-if="hostSelectMode === 'host'" class="host-select-content">
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
        
        <!-- 更新分页组件 -->
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
      
      <!-- 主机组选择模式 -->
      <div v-else class="group-select-content">
        <el-table
          :data="groupTableData"
          border
          height="400"
          @selection-change="handleGroupSelectionChange">
          <el-table-column type="selection" width="55"></el-table-column>
          <el-table-column prop="groupName" label="主机组名称"></el-table-column>
          <el-table-column prop="nodeCount" label="主机数量"></el-table-column>
        </el-table>
      </div>
      
      <div slot="footer">
        <el-button @click="hostSelectDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmHostSelection">确定</el-button>
      </div>
    </el-dialog>
    <!-- 脚本选择弹窗 -->
    <el-dialog
      title="选择脚本"
      :visible.sync="scriptDialogVisible"
      width="50%">
      <el-table
        :data="scriptList"
        border
        height="400"
        v-loading="scriptLoading"
        @row-click="handleScriptSelect">
        <el-table-column prop="name" label="脚本名称"></el-table-column>
        <el-table-column prop="type" label="脚本类型"></el-table-column>
      </el-table>
      
      <div slot="footer">
        <el-button @click="scriptDialogVisible = false">取消</el-button>
      </div>
    </el-dialog>
    </div>
  </template>
  
  <script>
  import { cronList, cronAdd, cronUpdate, cronDelete, cronEnable, cronInfo } from '@/api/exec/cron'
  import CodeMirror from 'codemirror'
  import 'codemirror/lib/codemirror.css'
  import 'codemirror/theme/darcula.css'
  import 'codemirror/mode/shell/shell.js'
  import 'codemirror/mode/python/python.js'
  import 'codemirror/mode/powershell/powershell.js'
  import { listLibrary } from '@/api/exec/script'
import { listNode } from '@/api/cmdb/node'
import { list_node_group } from '@/api/cmdb/node_group'

  export default {
    name: 'CronList',
    data() {
      return {
        params: { name: '', pageNum: 1, pageSize: 20 },
        tableData: [], total: 0, multipleSelection: [],
        dialogFormVisible: false, dialogFormTitle: '',
        dialogType: '',
        dialogFormData: {
          id: '', name: '', desc: '',
          cronType: 'interval', cronession: '', interval: 60, onceTime: '',
          cmdType: 'command', type: 'shell', content: '',
          timeout: 30, nodeIds: []
        },
        dialogFormRules: {
          name:      [{ required: true, message: '请输入任务名称', trigger: 'blur' }],
          cronType:  [{ required: true, message: '请选择任务类型', trigger: 'change' }],
          cronession:[{ required: true, message: '请输入Cron表达式', trigger: 'blur' }],
          interval:  [
            { required: true, message: '请输入间隔秒数', trigger: 'blur' },
            { type: 'number', min: 1, message: '间隔秒数必须大于0', trigger: 'blur' }
          ],
          onceTime:  [{ required: true, message: '请选择执行时间', trigger: 'change' }]
        },
        codeEditor: null,
        nodeOptions: [] // 假设你会在 created()/mounted() 填充主机列表
      }
    },
    filters: {
      cronTypeFilter(type) {
        const map = { interval: '间隔执行', once: '单次执行', cron: 'Cron表达式' }
        return map[type] || type
      }
    },
    created() {
      this.getTableData()
      // TODO: 拉取 this.nodeOptions
    },
    methods: {
      getCronTypeTag(type) {
        return { interval: 'success', once: 'warning', cron: 'primary' }[type] || ''
      },
      handleCronTypeChange() {
        // 切换任务类型时清空相关字段
        this.dialogFormData.cronession = ''
        this.dialogFormData.interval   = 60
        this.dialogFormData.onceTime   = ''
      },
      search() {
        this.params.pageNum = 1
        this.getTableData()
      },
      async getTableData() {
        const { data } = await cronList(this.params)
        this.tableData = data.list || []
        this.total     = data.total || 0
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      handleSizeChange(val) {
        this.params.pageSize = val
        this.getTableData()
      },
      handleCurrentChange(val) {
        this.params.pageNum = val
        this.getTableData()
      },
      create() {
        this.dialogFormTitle   = '新增定时任务'
        this.dialogType        = 'create'
        this.dialogFormVisible = true
        this.$nextTick(() => {
          this.$refs.dialogForm.resetFields()
        })
      },
      async update(row) {
        const { data } = await cronInfo({ id: row.id })
        Object.assign(this.dialogFormData, {
          id: data.id,
          name: data.name,
          desc: data.desc,
          cronType: data.cronType,
          cronession: data.cronession,
          interval: data.interval,
          onceTime: data.onceTime,
          cmdType: data.cmdType,
          type: data.type,
          content: data.content,
          timeout: data.timeout,
          nodeIds: data.nodeIds
        })
        this.dialogFormTitle   = '编辑定时任务'
        this.dialogType        = 'update'
        this.dialogFormVisible = true
      },
      async handleStatusChange(row) {
        try {
          await cronEnable({ id: row.id, enable: row.enable })
          this.$message.success('状态更新成功')
        } catch {
          row.enable = row.enable === 1 ? 0 : 1
        }
      },
      submitForm() {
        this.$refs.dialogForm.validate(async valid => {
          if (!valid) return
          try {
            if (this.dialogType === 'create') {
              await cronAdd(this.dialogFormData)
              this.$message.success('新增成功')
            } else {
              await cronUpdate(this.dialogFormData)
              this.$message.success('更新成功')
            }
            this.dialogFormVisible = false
            this.getTableData()
          } catch (err) {
            console.error(err)
          }
        })
      },
      singleDelete(id) {
        this.$confirm('确认删除该任务？', '提示', { type: 'warning' })
          .then(async () => {
            await cronDelete({ ids: [id] })
            this.$message.success('删除成功')
            this.getTableData()
          })
      },
      batchDelete() {
        if (!this.multipleSelection.length) {
          return this.$message.warning('请至少选择一条记录')
        }
        this.$confirm(
          `确认删除选中的 ${this.multipleSelection.length} 条任务？`,
          '提示',
          { type: 'warning' }
        ).then(async () => {
          await cronDelete({ ids: this.multipleSelection.map(i => i.id) })
          this.$message.success('删除成功')
          this.getTableData()
        })
      },
      handleCloseDrawer(done) {
        // 清理编辑器实例
        if (this.codeEditor) {
          this.codeEditor.toTextArea && this.codeEditor.toTextArea()
          this.codeEditor = null
        }
        done()
      },
      // 初始化 CodeMirror 编辑器
      initCodeMirror() {
        this.codeEditor = CodeMirror(this.$refs.codeEditor, {
          value: this.dialogFormData.content || '',
          mode: this.dialogFormData.type === 'python' ? 'python' : 'shell',
          theme: 'darcula',
          lineNumbers: true,
          tabSize: 2
        })
        this.codeEditor.on('change', () => {
          this.dialogFormData.content = this.codeEditor.getValue()
        })
      }
    },
    watch: {
      // 当抽屉打开时，初始化或重置 CodeMirror
      dialogFormVisible(val) {
        if (val) {
          this.$nextTick(() => {
            if (!this.codeEditor) {
              this.initCodeMirror()
            } else {
              this.codeEditor.setValue(this.dialogFormData.content || '')
            }
          })
        }
      },
      // 切换脚本类型时，更新编辑器模式
      'dialogFormData.type'(newType) {
        if (this.codeEditor) {
          const mode = newType === 'python' ? 'python' : 'shell'
          this.codeEditor.setOption('mode', mode)
        }
      }
    }
  }
  </script>
  
  <style scoped>
  .app-container {
    padding: 20px;
  }
  .drawer-form .el-drawer__body {
    padding: 0; /* 可根据需要调整 */
  }
  .drawer-scroll {
    height: calc(100vh - 80px);
    overflow-y: auto;
    padding: 20px;
  }  
  .code-editor-container {
    height: 300px;
    border: 1px solid #ddd;
  }
  </style>
  