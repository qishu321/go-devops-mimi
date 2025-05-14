<template>
    <div>
      <el-card class="container-card" shadow="always">
        <el-form
          ref="form"
          :model="form"
          :rules="rules"
          label-width="120px"
        >
          <!-- 上半部分：任务名称 & 执行参数 -->
          <el-row :gutter="20">
            <!-- 左半：任务名称 & 描述 -->
            <el-col :span="12">
              <el-form-item label="任务名称" prop="name">
                <el-input v-model="form.name" placeholder="请输入任务名称" />
              </el-form-item>
              <el-form-item label="描述信息" prop="desc">
                <el-input
                  v-model="form.desc"
                  type="textarea"
                  :rows="2"
                  placeholder="请输入描述信息"
                />
              </el-form-item>
            </el-col>
    
            <!-- 右半：执行参数 -->
            <el-col :span="12">
              <el-form-item label="执行参数">
                <el-button type="primary" size="small" @click="addParam">
                  添加参数
                </el-button>
                <el-table
                  :data="form.args"
                  border
                  size="small"
                  style="margin-top: 10px"
                >
                  <el-table-column label="参数名称" width="120">
                    <template #default="{ row }">
                      <el-input v-model="row.name" size="small" />
                    </template>
                  </el-table-column>
                  <el-table-column label="类型" width="100">
                    <template #default="{ row }">
                      <el-select v-model="row.type" size="small">
                        <el-option label="文本框" value="string" />
                        <el-option label="下拉选择" value="select" />
                      </el-select>
                    </template>
                  </el-table-column>
                  <el-table-column
                    v-if="form.args.some(i => i.type === 'select')"
                    label="选项值"
                  >
                    <template #default="{ row }">
                      <el-input
                        v-if="row.type === 'select'"
                        v-model="row.options"
                        type="textarea"
                        :rows="2"
                        placeholder="格式: value1:显示1\nvalue2:显示2"
                        size="small"
                      />
                    </template>
                  </el-table-column>
                  <el-table-column label="必填" width="80">
                    <template #default="{ row }">
                      <el-switch v-model="row.required" />
                    </template>
                  </el-table-column>
                  <el-table-column label="操作" width="80">
                    <template #default="{ $index }">
                      <el-button
                        type="danger"
                        size="mini"
                        @click="removeParam($index)"
                      >删除</el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </el-form-item>
            </el-col>
          </el-row>
    
          <!-- 子任务标题栏 -->
          <el-row
            :gutter="20"
            align="middle"
            style="margin-top:20px;"
          >
            <el-col :span="2">
              <span class="subtask-title">子任务</span>
            </el-col>
            <el-col :span="22">
              <el-button type="primary" size="small" @click="addSubTask">
                添加子任务
              </el-button>
            </el-col>
          </el-row>
    
          <!-- 子任务步骤 -->
          <el-collapse
            v-model="activeTasks"
            class="tasks-collapse"
          >
            <el-collapse-item
              v-for="(task, idx) in form.tasks"
              :key="task.tempId"
              :title="`步骤 ${idx + 1}：${task.name || '新建步骤'}`"
              :name="task.tempId"
            >
              <el-form-item :prop="`tasks.${idx}.name`" label="步骤名称">
                <el-input
                  v-model="task.name"
                  placeholder="请输入步骤名称"
                />
              </el-form-item>
    
              <el-form-item :prop="`tasks.${idx}.node_ids`" label="目标主机">
                <el-button
                  type="primary"
                  size="mini"
                  @click="openHostSelectDialog(idx)"
                >选择目标主机</el-button>
              
                <!-- 直接用 node_names 来渲染 tag -->
                <div v-if="task.node_names" style="margin-top: 10px;">
                  <el-tag
                    v-for="(name, i) in task.node_names.split(',')"
                    :key="i"
                    closable
                    @close="removeNodeId(idx, i)"
                    style="margin-right: 5px; margin-bottom: 5px;"
                  >
                    {{ name }}
                  </el-tag>
                </div>
              </el-form-item>
              <el-form-item label="脚本来源">
                <el-radio-group v-model="task.source">
                  <el-radio label="manual">手工输入</el-radio>
                  <el-radio label="repo">脚本仓库</el-radio>
                </el-radio-group>
              </el-form-item>
    
              <el-form-item
                v-if="task.source === 'repo'"
                label="选择脚本"
              >
                <el-select
                  v-model="task.scriptId"
                  placeholder="请选择脚本"
                  filterable
                  size="large"
                  @change="handleScriptSelect(idx)"
                >
                  <el-option
                    v-for="s in scriptList"
                    :key="s.id"
                    :label="s.name"
                    :value="s.id"
                  />
                </el-select>
              </el-form-item>
    
              <el-form-item label="脚本类型">
                <el-radio-group v-model="task.type">
                  <el-radio label="shell">Shell</el-radio>
                  <el-radio label="python">Python</el-radio>
                </el-radio-group>
              </el-form-item>
                    <!-- 新增：脚本超时时间 -->
                    <el-form-item :prop="`tasks.${idx}.timeout`" label="脚本超时时间(秒)">
                        <el-input-number
                        v-model="task.timeout"
                        :min="30" :max="3600"
                        size="small"
                        />
                    </el-form-item>
              <el-form-item label="脚本内容">
                <!-- 直接用原生 textarea 供 CodeMirror 挂载 -->
                <textarea :ref="`editor${idx}`" tyle="height: 200px; width: 100%;"></textarea>
              </el-form-item>    
              <el-button
                type="primary"
                size="small"
                @click="addSubTask"
              >继续添加子任务</el-button>
              <el-button
                type="danger"
                size="mini"
                icon="el-icon-delete"
                @click="removeSubTask(idx)"
              >删除本步</el-button>
            </el-collapse-item>
          </el-collapse>
    
          <!-- 主机/主机组选择对话框 -->
          <el-dialog
            title="选择执行主机"
            :visible.sync="hostSelectDialogVisible"
            width="70%"
            top="5vh"
          >
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
                    @keyup.enter.native="searchHosts"
                  >
                    <el-button slot="append" icon="el-icon-search" @click="searchHosts"></el-button>
                  </el-input>
                </el-form-item>
              </el-form>
              
              <el-table
                :data="hostTableData"
                border
                height="400"
                @selection-change="handleHostSelectionChange"
              >
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
                @selection-change="handleGroupSelectionChange"
              >
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
    
          <!-- 提交/取消 -->
          <el-form-item style="text-align: right; margin-top: 20px;">
            <el-button type="primary" @click="submitForm">提交</el-button>
            <el-button @click="cancel">取消</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </template>
  
  <script>
  import { addTaskManage } from '@/api/exec/task'
  import { listLibrary } from '@/api/exec/script'
  import { listNode } from '@/api/cmdb/node'
  import { list_node_group } from '@/api/cmdb/node_group'
  
  // CodeMirror 核心 & 模式
  import CodeMirror from 'codemirror'
  import 'codemirror/lib/codemirror.css'
  import 'codemirror/theme/darcula.css'
  import 'codemirror/mode/shell/shell.js'
  import 'codemirror/mode/python/python.js'
  // CodeMirror 插件
  import 'codemirror/addon/edit/matchbrackets.js'
  import 'codemirror/addon/edit/closebrackets.js'
  
  export default {
    name: 'TaskManageAdd',
    data() {
      return {
        form: {
          name: '',
          desc: '',
          args: [],
          tasks: []
        },
        nodeDialogVisible: false,
        currentTaskIndex: 0,
        selectedNodes: [],
        nodeList: [],
        nodeGroupList: [],
        groupTableData: [],
        activeTasks: [],
        scriptList: [],
        editors: [],   // 存放各步骤的 CodeMirror 实例
        rules: {
          name: [{ required: true, message: '请输入任务名称', trigger: 'blur' }]
          
        },
        hostSelectDialogVisible: false,
        hostSelectMode: 'host',
        hostSearchParams: {
          nodeName: '',
          pageNum: 1,
          pageSize: 10
        },
        hostTableData: [],
        hostTotal: 0,
        selectedHosts: [],
        selectedGroups: []
      }
    },
    mounted() {
      this.fetchScriptList()
      this.fetchNodeList()
      this.fetchNodeGroupList()
    },
    watch: {
      hostSelectMode(newVal) {
        if (newVal === 'group') {
          this.loadGroupTableData()
        }
      }
    },
    methods: {
      // —— 参数相关
      addParam() {
        this.form.args.push({ name: '', type: 'string', options: '', required: true })
      },
      removeParam(i) {
        this.form.args.splice(i, 1)
      },
      handleScriptSelect(index) {
        const task = this.form.tasks[index]
        const found = this.scriptList.find(s => s.id === task.scriptId)
        if (!found) {
          this.$message.error('未找到选中的脚本内容')
          return
        }
  
        // 1. 写到任务数据
        task.content = found.content
        // 2. 更新 CodeMirror 编辑器
        this.$nextTick(() => {
          const cm = this.editors[index]
          if (cm) cm.setValue(found.content)
        })
      },
      
      // —— 子任务相关
      addSubTask() {
        const tempId = Date.now()
        const idx = this.form.tasks.length
        this.form.tasks.push({
          tempId,
          name: '',
          source: 'manual',
          scriptId: null,
          type: 'shell',
          content: '',
          node_ids: '',
          node_names: '',
          timeout: 0            // 默认脚本超时时间（秒）
        })
        this.$nextTick(() => {
          this.activeTasks.push(tempId)
          this.initCodeEditor(idx)
        })
      },
      removeSubTask(i) {
        this.form.tasks.splice(i, 1)
        this.activeTasks.splice(i, 1)
        this.editors.splice(i, 1)
      },
  
      // —— 初始化单步 CodeMirror
      initCodeEditor(index) {
        let textarea = this.$refs[`editor${index}`]
        if (Array.isArray(textarea)) textarea = textarea[0]
        if (!textarea) return
  
        const cm = CodeMirror.fromTextArea(textarea, {
          mode: this.form.tasks[index].type,
          theme: 'darcula',
          lineNumbers: false,
          matchBrackets: true,
          autoCloseBrackets: true,
          indentUnit: 4
        })
        cm.on('change', () => {
          this.form.tasks[index].content = cm.getValue()
        })
        this.editors[index] = cm
      },
  
      // —— 主机选择
      openHostSelectDialog(idx) {
        this.currentTaskIndex = idx
        this.hostSelectDialogVisible = true
        if (this.hostSelectMode === 'host') {
          this.searchHosts()
        } else {
          this.loadGroupTableData()
        }
      },
      searchHosts() {
        listNode(this.hostSearchParams).then(res => {
          this.hostTableData = res.data.nodeLists
          this.hostTotal = res.data.total
        });
      },
      handleHostSelectionChange(val) {
        this.selectedHosts = val
      },
      handleGroupSelectionChange(val) {
        this.selectedGroups = val
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
        const task = this.form.tasks[this.currentTaskIndex];
    
        if (this.hostSelectMode === 'host') {
          // 单选主机：直接用 selectedHosts
          const ids   = this.selectedHosts.map(h => h.ID);
          const names = this.selectedHosts.map(h => h.nodeName);
    
          task.node_ids   = ids.join(',');
          task.node_names = names.join(',');
        } else {
          // 主机组模式：把组下面所有成员都“摊平”一遍
          const allNodes = this.selectedGroups.flatMap(g => g.t_node_group_s || []);
          const ids      = allNodes.map(n => n.ID);
          const names    = allNodes.map(n => n.nodeName);
    
          task.node_ids   = ids.join(',');
          task.node_names = names.join(',');
        }
    
        this.hostSelectDialogVisible = false;
      },
        
      loadGroupTableData() {
        if (this.nodeGroupList.length === 0) return
        this.groupTableData = this.nodeGroupList
      },
      getNodeName(id) {
        const nid = Number(id)
        const node = this.nodeList.find(n => n.ID === nid)
        if (node) return node.nodeName
        const group = this.nodeGroupList.find(g => g.ID === nid)
        if (group) return group.groupName
        return id
      },
  // —— 删除 tag 的时候，同时删掉 node_ids 和 node_names
  removeNodeId(taskIndex, tagIndex) {
    const task = this.form.tasks[taskIndex];
  
    const idsArr   = task.node_ids.split(',');
    const namesArr = task.node_names.split(',');
  
    idsArr.splice(tagIndex, 1);
    namesArr.splice(tagIndex, 1);
  
    task.node_ids   = idsArr.join(',');
    task.node_names = namesArr.join(',');
  },
  
      // —— 提交 / 取消
      async submitForm() {
        this.$refs.form.validate(async valid => {
          if (!valid) return
          const payload = {
            name: this.form.name,
            desc: this.form.desc,
            args: this.form.args,
            t_task_s: this.form.tasks.map((t, idx) => ({ // 关键修复点
                name: t.name,
                type: t.type,
                content: t.content,
                node_ids: t.node_ids,
                sort: idx + 1,     // ✅ 正确使用循环索引
                timeout: t.timeout
              }))
                  }
          await addTaskManage(payload)
          this.$message.success('操作成功')
          this.$router.push('/exec/task/manage')
        })
      },
      cancel() {
        this.$router.back()
      },
  
      // —— 拉列表
      async fetchScriptList() {
        const res = await listLibrary()
        this.scriptList = res.data.script_library_s.map(s => ({
          id:      s.ID,
          name:    s.name,
          content: s.content,
          type:    s.type
        }))
      },
        async fetchNodeList() {
        const res = await listNode()
        this.nodeList = res.data.nodeLists
      },
      async fetchNodeGroupList() {
        const res = await list_node_group()
        this.nodeGroupList = res.data.nodeGroupLists
        this.groupTableData = this.nodeGroupList
      }
    }
  }
  </script>
  
  <style scoped>
  .container-card {
    margin: 20px;
  }
  .tasks-collapse {
    margin-bottom: 20px;
  }
  
  .subtask-title {
    font-size: 14px;
    font-weight: bold;
  }
  </style>
  