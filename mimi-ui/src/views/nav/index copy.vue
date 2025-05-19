<template>
    <div class="nav-container">
      <div class="header">
        <el-button type="primary" @click="showAddDialog">添加类别</el-button>
      </div>
  
      <div class="nav-list">
        <div v-for="item in navList" :key="item.id" class="nav-item">
          <div class="nav-header">
            <span>{{ item.name }}</span>
            <div class="actions">
              <el-button size="mini" @click="showAddLinkDialog(item.id)">添加链接</el-button>
              <el-dropdown @command="cmd => handleCommand(cmd, item.id)">
                <span class="el-dropdown-link">
                  <i class="el-icon-more"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item :command="{ type: 'edit' }">编辑</el-dropdown-item>
                  <el-dropdown-item :command="{ type: 'delete' }" divided>删除</el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </div>
          </div>
  
          <div class="link-list" v-if="item.links && item.links.length">
            <div v-for="link in item.links" :key="link.id" class="link-item">
              <div class="link-content">
                <!-- 首字母头像 -->
                <div class="avatar" :style="{ backgroundColor: avatarColor(link.name) }">
                  {{ link.name.charAt(0).toUpperCase() }}
                </div>
                <div class="link-text">
                  <div class="link-name">{{ link.name }}</div>
                  <a :href="link.url" target="_blank" class="link-url">{{ link.url }}</a>
                </div>
              </div>
              <div class="link-actions">
                <el-button size="mini" @click="editLink(link)">编辑</el-button>
                <el-button size="mini" type="danger" @click="deleteLink(link.id)">删除</el-button>
              </div>
            </div>
          </div>
        </div>
      </div>
  
      <!-- 添加/编辑类别对话框 -->
      <el-dialog :title="dialogTitle" :visible.sync="dialogVisible" width="30%">
        <el-form :model="form" :rules="rules" ref="form">
          <el-form-item label="分类名称" prop="name">
            <el-input v-model="form.name" />
          </el-form-item>
          <el-form-item label="排序" prop="nav_sort">
            <el-input-number v-model="form.nav_sort" :min="0" />
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitForm">确定</el-button>
        </span>
      </el-dialog>
  
      <!-- 添加/编辑链接对话框 -->
      <el-dialog :title="linkDialogTitle" :visible.sync="linkDialogVisible" width="40%">
        <el-form :model="linkForm" :rules="linkRules" ref="linkForm">
          <el-form-item label="链接名称" prop="name">
            <el-input v-model="linkForm.name" />
          </el-form-item>
          <el-form-item label="链接描述" prop="desc">
            <el-input v-model="linkForm.desc" type="textarea" />
          </el-form-item>
          <el-form-item label="跳转地址" prop="url">
            <el-input v-model="linkForm.url" />
          </el-form-item>
          <el-form-item label="图标" prop="icon">
            <el-input v-model="linkForm.icon" placeholder="输入图标类名或URL" />
          </el-form-item>
          <el-form-item label="排序" prop="link_sort">
            <el-input-number v-model="linkForm.link_sort" :min="0" />
          </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
          <el-button @click="linkDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitLinkForm">确定</el-button>
        </span>
      </el-dialog>
    </div>
  </template>
  
  <script>
  import {
    listNav,
    addNav,
    updateNav,
    delAllNav,
    addLink,
    updateLink,
    delLink
  } from '@/api/nav/nav'
  
  export default {
    name: 'LetterAvaterNav',
    data() {
      return {
        navList: [],
        // 类别对话框
        dialogVisible: false,
        dialogTitle: '添加类别',
        form: { name: '', nav_sort: 0 },
        rules: { name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }] },
        isEdit: false,
        // 链接对话框
        linkDialogVisible: false,
        linkDialogTitle: '添加链接',
        linkForm: {
          name: '',
          desc: '',
          url: '',
          icon: '',
          link_sort: 0,
          navId: null,
          id: null
        },
        linkRules: {
          name: [{ required: true, message: '请输入链接名称', trigger: 'blur' }],
          url: [
            { required: true, message: '请输入跳转地址', trigger: 'blur' },
            { type: 'url', message: '请输入有效的URL地址', trigger: 'blur' }
          ]
        },
        isLinkEdit: false
      }
    },
    created() {
      this.getNavList()
    },
    methods: {
      async getNavList() {
        const res = await listNav()
        this.navList = res.data.navLists.map(item => ({ ...item, links: item.t_link_s }))
      },
      showAddDialog() {
        this.dialogTitle = '添加类别'
        this.isEdit = false
        this.form = { name: '', nav_sort: 0 }
        this.dialogVisible = true
      },
      editCategory(id) {
        const cat = this.navList.find(i => i.id === id)
        this.dialogTitle = '编辑类别'
        this.isEdit = true
        this.form = { ...cat }
        this.dialogVisible = true
      },
      async deleteCategory(id) {
        try {
          await this.$confirm('确认删除该分类吗?', '提示', { type: 'warning' })
          await delAllNav({ id })
          this.$message.success('删除成功')
          this.getNavList()
        } catch {}
      },
      handleCommand(cmd, id) {
        cmd.type === 'edit' ? this.editCategory(id) : this.deleteCategory(id)
      },
      submitForm() {
        this.$refs.form.validate(async valid => {
          if (!valid) return
          if (this.isEdit) await updateNav(this.form)
          else await addNav(this.form)
          this.$message.success(this.isEdit ? '更新成功' : '添加成功')
          this.dialogVisible = false
          this.getNavList()
        })
      },
      showAddLinkDialog(navId) {
        this.linkDialogTitle = '添加链接'
        this.isLinkEdit = false
        this.linkForm = { name: '', desc: '', url: '', icon: '', link_sort: 0, navId, id: null }
        this.linkDialogVisible = true
      },
      editLink(link) {
        this.linkDialogTitle = '编辑链接'
        this.isLinkEdit = true
        this.linkForm = { ...link, navId: link.nav_id || link.navId || this.findNavId(link.id) }
        this.linkDialogVisible = true
      },
      findNavId(linkId) {
        const nav = this.navList.find(n => n.t_link_s.some(l => l.id === linkId))
        return nav ? nav.id : null
      },
      async deleteLink(id) {
        try {
          await this.$confirm('确认删除该链接吗?', '提示', { type: 'warning' })
          await delLink({ id })
          this.$message.success('删除成功')
          this.getNavList()
        } catch {}
      },
      submitLinkForm() {
        this.$refs.linkForm.validate(async valid => {
          if (!valid) return
          if (this.isLinkEdit) await updateLink(this.linkForm)
          else await addLink(this.linkForm)
          this.$message.success(this.isLinkEdit ? '更新成功' : '添加成功')
          this.linkDialogVisible = false
          this.getNavList()
        })
      },
      // 根据名称生成头像背景色
      avatarColor(text) {
        const colors = ['#409EFF', '#67C23A', '#E6A23C', '#F56C6C', '#909399']
        const code = text.charCodeAt(0) || 0
        return colors[code % colors.length]
      }
    }
  }
  </script>
  
  <style scoped>
  .nav-container {
    padding: 20px;
    background: #f5f7fa;
  }
  .header {
    margin-bottom: 20px;
  }
  .nav-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(480px, 1fr));
    gap: 20px;
  }
  .nav-item {
    background: #fff;
    border-radius: 12px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
    padding: 16px;
    transition: box-shadow 0.3s;
  }
  .nav-item:hover {
    box-shadow: 0 4px 14px rgba(0, 0, 0, 0.1);
  }
  .nav-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
  }
  .nav-header span {
    font-weight: bold;
    font-size: 16px;
  }
  .actions {
    display: flex;
    align-items: center;
    gap: 8px;
  }
  .link-list {
    margin-top: 12px;
  }
  .link-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px;
    background: #f9f9f9;
    border-radius: 8px;
    margin-bottom: 8px;
    transition: background-color 0.3s;
  }
  .link-item:hover {
    background: #f0f0f0;
  }
  .link-content {
    display: flex;
    align-items: center;
    gap: 12px;
  }
  .avatar {
    width: 32px;
    height: 32px;
    border-radius: 50%;
    color: #fff;
    font-weight: bold;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 14px;
  }
  .link-text {
    display: flex;
    flex-direction: column;
  }
  .link-name {
    font-weight: 500;
    font-size: 14px;
  }
  .link-url {
    font-size: 12px;
    color: #999;
    text-decoration: none;
  }
  .link-actions {
    display: flex;
    gap: 6px;
  }
  .el-dropdown-link {
    cursor: pointer;
    color: #606266;
  }
  </style>
  