/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2020-09-21 09:05:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `name` varchar(50) NOT NULL COMMENT '部门名称',
  `pid` varchar(50) NOT NULL COMMENT '上级部门',
  `sort` tinyint(4) DEFAULT NULL COMMENT '排序',
  `create_at` int(13) DEFAULT NULL COMMENT '创建时间',
  `update_at` int(13) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of dept
-- ----------------------------
INSERT INTO `dept` VALUES ('18590aaa-ecf6-11ea-9953-4ccc6aee07d8', '产品运营部', '731bab7a-ecf4-11ea-acaf-4ccc6aee07d8', '1', '1599035230', '1599638110');
INSERT INTO `dept` VALUES ('4c7e1c5a-ecf4-11ea-acaf-4ccc6aee07d8', '中企', '0', '1', '1599034458', '1599554131');
INSERT INTO `dept` VALUES ('731bab7a-ecf4-11ea-acaf-4ccc6aee07d8', '北京', '4c7e1c5a-ecf4-11ea-acaf-4ccc6aee07d8', '1', '1599034523', '1599034523');
INSERT INTO `dept` VALUES ('7799445a-ecf4-11ea-acaf-4ccc6aee07d8', '上海', '4c7e1c5a-ecf4-11ea-acaf-4ccc6aee07d8', '2', '1599034531', '1599034566');
INSERT INTO `dept` VALUES ('7bdfc7aa-ecf4-11ea-acaf-4ccc6aee07d8', '广州', '4c7e1c5a-ecf4-11ea-acaf-4ccc6aee07d8', '3', '1599034538', '1599034572');

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS `dictionary`;
CREATE TABLE `dictionary` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `name` varchar(50) NOT NULL COMMENT '岗位名称',
  `description` varchar(255) NOT NULL COMMENT '排序',
  `create_at` int(13) DEFAULT NULL COMMENT '创建时间',
  `update_at` int(13) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of dictionary
-- ----------------------------

-- ----------------------------
-- Table structure for dictionary_detail
-- ----------------------------
DROP TABLE IF EXISTS `dictionary_detail`;
CREATE TABLE `dictionary_detail` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `dictionary_id` varchar(50) NOT NULL COMMENT 'ID',
  `label` varchar(50) NOT NULL COMMENT '岗位名称',
  `value` varchar(255) NOT NULL COMMENT '排序',
  `sort` varchar(50) DEFAULT NULL COMMENT '排序',
  `create_at` int(13) DEFAULT NULL COMMENT '创建时间',
  `update_at` int(13) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of dictionary_detail
-- ----------------------------

-- ----------------------------
-- Table structure for job
-- ----------------------------
DROP TABLE IF EXISTS `job`;
CREATE TABLE `job` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `name` varchar(50) NOT NULL COMMENT '岗位名称',
  `sort` varchar(50) NOT NULL COMMENT '排序',
  `create_at` int(13) DEFAULT NULL COMMENT '创建时间',
  `update_at` int(13) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of job
-- ----------------------------
INSERT INTO `job` VALUES ('0b345d0a-ed02-11ea-98bb-4ccc6aee07d8', '前端', '1', '1599040362', '1599554115');
INSERT INTO `job` VALUES ('459e6f3a-ed02-11ea-aba0-4ccc6aee07d8', 'java', '1', '1599040460', '1599554102');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `name` varchar(50) NOT NULL COMMENT '菜单名称',
  `icon` varchar(50) DEFAULT NULL COMMENT '图标',
  `url` varchar(50) DEFAULT NULL COMMENT '菜单链接',
  `component` varchar(50) DEFAULT NULL COMMENT '文件路径',
  `meta_data` varchar(200) DEFAULT NULL COMMENT '菜单数据 json string',
  `pid` varchar(50) DEFAULT NULL COMMENT '父级菜单ID',
  `sort` tinyint(4) unsigned zerofill DEFAULT NULL COMMENT '排序',
  `create_at` int(13) DEFAULT NULL COMMENT '创建时间',
  `update_at` int(13) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES ('1638c08a-ece7-11ea-9bbe-4ccc6aee07d8', '权限管理', '', '/system/permission', '/system/permission/index', '', '9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8', '0003', '1599028784', '1599028784');
