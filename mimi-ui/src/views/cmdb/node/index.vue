<template>
  <el-row :gutter="20">
    <!-- 左侧分组菜单 -->
    <el-col :span="4">
      <el-card class="group-card">
        <div style="display: flex; justify-content: space-between; align-items: center; padding: 0 10px;">
          <span>主机组列表</span>
          <!-- 使用文字型按钮更明显 -->
          <el-button type="primary" size="mini" icon="el-icon-plus" @click="onAddGroupClick"></el-button>
        </div>
        
        <el-menu
          :default-active="activeGroupId"
          @select="handleGroupSelect"
          class="el-menu-vertical-demo"
        >
          <el-menu-item index="all">全部主机</el-menu-item>
          <el-menu-item
            v-for="group in groups.nodeGroupLists"
            :key="group.ID"
            :index="group.ID.toString()"
          >
            <div style="display: flex; justify-content: space-between; align-items: center;">
              <span>{{ group.groupName }}</span>
              <el-dropdown trigger="click">
                <i class="el-icon-more"></i>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item @click.native="renameGroup(group)">重命名</el-dropdown-item>
                  <el-dropdown-item divided @click.native="deleteGroup(group)">删除</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </div>
          </el-menu-item>
        </el-menu>
        
        <!-- 添加主机组弹窗 -->
        <el-dialog title="添加主机组" :visible.sync="addGroupDialogVisible" width="400px" center>
          <el-form :model="newGroup" label-width="100px">
            <el-form-item label="主机组名称">
              <el-input v-model="newGroup.groupName" />
            </el-form-item>
          </el-form>
          <div slot="footer">
            <el-button @click="addGroupDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitAddGroup">确定</el-button>
          </div>
        </el-dialog>
              </el-card>
    </el-col>

    <!-- 右侧主机管理 -->
    <el-col :span="20">
      <el-card class="container-card" shadow="always">
        <el-form size="mini" :inline="true" :model="params" class="demo-form-inline">
          <el-form-item label="主机名称">
            <el-input
              v-model.trim="params.nodeName"
              clearable
              placeholder="主机名称"
              @change="search"
              @clear="search"
            />
          </el-form-item>
          <el-form-item label="IP">
            <el-input
              v-model.trim="params.publicIP"
              clearable
              placeholder="IP"
              @change="search"
              @clear="search"
            />
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" icon="el-icon-search" type="primary" @click="search">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button :loading="loading" icon="el-icon-plus" type="warning" @click="create">新增</el-button>
          </el-form-item>
          <el-form-item>
            <el-button
              :disabled="multipleSelection.length === 0"
              :loading="loading"
              icon="el-icon-edit"
              type="warning"
              @click="openBatchUpdateGroupDialog "
            >批量更新主机主机组</el-button>
          </el-form-item>

          <el-form-item>
            <el-button
              :disabled="multipleSelection.length === 0"
              :loading="loading"
              icon="el-icon-delete"
              type="danger"
              @click="batchDelete"
            >批量删除</el-button>
          </el-form-item>
        </el-form>
        <!-- 批量更新主机主机组的弹窗 -->
        <el-dialog title="批量更新主机主机组" :visible.sync="updateNodesGroupDialogVisible" width="400px" center>
          <el-form :model="batchUpdateForm">
            <el-form-item label="主机组" label-width="80px">
              <el-select
                v-model="batchUpdateForm.groupIds"
                multiple
                filterable
                placeholder="请选择主机组"
                style="width: 100%;"
              >
                <el-option
                  v-for="group in groupOptions"
                  :key="group.ID"
                  :label="group.groupName"
                  :value="group.ID"
                />
              </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="updateNodesGroupDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitBatchUpdateGroup">确定</el-button>
          </div>
        </el-dialog>

        <el-table :data="tableData" border auto-fit-column @selection-change="handleSelectionChange">
          <el-table-column type="selection" align="center" />
          <el-table-column show-overflow-tooltip align="center" prop="CreatedAt" label="创建时间" />
          <el-table-column show-overflow-tooltip align="center" prop="nodeName" label="主机名称" />
          <el-table-column show-overflow-tooltip align="center" prop="username" label="系统用户" />
          <el-table-column show-overflow-tooltip align="center" prop="publicIP" label="IP" />
          <!-- 新增主机组列 -->
          <el-table-column align="center"  label="主机组" width="180">
            <template slot-scope="scope">
              <div v-if="scope.row.t_node_group_s && scope.row.t_node_group_s.length">
                <el-tag
                  v-for="(item, index) in scope.row.t_node_group_s"
                  :key="index"
                  type="success"
                  size="mini"
                  style="margin-right: 4px;"
                >
                  {{ item.groupName }}
                </el-tag>
              </div>
              <span v-else>未绑定</span>
            </template>
          </el-table-column>
          <el-table-column show-overflow-tooltip align="center" prop="sshPort" label="端口" />
          <el-table-column show-overflow-tooltip align="center" prop="authmodel" label="密码/密钥" />
          <el-table-column align="center" label="连接状态" width="120">
            <template slot-scope="scope">
              <el-button
                :type="scope.row.status === 1 ? 'success' : 'danger'"
                size="mini"
                disabled>
                {{ scope.row.status === 1 ? 'RUNNING' : 'ERROR' }}
              </el-button>
            </template>
          </el-table-column>
          <el-table-column show-overflow-tooltip align="center" prop="label" label="标签" />
          <el-table-column show-overflow-tooltip align="center" prop="creator" label="创建人" />
          <el-table-column fixed="right" label="操作" align="center" width="120">
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
          :page-sizes="[20, 50, 100, 300]"
          layout="total, prev, pager, next, sizes"
          background
          style="margin-top: 10px; float: right; margin-bottom: 10px;"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />

        <el-dialog :title="dialogFormTitle" :visible.sync="dialogFormVisible">
          <el-form ref="dialogForm" size="small" :model="dialogFormData" :rules="dialogFormRules" label-width="120px">
            <el-form-item label="主机名称" prop="nodeName">
              <el-input v-model.trim="dialogFormData.nodeName" style="width: 240px" placeholder="主机名称" />
            </el-form-item>
            <el-form-item label="系统用户" prop="username">
              <el-input v-model.trim="dialogFormData.username" style="width: 240px" placeholder="系统用户" />
            </el-form-item>
            <el-form-item label="密码/密钥" prop="authmodel">
              <el-radio-group v-model="dialogFormData.authmodel">
                <el-radio size="large" border label="password">密码</el-radio>
                <el-radio size="large" border label="private_key">密钥</el-radio>
              </el-radio-group>
              <el-input
                v-if="dialogFormData.authmodel === 'password'"
                v-model="dialogFormData.password"
                show-password
                type="password"
                placeholder="请输入密码"
                style="margin-top: 10px; width: 300px"
              ></el-input>
              <el-input
                v-else-if="dialogFormData.authmodel === 'private_key'"
                v-model="dialogFormData.private_key"
                type="textarea"
                placeholder="请输入密钥内容"
                :rows="4"
                style="margin-top: 10px; width: 300px"
              ></el-input>
            </el-form-item>
            <el-form-item label="IP地址" prop="publicIP">
              <el-input v-model.trim="dialogFormData.publicIP" style="width: 240px" placeholder="请输入IP地址" show-word-limit />
            </el-form-item>
            <el-form-item label="ssh端口" prop="sshPort">
              <el-input-number v-model.number="dialogFormData.sshPort" placeholder="sshPort" show-word-limit />
            </el-form-item>
            <el-form-item label="连接超时" prop="timeout">
              <el-input-number v-model.number="dialogFormData.timeout" placeholder="timeout" show-word-limit />
            </el-form-item>
            <el-form-item label="主机组" prop="groupId">
              <el-select
                v-model="dialogFormData.groupId"
                multiple
                filterable
                placeholder="请选择主机组"
                style="width: 240px;"
              >
                <el-option
                  v-for="group in groupOptions"
                  :key="group.ID"
                  :label="group.groupName"
                  :value="group.ID"
                ></el-option>
              </el-select>
            </el-form-item>
            
            <el-form-item label="标签" prop="label">
              <el-input v-model.trim="dialogFormData.label" style="width: 240px" placeholder="标签"  />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button size="mini" @click="cancelForm">取 消</el-button>
            <el-button size="mini" :loading="submitLoading" type="primary" @click="submitForm">确 定</el-button>
          </div>
        </el-dialog>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import { listNode, addNode, updateNode, delNode,addNodesGroup } from '@/api/cmdb/node'
