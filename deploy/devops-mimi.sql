/*
Navicat MySQL Data Transfer

Source Server         : 192.168.2.112
Source Server Version : 80300
Source Host           : 192.168.2.112:3310
Source Database       : devops-mimi

Target Server Type    : MYSQL
Target Server Version : 80300
File Encoding         : 65001

Date: 2025-05-20 16:19:41
*/
CREATE DATABASE IF NOT EXISTS `devops-mimi` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `devops-mimi`;
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for apis
-- ----------------------------
DROP TABLE IF EXISTS `apis`;
CREATE TABLE `apis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `method` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''请求方式''',
  `path` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''访问路径''',
  `category` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''所属类别''',
  `remark` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''备注''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  KEY `idx_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of apis
-- ----------------------------
INSERT INTO `apis` VALUES ('1', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/base/login', 'base', '用户登录', '系统');
INSERT INTO `apis` VALUES ('2', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/base/logout', 'base', '用户登出', '系统');
INSERT INTO `apis` VALUES ('3', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/base/refreshToken', 'base', '刷新JWT令牌', '系统');
INSERT INTO `apis` VALUES ('4', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/user/info', 'user', '获取当前登录用户信息', '系统');
INSERT INTO `apis` VALUES ('5', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/user/list', 'user', '获取用户列表', '系统');
INSERT INTO `apis` VALUES ('6', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/user/changePwd', 'user', '更新用户登录密码', '系统');
INSERT INTO `apis` VALUES ('7', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/user/add', 'user', '创建用户', '系统');
INSERT INTO `apis` VALUES ('8', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/user/update', 'user', '更新用户', '系统');
INSERT INTO `apis` VALUES ('9', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/user/delete', 'user', '批量删除用户', '系统');
INSERT INTO `apis` VALUES ('10', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/user/changeUserStatus', 'user', '更改用户状态', '系统');
INSERT INTO `apis` VALUES ('11', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/group/list', 'group', '获取分组列表', '系统');
INSERT INTO `apis` VALUES ('12', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/group/tree', 'group', '获取分组列表树', '系统');
INSERT INTO `apis` VALUES ('13', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/group/add', 'group', '创建分组', '系统');
INSERT INTO `apis` VALUES ('14', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/group/update', 'group', '更新分组', '系统');
INSERT INTO `apis` VALUES ('15', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/group/delete', 'group', '批量删除分组', '系统');
INSERT INTO `apis` VALUES ('16', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/group/adduser', 'group', '添加用户到分组', '系统');
INSERT INTO `apis` VALUES ('17', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/group/removeuser', 'group', '将用户从分组移出', '系统');
INSERT INTO `apis` VALUES ('18', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/group/useringroup', 'group', '获取在分组内的用户列表', '系统');
INSERT INTO `apis` VALUES ('19', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/group/usernoingroup', 'group', '获取不在分组内的用户列表', '系统');
INSERT INTO `apis` VALUES ('20', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/role/list', 'role', '获取角色列表', '系统');
INSERT INTO `apis` VALUES ('21', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/role/add', 'role', '创建角色', '系统');
INSERT INTO `apis` VALUES ('22', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/role/update', 'role', '更新角色', '系统');
INSERT INTO `apis` VALUES ('23', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/role/getmenulist', 'role', '获取角色的权限菜单', '系统');
INSERT INTO `apis` VALUES ('24', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/role/updatemenus', 'role', '更新角色的权限菜单', '系统');
INSERT INTO `apis` VALUES ('25', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/role/getapilist', 'role', '获取角色的权限接口', '系统');
INSERT INTO `apis` VALUES ('26', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/role/updateapis', 'role', '更新角色的权限接口', '系统');
INSERT INTO `apis` VALUES ('27', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/role/delete', 'role', '批量删除角色', '系统');
INSERT INTO `apis` VALUES ('28', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/menu/list', 'menu', '获取菜单列表', '系统');
INSERT INTO `apis` VALUES ('29', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/menu/tree', 'menu', '获取菜单树', '系统');
INSERT INTO `apis` VALUES ('30', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/menu/access/tree', 'menu', '获取用户菜单树', '系统');
INSERT INTO `apis` VALUES ('31', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/menu/add', 'menu', '创建菜单', '系统');
INSERT INTO `apis` VALUES ('32', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/menu/update', 'menu', '更新菜单', '系统');
INSERT INTO `apis` VALUES ('33', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/menu/delete', 'menu', '批量删除菜单', '系统');
INSERT INTO `apis` VALUES ('34', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/api/list', 'api', '获取接口列表', '系统');
INSERT INTO `apis` VALUES ('35', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/api/tree', 'api', '获取接口树', '系统');
INSERT INTO `apis` VALUES ('36', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/api/add', 'api', '创建接口', '系统');
INSERT INTO `apis` VALUES ('37', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/api/update', 'api', '更新接口', '系统');
INSERT INTO `apis` VALUES ('38', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/api/delete', 'api', '批量删除接口', '系统');
INSERT INTO `apis` VALUES ('39', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/system/log/operation/list', 'log', '获取操作日志列表', '系统');
INSERT INTO `apis` VALUES ('40', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/system/log/operation/delete', 'log', '批量删除操作日志', '系统');
INSERT INTO `apis` VALUES ('41', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'GET', '/example/cloudaccount/list', 'cloudAccount', '获取云账户列表', '系统');
INSERT INTO `apis` VALUES ('42', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/example/cloudaccount/add', 'cloudAccount', '添加云账户', '系统');
INSERT INTO `apis` VALUES ('43', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/example/cloudaccount/update', 'cloudAccount', '更新云账户信息', '系统');
INSERT INTO `apis` VALUES ('44', '2025-04-07 11:03:12.882', '2025-04-07 11:03:12.882', null, 'POST', '/example/cloudaccount/delete', 'cloudAccount', '批量删除云账户', '系统');
INSERT INTO `apis` VALUES ('45', '2025-04-07 11:52:45.002', '2025-04-07 11:52:45.002', null, 'GET', '/cmdb/node/list', 'node', '获取主机列表', 'admin');
INSERT INTO `apis` VALUES ('46', '2025-04-07 11:52:58.152', '2025-04-07 11:52:58.152', null, 'POST', '/cmdb/node/add', 'node', '新增主机', 'admin');
INSERT INTO `apis` VALUES ('47', '2025-04-07 11:53:11.992', '2025-04-07 11:53:11.992', null, 'POST', '/cmdb/node/update', 'node', '修改主机', 'admin');
INSERT INTO `apis` VALUES ('48', '2025-04-07 11:53:23.251', '2025-04-07 11:53:23.251', null, 'POST', '/cmdb/node/delete', 'node', '删除主机', 'admin');
INSERT INTO `apis` VALUES ('49', '2025-04-07 11:53:41.113', '2025-04-07 11:53:41.113', null, 'GET', '/cmdb/node_group/list', 'node', '获取主机组列表', 'admin');
INSERT INTO `apis` VALUES ('50', '2025-04-07 11:53:52.167', '2025-04-07 11:53:52.167', null, 'POST', '/cmdb/node_group/add', 'node', '新增主机组', 'admin');
INSERT INTO `apis` VALUES ('51', '2025-04-07 11:54:03.712', '2025-04-07 11:54:03.712', null, 'POST', '/cmdb/node_group/update', 'node', '修改主机组', 'admin');
INSERT INTO `apis` VALUES ('52', '2025-04-07 11:54:14.897', '2025-04-07 11:54:14.897', null, 'POST', '/cmdb/node_group/delete', 'node', '删除主机组', 'admin');
INSERT INTO `apis` VALUES ('53', '2025-04-07 11:54:26.383', '2025-04-07 11:54:26.383', null, 'POST', '/cmdb/node_group/add_node_to_group', 'node', '添加主机到主机组', 'admin');
INSERT INTO `apis` VALUES ('54', '2025-04-07 11:54:36.999', '2025-04-07 11:54:36.999', null, 'POST', '/cmdb/node_group/remonv_node_to_group', 'node', '从主机组移除主机', 'admin');
INSERT INTO `apis` VALUES ('55', '2025-04-09 17:15:39.628', '2025-04-09 17:15:39.628', null, 'POST', '/cmdb/node/add_nodes_group', 'node', '多对多添加主机到主机组', 'admin');
INSERT INTO `apis` VALUES ('56', '2025-04-15 16:37:08.322', '2025-04-15 16:38:54.873', null, 'POST', '/exec/script/add_run', 'exec', '执行脚本', 'admin');
INSERT INTO `apis` VALUES ('57', '2025-04-16 17:39:02.345', '2025-04-16 17:39:02.345', null, 'GET', '/exec/script/list', 'exec', '查看执行脚本记录', 'admin');
INSERT INTO `apis` VALUES ('58', '2025-04-16 17:39:23.544', '2025-04-16 17:39:23.544', null, 'POST', '/exec/script_library/add', 'exec', '新增脚本', 'admin');
INSERT INTO `apis` VALUES ('59', '2025-04-16 17:39:41.572', '2025-04-16 17:39:41.572', null, 'POST', '/exec/script_library/update', 'exec', '修改脚本', 'admin');
INSERT INTO `apis` VALUES ('60', '2025-04-16 17:40:20.559', '2025-04-16 17:40:20.559', null, 'POST', '/exec/script_library/delete', 'exec', '删除脚本', 'admin');
INSERT INTO `apis` VALUES ('61', '2025-04-16 17:40:48.166', '2025-04-16 17:40:48.166', null, 'GET', '/exec/script_library/list', 'exec', '查看脚本列表', 'admin');
INSERT INTO `apis` VALUES ('62', '2025-04-16 17:41:09.807', '2025-04-16 17:41:09.807', null, 'GET', '/exec/script_library/info', 'exec', '查看指定脚本详情', 'admin');
INSERT INTO `apis` VALUES ('63', '2025-04-21 17:09:00.146', '2025-04-21 17:09:00.146', null, 'POST', '/exec/transfer/add_run', 'exec', '执行文件分发', 'admin');
INSERT INTO `apis` VALUES ('64', '2025-04-21 17:09:24.522', '2025-04-21 17:09:24.522', null, 'POST', '/exec/transfer/upload', 'exec', '上传文件', 'admin');
INSERT INTO `apis` VALUES ('65', '2025-04-21 17:09:46.608', '2025-04-21 17:09:46.608', null, 'GET', '/exec/transfer/info', 'exec', '查看指定分发记录详情', 'admin');
INSERT INTO `apis` VALUES ('66', '2025-04-21 17:10:08.063', '2025-04-21 17:10:08.063', null, 'GET', '/exec/transfer/list', 'exec', '查看分发记录列表', 'admin');
INSERT INTO `apis` VALUES ('67', '2025-04-23 16:09:14.096', '2025-04-23 16:09:14.096', null, 'POST', '/exec/task_manage/add', 'exec', '新建任务管理', 'admin');
INSERT INTO `apis` VALUES ('68', '2025-04-23 16:09:33.219', '2025-04-23 16:09:33.219', null, 'POST', '/exec/task_manage/update', 'exec', '修改任务管理', 'admin');
INSERT INTO `apis` VALUES ('69', '2025-04-23 16:09:59.498', '2025-04-23 16:09:59.498', null, 'GET', '/exec/task_manage/list', 'exec', '查看任务列表', 'admin');
INSERT INTO `apis` VALUES ('70', '2025-04-23 16:10:20.323', '2025-04-23 16:10:20.323', null, 'GET', '/exec/task_manage/info', 'exec', '查看指定任务详情', 'admin');
INSERT INTO `apis` VALUES ('71', '2025-04-23 16:33:57.920', '2025-04-23 16:33:57.920', null, 'POST', '/exec/task_manage/delete', 'exec', '删除任务列表', 'admin');
INSERT INTO `apis` VALUES ('72', '2025-04-25 17:08:43.149', '2025-04-25 17:08:43.149', null, 'GET', '/exec/run_task_manage/run', 'exec', '执行任务', 'admin');
INSERT INTO `apis` VALUES ('73', '2025-04-30 11:47:31.916', '2025-04-30 11:48:13.411', null, 'GET', '/exec/run_task_manage/run_info', 'exec', '查看执行任务日志', 'admin');
INSERT INTO `apis` VALUES ('74', '2025-04-30 11:48:05.846', '2025-04-30 11:48:05.846', null, 'POST', '/exec/run_task_manage/add_run', 'exec', '执行任务新', 'admin');
INSERT INTO `apis` VALUES ('75', '2025-04-30 18:42:44.614', '2025-04-30 18:42:44.614', null, 'GET', '/exec/run_task_manage/run_info_webSocket', 'exec', 'run_info_webSocket', 'admin');
INSERT INTO `apis` VALUES ('76', '2025-05-07 11:56:19.089', '2025-05-07 11:56:19.089', null, 'GET', '/exec/run_task_manage/run_list', 'exec', '获取任务执行记录', 'admin');
INSERT INTO `apis` VALUES ('77', '2025-05-09 10:44:49.272', '2025-05-09 10:44:49.272', null, 'POST', '/exec/cron/add', 'cron', '新增定时任务', 'admin');
INSERT INTO `apis` VALUES ('78', '2025-05-09 10:45:22.319', '2025-05-09 10:45:22.319', null, 'POST', '/exec/cron/update', 'cron', '修改定时任务', 'admin');
INSERT INTO `apis` VALUES ('79', '2025-05-09 10:45:40.067', '2025-05-09 10:45:40.067', null, 'POST', '/exec/cron/delete', 'cron', '删除定时任务', 'admin');
INSERT INTO `apis` VALUES ('80', '2025-05-09 10:46:13.700', '2025-05-09 10:46:13.700', null, 'POST', '/exec/cron/enable', 'cron', '启用定时任务', 'admin');
INSERT INTO `apis` VALUES ('81', '2025-05-09 10:47:23.078', '2025-05-09 10:47:23.078', null, 'GET', '/exec/cron/list', 'cron', '定时任务列表', 'admin');
INSERT INTO `apis` VALUES ('82', '2025-05-09 10:48:12.343', '2025-05-09 10:48:12.343', null, 'GET', '/exec/cron/info', 'cron', '定时任务详情', 'admin');
INSERT INTO `apis` VALUES ('83', '2025-05-13 17:03:18.217', '2025-05-13 17:03:18.217', null, 'GET', '/exec/cron/log/list', 'cron', '查看定时任务执行日志', 'admin');
INSERT INTO `apis` VALUES ('84', '2025-05-15 17:25:30.884', '2025-05-15 17:25:30.884', null, 'GET', '/nav/list', 'nav', '导航分类', 'admin');
INSERT INTO `apis` VALUES ('85', '2025-05-15 17:25:50.587', '2025-05-15 17:25:50.587', null, 'POST', '/nav/add', 'nav', '新增导航分类', 'admin');
INSERT INTO `apis` VALUES ('86', '2025-05-15 17:26:13.874', '2025-05-15 17:26:13.874', null, 'GET', '/nav/info', 'nav', '指定导航详情', 'admin');
INSERT INTO `apis` VALUES ('87', '2025-05-15 17:26:38.912', '2025-05-15 17:26:38.912', null, 'POST', '/nav/update', 'nav', '修改导航分类', 'admin');
INSERT INTO `apis` VALUES ('88', '2025-05-15 17:27:00.653', '2025-05-15 17:27:00.653', null, 'POST', '/nav/delete', 'nav', '删除导航分类', 'admin');
INSERT INTO `apis` VALUES ('89', '2025-05-15 17:28:11.787', '2025-05-15 17:28:11.787', null, 'POST', '/nav/delete_all', 'nav', '删除导航分类及关联链接', 'admin');
INSERT INTO `apis` VALUES ('90', '2025-05-15 17:54:19.524', '2025-05-15 17:54:19.524', null, 'POST', '/nav/link/add', 'nav', '新增link', 'admin');
INSERT INTO `apis` VALUES ('91', '2025-05-15 17:54:34.930', '2025-05-15 17:54:34.930', null, 'GET', '/nav/link/info', 'nav', '获取link详情', 'admin');
INSERT INTO `apis` VALUES ('92', '2025-05-15 17:54:51.780', '2025-05-15 17:54:51.780', null, 'POST', '/nav/link/update', 'nav', '更新link', 'admin');
INSERT INTO `apis` VALUES ('93', '2025-05-15 17:55:16.803', '2025-05-15 17:55:16.803', null, 'POST', '/nav/link/delete', 'nav', '删除link', 'admin');

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_index` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=2367 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('2319', 'p', 'admin', '/cmdb/node/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2328', 'p', 'admin', '/cmdb/node/add_nodes_group', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2321', 'p', 'admin', '/cmdb/node/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2318', 'p', 'admin', '/cmdb/node/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2320', 'p', 'admin', '/cmdb/node/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2323', 'p', 'admin', '/cmdb/node_group/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2326', 'p', 'admin', '/cmdb/node_group/add_node_to_group', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2325', 'p', 'admin', '/cmdb/node_group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2322', 'p', 'admin', '/cmdb/node_group/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2327', 'p', 'admin', '/cmdb/node_group/remonv_node_to_group', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2324', 'p', 'admin', '/cmdb/node_group/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2315', 'p', 'admin', '/example/cloudaccount/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2317', 'p', 'admin', '/example/cloudaccount/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2314', 'p', 'admin', '/example/cloudaccount/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2316', 'p', 'admin', '/example/cloudaccount/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2350', 'p', 'admin', '/exec/cron/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2352', 'p', 'admin', '/exec/cron/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2353', 'p', 'admin', '/exec/cron/enable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2355', 'p', 'admin', '/exec/cron/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2354', 'p', 'admin', '/exec/cron/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2356', 'p', 'admin', '/exec/cron/log/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2351', 'p', 'admin', '/exec/cron/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2347', 'p', 'admin', '/exec/run_task_manage/add_run', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2345', 'p', 'admin', '/exec/run_task_manage/run', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2346', 'p', 'admin', '/exec/run_task_manage/run_info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2348', 'p', 'admin', '/exec/run_task_manage/run_info_webSocket', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2349', 'p', 'admin', '/exec/run_task_manage/run_list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2329', 'p', 'admin', '/exec/script/add_run', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2330', 'p', 'admin', '/exec/script/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2331', 'p', 'admin', '/exec/script_library/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2333', 'p', 'admin', '/exec/script_library/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2335', 'p', 'admin', '/exec/script_library/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2334', 'p', 'admin', '/exec/script_library/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2332', 'p', 'admin', '/exec/script_library/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2340', 'p', 'admin', '/exec/task_manage/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2344', 'p', 'admin', '/exec/task_manage/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2343', 'p', 'admin', '/exec/task_manage/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2342', 'p', 'admin', '/exec/task_manage/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2341', 'p', 'admin', '/exec/task_manage/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2336', 'p', 'admin', '/exec/transfer/add_run', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2338', 'p', 'admin', '/exec/transfer/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2339', 'p', 'admin', '/exec/transfer/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2337', 'p', 'admin', '/exec/transfer/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2358', 'p', 'admin', '/nav/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2361', 'p', 'admin', '/nav/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2362', 'p', 'admin', '/nav/delete_all', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2359', 'p', 'admin', '/nav/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2363', 'p', 'admin', '/nav/link/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2366', 'p', 'admin', '/nav/link/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2364', 'p', 'admin', '/nav/link/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2365', 'p', 'admin', '/nav/link/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2357', 'p', 'admin', '/nav/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2360', 'p', 'admin', '/nav/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2309', 'p', 'admin', '/system/api/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2311', 'p', 'admin', '/system/api/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2307', 'p', 'admin', '/system/api/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2308', 'p', 'admin', '/system/api/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2310', 'p', 'admin', '/system/api/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2274', 'p', 'admin', '/system/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2275', 'p', 'admin', '/system/base/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2276', 'p', 'admin', '/system/base/refreshToken', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2286', 'p', 'admin', '/system/group/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2289', 'p', 'admin', '/system/group/adduser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2288', 'p', 'admin', '/system/group/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2284', 'p', 'admin', '/system/group/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2290', 'p', 'admin', '/system/group/removeuser', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2285', 'p', 'admin', '/system/group/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2287', 'p', 'admin', '/system/group/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2291', 'p', 'admin', '/system/group/useringroup', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2292', 'p', 'admin', '/system/group/usernoingroup', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2313', 'p', 'admin', '/system/log/operation/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2312', 'p', 'admin', '/system/log/operation/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2303', 'p', 'admin', '/system/menu/access/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2304', 'p', 'admin', '/system/menu/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2306', 'p', 'admin', '/system/menu/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2301', 'p', 'admin', '/system/menu/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2302', 'p', 'admin', '/system/menu/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2305', 'p', 'admin', '/system/menu/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2294', 'p', 'admin', '/system/role/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2300', 'p', 'admin', '/system/role/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2298', 'p', 'admin', '/system/role/getapilist', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2296', 'p', 'admin', '/system/role/getmenulist', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2293', 'p', 'admin', '/system/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2295', 'p', 'admin', '/system/role/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2299', 'p', 'admin', '/system/role/updateapis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2297', 'p', 'admin', '/system/role/updatemenus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2280', 'p', 'admin', '/system/user/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2279', 'p', 'admin', '/system/user/changePwd', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2283', 'p', 'admin', '/system/user/changeUserStatus', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2282', 'p', 'admin', '/system/user/delete', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2277', 'p', 'admin', '/system/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2278', 'p', 'admin', '/system/user/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2281', 'p', 'admin', '/system/user/update', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('2', 'p', 'user', '/system/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('4', 'p', 'user', '/system/base/logout', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('6', 'p', 'user', '/system/base/refreshToken', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('46', 'p', 'user', '/system/log/operation/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('36', 'p', 'user', '/system/menu/access/tree', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('11', 'p', 'user', '/system/user/changePwd', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('8', 'p', 'user', '/system/user/info', 'GET', '', '', '');

-- ----------------------------
-- Table structure for cloud_account
-- ----------------------------
DROP TABLE IF EXISTS `cloud_account`;
CREATE TABLE `cloud_account` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `cloud_name` varchar(24) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '云账号名称',
  `cloud_type` varchar(10) COLLATE utf8mb4_general_ci DEFAULT 'tx' COMMENT '云厂商名称',
  `secret_id` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '访问秘钥ID',
  `secret_key` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '访问秘钥key',
  `remark` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '说明',
  PRIMARY KEY (`id`),
  KEY `idx_cloud_account_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of cloud_account
-- ----------------------------

-- ----------------------------
-- Table structure for groups
-- ----------------------------
DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''分组名称''',
  `remark` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''分组中文说明''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '''父组编号(编号为0时表示根组)''',
  `source` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''来源：dingTalk、weCom、ldap、platform''',
  PRIMARY KEY (`id`),
  KEY `idx_groups_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of groups
-- ----------------------------
INSERT INTO `groups` VALUES ('1', '2025-04-07 11:03:13.022', '2025-04-07 17:25:28.193', null, 'root', '根部门', 'system', '0', 'system');
INSERT INTO `groups` VALUES ('2', '2025-04-07 11:03:13.022', '2025-04-07 11:03:13.022', null, 'backend', '后端部', 'system', '1', 'system');
INSERT INTO `groups` VALUES ('3', '2025-04-07 11:03:13.022', '2025-04-07 11:03:13.022', null, 'test', '测试部', 'system', '1', 'system');
INSERT INTO `groups` VALUES ('4', '2025-04-07 11:03:13.022', '2025-04-07 11:03:13.022', null, 'ops', '运维部', 'system', '1', 'system');

-- ----------------------------
-- Table structure for group_users
-- ----------------------------
DROP TABLE IF EXISTS `group_users`;
CREATE TABLE `group_users` (
  `group_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`group_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of group_users
-- ----------------------------
INSERT INTO `group_users` VALUES ('1', '1');

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''菜单名称(英文名, 可用于国际化)''',
  `title` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''菜单标题(无法国际化时使用)''',
  `icon` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''菜单图标''',
  `path` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''菜单访问路径''',
  `redirect` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''重定向路径''',
  `component` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''前端组件路径''',
  `sort` int DEFAULT '999' COMMENT '''菜单顺序(1-999)''',
  `status` tinyint(1) DEFAULT '1' COMMENT '''菜单状态(正常/禁用, 默认正常)''',
  `hidden` tinyint(1) DEFAULT '2' COMMENT '''菜单在侧边栏隐藏(1隐藏，2显示)''',
  `no_cache` tinyint(1) DEFAULT '2' COMMENT '''菜单是否被 <keep-alive> 缓存(1不缓存，2缓存)''',
  `always_show` tinyint(1) DEFAULT '2' COMMENT '''忽略之前定义的规则，一直显示根路由(1忽略，2不忽略)''',
  `breadcrumb` tinyint(1) DEFAULT '1' COMMENT '''面包屑可见性(可见/隐藏, 默认可见)''',
  `active_menu` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''在其它路由时，想在侧边栏高亮的路由''',
  `parent_id` bigint unsigned DEFAULT '0' COMMENT '''父菜单编号(编号为0时表示根菜单)''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  KEY `idx_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES ('1', '2025-04-07 11:03:12.758', '2025-05-15 18:30:32.242', null, 'System', '系统管理', 'system', '/system', '/system/user', 'Layout', '999', '1', '2', '2', '2', '1', '', '0', 'admin');
INSERT INTO `menus` VALUES ('2', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'User', '用户管理', 'user', 'user', '', '/system/user/index', '2', '1', '2', '2', '2', '1', '', '1', '系统');
INSERT INTO `menus` VALUES ('3', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'Group', '分组管理', 'peoples', 'group', '', '/system/group/index', '3', '1', '2', '1', '2', '1', '', '1', '系统');
INSERT INTO `menus` VALUES ('4', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'Role', '角色管理', 'eye-open', 'role', '', '/system/role/index', '4', '1', '2', '2', '2', '1', '', '1', '系统');
INSERT INTO `menus` VALUES ('5', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'Menu', '菜单管理', 'tree-table', 'menu', '', '/system/menu/index', '5', '1', '2', '2', '2', '1', '', '1', '系统');
INSERT INTO `menus` VALUES ('6', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'Api', '接口管理', 'tree', 'api', '', '/system/api/index', '6', '1', '2', '2', '2', '1', '', '1', '系统');
INSERT INTO `menus` VALUES ('7', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'Log', '日志管理', 'log', 'log', '/system/log/operationLog', '/system/log/index', '7', '1', '2', '2', '2', '1', '', '1', '系统');
INSERT INTO `menus` VALUES ('8', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'OperationLog', '操作日志', 'documentation', 'operationLog', '', '/system/log/operationLog/index', '8', '1', '2', '2', '2', '1', '', '7', '系统');
INSERT INTO `menus` VALUES ('9', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'Example', '示例模块', 'example', '/example', '', 'Layout', '9', '1', '2', '2', '2', '1', '', '0', '系统');
INSERT INTO `menus` VALUES ('10', '2025-04-07 11:03:12.758', '2025-04-07 11:03:12.758', null, 'CloudAccount', '云账户', 'peoples', 'CloudAccount', '', '/example/cloudAccount/index', '10', '1', '2', '2', '2', '1', '', '9', '系统');
INSERT INTO `menus` VALUES ('11', '2025-04-07 14:46:41.577', '2025-04-07 14:46:41.577', null, 'Cmdb', '基础设施', 'build', '/cmdb', '', 'Layout', '10', '1', '2', '2', '2', '1', '', '0', 'admin');
INSERT INTO `menus` VALUES ('12', '2025-04-07 14:48:03.881', '2025-04-07 14:48:03.881', null, 'Node', '主机管理', 'list', 'node', '', '/cmdb/node/index', '1', '1', '2', '2', '2', '1', '', '11', 'admin');
INSERT INTO `menus` VALUES ('13', '2025-04-16 18:09:25.781', '2025-04-16 18:16:16.890', null, 'Exec', '作业管理', 'cascader', '/exec', '/exec/script/command', 'Layout', '3', '1', '2', '2', '2', '1', '', '0', 'admin');
INSERT INTO `menus` VALUES ('14', '2025-04-16 18:12:03.859', '2025-04-16 18:15:34.734', null, 'Script', '批量执行', 'build', 'script', '/exec/script/command', '/exec/script/index', '1', '1', '2', '2', '2', '1', '', '13', 'admin');
INSERT INTO `menus` VALUES ('15', '2025-04-16 18:15:18.309', '2025-04-16 18:15:18.309', null, 'Cmd', '脚本执行', 'input', 'command', '', '/exec/script/command/index', '10', '1', '2', '2', '2', '1', '', '14', 'admin');
INSERT INTO `menus` VALUES ('16', '2025-04-16 18:19:43.860', '2025-04-16 18:19:43.860', null, 'Library', '脚本库', 'dict', 'library', '', '/exec/script/library/index', '11', '1', '2', '2', '2', '1', '', '14', 'admin');
INSERT INTO `menus` VALUES ('17', '2025-04-16 18:28:16.566', '2025-04-16 18:33:20.138', null, 'Script_log', '脚本执行日志', 'log', 'script_log', '', '/exec/script/script_log/index', '19', '1', '2', '2', '2', '1', '', '14', 'admin');
INSERT INTO `menus` VALUES ('18', '2025-04-21 17:13:43.937', '2025-04-21 17:13:43.937', null, 'Transfer', '文件分发', 'drag copy', 'transfer', '', '/exec/transfer/index', '2', '1', '2', '2', '2', '1', '', '13', 'admin');
INSERT INTO `menus` VALUES ('19', '2025-04-23 16:27:44.949', '2025-04-23 17:05:49.426', null, 'Task', '任务中心', 'calendar', 'task', '/exec/task/manage', '/exec/task/index', '3', '1', '2', '1', '2', '1', '', '13', 'admin');
INSERT INTO `menus` VALUES ('20', '2025-04-23 16:29:27.318', '2025-04-23 17:05:10.842', null, 'Manage', '任务管理', 'people', 'manage', '', '/exec/task/manage/index', '30', '1', '2', '1', '2', '1', '', '19', 'admin');
INSERT INTO `menus` VALUES ('21', '2025-04-27 11:33:37.828', '2025-05-07 11:29:51.949', null, 'ManageLog', '任务执行日志', 'build', 'manageLog', '', '/exec/task/manage/manage_log/index', '31', '1', '2', '2', '2', '1', '', '19', 'admin');
INSERT INTO `menus` VALUES ('22', '2025-05-09 11:32:46.971', '2025-05-09 11:32:46.971', null, 'Cron', '定时任务', 'international', 'cron', '/exec/cron/list', '/exec/cron/index', '4', '1', '2', '2', '2', '1', '', '13', 'admin');
INSERT INTO `menus` VALUES ('23', '2025-05-09 11:35:09.641', '2025-05-09 11:35:09.641', null, 'List', '定时任务管理', 'list', 'list', '', '/exec/cron/list/index', '41', '1', '2', '2', '2', '1', '', '22', 'admin');
INSERT INTO `menus` VALUES ('24', '2025-05-13 17:05:08.530', '2025-05-13 17:05:08.530', null, 'cron_log', '定时任务日志', 'logininfor', 'log', '', '/exec/cron/log/index', '49', '1', '2', '1', '2', '1', '', '22', 'admin');
INSERT INTO `menus` VALUES ('25', '2025-05-15 18:32:09.803', '2025-05-15 18:52:47.930', null, 'Nav', '便捷导航', 'icon', '/nav', '', 'Layout', '1', '1', '2', '1', '2', '1', '', '0', 'admin');
INSERT INTO `menus` VALUES ('26', '2025-05-15 18:54:30.203', '2025-05-16 18:40:59.450', null, 'Link', '导航管理', 'icon', 'Link', '', '/nav/index', '2', '1', '2', '2', '2', '1', '', '25', 'admin');

-- ----------------------------
-- Table structure for operation_logs
-- ----------------------------
DROP TABLE IF EXISTS `operation_logs`;
CREATE TABLE `operation_logs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''用户登录名''',
  `ip` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''Ip地址''',
  `ip_location` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''Ip所在地''',
  `method` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''请求方式''',
  `path` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''访问路径''',
  `remark` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''备注''',
  `status` int DEFAULT NULL COMMENT '''响应状态码''',
  `start_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''发起时间''',
  `time_cost` int DEFAULT NULL COMMENT '''请求耗时(ms)''',
  `user_agent` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''浏览器标识''',
  PRIMARY KEY (`id`),
  KEY `idx_operation_logs_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6491 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of operation_logs
-- ----------------------------

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `keyword` varchar(20) COLLATE utf8mb4_general_ci NOT NULL,
  `remark` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''备注''',
  `status` tinyint(1) DEFAULT '1' COMMENT '''1正常, 2禁用''',
  `sort` int DEFAULT '999' COMMENT '''角色排序(排序越大权限越低, 不能查看比自己序号小的角色, 不能编辑同序号用户权限, 排序为1表示超级管理员)''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `keyword` (`keyword`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES ('1', '2025-04-07 11:03:12.725', '2025-05-15 18:54:38.433', null, '管理员', 'admin', '', '1', '1', '系统');
INSERT INTO `roles` VALUES ('2', '2025-04-07 11:03:12.725', '2025-04-07 11:03:12.725', null, '普通用户', 'user', '', '1', '3', '系统');
INSERT INTO `roles` VALUES ('3', '2025-04-07 11:03:12.725', '2025-04-07 11:03:12.725', null, '访客', 'guest', '', '1', '5', '系统');

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus` (
  `menu_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`menu_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of role_menus
-- ----------------------------
INSERT INTO `role_menus` VALUES ('1', '1');
INSERT INTO `role_menus` VALUES ('2', '1');
INSERT INTO `role_menus` VALUES ('3', '1');
INSERT INTO `role_menus` VALUES ('4', '1');
INSERT INTO `role_menus` VALUES ('5', '1');
INSERT INTO `role_menus` VALUES ('6', '1');
INSERT INTO `role_menus` VALUES ('7', '1');
INSERT INTO `role_menus` VALUES ('7', '2');
INSERT INTO `role_menus` VALUES ('8', '1');
INSERT INTO `role_menus` VALUES ('8', '2');
INSERT INTO `role_menus` VALUES ('9', '2');
INSERT INTO `role_menus` VALUES ('10', '1');
INSERT INTO `role_menus` VALUES ('10', '2');
INSERT INTO `role_menus` VALUES ('11', '1');
INSERT INTO `role_menus` VALUES ('12', '1');
INSERT INTO `role_menus` VALUES ('13', '1');
INSERT INTO `role_menus` VALUES ('14', '1');
INSERT INTO `role_menus` VALUES ('15', '1');
INSERT INTO `role_menus` VALUES ('16', '1');
INSERT INTO `role_menus` VALUES ('17', '1');
INSERT INTO `role_menus` VALUES ('18', '1');
INSERT INTO `role_menus` VALUES ('19', '1');
INSERT INTO `role_menus` VALUES ('20', '1');
INSERT INTO `role_menus` VALUES ('21', '1');
INSERT INTO `role_menus` VALUES ('22', '1');
INSERT INTO `role_menus` VALUES ('23', '1');
INSERT INTO `role_menus` VALUES ('24', '1');
INSERT INTO `role_menus` VALUES ('25', '1');
INSERT INTO `role_menus` VALUES ('26', '1');

-- ----------------------------
-- Table structure for t_cron
-- ----------------------------
DROP TABLE IF EXISTS `t_cron`;
CREATE TABLE `t_cron` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `desc` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `cron_type` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `cronession` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `interval` bigint DEFAULT '0' COMMENT '''间隔秒数''',
  `once_at` datetime(3) DEFAULT NULL COMMENT '''一次性任务执行时间''',
  `node_ids` longtext COLLATE utf8mb4_general_ci COMMENT '''执行主机ID列表''',
  `cmd_type` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''类型，如脚本执行日志或命令执行日志''',
  `type` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''脚本类型，如 shell、python 等''',
  `content` text COLLATE utf8mb4_general_ci COMMENT '''脚本内容''',
  `timeout` bigint DEFAULT NULL COMMENT '''脚本超时时间（秒）''',
  `enable` tinyint(1) DEFAULT '0' COMMENT '''是否可用 1：是 0： 否''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_t_cron_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_cron
-- ----------------------------
INSERT INTO `t_cron` VALUES ('1', '2025-05-12 16:35:05.390', '2025-05-13 17:13:05.599', null, 'test', 'test', 'interval', '', '10', null, '[6,5]', 'script', 'bash', '#!/bin/bash\nhostname && date', '30', '0', 'admin');
INSERT INTO `t_cron` VALUES ('2', '2025-05-12 17:38:23.128', '2025-05-13 16:09:14.618', null, 'test2', 'test2', 'once', '', '60', '2025-05-13 11:05:04.000', '[4,5]', 'command', 'shell', 'hostname', '30', '0', 'admin');
INSERT INTO `t_cron` VALUES ('3', '2025-05-12 17:40:41.656', '2025-05-13 17:11:54.804', null, 'test3', 'test3', 'cron', '1 * * * * ', '60', null, '[4,5]', 'command', 'shell', 'date', '30', '0', 'admin');

-- ----------------------------
-- Table structure for t_cron_log
-- ----------------------------
DROP TABLE IF EXISTS `t_cron_log`;
CREATE TABLE `t_cron_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `desc` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `cron_type` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `cronession` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `interval` bigint DEFAULT '0' COMMENT '''间隔秒数''',
  `once_at` datetime(3) DEFAULT NULL COMMENT '''一次性任务执行时间''',
  `node_ids` longtext COLLATE utf8mb4_general_ci COMMENT '''执行主机ID列表''',
  `status` tinyint DEFAULT NULL COMMENT '''1:成功,2：失败''',
  `start_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''发起时间''',
  `end_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''执行结束时间''',
  `time_cost` int DEFAULT NULL COMMENT '''执行耗时（毫秒）''',
  PRIMARY KEY (`id`),
  KEY `idx_t_cron_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_cron_log
-- ----------------------------

-- ----------------------------
-- Table structure for t_cron_script_log_s
-- ----------------------------
DROP TABLE IF EXISTS `t_cron_script_log_s`;
CREATE TABLE `t_cron_script_log_s` (
  `cron_log_id` bigint unsigned NOT NULL,
  `script_log_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`cron_log_id`,`script_log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_cron_script_log_s
-- ----------------------------

-- ----------------------------
-- Table structure for t_link
-- ----------------------------
DROP TABLE IF EXISTS `t_link`;
CREATE TABLE `t_link` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `desc` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_link
-- ----------------------------
INSERT INTO `t_link` VALUES ('1', 'Grafana1', '内网grafana1', 'http://192.168.2.239:3000/login');
INSERT INTO `t_link` VALUES ('2', 'harbor', '内网harbor', 'http://192.168.2.239:80');
INSERT INTO `t_link` VALUES ('4', '测试1', '测试1', 'http://localhost:8090/#/nav/Link');
INSERT INTO `t_link` VALUES ('5', '1231', '随机测试', 'http://localhost:8090/#/nav/Link');
INSERT INTO `t_link` VALUES ('6', '夜莺', '内网夜莺', 'http://localhost:8090/#/nav/Link');
INSERT INTO `t_link` VALUES ('7', 'jumpserver', '内网跳板机', 'http://localhost:8090/#/nav/Link');

-- ----------------------------
-- Table structure for t_link_s
-- ----------------------------
DROP TABLE IF EXISTS `t_link_s`;
CREATE TABLE `t_link_s` (
  `nav_id` bigint unsigned NOT NULL,
  `link_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`nav_id`,`link_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_link_s
-- ----------------------------
INSERT INTO `t_link_s` VALUES ('1', '3');
INSERT INTO `t_link_s` VALUES ('2', '1');
INSERT INTO `t_link_s` VALUES ('2', '2');
INSERT INTO `t_link_s` VALUES ('2', '5');
INSERT INTO `t_link_s` VALUES ('2', '6');
INSERT INTO `t_link_s` VALUES ('2', '7');
INSERT INTO `t_link_s` VALUES ('3', '4');

-- ----------------------------
-- Table structure for t_manage_log
-- ----------------------------
DROP TABLE IF EXISTS `t_manage_log`;
CREATE TABLE `t_manage_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''任务组执行名称''',
  `args` text COLLATE utf8mb4_general_ci COMMENT '''执行时传入的参数或全局变量，建议为JSON格式''',
  `desc` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''任务组描述或备注信息''',
  `status` varchar(32) COLLATE utf8mb4_general_ci DEFAULT 'draft' COMMENT '''任务状态（pending、running、success、failed）''',
  `start_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''发起时间''',
  `end_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''执行结束时间''',
  `time_cost` int DEFAULT NULL COMMENT '''执行耗时（毫秒）''',
  PRIMARY KEY (`id`),
  KEY `idx_t_manage_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_manage_log
-- ----------------------------

-- ----------------------------
-- Table structure for t_nav
-- ----------------------------
DROP TABLE IF EXISTS `t_nav`;
CREATE TABLE `t_nav` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `nav_sort` bigint DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_nav
-- ----------------------------
INSERT INTO `t_nav` VALUES ('1', 'test1', '0');
INSERT INTO `t_nav` VALUES ('2', '大屏', '1');
INSERT INTO `t_nav` VALUES ('3', '测试3', '10');

-- ----------------------------
-- Table structure for t_nodes
-- ----------------------------
DROP TABLE IF EXISTS `t_nodes`;
CREATE TABLE `t_nodes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `node_name` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''CMDB名称''',
  `username` longtext COLLATE utf8mb4_general_ci COMMENT '''服务器用户名称''',
  `public_ip` longtext COLLATE utf8mb4_general_ci COMMENT '''IP地址''',
  `ssh_port` bigint DEFAULT NULL COMMENT '''SSH端口号''',
  `authmodel` longtext COLLATE utf8mb4_general_ci COMMENT '''连接服务器所使用的是密钥还是密码''',
  `password` longtext COLLATE utf8mb4_general_ci COMMENT '''password''',
  `private_key` longtext COLLATE utf8mb4_general_ci COMMENT '''私钥''',
  `status` tinyint DEFAULT NULL COMMENT '''1:SSH连接成功,2：SSH连接失败''',
  `timeout` bigint DEFAULT NULL COMMENT '''超时时间''',
  `label` longtext COLLATE utf8mb4_general_ci COMMENT '''标签''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_t_nodes_node_name` (`node_name`),
  KEY `idx_t_nodes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_nodes
-- ----------------------------
INSERT INTO `t_nodes` VALUES ('4', '2025-04-08 15:52:24.830', '2025-05-06 15:04:47.291', null, 'kafka', 'root', '192.168.2.115', '22', 'password', 'YWRtaW5AMTIz', '', '1', '5', 'test', 'admin');
INSERT INTO `t_nodes` VALUES ('5', '2025-04-08 15:52:32.808', '2025-05-06 15:04:53.587', null, 'node3', 'root', '192.168.2.114', '22', 'password', 'YWRtaW5AMTIz', '', '1', '6', 'test', 'admin');
INSERT INTO `t_nodes` VALUES ('6', '2025-04-08 18:09:33.394', '2025-05-06 15:04:21.458', null, 'node1', 'root', '192.168.2.112', '22', 'password', 'YWRtaW5AMTIz', '', '1', '5', 'test', 'admin');

-- ----------------------------
-- Table structure for t_node_group
-- ----------------------------
DROP TABLE IF EXISTS `t_node_group`;
CREATE TABLE `t_node_group` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `group_name` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `desc` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''描述''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_t_node_group_group_name` (`group_name`),
  KEY `idx_t_node_group_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_node_group
-- ----------------------------
INSERT INTO `t_node_group` VALUES ('4', '2025-04-09 11:19:15.255', '2025-04-09 11:19:15.255', null, 'test2', '', 'admin');
INSERT INTO `t_node_group` VALUES ('6', '2025-04-17 18:56:48.011', '2025-04-17 18:56:48.011', null, 'test', '', 'admin');

-- ----------------------------
-- Table structure for t_node_group_s
-- ----------------------------
DROP TABLE IF EXISTS `t_node_group_s`;
CREATE TABLE `t_node_group_s` (
  `node_group_id` bigint unsigned NOT NULL,
  `nodes_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`node_group_id`,`nodes_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_node_group_s
-- ----------------------------
INSERT INTO `t_node_group_s` VALUES ('4', '4');
INSERT INTO `t_node_group_s` VALUES ('4', '5');
INSERT INTO `t_node_group_s` VALUES ('4', '6');
INSERT INTO `t_node_group_s` VALUES ('6', '5');
INSERT INTO `t_node_group_s` VALUES ('6', '6');

-- ----------------------------
-- Table structure for t_script
-- ----------------------------
DROP TABLE IF EXISTS `t_script`;
CREATE TABLE `t_script` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''任务或脚本执行的名称''',
  `node_ids` longtext COLLATE utf8mb4_general_ci COMMENT '''执行主机ID列表''',
  `status` tinyint DEFAULT NULL COMMENT '''1:成功,2：失败''',
  `cmd_type` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''类型，如脚本执行日志或命令执行日志''',
  `start_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''发起时间''',
  `end_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''执行结束时间''',
  `time_cost` int DEFAULT NULL COMMENT '''执行耗时（毫秒）''',
  `desc` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''额外描述或备注''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''执行人''',
  PRIMARY KEY (`id`),
  KEY `idx_t_script_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_script
-- ----------------------------
INSERT INTO `t_script` VALUES ('1', '2025-04-15 17:04:04.457', '2025-04-15 17:04:04.457', null, 'test', '[4,5,6,7]', '1', 'command', '2025-04-15T17:04:03+08:00', '2025-04-15T17:04:04+08:00', '718', '', 'admin');
INSERT INTO `t_script` VALUES ('2', '2025-04-15 17:36:09.818', '2025-04-15 17:36:09.818', null, 'test-2025-04-15 17:36:09', '[4,5,6,7]', '1', 'command', '2025-04-15T17:36:09+08:00', '2025-04-15T17:36:09+08:00', '713', '', 'admin');
INSERT INTO `t_script` VALUES ('3', '2025-04-18 10:03:51.868', '2025-04-18 10:03:51.868', null, '临时命令-20250418-10:03:51', '[4,5]', '1', 'command', '2025-04-18T10:03:51+08:00', '2025-04-18T10:03:51+08:00', '386', '', 'admin');
INSERT INTO `t_script` VALUES ('4', '2025-04-18 10:41:13.241', '2025-04-18 10:41:13.241', null, '临时命令-20250418-10:41:13', '[4,5]', '1', 'command', '2025-04-18T10:41:12+08:00', '2025-04-18T10:41:13+08:00', '386', '', 'admin');
INSERT INTO `t_script` VALUES ('5', '2025-04-18 11:26:01.790', '2025-04-18 11:26:01.790', null, '临时命令-20250418-11:26:01', '[6]', '1', 'command', '2025-04-18T11:26:01+08:00', '2025-04-18T11:26:01+08:00', '165', '', 'admin');
INSERT INTO `t_script` VALUES ('6', '2025-04-18 14:25:56.207', '2025-04-18 14:25:56.207', null, '临时命令123-20250418-14:25:56', '[6]', '1', 'command', '2025-04-18T14:25:56+08:00', '2025-04-18T14:25:56+08:00', '167', '', 'admin');
INSERT INTO `t_script` VALUES ('7', '2025-04-18 14:32:53.649', '2025-04-18 14:32:53.649', null, '临时命令-20250418-14:32:53', '[4,5]', '1', 'command', '2025-04-18T14:32:53+08:00', '2025-04-18T14:32:53+08:00', '326', '', 'admin');
INSERT INTO `t_script` VALUES ('8', '2025-04-18 14:37:24.312', '2025-04-18 14:37:24.312', null, '临时命令-20250418-14:37:24', '[4,5]', '2', 'script', '2025-04-18T14:37:23+08:00', '2025-04-18T14:37:24+08:00', '376', '', 'admin');
INSERT INTO `t_script` VALUES ('9', '2025-04-18 16:13:43.974', '2025-04-18 16:13:43.974', null, '临时命令-20250418-16:13:43', '[6,5,4]', '2', 'script', '2025-04-18T16:13:43+08:00', '2025-04-18T16:13:43+08:00', '524', '1', 'admin');
INSERT INTO `t_script` VALUES ('10', '2025-04-18 16:15:52.687', '2025-04-18 16:15:52.687', null, '临时命令-20250418-16:15:52', '[6,5,4]', '2', 'script', '2025-04-18T16:15:52+08:00', '2025-04-18T16:15:52+08:00', '507', '', 'admin');
INSERT INTO `t_script` VALUES ('11', '2025-04-18 16:19:45.024', '2025-04-18 16:19:45.024', null, '临时命令-20250418-16:19:35', '[6,5,4]', '2', 'script', '2025-04-18T16:16:55+08:00', '2025-04-18T16:19:36+08:00', '161820', '', 'admin');
INSERT INTO `t_script` VALUES ('12', '2025-04-18 16:22:07.804', '2025-04-18 16:22:07.804', null, '临时命令-20250418-16:22:07', '[6,5,4]', '2', 'script', '2025-04-18T16:20:14+08:00', '2025-04-18T16:22:07+08:00', '112804', '', 'admin');
INSERT INTO `t_script` VALUES ('13', '2025-04-18 16:30:31.570', '2025-04-18 16:30:31.570', null, '临时命令-20250418-16:30:31', '[6,5,4]', '2', 'script', '2025-04-18T16:24:28+08:00', '2025-04-18T16:30:31+08:00', '362696', '', 'admin');
INSERT INTO `t_script` VALUES ('14', '2025-04-18 16:30:47.882', '2025-04-18 16:30:47.882', null, '临时命令-20250418-16:30:47', '[6,5,4]', '1', 'script', '2025-04-18T16:30:47+08:00', '2025-04-18T16:30:47+08:00', '595', '', 'admin');
INSERT INTO `t_script` VALUES ('15', '2025-04-18 16:32:41.703', '2025-04-18 16:32:41.703', null, '临时命令-20250418-16:32:41', '[6,5]', '2', 'script', '2025-04-18T16:32:41+08:00', '2025-04-18T16:32:41+08:00', '348', '123', 'admin');
INSERT INTO `t_script` VALUES ('16', '2025-04-18 16:33:48.756', '2025-04-18 16:33:48.756', null, '临时命令-20250418-16:33:48', '[6]', '2', 'script', '2025-04-18T16:33:48+08:00', '2025-04-18T16:33:48+08:00', '144', '', 'admin');
INSERT INTO `t_script` VALUES ('17', '2025-04-18 16:34:58.181', '2025-04-18 16:34:58.181', null, '临时命令-20250418-16:34:58', '[6]', '2', 'script', '2025-04-18T16:34:58+08:00', '2025-04-18T16:34:58+08:00', '156', '', 'admin');
INSERT INTO `t_script` VALUES ('18', '2025-04-18 16:35:15.389', '2025-04-18 16:35:15.389', null, '临时命令-20250418-16:35:15', '[6]', '1', 'script', '2025-04-18T16:35:15+08:00', '2025-04-18T16:35:15+08:00', '187', '', 'admin');
INSERT INTO `t_script` VALUES ('19', '2025-04-18 18:15:53.044', '2025-04-18 18:15:53.044', null, '临时命令-20250418-18:15:53', '[6,5]', '1', 'command', '2025-04-18T18:15:52+08:00', '2025-04-18T18:15:53+08:00', '400', '', 'admin');
INSERT INTO `t_script` VALUES ('20', '2025-04-18 18:21:21.936', '2025-04-18 18:21:21.936', null, '临时命令-20250418-18:21:21', '[6,5]', '1', 'command', '2025-04-18T18:21:11+08:00', '2025-04-18T18:21:21+08:00', '10365', '', 'admin');
INSERT INTO `t_script` VALUES ('21', '2025-04-18 18:24:51.227', '2025-04-18 18:24:51.227', null, '临时命令-20250418-18:24:51', '[6,5]', '1', 'command', '2025-04-18T18:24:40+08:00', '2025-04-18T18:24:51+08:00', '10335', '', 'admin');
INSERT INTO `t_script` VALUES ('22', '2025-04-18 18:27:02.900', '2025-04-18 18:27:02.900', null, '临时命令-20250418-18:27:02', '[6]', '1', 'command', '2025-04-18T18:26:57+08:00', '2025-04-18T18:27:02+08:00', '5180', '', 'admin');
INSERT INTO `t_script` VALUES ('23', '2025-04-18 18:30:28.649', '2025-04-18 18:30:28.649', null, '临时命令-20250418-18:30:28', '[6]', '1', 'command', '2025-04-18T18:30:23+08:00', '2025-04-18T18:30:28+08:00', '5187', '', 'admin');
INSERT INTO `t_script` VALUES ('24', '2025-04-18 18:33:26.258', '2025-04-18 18:33:26.258', null, '临时命令-20250418-18:33:26', '[6]', '1', 'command', '2025-04-18T18:33:23+08:00', '2025-04-18T18:33:26+08:00', '3165', '', 'admin');
INSERT INTO `t_script` VALUES ('25', '2025-04-18 18:38:17.599', '2025-04-18 18:38:17.599', null, '临时命令-20250418-18:38:17', '[5,4,6]', '1', 'command', '2025-04-18T18:37:01+08:00', '2025-04-18T18:38:17+08:00', '75619', '', 'admin');
INSERT INTO `t_script` VALUES ('26', '2025-04-18 18:40:05.138', '2025-04-18 18:40:05.138', null, '临时命令-20250418-18:40:05', '[5,4,6]', '1', 'command', '2025-04-18T18:39:24+08:00', '2025-04-18T18:40:05+08:00', '40236', '', 'admin');
INSERT INTO `t_script` VALUES ('27', '2025-04-18 18:46:35.566', '2025-04-18 18:46:35.566', null, '临时命令-20250418-18:46:35', '[6]', '1', 'command', '2025-04-18T18:45:50+08:00', '2025-04-18T18:46:35+08:00', '45174', '', 'admin');
INSERT INTO `t_script` VALUES ('28', '2025-04-18 18:48:28.758', '2025-04-18 18:48:28.758', null, 'test-20250418-18:48:28', '[6]', '1', 'command', '2025-04-18T18:47:43+08:00', '2025-04-18T18:48:28+08:00', '45174', '', 'admin');
INSERT INTO `t_script` VALUES ('29', '2025-04-18 18:51:03.126', '2025-04-18 18:51:03.126', null, 'test-20250418-18:51:03', '[6]', '1', 'command', '2025-04-18T18:50:17+08:00', '2025-04-18T18:51:03+08:00', '45190', '', 'admin');
INSERT INTO `t_script` VALUES ('30', '2025-04-24 11:59:37.294', '2025-04-24 11:59:37.294', null, '临时命令-20250424-11:59:37', '[6,5,4]', '1', 'command', '2025-04-24T11:59:37+08:00', '2025-04-24T11:59:37+08:00', '184', '', 'admin');

-- ----------------------------
-- Table structure for t_script_library
-- ----------------------------
DROP TABLE IF EXISTS `t_script_library`;
CREATE TABLE `t_script_library` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''模板名称，用于标识和展示模板''',
  `content` text COLLATE utf8mb4_general_ci COMMENT '''脚本内容''',
  `desc` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''模板描述信息''',
  `type` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''脚本类型，例如 shell、powershell''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_t_script_library_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_script_library
-- ----------------------------
INSERT INTO `t_script_library` VALUES ('4', '2025-04-17 11:27:51.156', '2025-04-17 14:54:30.644', null, 'test', '#!/bin/bash\nhostname', '123', 'shell', 'admin');
INSERT INTO `t_script_library` VALUES ('5', '2025-04-17 14:49:53.100', '2025-04-18 16:15:33.776', null, 'test2', '#!/bin/bash\n# Description:  获取Linux 本机被连接的IP Top 5\n\n# Get Data/Time\nCTIME=$(date \"+%Y-%m-%d %H-%M-%S\")\necho -e \"------ $CTIME connect Top 10 ------\"\nnetstat -ant|awk \'NR > 2  {print $5}\'|awk -F \":\" \'{print $1}\'|grep -v \"^$\"|sort -nr|uniq -c|sort -nr\n', '测试2', 'shell', 'admin');

-- ----------------------------
-- Table structure for t_script_log
-- ----------------------------
DROP TABLE IF EXISTS `t_script_log`;
CREATE TABLE `t_script_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''脚本名称，用于标识和展示''',
  `node_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''节点名称，标识执行节点''',
  `type` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''脚本类型，如 shell、python 等''',
  `content` text COLLATE utf8mb4_general_ci COMMENT '''脚本内容''',
  `status` tinyint DEFAULT NULL COMMENT '''1:成功,2：失败''',
  `timeout` bigint DEFAULT NULL COMMENT '''脚本超时时间（秒）''',
  `run_log` longtext COLLATE utf8mb4_general_ci COMMENT '''执行结果''',
  `start_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''发起时间''',
  `end_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''执行结束时间''',
  `time_cost` int DEFAULT NULL COMMENT '''执行耗时（毫秒）''',
  PRIMARY KEY (`id`),
  KEY `idx_t_script_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=151 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_script_log
-- ----------------------------

-- ----------------------------
-- Table structure for t_script_s
-- ----------------------------
DROP TABLE IF EXISTS `t_script_s`;
CREATE TABLE `t_script_s` (
  `script_id` bigint unsigned NOT NULL,
  `script_log_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`script_id`,`script_log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_script_s
-- ----------------------------
INSERT INTO `t_script_s` VALUES ('1', '1');
INSERT INTO `t_script_s` VALUES ('1', '2');
INSERT INTO `t_script_s` VALUES ('1', '3');
INSERT INTO `t_script_s` VALUES ('1', '4');
INSERT INTO `t_script_s` VALUES ('2', '5');
INSERT INTO `t_script_s` VALUES ('2', '6');
INSERT INTO `t_script_s` VALUES ('2', '7');
INSERT INTO `t_script_s` VALUES ('2', '8');
INSERT INTO `t_script_s` VALUES ('3', '9');
INSERT INTO `t_script_s` VALUES ('3', '10');
INSERT INTO `t_script_s` VALUES ('4', '11');
INSERT INTO `t_script_s` VALUES ('4', '12');
INSERT INTO `t_script_s` VALUES ('5', '13');
INSERT INTO `t_script_s` VALUES ('6', '14');
INSERT INTO `t_script_s` VALUES ('7', '15');
INSERT INTO `t_script_s` VALUES ('7', '16');
INSERT INTO `t_script_s` VALUES ('8', '17');
INSERT INTO `t_script_s` VALUES ('8', '18');
INSERT INTO `t_script_s` VALUES ('9', '19');
INSERT INTO `t_script_s` VALUES ('9', '20');
INSERT INTO `t_script_s` VALUES ('9', '21');
INSERT INTO `t_script_s` VALUES ('10', '22');
INSERT INTO `t_script_s` VALUES ('10', '23');
INSERT INTO `t_script_s` VALUES ('10', '24');
INSERT INTO `t_script_s` VALUES ('11', '25');
INSERT INTO `t_script_s` VALUES ('11', '26');
INSERT INTO `t_script_s` VALUES ('11', '27');
INSERT INTO `t_script_s` VALUES ('12', '28');
INSERT INTO `t_script_s` VALUES ('12', '29');
INSERT INTO `t_script_s` VALUES ('12', '30');
INSERT INTO `t_script_s` VALUES ('13', '31');
INSERT INTO `t_script_s` VALUES ('13', '32');
INSERT INTO `t_script_s` VALUES ('13', '33');
INSERT INTO `t_script_s` VALUES ('14', '34');
INSERT INTO `t_script_s` VALUES ('14', '35');
INSERT INTO `t_script_s` VALUES ('14', '36');
INSERT INTO `t_script_s` VALUES ('15', '37');
INSERT INTO `t_script_s` VALUES ('15', '38');
INSERT INTO `t_script_s` VALUES ('16', '39');
INSERT INTO `t_script_s` VALUES ('17', '40');
INSERT INTO `t_script_s` VALUES ('18', '41');
INSERT INTO `t_script_s` VALUES ('19', '42');
INSERT INTO `t_script_s` VALUES ('19', '43');
INSERT INTO `t_script_s` VALUES ('20', '44');
INSERT INTO `t_script_s` VALUES ('20', '45');
INSERT INTO `t_script_s` VALUES ('21', '46');
INSERT INTO `t_script_s` VALUES ('21', '47');
INSERT INTO `t_script_s` VALUES ('22', '48');
INSERT INTO `t_script_s` VALUES ('23', '49');
INSERT INTO `t_script_s` VALUES ('24', '50');
INSERT INTO `t_script_s` VALUES ('25', '51');
INSERT INTO `t_script_s` VALUES ('25', '52');
INSERT INTO `t_script_s` VALUES ('25', '53');
INSERT INTO `t_script_s` VALUES ('26', '54');
INSERT INTO `t_script_s` VALUES ('26', '55');
INSERT INTO `t_script_s` VALUES ('26', '56');
INSERT INTO `t_script_s` VALUES ('27', '57');
INSERT INTO `t_script_s` VALUES ('28', '58');
INSERT INTO `t_script_s` VALUES ('29', '59');
INSERT INTO `t_script_s` VALUES ('30', '60');
INSERT INTO `t_script_s` VALUES ('30', '61');
INSERT INTO `t_script_s` VALUES ('30', '62');

-- ----------------------------
-- Table structure for t_task
-- ----------------------------
DROP TABLE IF EXISTS `t_task`;
CREATE TABLE `t_task` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''子任务名称或类型''',
  `type` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''脚本类型，如 shell、python 等''',
  `content` text COLLATE utf8mb4_general_ci COMMENT '''执行命令或脚本内容''',
  `sort` bigint DEFAULT '999' COMMENT '''执行顺序(1-999)''',
  `timeout` bigint DEFAULT NULL COMMENT '''脚本超时时间（秒）''',
  `node_ids` longtext COLLATE utf8mb4_general_ci COMMENT '''执行主机ID列表''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  KEY `idx_t_task_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_task
-- ----------------------------
INSERT INTO `t_task` VALUES ('1', '2025-04-24 16:56:37.051', '2025-05-07 18:11:51.719', null, '子任务一', 'shell', '#!/bin/bash\necho \"${env_task_s[ENV]}\"\necho \"${env_task_s[NAME]}\"\nhostname\n\n\n\n', '1', '30', '6,5,4', 'admin');
INSERT INTO `t_task` VALUES ('2', '2025-04-24 16:56:37.051', '2025-04-24 18:17:20.288', null, '2', 'shell', '#!/bin/bash\nhostname', '2', '30', '5', 'admin');
INSERT INTO `t_task` VALUES ('5', '2025-04-25 18:34:28.688', '2025-04-25 18:34:28.688', null, '2', 'shell', 'hostname', '2', '30', '6,5,4', 'admin');
INSERT INTO `t_task` VALUES ('6', '2025-04-28 18:20:51.663', '2025-05-07 18:11:51.722', null, '子任务2', 'shell', 'date && hostname', '2', '30', '6', 'admin');
INSERT INTO `t_task` VALUES ('7', '2025-05-07 18:14:06.668', '2025-05-07 18:14:06.668', null, 'test1', 'shell', '#!/bin/bash\nhostname', '1', '30', '4,5', 'admin');

-- ----------------------------
-- Table structure for t_task_log
-- ----------------------------
DROP TABLE IF EXISTS `t_task_log`;
CREATE TABLE `t_task_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''子任务名称''',
  `content` text COLLATE utf8mb4_general_ci COMMENT '''执行命令或脚本内容''',
  `sort` bigint DEFAULT '999' COMMENT '''执行顺序(1-999)''',
  `node_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''节点名称，标识执行节点''',
  `run_log` text COLLATE utf8mb4_general_ci COMMENT '''执行结果日志''',
  `status` varchar(32) COLLATE utf8mb4_general_ci DEFAULT 'pending' COMMENT '''子任务执行状态（pending、running、success、failed）''',
  `start_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''发起时间''',
  `end_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''执行结束时间''',
  `time_cost` int DEFAULT NULL COMMENT '''执行耗时（毫秒）''',
  PRIMARY KEY (`id`),
  KEY `idx_t_task_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=190 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_task_log
-- ----------------------------

-- ----------------------------
-- Table structure for t_task_log_s
-- ----------------------------
DROP TABLE IF EXISTS `t_task_log_s`;
CREATE TABLE `t_task_log_s` (
  `task_manage_log_id` bigint unsigned NOT NULL,
  `task_log_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`task_manage_log_id`,`task_log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_task_log_s
-- ----------------------------

-- ----------------------------
-- Table structure for t_task_manage
-- ----------------------------
DROP TABLE IF EXISTS `t_task_manage`;
CREATE TABLE `t_task_manage` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''任务组名称''',
  `args` text COLLATE utf8mb4_general_ci COMMENT '''执行时传入的参数或全局变量，建议为JSON格式''',
  `desc` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''任务组描述信息或备注''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  KEY `idx_t_task_manage_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_task_manage
-- ----------------------------
INSERT INTO `t_task_manage` VALUES ('3', '2025-04-24 16:56:37.049', '2025-05-07 18:11:51.716', null, 'test', '[{\"name\":\"ENV\",\"type\":\"select\",\"options\":\"test:测试环境\\nprod:生产环境\",\"required\":true},{\"name\":\"NAME\",\"type\":\"string\",\"options\":\"\",\"required\":false}]', '123', 'admin');
INSERT INTO `t_task_manage` VALUES ('5', '2025-05-07 18:14:06.665', '2025-05-07 18:14:06.665', null, '121', '[]', '312132', 'admin');

-- ----------------------------
-- Table structure for t_task_manage_log
-- ----------------------------
DROP TABLE IF EXISTS `t_task_manage_log`;
CREATE TABLE `t_task_manage_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''任务组执行名称''',
  `status` varchar(32) COLLATE utf8mb4_general_ci DEFAULT 'draft' COMMENT '''任务状态（pending、running、success、failed）''',
  `start_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''发起时间''',
  `end_time` varchar(2048) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''执行结束时间''',
  `time_cost` int DEFAULT NULL COMMENT '''执行耗时（毫秒）''',
  PRIMARY KEY (`id`),
  KEY `idx_t_task_manage_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_task_manage_log
-- ----------------------------

-- ----------------------------
-- Table structure for t_task_manage_log_s
-- ----------------------------
DROP TABLE IF EXISTS `t_task_manage_log_s`;
CREATE TABLE `t_task_manage_log_s` (
  `manage_log_id` bigint unsigned NOT NULL,
  `task_manage_log_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`manage_log_id`,`task_manage_log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_task_manage_log_s
-- ----------------------------

-- ----------------------------
-- Table structure for t_task_s
-- ----------------------------
DROP TABLE IF EXISTS `t_task_s`;
CREATE TABLE `t_task_s` (
  `task_manage_id` bigint unsigned NOT NULL,
  `task_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`task_manage_id`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_task_s
-- ----------------------------
INSERT INTO `t_task_s` VALUES ('3', '1');
INSERT INTO `t_task_s` VALUES ('3', '6');
INSERT INTO `t_task_s` VALUES ('5', '7');

-- ----------------------------
-- Table structure for t_transfer
-- ----------------------------
DROP TABLE IF EXISTS `t_transfer`;
CREATE TABLE `t_transfer` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''分发名称，用于标识和展示''',
  `source_path` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `target_path` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `node_ids` longtext COLLATE utf8mb4_general_ci COMMENT '''目标主机''',
  `status` tinyint DEFAULT NULL COMMENT '''1:成功,2：失败''',
  `run_log` text COLLATE utf8mb4_general_ci COMMENT '''执行结果''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建人''',
  PRIMARY KEY (`id`),
  KEY `idx_t_transfer_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of t_transfer
-- ----------------------------
INSERT INTO `t_transfer` VALUES ('1', '2025-04-22 09:55:18.199', '2025-04-22 09:55:18.199', null, '1-20250422-095518', './uploads/1745286914_black.xlsx', '/tmp', '[6,5,4]', '1', '[node3] 上传成功，目标路径：/tmp/1745286914_black.xlsx\n[node4] 上传成功，目标路径：/tmp/1745286914_black.xlsx\n[node1] 上传成功，目标路径：/tmp/1745286914_black.xlsx\n', 'admin');
INSERT INTO `t_transfer` VALUES ('2', '2025-04-22 10:02:40.996', '2025-04-22 10:02:40.996', null, '2-20250422-100240', './uploads/1745287337_black.xlsx', '/tmp', '[6,5,4]', '1', '[node1] 上传成功，目标路径：/tmp/1745287337_black.xlsx\n[node3] 上传成功，目标路径：/tmp/1745287337_black.xlsx\n[node4] 上传成功，目标路径：/tmp/1745287337_black.xlsx\n', 'admin');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''用户名''',
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''用户密码''',
  `nickname` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''中文名''',
  `given_name` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''花名''',
  `mail` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''邮箱''',
  `job_number` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''工号''',
  `mobile` varchar(15) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''手机号''',
  `avatar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''头像''',
  `postal_address` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''地址''',
  `departments` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''部门''',
  `position` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''职位''',
  `introduction` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''个人简介''',
  `status` tinyint(1) DEFAULT '1' COMMENT '''状态:1在职, 2离职''',
  `creator` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''创建者''',
  `source` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '''用户来源：dingTalk、wecom、feishu、ldap、platform''',
  `department_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '''部门id''',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `mobile` (`mobile`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', '2025-04-07 11:03:12.783', '2025-04-07 17:25:38.230', null, 'admin', 'LTuAAYEYJMDnZ4JFLyf5rl5QVcGPzBW2/zSAAnHavOnhM7sUf3+SZCj2P4wZIrxoGqofVXo1OTEySWxWKueYbdgp5cBjhh1kpMpv1+myFG13s79TNZtvGh9NsIaXovp1tokbQR9VvQMokXGVGnhNGX4A9oMU4wQVWlGudgO54MQ=', '管理员', '最强后台', 'admin@eryajf.net', '0000', '18888888888', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', '中国河南省南阳市', 'root', '系统管理员', '最强后台的管理员', '1', 'admin', '', '1');

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `role_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_roles
-- ----------------------------
INSERT INTO `user_roles` VALUES ('1', '1');