INSERT INTO `menu` VALUES ('552efc7a-ece5-11ea-b2e8-4ccc6aee07d8', '菜单管理', '', '/system/menu', '/system/menu/index', '', '9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8', '0004', '1599028030', '1599028030');
INSERT INTO `menu` VALUES ('758444da-ece5-11ea-b2e8-4ccc6aee07d8', '岗位管理', '', '/system/job', '/system/job/index', '', '9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8', '0006', '1599028085', '1599638437');
INSERT INTO `menu` VALUES ('83283e2a-ece5-11ea-b2e8-4ccc6aee07d8', '字典管理', '', '/system/dictionary', '/system/dictionary/index', '', '9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8', '0007', '1599028108', '1599638443');
INSERT INTO `menu` VALUES ('9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8', '系统管理', '', '', '', '', '0', '0001', '1599016550', '1599638416');
INSERT INTO `menu` VALUES ('b3b4a67a-ece8-11ea-9bbb-4ccc6aee07d8', '部门管理', '', '/system/dept', '/system/dept/index', '', '9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8', '0005', '1599029477', '1599638433');
INSERT INTO `menu` VALUES ('b4a4323a-ecca-11ea-8a5e-4ccc6aee07d8', '用户管理', '', '/system/user', '/system/user/index', '', '9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8', '0001', '1599016594', '1599016594');
INSERT INTO `menu` VALUES ('cc5ce39a-ecca-11ea-8a5e-4ccc6aee07d8', '角色管理', '', '/system/role', '/system/role/index', '', '9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8', '0002', '1599016634', '1599016634');

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` varchar(50) NOT NULL,
  `code` varchar(50) NOT NULL COMMENT '权限code',
  `name` varchar(50) DEFAULT NULL COMMENT '权限名称',
  `pid` varchar(50) NOT NULL COMMENT '父级类名ID',
  `sort` tinyint(4) DEFAULT NULL COMMENT '排序',
  `create_at` int(13) DEFAULT NULL COMMENT '创建时间',
  `update_at` int(13) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of permission
-- ----------------------------
INSERT INTO `permission` VALUES ('022a92ef-ed95-11ea-9046-4ccc6aee07d8', 'user_create', '用户创建', 'e2931c4f-ed94-11ea-9046-4ccc6aee07d8', '1', '1599103483', '1599639161');
INSERT INTO `permission` VALUES ('0a4d3eff-ed95-11ea-9046-4ccc6aee07d8', 'user_update', '用户修改', 'e2931c4f-ed94-11ea-9046-4ccc6aee07d8', '1', '1599103496', '1599639156');
INSERT INTO `permission` VALUES ('149785ff-ed95-11ea-9046-4ccc6aee07d8', 'user_delete', '用户删除', 'e2931c4f-ed94-11ea-9046-4ccc6aee07d8', '1', '1599103513', '1599639151');
INSERT INTO `permission` VALUES ('25ea567f-ed95-11ea-9046-4ccc6aee07d8', 'role_manage', '角色管理', '0', '10', '1599103543', '1599638647');
INSERT INTO `permission` VALUES ('e2931c4f-ed94-11ea-9046-4ccc6aee07d8', 'user_manage', '用户管理', '0', '1', '1599103430', '1599103558');
INSERT INTO `permission` VALUES ('e78e7b3f-ed95-11ea-949d-4ccc6aee07d8', 'role_create', '角色创建', '25ea567f-ed95-11ea-9046-4ccc6aee07d8', '5', '1599103867', '1599638687');
INSERT INTO `permission` VALUES ('f08583bf-ed94-11ea-9046-4ccc6aee07d8', 'user_select', '用户查询', 'e2931c4f-ed94-11ea-9046-4ccc6aee07d8', '1', '1599103453', '1599639165');
INSERT INTO `permission` VALUES ('f466990f-ed95-11ea-949d-4ccc6aee07d8', 'role_update', '角色修改', '25ea567f-ed95-11ea-9046-4ccc6aee07d8', '5', '1599103889', '1599638713');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `status` varchar(20) NOT NULL DEFAULT 'active' COMMENT '状态：active、ban',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `create_at` int(13) DEFAULT NULL COMMENT '创建时间',
  `update_at` int(13) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES ('5859694a-f180-11ea-80ea-4ccc6aee07d8', '管理员', 'active', '管理员', '1599534412', '1599554196');
INSERT INTO `role` VALUES ('d2dd335a-f17f-11ea-9d24-4ccc6aee07d8', 'role-test', 'active', 'role-test', '1599534188', '1599554201');

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `role_id` varchar(50) NOT NULL COMMENT '角色ID',
  `menu_id` varchar(50) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES ('8d4eec8a-f18d-11ea-926a-4ccc6aee07d8', 'd2dd335a-f17f-11ea-9d24-4ccc6aee07d8', 'b4a4323a-ecca-11ea-8a5e-4ccc6aee07d8');
INSERT INTO `role_menu` VALUES ('8d4eec8a-f18d-11ea-926b-4ccc6aee07d8', 'd2dd335a-f17f-11ea-9d24-4ccc6aee07d8', '9a23ae4a-ecca-11ea-8a5e-4ccc6aee07d8');

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission` (
  `id` varchar(50) DEFAULT NULL COMMENT 'ID',
  `role_id` varchar(50) DEFAULT NULL COMMENT '角色ID',
  `permission_id` varchar(50) DEFAULT NULL COMMENT '权限ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of role_permission
-- ----------------------------
INSERT INTO `role_permission` VALUES ('8d7b059a-f18d-11ea-926c-4ccc6aee07d8', 'd2dd335a-f17f-11ea-9d24-4ccc6aee07d8', 'f107998a-f18c-11ea-9264-4ccc6aee07d8');
INSERT INTO `role_permission` VALUES ('8d7b059a-f18d-11ea-926d-4ccc6aee07d8', 'd2dd335a-f17f-11ea-9d24-4ccc6aee07d8', 'f466990f-ed95-11ea-949d-4ccc6aee07d8');
INSERT INTO `role_permission` VALUES ('8d7b059a-f18d-11ea-926e-4ccc6aee07d8', 'd2dd335a-f17f-11ea-9d24-4ccc6aee07d8', 'e78e7b3f-ed95-11ea-949d-4ccc6aee07d8');
INSERT INTO `role_permission` VALUES ('8d7b059a-f18d-11ea-926f-4ccc6aee07d8', 'd2dd335a-f17f-11ea-9d24-4ccc6aee07d8', '25ea567f-ed95-11ea-9046-4ccc6aee07d8');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `account` varchar(50) NOT NULL COMMENT '账号',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `status` varchar(20) DEFAULT 'active' COMMENT '状态：active，ban',
  `nickname` varchar(50) DEFAULT NULL COMMENT '昵称',
  `phone` varchar(50) DEFAULT NULL COMMENT '手机号',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `dept_id` varchar(50) DEFAULT NULL COMMENT '部门ID ',
  `job_id` varchar(50) DEFAULT NULL COMMENT '岗位ID ',
  `create_at` int(13) DEFAULT NULL COMMENT '创建时间',
  `update_at` int(13) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('0', 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'active', '', null, null, null, null, null, null);
INSERT INTO `user` VALUES ('eb57754a-f1a6-11ea-8278-4ccc6aee07d8', 'a', '', 'active', 'a', 'a', 'a@163.com', '18590aaa-ecf6-11ea-9953-4ccc6aee07d8', '0b345d0a-ed02-11ea-98bb-4ccc6aee07d8', '1599550980', '1599638150');

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` varchar(50) NOT NULL COMMENT 'ID',
  `user_id` varchar(50) NOT NULL COMMENT '用户ID',
  `role_id` varchar(50) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='叶文祥的前端博客：用户表';

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES ('e08d6325-f271-11ea-820c-4ccc6aee07d8', 'eb57754a-f1a6-11ea-8278-4ccc6aee07d8', '5859694a-f180-11ea-80ea-4ccc6aee07d8');
