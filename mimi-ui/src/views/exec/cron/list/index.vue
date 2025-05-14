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
          <el-table-column label="任务执行规则" width="200" align="center">
            <template slot-scope="scope">
              <!-- 根据不同类型，展示规则 -->
              <span style="margin-left: 8px; color: #606266;">
                <template v-if="scope.row.cronType === 'interval'">
                  {{ scope.row.interval }}秒
                </template>
                <template v-else-if="scope.row.cronType === 'once'">
                  {{ formatTime(scope.row.onceTime) }}
                </template>
                <template v-else>
                  {{ scope.row.cronession }}
                </template>
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="enable" width="120" label="是否启用" align="center">
            <template #default="{ row }">
              <el-tag :type="row.enable === 1 ? 'success' : 'danger'">
                {{ row.enable === 1 ? '已启用' : '未启用' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="CreatedAt" label="创建时间" width="180" align="center">
            <template #default="{ row }">
              {{ formatTime(row.CreatedAt) }}
            </template>
          </el-table-column>    
          <el-table-column label="操作" width="360" align="center" fixed="right">
            <template #default="{ row }">
              <!-- 根据当前 enable 值，展示不同按钮 -->
              <el-button
                size="mini"
                :type="row.enable === 0 ? 'success' : 'warning'"
                :icon="row.enable === 0 ? 'el-icon-cpu' : 'el-icon-close'"
                @click="toggleEnable(row)"
              >
                {{ row.enable === 0 ? '启用定时任务' : '关闭执行任务' }}
              </el-button>
          
              <el-button
              size="mini"
              type="primary"
              icon="el-icon-edit"
              @click="row.enable === 1 ? $message.error('已启用的定时任务，无法编辑，请关闭任务后再编辑') : update(row)"
            >编辑</el-button>          
              <el-popconfirm
                title="确认删除该任务？"
                @confirm="singleDelete(row.id)"
              >
                <el-button
                  slot="reference"
                  size="mini"
                  type="danger"
                  icon="el-icon-delete"
                >删除</el-button>
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
              v-model="dialogFormData.cronession" style="width: 240px"
              placeholder="请输入Cron表达式"
            />
          </el-form-item>
  
          <el-form-item label="命令类型"  prop="cmd_type" required>
            <el-select v-model="dialogFormData.cmd_type" placeholder="请选择" @change="handlecmd_typeChange">
              <el-option label="命令" value="command"></el-option>
              <el-option label="脚本" value="script"></el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="脚本类型" prop="type">
            <el-radio-group v-model="dialogFormData.type">
              <el-radio-button label="bash">Shell</el-radio-button>
              <el-radio-button label="py">Python</el-radio-button>
            </el-radio-group>
          </el-form-item>
  
          <!-- 这里用 CodeMirror 替代 el-input -->
          <el-form-item label="内容" prop="content" required>
            <div ref="codeEditor" class="code-editor-container"></div>
          </el-form-item>
  
          <el-form-item label="超时时间(秒)" prop="timeout">
            <el-input-number v-model.number="dialogFormData.timeout" :min="1" />
          </el-form-item>
  
          <el-form-item label="执行主机" prop="node_ids" required>
            <el-button 
              type="primary" 
              size="small" 
              @click="openHostSelectDialog">
              选择执行主机
            </el-button>
            <div v-if="dialogFormData.node_ids.length > 0" style="margin-top: 10px;">
              <el-tag 
                v-for="(nodeId, index) in dialogFormData.node_ids" 
                :key="index"
                closable
                @close="dialogFormData.node_ids.splice(index, 1)"
                style="margin-right: 5px; margin-bottom: 5px;">
                {{ getNodeName(nodeId) }}
              </el-tag>
            </div>
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
    </el-dialog>    <!-- 脚本选择弹窗 -->
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
import dayjs from 'dayjs'

  export default {
    name: 'CronList',
    data() {
      return {
        params: { name: '', pageNum: 1, pageSize: 20 },
        tableData: [], total: 0, multipleSelection: [],
        dialogFormVisible: false, dialogFormTitle: '',
        scriptDialogVisible: false,
        scriptList: [],
        scriptLoading: false,  
        hostSelectDialogVisible: false,
        hostSelectMode: 'host',
        hostSearchParams: {
          nodeName: '',
          pageNum: 1,
          pageSize: 10
        },
        hostTableData: [],
        hostTotal: 0,
        groupTableData: [],
        selectedHosts: [],
        selectedGroups: [],
      
        dialogType: '',
        dialogFormData: {
          id: '', name: '', desc: '',
          cronType: 'interval', cronession: '', interval: 60, onceTime: '',
          cmd_type: 'command', type: 'shell', content: '',
          timeout: 30, node_ids: []
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
    mounted() {
        this.loadGroupTableData()
        this.loadHostTableData()
      },
    
    methods: {
        async loadScriptList() {
            this.scriptLoading = true;
            try {
              const res = await listLibrary({});
              this.scriptList = res.data.script_library_s || [];
            } catch (error) {
              console.error('加载脚本列表失败:', error);
              this.$message.error('加载脚本列表失败');
            } finally {
              this.scriptLoading = false;
            }
          },
          
          handleScriptSelect(row) {
            this.dialogFormData.content = row.content; // 修改为设置dialogFormData.content
            if (this.codeEditor) {
                this.codeEditor.setValue(row.content);
            }
            this.scriptDialogVisible = false;
            this.dialogFormData.type = row.type; // 自动设置脚本类型
        },
                  handlecmd_typeChange(val) {
            if (val === 'script') {
              this.loadScriptList();
              this.scriptDialogVisible = true;
            }
          },
          getNodeName(nodeId) {
            const host = this.hostTableData.find(item => item.ID === nodeId);
            if (host) return `${host.nodeName} (${host.publicIP})`;
            
            for (const group of this.groupTableData) {
              if (group.t_node_group_s) {
                const node = group.t_node_group_s.find(item => item.ID === nodeId);
                if (node) return `${node.nodeName} (${node.publicIP})`;
              }
            }
            
            return `ID: ${nodeId}`;
          },
          openHostSelectDialog() {
            this.hostSelectDialogVisible = true;
            if (this.hostSelectMode === 'host') {
              this.searchHosts();
            } else {
              // 切到“主机组”就给它赋值
              this.groupTableData = this.groups.nodeGroupLists || [];
            }
            },
                
          async loadHostTableData() {
            try {
              const res = await listNode({
                nodeName: this.hostSearchParams.nodeName,
                pageNum: this.hostSearchParams.pageNum,
                pageSize: this.hostSearchParams.pageSize
              })
              this.hostTableData = res.data.nodeLists || []
              this.hostTotal = res.data.total || 0
            } catch (error) {
              console.error('加载主机列表失败:', error)
              this.$message.error('加载主机列表失败')
            }
          },
          
          async loadGroupTableData() {
            try {
              const res = await list_node_group()
              this.groupTableData = res.data.nodeGroupLists || []
            } catch (error) {
              console.error('加载主机组列表失败:', error)
              this.$message.error('加载主机组列表失败')
            }
          },
          
          searchHosts() {
            this.hostSearchParams.pageNum = 1
            this.loadHostTableData()
          },
          
          handleHostSizeChange(val) {
            this.hostSearchParams.pageSize = val
            this.hostSearchParams.pageNum = 1
            this.loadHostTableData()
          },
          
          handleHostPageChange(page) {
            this.hostSearchParams.pageNum = page
            this.loadHostTableData()
          },
          
          handleHostSelectionChange(selection) {
            this.selectedHosts = selection
          },
          
          handleGroupSelectionChange(selection) {
            this.selectedGroups = selection
          },
          
          confirmHostSelection() {
            try {
              const hostIds = this.selectedHosts.map(host => host.ID);
              const groupHostIds = this.selectedGroups.flatMap(group => 
                group.t_node_group_s ? group.t_node_group_s.map(node => node.ID) : []
              );
              
              this.dialogFormData.node_ids = [...new Set([...hostIds, ...groupHostIds])];
              this.hostSelectDialogVisible = false;
            } catch (error) {
              console.error('确认选择出错:', error);
              this.$message.error('确认选择失败');
            }
          },
        
  
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
        this.tableData = data.cron_s || []
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
      formatTime(time) {
        return dayjs(time).format('YYYY-MM-DD HH:mm:ss')
      },
   
      async update(row) {
        const { data } = await cronInfo({ id: row.ID })
        Object.assign(this.dialogFormData, {
          id: data.ID,
          name: data.name,
          desc: data.desc,
          cronType: data.cronType,
          cronession: data.cronession,
          interval: data.interval,
          onceTime: data.onceTime,
          cmd_type: data.cmd_type,
          type: data.type,
          content: data.content,
          timeout: data.timeout,
          node_ids: typeof data.node_ids === 'string'
          ? JSON.parse(data.node_ids)
          : data.node_ids
                })
        this.dialogFormTitle   = '编辑定时任务'
        this.dialogType        = 'update'
        this.dialogFormVisible = true
      },
  // 切换启用/关闭
  async toggleEnable(row) {
    try {
      // 后端接口里：enable:1 表示启用；enable:0 表示关闭
      const newEnable = row.enable === 0 ? 1 : 0;
      await cronEnable({ id: row.ID, enable: newEnable });
      // 前端状态同步
      row.enable = newEnable;
      this.$message.success(newEnable === 1 ? '任务已启用' : '任务已关闭');
    } catch (err) {
      this.$message.error('操作失败');
      console.error(err);
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

  