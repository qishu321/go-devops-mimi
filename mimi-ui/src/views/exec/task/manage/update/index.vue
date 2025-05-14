<template>
  <div>
    <el-card class="container-card" shadow="always">
      <el-form
        ref="form"
        :model="form"
        :rules="rules"
        label-width="120px"
      >
        <!-- --- 任务基本信息 --- -->
        <el-row :gutter="20">
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
          <el-col :span="12">
            <el-form-item label="执行参数" prop="args">
              <el-button type="primary" size="small" @click="addParam">
                添加参数
              </el-button>
              <el-table
                :data="form.args"
                border
                size="small"
                style="margin-top:10px"
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
                  v-if="form.args.some(i => i.type==='select')"
                  label="选项值"
                >
                  <template #default="{ row }">
                    <el-input
                      v-if="row.type==='select'"
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

        <!-- --- 子任务步骤 --- -->
        <el-row :gutter="20" align="middle" style="margin-top:20px;">
          <el-col :span="2">
            <span class="subtask-title">子任务</span>
          </el-col>
          <el-col :span="22">
            <el-button type="primary" size="small" @click="addSubTask">
              添加子任务
            </el-button>
          </el-col>
        </el-row>

        <el-collapse
          v-model="activeTasks"
          class="tasks-collapse"
          @change="onCollapseChange"
        >
          <el-collapse-item
            v-for="(task, idx) in form.tasks"
            :key="task.tempId"
            :name="task.tempId"
            :title="`步骤 ${idx+1}：${task.name||'新建步骤'}`"
          >
          <el-button
          size="mini"
          type="danger"
          icon="el-icon-delete"
          @click="removeSubTask(idx)"
        >删除本步</el-button>
            <el-form-item :prop="`tasks.${idx}.name`" label="步骤名称">
              <el-input v-model="task.name" placeholder="请输入步骤名称" />
            </el-form-item>
            <el-form-item :prop="`tasks.${idx}.node_ids`" label="目标主机">
              <el-button size="mini" type="primary" @click="openHostSelectDialog(idx)">
                选择目标主机
              </el-button>
              <div v-if="task.node_names" style="margin-top:10px;">
                <el-tag
                  v-for="(name,i) in task.node_names.split(',')"
                  :key="i"
                  closable
                  @close="removeNodeId(idx,i)"
                  style="margin-right:5px; margin-bottom:5px;"
                >{{name}}</el-tag>
              </div>
            </el-form-item>
            <el-form-item label="脚本来源">
              <el-radio-group v-model="task.source">
                <el-radio label="manual">手工输入</el-radio>
                <el-radio label="repo">脚本仓库</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item v-if="task.source==='repo'" label="选择脚本">
              <el-select
                v-model="task.scriptId"
                placeholder="请选择脚本"
                filterable
                size="small"
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
            <el-form-item :prop="`tasks.${idx}.timeout`" label="脚本超时时间(秒)">
              <el-input-number
                v-model="task.timeout"
                :min="30"
                :max="3600"
                size="small"
              />
            </el-form-item>
            <el-form-item label="脚本内容">
              <div class="editor-wrapper">
                <textarea :ref="`editor${idx}`"></textarea>
              </div>
            </el-form-item>
            <el-button size="small" type="primary" @click="addSubTask">继续添加子任务</el-button>
          </el-collapse-item>
        </el-collapse>

        <!-- 主机选择对话框 -->
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

        <!-- 操作按钮 -->
        <el-form-item style="text-align:right; margin-top:20px;">
          <el-button type="primary" @click="submitForm">保存</el-button>
          <el-button @click="cancel">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { infoTaskManage, updateTaskManage } from '@/api/exec/task'
import { listLibrary } from '@/api/exec/script'
import { listNode } from '@/api/cmdb/node'
import { list_node_group } from '@/api/cmdb/node_group'
import CodeMirror from 'codemirror'
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/darcula.css'
import 'codemirror/mode/shell/shell.js'
import 'codemirror/addon/edit/matchbrackets.js'
import 'codemirror/addon/edit/closebrackets.js'

