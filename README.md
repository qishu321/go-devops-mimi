<div align="center">
<img src="https://cdn.nlark.com/yuque/0/2025/jpeg/28693706/1747205373498-713adaed-dee9-4b2d-9be2-599da02b7bc4.jpeg?x-oss-process=image%2Fformat%2Cwebp" width="80"  height="80">

<h1 align="center">Go-devops-mini</h1>
  <p align="center">
   本项目使用gin、gorm和ssh开发。提供完善的批量执行、作业管理、基础设施管理等功能，帮助基础运维同学快速、低成本、可视化、自动化的运维平台项目
    <br />
  </p>
    <br />


<p> 🐉 特别感谢xirang开源项目
     <br />
-   [(xirang)](https://github.com/eryajf/xirang.git)

</p>


<img src="https://cdn.jsdelivr.net/gh/eryajf/tu@main/img/image_20240420_214408.gif" width="800"  height="3">
</div><br>

## 功能
 - 便捷导航
  - 分类和导航的增删改查
  - 导航页的实现
 - 基础设施管理
   - 主机组的增删改
   - 主机的增删改查
   - 支持主机和主机组多对多
 - 作业管理
    - 批量执行：多主机或主机组的批量执行命令
    - 脚本执行：多主机或主机组的批量执行脚本
    - 文件分发：支持服务器文件和本地上传文件分发到多主机
    - 定时任务：可支持单次执行、间隔执行、cron表达式执行
    - 作业编排：工作流的批量执行
 - 权限控制
   - 用户权限
```
后期更新计划：
1、新增个便捷导航，可以把运维资源（跳板机、监控、jenkins、harbor等资源，添加到导航里。）（已实现）
2、k8s管理，实现多集群的切换
```

## 📖 目录结构

```
go-devops-mimi/server
├── config----------------配置文件读取
├── controller------------控制层
├── logic-----------------逻辑层
├── middleware------------中间件
├── model-----------------对象定义
├── public----------------一些公共组件与工具
├── routers---------------路由
├── service---------------服务层
├── test------------------一些测试
├── config.yml------------配置文件
└── main.go---------------程序入口
```

## 👀 功能概览

|  ![命令执行](https://cdn.nlark.com/yuque/0/2025/png/28693706/1747190533086-b379d2f4-1fb0-4f7d-8acc-1daa4ce1cacd.png?x-oss-process=image%2Fformat%2Cwebp)  | ![脚本详情](https://cdn.nlark.com/yuque/0/2025/png/28693706/1747190649953-c8b8564b-6165-40e7-a2db-7139f31cf203.png?x-oss-process=image%2Fformat%2Cwebp)     |
| :------------------------------------------------------------------------------: | -------------------------------------------------------------------------------- |
| ![查看脚本执行日志](https://cdn.nlark.com/yuque/0/2025/png/28693706/1747190696776-f638f804-01c1-457e-8de2-ccaf56ceeedc.png?x-oss-process=image%2Fformat%2Cwebp) | ![文件分发](https://cdn.nlark.com/yuque/0/2025/png/28693706/1747190931557-ff3158ec-15f9-4b50-8063-45bcc8c7332b.png?x-oss-process=image%2Fformat%2Cwebp) |
| ![创建任务，添加子任务，绑定主机，可以添加环境变量](https://cdn.nlark.com/yuque/0/2025/png/28693706/1747190970578-2177cc28-24a0-4626-b9e5-4da7f1f5fe9a.png?x-oss-process=image%2Fformat%2Cwebp) | ![定时任务](https://cdn.nlark.com/yuque/0/2025/png/28693706/1747191018479-18527802-045f-4a48-b4f5-7a05b8cbbab0.png?x-oss-process=image%2Fformat%2Cwebp) |
| ![分类和导航的增删改查](https://cdn.nlark.com/yuque/0/2025/png/28693706/1747620842841-0b7e457e-13a5-4d4a-aa7c-1874581f2e11.png?x-oss-process=image%2Fformat%2Cwebp) | ![导航页](https://cdn.nlark.com/yuque/0/2025/png/28693706/1747620817372-d3635ec6-8eb5-4eed-a96d-5f666d569026.png?x-oss-process=image%2Fformat%2Cwebp) |

## 🚀 快速开始

go-devops-mimi 项目的基础依赖项只有 MySQL，本地准备好这个服务之后，就可以启动项目，进行调试。


### 拉取代码

```sh
# 后端代码
$ git clone https://github.com/qishu321/go-devops-mimi.git

# 前端代码
$ git clone https://github.com/qishu321/go-devops-mimi-ui.git
```

### 更改配置

```sh
# 修改后端配置
$ cd go-devops-mimi/server
# 文件路径 config.yml, 根据自己本地的情况，调整数据库等配置信息。
$ vim config.yml
```

### 启动服务

```sh
# 启动后端
$ cd go-devops-mimi/server
$ go mod tidy
$ make run

# 启动前端
$ cd go-devops-mimi-ui
$ git config --global url."https://".insteadOf git://
$ npm install --registry=http://registry.npmmirror.com
$ yarn dev
```

本地访问：http://localhost:8090，用户名/密码：admin/123456
