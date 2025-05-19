<template>
    <div class="dashboard-container">
      <!-- 顶部导航 -->
      <div class="header-nav">
        <div 
          v-for="nav in topNavList"
          :key="nav.id"
          class="nav-tab"
          :style="{ borderColor: getColor(nav.id, 'nav') }"
        >
          {{ nav.name }}
        </div>
      </div>
  
      <!-- 主内容区 -->
      <div class="main-content">
        <div 
          v-for="category in categoryList"
          :key="category.id"
          class="category-group"
        >
          <h3 class="group-title">{{ category.name }}</h3>
          <div class="tool-grid">
            <div
              v-for="tool in category.links"
              :key="tool.id"
              class="tool-card"
              @click="openTool(tool.url)"
            >
              <div 
                class="tool-avatar"
                :style="{ backgroundColor: getColor(tool.name, 'icon') }"
              >
                {{ tool.name.charAt(0).toUpperCase() }}
              </div>
              <div class="tool-info">
                <div class="tool-name">{{ tool.name }}</div>
                <div class="tool-desc">{{ tool.desc }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { listNav } from '@/api/nav/nav'
  
  export default {
    name: 'ToolDashboard',
    data() {
      return {
        navData: []
      }
    },
    computed: {
      topNavList() {
        return this.navData.filter(item => item.isTop) || []
      },
      categoryList() {
        return this.navData.filter(item => !item.isTop) || []
      }
    },
    async created() {
      await this.fetchData()
    },
    methods: {
      async fetchData() {
        try {
          const res = await listNav()
          this.navData = res.data.navLists.map(item => ({
            ...item,
            links: item.t_link_s || []
          }))
        } catch (error) {
          console.error('数据加载失败:', error)
        }
      },
      getColor(target, type) {
        const navColors = ['#F56C6C', '#67C23A']
        const iconColors = ['#409EFF', '#67C23A', '#E6A23C', '#F56C6C']
        
        if (type === 'nav') {
          return navColors[target % navColors.length]
        }
        return iconColors[target.charCodeAt(0) % iconColors.length]
      },
      openTool(url) {
        window.open(url, '_blank')
      }
    }
  }
  </script>
  
  <style scoped>
  .dashboard-container {
    background: #ffffff;
    min-height: 100vh;
  }
  
  /* 顶部导航 */
  .header-nav {
    display: flex;
    padding: 16px 24px;
    background: #f8f9fa;
    border-bottom: 1px solid #e4e7ed;
  }
  
  .nav-tab {
    padding: 8px 16px;
    margin-right: 24px;
    border-left: 3px solid;
    font-size: 15px;
    font-weight: 500;
    color: #303133;
  }
  
  /* 主内容区 */
  .main-content {
    padding: 24px;
    max-width: 1600px;
    margin: 0 auto;
  }
  
  /* 分类组 */
  .category-group {
    margin-bottom: 32px;
  }
  
  .group-title {
    font-size: 18px;
    font-weight: 600;
    color: #303133;
    margin-bottom: 16px;
    padding-bottom: 8px;
    border-bottom: 1px solid #ebeef5;
  }
  
  /* 工具网格 */
  .tool-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 16px;
  }
  
  /* 工具卡片 */
  .tool-card {
    display: flex;
    align-items: center;
    padding: 16px;
    border: 1px solid #ebeef5;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
    background: #fff;
  }
  
  .tool-card:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
    transform: translateY(-2px);
    border-color: #409EFF;
  }
  
  /* 工具图标 */
  .tool-avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    color: white;
    font-weight: bold;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 16px;
    flex-shrink: 0;
    margin-right: 16px;
  }
  
  /* 工具信息 */
  .tool-info {
    flex: 1;
    min-width: 0;
  }
  
  .tool-name {
    font-size: 16px;
    font-weight: 600;
    color: #303133;
    line-height: 1.4;
    margin-bottom: 4px;
  }
  
  .tool-desc {
    font-size: 13px;
    color: #606266;
    line-height: 1.5;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
  }
  
  @media (max-width: 768px) {
    .tool-grid {
      grid-template-columns: 1fr;
    }
    
    .main-content {
      padding: 16px;
    }
    
    .tool-card {
      padding: 12px;
    }
  }
  </style>