export default {
  name: 'TaskManageUpdate',
  data() {
    return {
      form: { name: '', desc: '', args: [], tasks: [] },
      transParams: { id: '', name: '' },
      activeTasks: [],
      scriptList: [],
      nodeList: [],
      nodeGroupList: [],
      groupTableData: [],
      hostSelectDialogVisible: false,
      hostSelectMode: 'host',
      hostSearchParams: { nodeName: '', pageNum: 1, pageSize: 10 },
      hostTableData: [],
      hostTotal: 0,
      selectedHosts: [],
      selectedGroups: [],
      editors: [],
      rules: { name: [{ required: true, message: '请输入任务名称', trigger: 'blur' }] }
    }
  },
  async created() {
    await Promise.all([this.fetchScriptList(), this.fetchNodeList(), this.fetchNodeGroupList()])
    const { id, name } = this.$route.query || {}
    if (id && name) {
      this.transParams = { id, name }
      this.showInfo(id, name)
    }
  },
  methods: {
    async showInfo(id, name) {
      try {
        const res = await infoTaskManage({ id, name })
        const d = res.data.data || res.data
        this.form.name = d.name
        this.form.desc = d.desc
        this.form.args = typeof d.args === 'string' ? JSON.parse(d.args) : d.args || []
        this.form.tasks = (d.t_task_s || []).map((t, i) => ({
          id: t.ID,
          tempId: `step-${i}-${Date.now()}`,
          name: t.name,
          type: t.type,
          content: t.content,
          source: t.script_id ? 'repo' : 'manual',
          scriptId: t.script_id,
          timeout: t.timeout,
          node_ids: t.node_ids,
          node_names: t.node_ids
            .split(',')
            .map(id => this.getNodeName(id))
            .join(',')
        }))
        this.activeTasks = this.form.tasks.map(t => t.tempId)
        this.$nextTick(() => {
          this.form.tasks.forEach((_, idx) => this.initCodeEditor(idx))
        })
      } catch (e) {
        console.error('获取详情失败', e)
      }
    },
    initCodeEditor(index) {
      const ref = this.$refs[`editor${index}`]
      const textarea = Array.isArray(ref) ? ref[0] : ref
      if (!textarea) return
      const cm = CodeMirror.fromTextArea(textarea, {
        mode: 'shell',
        theme: 'darcula',
        lineNumbers: true,
        matchBrackets: true,
        autoCloseBrackets: true,
        lineWrapping: true,
        viewportMargin: Infinity
      })
      cm.setValue(this.form.tasks[index].content || '')
      cm.on('change', () => {
        this.form.tasks[index].content = cm.getValue()
      })
      this.editors[index] = cm
    },
    onCollapseChange(names) {
      this.$nextTick(() => {
        names.forEach(id => {
          const i = this.form.tasks.findIndex(t => t.tempId === id)
          this.editors[i]?.refresh()
        })
      })
    },
    addParam() {
      this.form.args.push({ name: '', type: 'string', options: '', required: true })
    },
    removeParam(i) {
      this.form.args.splice(i, 1)
    },
    handleScriptSelect(idx) {
      const t = this.form.tasks[idx]
      const f = this.scriptList.find(s => s.id === t.scriptId)
      if (f) {
        t.content = f.content
        this.editors[idx]?.setValue(f.content)
      }
    },
    addSubTask() {
      const i = this.form.tasks.length
      const ts = Date.now()
      this.form.tasks.push({
        id: 0,
        tempId: `step-${i}-${ts}`,
        name: '',
        type: 'shell',
        content: '',
        source: 'manual',
        scriptId: null,
        timeout: 30,
        node_ids: '',
        node_names: ''
      })
      this.activeTasks.push(this.form.tasks[i].tempId)
      this.$nextTick(() => this.initCodeEditor(i))
    },
    removeSubTask(i) {
      this.form.tasks.splice(i, 1)
      this.activeTasks.splice(i, 1)
      this.editors.splice(i, 1)
    },
    removeNodeId(taskIndex, tagIndex) {
      const task = this.form.tasks[taskIndex]
      const idsArr = task.node_ids.split(',')
      const namesArr = task.node_names.split(',')
      idsArr.splice(tagIndex, 1)
      namesArr.splice(tagIndex, 1)
      task.node_ids = idsArr.join(',')
      task.node_names = namesArr.join(',')
    },
    openHostSelectDialog(idx) {
      this.currentTaskIndex = idx
      this.hostSelectDialogVisible = true
      this.hostSelectMode === 'host' ? this.searchHosts() : this.loadGroupTableData()
    },
    searchHosts() {
      listNode(this.hostSearchParams).then(r => {
        this.hostTableData = r.data.nodeLists
        this.hostTotal = r.data.total
      })
    },
    handleHostSelectionChange(v) {
      this.selectedHosts = v
    },
    handleGroupSelectionChange(v) {
      this.selectedGroups = v
    },
    handleHostSizeChange(s) {
      this.hostSearchParams.pageSize = s
      this.searchHosts()
    },
    handleHostPageChange(p) {
      this.hostSearchParams.pageNum = p
      this.searchHosts()
    },
    confirmHostSelection() {
      const t = this.form.tasks[this.currentTaskIndex]
      if (this.hostSelectMode === 'host') {
        const ids = this.selectedHosts.map(h => h.ID)
        const names = this.selectedHosts.map(h => h.nodeName)
        t.node_ids = ids.join(',')
        t.node_names = names.join(',')
      } else {
        const nodes = this.selectedGroups.flatMap(g => g.t_node_group_s || [])
        t.node_ids = nodes.map(n => n.ID).join(',')
        t.node_names = nodes.map(n => n.nodeName).join(',')
      }
      this.hostSelectDialogVisible = false
    },
    loadGroupTableData() {
      this.groupTableData = this.nodeGroupList
    },
    getNodeName(id) {
      const nid = Number(id)
      const nd = this.nodeList.find(n => n.ID === nid)
      if (nd) return nd.nodeName
      const gp = this.nodeGroupList.find(g => g.ID === nid)
      if (gp) return gp.groupName
      return id
    },
    async submitForm() {
      this.$refs.form.validate(async v => {
        if (!v) return
        this.form.tasks.forEach((t, i) => {
          if (this.editors[i]) t.content = this.editors[i].getValue()
        })
        const pw = {
          id: parseInt(this.transParams.id),
          name: this.form.name,
          desc: this.form.desc,
          args: JSON.stringify(this.form.args),
          t_task_s: this.form.tasks.map((t, i) => ({
            id: t.id,
            name: t.name,
            type: t.type,
            content: t.content,
            node_ids: t.node_ids,
            timeout: t.timeout,
            sort: i + 1,
            script_id: t.scriptId
          }))
        }
        try {
          await updateTaskManage(pw)
          this.$message.success('更新成功')
          this.$router.push('/exec/task/manage')
        } catch (e) {
          this.$message.error('更新失败')
          console.error(e)
        }
      })
    },
    cancel() {
      this.$router.back()
    },
    async fetchScriptList() {
      const r = await listLibrary()
      this.scriptList = r.data.script_library_s.map(s => ({
        id: s.ID,
        name: s.name,
        content: s.content,
        type: s.type
      }))
    },
    async fetchNodeList() {
      const r = await listNode()
      this.nodeList = r.data.nodeLists
    },
    async fetchNodeGroupList() {
      const r = await list_node_group()
      this.nodeGroupList = r.data.nodeGroupLists
      this.groupTableData = this.nodeGroupList
    }
  }
}
</script>

<style scoped>
.container-card{margin:20px}
.tasks-collapse{margin-bottom:20px}
.subtask-title{font-size:14px;font-weight:bold}
.editor-wrapper{height:200px;width:100%}
.editor-wrapper .CodeMirror{height:100% !important}
</style>
