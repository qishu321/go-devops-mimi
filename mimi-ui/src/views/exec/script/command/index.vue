<template>
  <div class="command-exec-container">
    <div class="header">
      <h2>命令执行</h2>
      <div class="command-form">
        <el-form :model="executeForm" label-width="100px">
          <el-form-item label="命令名称" required>
            <el-input v-model.trim="executeForm.name" style="width: 400px;" />
          </el-form-item>
          <el-form-item label="执行描述">
            <el-input type="textarea" v-model="executeForm.desc" :rows="3" placeholder="请输入执行描述信息" style="width: 400px;"></el-input>
          </el-form-item>
          <el-form-item label="命令类型" required>
            <el-select v-model="executeForm.cmd_type" placeholder="请选择" @change="handleCmdTypeChange">
              <el-option label="命令" value="command"></el-option>
              <el-option label="脚本" value="script"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="脚本类型" required>
            <el-select v-model="executeForm.type" placeholder="请选择">
              <el-option label="bash" value="bash"></el-option>
              <el-option label="py" value="py"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="会话超时(秒)">
            <el-input-number v-model="executeForm.timeout" :min="10" :max="3600"></el-input-number>
          </el-form-item>
          <!-- 替换原有的执行主机选择部分 -->
          <el-form-item label="执行主机" required>
            <el-button 
              type="primary" 
              size="small" 
              @click="openHostSelectDialog">
              选择执行主机
            </el-button>
            <div v-if="executeForm.node_ids.length > 0" style="margin-top: 10px;">
              <el-tag 
                v-for="(nodeId, index) in executeForm.node_ids" 
                :key="index"
                closable
                @close="executeForm.node_ids.splice(index, 1)"
                style="margin-right: 5px; margin-bottom: 5px;">
                {{ getNodeName(nodeId) }}
              </el-tag>
            </div>
          </el-form-item>
          <el-form-item label="命令内容" required>
            <div class="editor-container">
              <textarea ref="editor"></textarea>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" :loading="loading" :disabled="loading" @click="handleExecute">执行</el-button>
            <el-button @click="resetForm">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    
    <!-- 主机选择弹窗 -->
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
    
    <!-- 执行结果抽屉 -->
    <el-drawer
      title="执行结果"
      :visible.sync="resultDrawerVisible"
      size="50%"
      :before-close="() => { this.resultDrawerVisible = false }">
      <div v-loading="loading" style="padding: 20px;">
        <el-form :model="executionResult" label-width="100px">
          <el-form-item label="任务名称">
            <el-input v-model="executionResult.name" readonly />
          </el-form-item>
          <el-form-item label="执行状态">
            <el-tag :type="executionResult.status === 1 ? 'success' : 'danger'">
              {{ executionResult.status === 1 ? '成功' : '失败' }}
            </el-tag>
          </el-form-item>
          <el-form-item label="执行耗时">
            <el-input v-model="executionResult.timeCost" readonly style="width: 400px;">
              <template slot="append">ms</template>
            </el-input>
          </el-form-item>
        </el-form>
        
        <el-divider></el-divider>
        
        <h3>执行日志</h3>
        <el-table :data="executionResult.t_script_log_s" border style="margin-top: 20px;">
          <el-table-column prop="node_name" label="主机名称" width="120"></el-table-column>
          <el-table-column prop="status" label="状态" width="80">
            <template slot-scope="scope">
              <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
                {{ scope.row.status === 1 ? '成功' : '失败' }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="timeCost" label="耗时(ms)" width="100"></el-table-column>
          <el-table-column prop="run_log" label="执行结果">
            <template slot-scope="scope">
              <pre style="margin: 0; white-space: pre-wrap;">{{ scope.row.run_log }}</pre>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-drawer>
  </div>
</template>

<script>
import CodeMirror from 'codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/darcula.css'
import 'codemirror/mode/shell/shell.js'
import 'codemirror/mode/python/python.js'
import 'codemirror/mode/powershell/powershell.js'
import { addRun,listLibrary } from '@/api/exec/script'
import { listNode } from '@/api/cmdb/node'
import { list_node_group } from '@/api/cmdb/node_group'

export default {
  name: 'CommandExec',
  data() {
    return {
      groups: {},
      executeForm: {
        name: '临时命令',
        cmd_type: 'command',
        type: 'bash',
        command: '',
        timeout: 30,
        desc: '',
        node_ids: []
      },
      editor: null,
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
      scriptDialogVisible: false,
      scriptList: [],
      scriptLoading: false,
      resultDrawerVisible: false,
      executionResult: {},
      loading: false
    }
  },
  mounted() {
    this.initEditor()
    this.getGroups()
    this.getNodes()
  },
  // 将 watch 移到 methods 外部
  watch: {
    hostSelectMode(newVal) {
      if (newVal === 'group' && this.groupTableData.length === 0) {
        this.loadGroupTableData()
      }
    },
    'executeForm.type'(newVal) {
      if (this.editor) {
        this.editor.setOption('mode', newVal)
      }
    }
  },
  
  methods: {
    initEditor() {
      this.editor = CodeMirror.fromTextArea(this.$refs.editor, {
        mode: this.executeForm.type,
        theme: 'darcula',
        lineNumbers: true,
        matchBrackets: true,
        autoCloseBrackets: true
      })
      
      this.editor.on('change', (cm) => {
        this.executeForm.command = cm.getValue()
      })
    },
            async getNodes() {
              try {
                await listNode()
              } catch (error) {
                Message({
                  message: '获取节点列表失败',
                  type: 'error'
                })
              }
            },
        async getGroups() {
          try {
            const res = await list_node_group()
            this.groups = res.data
          } catch (error) {
            Message({
              message: '获取分组失败',
              type: 'error'
            })
          }
        },
    
        async handleExecute() {
          let loadingInstance = null;
          try {
            if (!this.executeForm.node_ids.length) {
              this.$message.warning('请选择执行主机')
              return
            }
            if (!this.executeForm.command.trim()) {
              this.$message.warning('请输入命令内容')
              return
            }
        
            this.$message.success('执行请求已提交')
        
            loadingInstance = this.$loading({
              lock: true,
              text: '命令执行中...',
              spinner: 'el-icon-loading',
              background: 'rgba(0, 0, 0, 0.7)'
            });
        
            this.loading = true
        
            const res = await addRun(this.executeForm)
            this.executionResult = res.data
            this.resultDrawerVisible = true
          } catch (error) {
            console.error('执行失败:', error)
            this.$message.error('执行失败')
          } finally {
            if (loadingInstance) {
              loadingInstance.close()
            }
            this.loading = false
          }
        },
            resetForm() {
      this.executeForm = {
        name: '临时命令',
        cmd_type: 'command',
        type: 'bash',
        command: '',
        timeout: 30,
        desc: '',
        node_ids: []
      }
      this.editor.setValue('')
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
        handleCmdTypeChange(val) {
          if (val === 'script') {
            this.loadScriptList();
            this.scriptDialogVisible = true;
          }
        },
        
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
          this.executeForm.command = row.content; // 假设接口返回的脚本内容字段是content
          this.editor.setValue(row.content);
          this.scriptDialogVisible = false;
          this.executeForm.type = row.type; // 自动设置脚本类型
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
      this.hostSearchParams.pageNum = 1 // 切换每页条数时重置页码
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
    
    getNodeName(nodeId) {
      // 先在主机列表中查找
      const host = this.hostTableData.find(item => item.ID === nodeId);
      if (host) return `${host.nodeName} (${host.publicIP})`;
      
      // 再在主机组的主机中查找
      for (const group of this.groupTableData) {
        if (group.t_node_group_s) {
          const node = group.t_node_group_s.find(item => item.ID === nodeId);
          if (node) return `${node.nodeName} (${node.publicIP})`;
        }
      }
      
      // 如果都没找到，返回ID
      return `ID: ${nodeId}`;
    },
    
    confirmHostSelection() {
      try {
        // 合并主机和主机组中的主机ID
        const hostIds = this.selectedHosts.map(host => host.ID);
        const groupHostIds = this.selectedGroups.flatMap(group => 
          group.t_node_group_s ? group.t_node_group_s.map(node => node.ID) : []
        );
        
        // 去重
        this.executeForm.node_ids = [...new Set([...hostIds, ...groupHostIds])];
        this.hostSelectDialogVisible = false;
      } catch (error) {
        console.error('确认选择出错:', error);
        this.$message.error('确认选择失败');
      }
    }
    
  }
}


</script>

<style scoped>
.command-exec-container {
  padding: 20px;
}
.header {
  margin-bottom: 20px;
}
.command-form {
  margin-top: 20px;
  padding: 15px;
  background: #f5f7fa;
  border-radius: 4px;
}
.editor-container {
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  overflow: hidden;
}
.editor-container textarea {
  width: 100%;
  min-height: 300px;
}
.host-select-tabs {
  margin-bottom: 20px;
}
.host-select-content, .group-select-content {
  margin-top: 10px;
}
</style>