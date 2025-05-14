<template>
  <el-card>
    <el-row :gutter="16">
      <!-- 左侧节点列表 -->
      <el-col :span="4">
        <el-menu
          :default-active="activeNode"
          @select="selectNode"
          :router="false"
          class="node-menu"
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
  </el-card>
</template>

<script>
import { Terminal } from 'xterm'
import 'xterm/css/xterm.css'

export default {
  name: 'TaskLogViewer',
  data() {
    return {
      taskParams: {
        id: this.$route.query.id || '',
        name: this.$route.query.name || '',
        env_task_s: this.$route.query.env_task_s || {}
      },
      ws: null,
      wsStatus: false,
      term: null,
      reconnectAttempts: 0,
      nodeList: [],
      logsByNode: {},
      activeNode: ''
    }
  },
  mounted() {
    console.log('路由参数:', this.$route.query)
    console.log('解析后的env_task_s:', this.taskParams.env_task_s)  
    this.initTerminal()
    this.connectWS()
  },
  methods: {
    initTerminal() {
      this.term = new Terminal({
        cursorBlink: true,
        rows: 20,
        cols: 80
      })
      this.term.open(this.$refs.xtermContainer)
      this.term.clear()
    },
    connectWS() {
      if (this.ws) return

      const baseUrl = process.env.VUE_APP_WS_API.replace(/\/$/, '')
      const apiPath = '/api/exec/run_task_manage/run'
      // 将env_task_s转换为URL查询参数
      const envParams = this.taskParams.env_task_s
      // 修改参数处理方式，直接使用原始JSON字符串
      const envQuery = `env_task_s=${encodeURIComponent(this.$route.query.env_task_s)}`
      const wsUrl = `${baseUrl}${apiPath}?id=${this.taskParams.id}&name=${encodeURIComponent(this.taskParams.name)}&${envQuery}`
      console.log('修正后的WebSocket连接URL:', wsUrl)
      this.term.writeln('⚡ 正在连接 WebSocket...')
      this.ws = new WebSocket(wsUrl)

      this.ws.onopen = () => {
        this.wsStatus = true
        this.reconnectAttempts = 0
        this.term.writeln('✅ 已连接，开始接收日志')
      }

      this.ws.onmessage = ({ data }) => {
        let msg
        try {
          msg = JSON.parse(data)
        } catch (e) {
          this.term.write(data.replace(/\n/g, '\r\n'))
          return
        }
        const { node, data: line } = msg

        if (!this.logsByNode[node]) {
          this.$set(this.logsByNode, node, [])
          this.nodeList.push(node)
          if (!this.activeNode) {
            this.activeNode = node
          }
        }
        this.logsByNode[node].push(line)

        if (node === this.activeNode) {
          this.term.write(line.replace(/\n/g, '\r\n'))
        }
      }

      this.ws.onerror = (err) => {
        console.error('WebSocket 错误:', err)
        this.term.writeln('\x1b[31m⚠ WebSocket 错误...\x1b[0m')
      }

      this.ws.onclose = () => {
        this.wsStatus = false
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
  height: 400px;
  background: #000;
  font-size: 14px;
  color: #fff;
  overflow: hidden;
  border-radius: 4px;
}
.node-menu {
  background: #f5f7fa;
  border-radius: 4px;
  height: 400px;
  overflow-y: auto;
}
</style>
