<template>
  <div class="script-info-container">
    <div class="header">
      <h2>脚本详情 - {{ scriptData.name }}</h2>
      <div class="script-meta">
        <p><span class="label">创建时间：</span>{{ formatTime(scriptData.CreatedAt) }}</p>
        <p><span class="label">更新时间：</span>{{ formatTime(scriptData.UpdatedAt) }}</p>
        <p><span class="label">脚本类型：</span>{{ scriptData.type }}</p>
        <p><span class="label">创建用户：</span>{{ scriptData.creator }}</p>
        <p><span class="label">描述：</span>{{ scriptData.desc }}</p>
        <p><span class="label">脚本内容：</span></p>
    <span class="editor-container">
      <textarea ref="editor" style="width:100%; height:800px;"></textarea>
    </span>
</div>
</div>

  </div>
</template>

<script>
import CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';
import 'codemirror/theme/darcula.css';
import 'codemirror/mode/shell/shell.js';

import { infoLibrary } from '@/api/exec/script';

export default {
  name: 'ScriptInfo',
  data() {
    return {
      dialogVisible: false,
      dialogTitle: '',
      editor: null,
      scriptData: {
        ID: '',
        name: '',
        type: '',
        content: '',
        desc: ''
      },
      transParams: {
        id: '',
        name: ''
      }
    };
  },
  created() {
    // 从路由参数中获取传递过来的数据
    if (this.$route.query) {
      this.transParams.id = this.$route.query.id || '';
      this.transParams.name = this.$route.query.name || '';
      this.showInfo(this.transParams.id, this.transParams.name);
    }
  },
  methods: {
    async showInfo(id, name) {
      try {
        const { data } = await infoLibrary({ id, name });
        this.scriptData = data;
        this.dialogTitle = `脚本详情 - ${data.name}`;
        this.dialogVisible = true;
        
        this.$nextTick(() => {
          if (!this.editor) {
            this.editor = CodeMirror.fromTextArea(this.$refs.editor, {
              mode: this.scriptData.type || 'shell',
              theme: 'darcula',
              lineNumbers: true,
              matchBrackets: true,
              autoCloseBrackets: true,
              readOnly: true
            });
          }
          this.editor.setOption('mode', this.scriptData.type || 'shell');
          this.editor.setValue(this.scriptData.content || '');
        });
      } catch (error) {
        console.error(error);
      }
    },
    formatTime(timeStr) {
      if (!timeStr) return '';
      const date = new Date(timeStr);
      return date.toLocaleString('zh-CN', {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false
      }).replace(/\//g, '-');
    },
  }
};
</script>

<style scoped>
.script-info-container {
  padding: 20px;
}
.header {
  margin-bottom: 20px;
}
.script-meta {
  margin-top: 20px;
  padding: 15px;
  background: #f5f7fa;
  border-radius: 4px;
}
.label {
  font-weight: bold;
  color: #606266;
  margin-right: 10px;
}
.editor-container {
  margin: 20px 0;
}
</style>