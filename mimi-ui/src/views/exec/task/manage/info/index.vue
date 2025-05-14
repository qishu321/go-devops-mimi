<template>
  <el-card>
    <!-- 返回按钮 -->
    <div style="margin-bottom: 16px">
      <el-button @click="backToManage" type="primary">返回任务管理</el-button>
    </div>
    
    <!-- 任务列表表格 -->
    <el-table
      v-if="!showDetail"
      :data="taskList"
      style="width: 100%"
    >
      <el-table-column prop="name" label="子任务名称" />
      <el-table-column prop="status" label="状态" />
      <el-table-column prop="timeCost" label="耗时(ms)" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button
            size="small"
            @click="showTaskDetail(scope.row)"
          >查看执行详情</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 日志详情视图 -->
    <div v-else>
      <el-button @click="backToList">返回列表</el-button>
      <el-row :gutter="16">
        <!-- 左侧节点列表 -->
        <el-col :span="4">
          <el-menu
            :default-active="activeNode"
            @select="selectNode"
            class="node-menu"
            :router="false"
          >
            <el-menu-item
              v-for="node in nodeList"
              :key="node"
              :index="node"
            >
              {{ node }}
            </el-menu-item>
          </el-menu>
        </el-col>
  
        <!-- 右侧终端输出 -->
        <el-col :span="20">
          <div ref="xtermContainer" class="terminal-container"></div>
        </el-col>
      </el-row>
    </div>
  </el-card>
</template>

<script>
import { Terminal } from 'xterm'
import { runInfo } from '@/api/exec/task'
import 'xterm/css/xterm.css'

export default {
  name: 'TaskLogViewer',
  data() {
    return {
      runID: this.$route.query.runID,
      ws: null,
      term: null,
      taskList: [], // 子任务列表
      showDetail: false, // 是否显示详情
      nodeList: [],
      logsByNode: {},
      activeNode: ''
    }
  },
  async mounted() {
    await this.fetchTaskInfo()
  },
  methods: {
    async fetchTaskInfo() {
      try {
        const res = await runInfo({ run_id: this.runID })
        if (res.code === 0) {
          this.taskList = res.data.t_task_manage_log_s || []
        }
      } catch (err) {
        console.error('获取任务信息失败:', err)
      }
    },
    showTaskDetail(task) {
      this.showDetail = true
      this.$nextTick(() => {
        this.initTerminal()
        this.connectWS(task.ID) // 传入子任务ID
      })
    },
    backToList() {
      this.showDetail = false
      if (this.ws) {
        this.ws.close()
        this.ws = null
      }
    },
    backToManage() {
      this.$router.push('/exec/task/manage')
    },
    initTerminal() {
      this.term = new Terminal({
        cursorBlink: true,
        rows: 40,
        cols: 140,
        fontSize: 14,
        lineHeight: 1.2
      })
      this.term.open(this.$refs.xtermContainer)
      this.term.clear()
      this.term.writeln(`⚡ 正在连接 WebSocket，runID=${this.runID}...`)
    },
    connectWS(taskId) {
      const baseUrl = process.env.VUE_APP_WS_API.replace(/\/$/, '')
      const apiPath = '/api/exec/run_task_manage/run_info_webSocket'
      const wsUrl = `${baseUrl}${apiPath}?task_id=${encodeURIComponent(taskId)}`

      console.log('WS 连接地址：', wsUrl)
      this.ws = new WebSocket(wsUrl)

      this.ws.onopen = () => {
        this.term.writeln('✅ 已连接，开始接收日志')
      }

      this.ws.onmessage = ({ data }) => {
        // 后端推送的是 {"node":"node1","data":"...output..."}
        let msg
        try {
          msg = JSON.parse(data)
        } catch {
          // 如果不是 JSON 格式，就当纯文本
          this.term.write(data.replace(/\n/g, '\r\n'))
          return
        }
        const { node, data: line } = msg

        // 第一次看到这个节点，初始化数据结构
        if (!this.logsByNode[node]) {
          this.$set(this.logsByNode, node, [])
          this.nodeList.push(node)
          if (!this.activeNode) {
            this.activeNode = node
          }
        }
        // 保存所有行
        this.logsByNode[node].push(line)

        // 只有当前激活的节点，才往 xterm 写
        if (node === this.activeNode) {
          this.term.write(line.replace(/\n/g, '\r\n'))
        }
      }

      this.ws.onerror = (err) => {
        console.error('WebSocket 错误:', err)
        this.term.writeln('\x1b[31m⚠ WebSocket 错误，请检查后端服务。\x1b[0m')
      }

      this.ws.onclose = () => {
        this.term.writeln('ℹ️ WebSocket 已关闭')
        this.ws = null
      }
    },
    selectNode(node) {
      this.activeNode = node
      this.term.clear()
      const lines = this.logsByNode[node] || []
      lines.forEach(line => {
        this.term.write(line.replace(/\n/g, '\r\n'))
      })
    }
  }
}
</script>

<style scoped>
.terminal-container {
  height: 800px;
  background: #000;
  width: 100%;
  overflow: hidden;
  border-radius: 4px;
}
.node-menu {
  background: #f5f7fa;
  border-radius: 4px;
  height: 800px;
  overflow-y: auto;
}
</style>
