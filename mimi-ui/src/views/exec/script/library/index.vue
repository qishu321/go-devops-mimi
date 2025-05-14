<template>
  <div>
    <el-card class="container-card" shadow="always">
      <el-form size="mini" :inline="true" :model="params" class="demo-form-inline">
        <el-form-item label="脚本名称">
          <el-input v-model.trim="params.name" clearable placeholder="脚本名称" @change="search" @clear="search" />
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

      <el-table :data="tableData" border @selection-change="handleSelectionChange">
        <el-table-column type="selection" align="center" width="50" />
        <el-table-column prop="CreatedAt" label="创建时间" align="center" width="160" />
        <el-table-column prop="name" label="脚本名称" align="center" min-width="100" />
        <el-table-column prop="type" label="脚本类型" align="center" width="100">
          <template slot-scope="scope">
            <el-tag :type="getScriptTypeTag(scope.row.type)">{{ scope.row.type }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="desc" label="描述" align="center" min-width="150" show-overflow-tooltip />
        <el-table-column prop="creator" label="最新编辑用户"  align="center" min-width="80" show-overflow-tooltip />
        <el-table-column fixed="right" label="查看脚本详情" align="center" width="100">
          <template slot-scope="scope">
            <el-tooltip  content="详情" effect="dark" placement="top">
              <el-button size="mini" icon="el-icon-setting" circle type="info" @click="addUp(scope.row)" />
            </el-tooltip>
          </template>
        </el-table-column>

        <el-table-column fixed="right" label="操作"  align="center" width="160">  
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
        :page-sizes="[10, 20, 50, 100]"
        layout="total, sizes, prev, pager, next, jumper"
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />

      <el-dialog :title="dialogFormTitle" :visible.sync="dialogFormVisible" width="1200px">
        <el-form ref="dialogForm" :model="dialogFormData" :rules="dialogFormRules" label-width="120px">
          <el-form-item label="脚本名称" prop="name">
            <el-input v-model.trim="dialogFormData.name" placeholder="请输入脚本名称" />
          </el-form-item>
          <el-form-item label="脚本类型" prop="type">
            <el-select v-model="dialogFormData.type" placeholder="请选择脚本类型">
              <el-option v-for="item in scriptTypes" :key="item.value" :label="item.label" :value="item.value" />
            </el-select>
          </el-form-item>
          
          <el-form-item label="脚本内容" prop="content">
            <!-- 这里的 ref="editor" 用来初始化 CodeMirror -->
            <textarea ref="editor" style="width:100%; height:800px;"></textarea>
          </el-form-item>
          
          <el-form-item label="描述" prop="desc">
            <el-input v-model="dialogFormData.desc" type="textarea" :rows="3" placeholder="请输入描述信息" />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="cancelForm">取 消</el-button>
          <el-button type="primary" @click="submitForm" :loading="submitLoading">确 定</el-button>
        </div>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
import { listLibrary, deleteLibrary, updateLibrary, addLibrary, infoLibrary } from '@/api/exec/script'
import { Message } from 'element-ui'
import CodeMirror from 'codemirror';
import 'codemirror/lib/codemirror.css';
import 'codemirror/theme/darcula.css';
import 'codemirror/mode/shell/shell.js';


export default {
  name: 'ScriptLibrary',
  data() {
    return {
      editor: null,
      editorInited: false,  
      params: {
        name: '',
        pageNum: 1,
        pageSize: 10
      },
      tableData: [],
      total: 0,
      loading: false,
      multipleSelection: [],
      
      dialogFormVisible: false,
      dialogFormTitle: '',
      dialogType: '',
      submitLoading: false,
      dialogFormData: {
        ID: '',
        name: '',
        type: '',
        content: '',
        desc: ''
      },
      transParams: {
        id: '',
        name: ''
      },
      dialogFormRules: {
        name: [
          { required: true, message: '请输入脚本名称', trigger: 'blur' },
          { min: 1, max: 64, message: '长度在1到64个字符', trigger: 'blur' }
        ],
        type: [
          { required: true, message: '请选择脚本类型', trigger: 'change' }
        ],
        content: [
          { required: true, message: '请输入脚本内容', trigger: 'blur' }
        ]
      },
      scriptTypes: [
        { label: 'Shell', value: 'shell' },
        { label: 'Python', value: 'python' },
        { label: 'PowerShell', value: 'powershell' }
      ]
    }
  },
  mounted() {
    // 初始化一次 CodeMirror
    this.editor = CodeMirror.fromTextArea(this.$refs.editor, {
      mode: 'shell',           // 根据脚本类型改 mode
      theme: 'darcula',
      lineNumbers: true,
      matchBrackets: true,
      autoCloseBrackets: true
    });
  },
  watch: {
    dialogFormVisible(val) {
      if (!val) return;
      this.$nextTick(() => {
        // 第一次打开时新建
        if (!this.editorInited) {
          this.editor = CodeMirror.fromTextArea(this.$refs.editor, {
            mode: this.dialogFormData.type || 'shell',
            theme: 'darcula',
            lineNumbers: true,
            matchBrackets: true,
            autoCloseBrackets: true
          });
          this.editorInited = true;
        }
        // 每次打开都要加载最新 content
        this.editor.setOption('mode', this.dialogFormData.type || 'shell');
        this.editor.setValue(this.dialogFormData.content || '');
      });
    }
  },
    created() {
    this.getTableData()
  },
  methods: {
     // 穿梭框
     addUp(row) {
      this.dialogTransfer = '脚本详情'
       this.dialogTransferVisible = true
       this.transParams.id = row.ID
       this.transParams.name = row.name
       this.$router.push({ path: '/libraryInfo', query: { id: row.ID,name:row.name } })
    },

    getScriptTypeTag(type) {
      const map = {
        shell: '',
        python: 'success',
        powershell: 'warning'
      }
      return map[type] || ''
    },
    
    async getTableData() {
      this.loading = true
      try {
        const { data } = await listLibrary(this.params)
        this.tableData = data.script_library_s || []
        this.total = data.total || 0
      } catch (error) {
        console.error(error)
      } finally {
        this.loading = false
      }
    },
    
    search() {
      this.params.pageNum = 1
      this.getTableData()
    },
    
    create() {
      this.dialogFormTitle = '新增脚本';
      this.dialogType = 'create';
      this.dialogFormData = { ID:'', name:'', type:'shell', content:'', desc:'' };
      this.dialogFormVisible = true;
    },
      
    async update(row) {
      const { data } = await infoLibrary({ id: row.ID,name:row.name });
      this.dialogFormData = {
        ID: row.ID,
        name: data.name,
        type: data.type,
        content: data.content,
        desc: data.desc
      };
      this.dialogFormTitle = '编辑脚本';
      this.dialogType = 'update';
      this.dialogFormVisible = true;
    },
      
  // 提交前同步内容并提交
  submitForm() {
    // 把编辑器内容回写到 model
    this.dialogFormData.content = this.editor.getValue();
    this.$refs.dialogForm.validate(async valid => {
      if (!valid) return;
      this.submitLoading = true;
      try {
        if (this.dialogType === 'create') {
          await addLibrary(this.dialogFormData);
        } else {
          await updateLibrary(this.dialogFormData);
        }
        Message.success('操作成功');
        this.resetForm();
        this.getTableData();
      } finally {
        this.submitLoading = false;
      }
    });
  },
  resetForm() {
    this.$refs.dialogForm.resetFields();
    this.dialogFormVisible = false;
  },
    
    cancelForm() {
      this.resetForm()
    },
    
    batchDelete() {
      if (this.multipleSelection.length === 0) {
        Message.warning('请至少选择一项')
        return
      }
      
      this.$confirm('确定删除选中的脚本吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        this.loading = true
        try {
          await deleteLibrary({ 
            ids: this.multipleSelection.map(item => item.ID) 
          })
          Message.success('删除成功')
          this.getTableData()
        } catch (error) {
          console.error(error)
        } finally {
          this.loading = false
        }
      }).catch(() => {})
    },
        // 单个删除主机
        async singleDelete(Id) {
          this.loading = true
          try {
            await deleteLibrary({ Ids: [Id] }).then(res => {
              this.getTableData()
              Message.success('删除成功')            })
          } finally {
            this.loading = false
          }
          if (this.activeGroupId === 'all') {
            this.getTableData()
          }
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
    }
  }
}
</script>

<style scoped>
.container-card {
  margin: 10px;
  margin-bottom: 20px;
}

</style>