import {list_node_group,add_node_group,update_node_group, del_node_group} from '@/api/cmdb/node_group'
import { Message } from 'element-ui'

export default {
  name: 'Node',
  data() {
    return {
      // 分组相关数据（接口返回结构为 { nodeGroupLists: [...] }）
      groups: {},
      activeGroupId: 'all',
      addGroupDialogVisible: false,
      newGroup: {
        groupName: ''
      },
  
      // 查询参数，仅对“全部主机”有效
      params: {
        nodeName: '',
        username: '',
        publicIP: '',
        sshPort: 22,
        authmodel: '',
        status: 0,
        CreatedAt: '',
        label: '',
        creator: '',
        pageNum: 1,
        pageSize: 10,
      },
      // 表格数据（服务器列表）
      tableData: [],
      allNodes: [],
      total: 0,
      loading: false,
      updateNodesGroupDialogVisible: false,
      batchUpdateForm: {
        groupIds: []  // 用于保存选中的主机组ID数组
      },
  
      // 主机新增/编辑对话框数据
      submitLoading: false,
      dialogFormTitle: '',
      dialogType: '',
      dialogFormVisible: false,
      groupOptions: [],  // 用于存放主机组列表
      dialogFormData: {
        ID: '',
        nodeName: '',
        username: '',
        publicIP: '',
        sshPort: 22,
        authmodel: 'password',
        password: '',
        private_key: '',
        timeout: 5,
        label: '',
        groupId: [],     // 新增属性，保存已选主机组ID数组
        creator: '',
      },
      dialogFormRules: {
        nodeName: [
          { required: true, message: '请输入主机名称', trigger: 'blur' },
          { min: 1, max: 100, message: '长度在 1 到 100 个字符', trigger: 'blur' }
        ],
        username: [
          { required: true, message: '请输入系统用户名称', trigger: 'blur' },
          { min: 1, max: 50, message: '长度在 1 到 50 个字符', trigger: 'blur' }
        ],
        authmodel: [
          { required: true,  trigger: 'blur' },
        ],
        publicIP: [
          { required: true, message: '请输入ip地址', trigger: 'change' }
        ],
      },
      popoverVisible: false,
      multipleSelection: []
    }
  },
  created() {
    // 加载分组数据和全部主机数据
    this.getGroups()
    this.getTableData()
  },
  methods: {
    openBatchUpdateGroupDialog() {
      this.batchUpdateForm.groupIds = [];  // 每次打开前重置
      this.updateNodesGroupDialogVisible = true;
    },
    async submitBatchUpdateGroup() {
      // 不再判断 groupIds 是否为空，因为主机组可以为空
      try {
        // 构建请求数据，取出批量选中的主机的ID
        const payload = {
          groupIds: this.batchUpdateForm.groupIds, // 如果用户没选则为 []
          nodeIds: this.multipleSelection.map(item => item.ID)
        };
        await addNodesGroup(payload);
        this.$message.success("更新成功");
        this.updateNodesGroupDialogVisible = false;
        this.getTableData(); // 刷新表格数据
      } catch (error) {
        this.$message.error("更新失败");
      }
    },
    
    onAddGroupClick() {
      this.newGroup.groupName = ''
      this.addGroupDialogVisible = true
    },
    async submitAddGroup() {
      try {
        // 调用添加主机组接口（请根据实际情况修改调用方式）
        await add_node_group(this.newGroup)
        this.$message.success("添加成功")
        this.addGroupDialogVisible = false
        this.getGroups()  // 添加成功后重新加载分组列表
      } catch (err) {
        this.$message.error("添加失败")
      }
    },
    renameGroup(group) {
      this.$prompt('请输入新的主机组名称', '重命名', {
        inputValue: group.groupName
      }).then(({ value }) => {
        update_node_group({ id: group.ID, groupName: value }).then(() => {
          this.$message.success("修改成功")
          this.getGroups()
        })
      })
    },
    deleteGroup(group) {
      this.$confirm('确认删除该主机组？', '提示', {
        type: 'warning'
      }).then(() => {
        del_node_group({ Ids: [group.ID] }).then(() => {
          this.$message.success("删除成功")
          this.getGroups()
        })
      })
    },
    // 获取分组列表
    async getGroups() {
      try {
        const res = await list_node_group()
        // 假设接口返回格式 { nodeGroupLists: [ ... ] }
        this.groups = res.data
        this.groupOptions = res.data.nodeGroupLists || []
      } catch (error) {
        Message({
          message: '获取分组失败',
          type: 'error'
        })
      }
    },
    // 分组选择处理：如果选中“全部主机”，调用接口加载全部；否则直接从 groups 找对应的服务器列表
    handleGroupSelect(groupId) {
      this.activeGroupId = groupId
      this.params.pageNum = 1
      if (groupId === 'all') {
        // “全部主机”状态下，显示全部数据
        this.tableData = this.allNodes || []
        this.total = this.tableData.length
      } else {
        // 从全部主机数据中筛选出绑定了该主机组的主机
        const filteredNodes = (this.allNodes || []).filter(node =>
          node.t_node_group_s &&
          node.t_node_group_s.some(item => item.ID.toString() === groupId.toString())
        )
        this.tableData = filteredNodes
        this.total = filteredNodes.length
      }
    },
        // 针对“全部主机”状态下的查询操作
    search() {
      this.params.pageNum = 1
      this.getTableData()
    },
    // 获取全部主机数据
    async getTableData() {
      if (this.activeGroupId !== 'all') return
      this.loading = true
      try {
        const { data } = await listNode(this.params)
        // 缓存全部数据以便后续过滤
        this.allNodes = data.nodeLists
        this.tableData = data.nodeLists
        this.total = data.total
      } finally {
        this.loading = false
      }
    },
        // 新增主机记录
    create() {
      this.dialogFormTitle = '新增接口'
      this.dialogType = 'create'
      this.dialogFormVisible = true
    },
    // 编辑主机记录
    update(row) {
      // 创建一个新的对象，确保数据是全新的，保证响应性更新
      this.dialogFormData = {
        ID: row.ID,
        nodeName: row.nodeName,
        username: row.username,
        publicIP: row.publicIP,
        sshPort: row.sshPort,
        authmodel: row.authmodel,
        password: row.password,
        private_key: row.private_key,
        timeout: row.timeout,
        label: row.label,
        creator: row.creator,
        // 将绑定的主机组转换为ID数组
        groupId: row.t_node_group_s ? row.t_node_group_s.map(item => item.ID) : []
      }
      this.dialogFormTitle = '修改接口'
      this.dialogType = 'update'
      this.dialogFormVisible = true
    },
        // 判断接口返回结果（主机相关）
    judgeResult(res) {
      if (res.code == 0) {
        Message({
          showClose: true,
          message: "操作成功",
          type: 'success'
        })
      }
    },
    // 提交新增/修改主机表单
    submitForm() {
      this.$refs['dialogForm'].validate(async valid => {
        if (valid) {
          this.submitLoading = true
          try {
            if (this.dialogType === 'create') {
              await addNode(this.dialogFormData).then(res => {
                this.judgeResult(res)
              })
            } else {
              await updateNode(this.dialogFormData).then(res => {
                this.judgeResult(res)
              })
            }
          } finally {
            this.submitLoading = false
          }
          this.resetForm()
          if (this.activeGroupId === 'all') {
            this.getTableData()
          }
        } else {
          Message({
            showClose: true,
            message: '表单校验失败',
            type: 'error'
          })
          return false
        }
      })
    },
    // 关闭对话框并重置表单
    cancelForm() {
      this.resetForm()
    },
    resetForm() {
      this.dialogFormVisible = false
      this.$refs['dialogForm'].resetFields()
      this.dialogFormData = {
        nodeName: '',
        username: '',
        publicIP: '',
        sshPort: 22,
        authmodel: 'password',
        status: 0,
        CreatedAt: '',
        label: '',
        creator: '',
      }
    },
    // 批量删除主机
    batchDelete() {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        this.loading = true
        const Ids = this.multipleSelection.map(x => x.ID)
        try {
          await delNode({ Ids: Ids }).then(res => {
            this.judgeResult(res)
          })
        } finally {
          this.loading = false
        }
        if (this.activeGroupId === 'all') {
          this.getTableData()
        }
      }).catch(() => {
        Message({
          showClose: true,
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    // 表格多选改变
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    // 单个删除主机
    async singleDelete(Id) {
      this.loading = true
      try {
        await delNode({ Ids: [Id] }).then(res => {
          this.judgeResult(res)
        })
      } finally {
        this.loading = false
      }
      if (this.activeGroupId === 'all') {
        this.getTableData()
      }
    },
    // 分页事件
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
  margin-bottom: 100px;
}
.group-card {
  margin: 10px;
}
.delete-popover {
  margin-left: 10px;
}
</style>
