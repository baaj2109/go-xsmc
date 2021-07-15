/*
Navicat MySQL Data Transfer
Source Server         : local
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : xcms
Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for data
-- ----------------------------
DROP TABLE IF EXISTS `data`;
CREATE TABLE `data` (
  `did` int(11) NOT NULL AUTO_INCREMENT,
  `mid` int(11) NOT NULL DEFAULT '0',
  `parent` int(11) NOT NULL DEFAULT '0',
  `name` varchar(60) NOT NULL DEFAULT '',
  `content` varchar(2048) NOT NULL DEFAULT '{}',
  `seq` int(11) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `update_time` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`did`),
  KEY `data_seq` (`seq`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of data
-- ----------------------------
INSERT INTO `data` VALUES ('1', '9', '0', '小明22', '{\n  \"parent\": 0,\n  \"name\": \"小明22\",\n  \"seq\": \"0\",\n  \"status\": \"0\",\n  \"username\": \"学生姓名\",\n  \"sex\": \"male\",\n  \"class\": 3\n}', '0', '0', '1593919514');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `mid` int(11) NOT NULL AUTO_INCREMENT,
  `parent` int(11) NOT NULL DEFAULT '0',
  `name` varchar(45) NOT NULL DEFAULT '',
  `seq` int(11) NOT NULL DEFAULT '0',
  `format` varchar(2048) NOT NULL DEFAULT '{}',
  PRIMARY KEY (`mid`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES ('1', '0', '商城', '0', '{\"schema\":{\"name\":{\"type\":\"string\",\"title\":\"姓名\"},\"class\":{\"type\":\"number\",\"title\":\"班级\"},\"sex\":{\"type\":\"string\",\"title\":\"班级\",\"enum\":[\"unkown\",\"male\",\"female\"]}},\"form\":[{\"key\":\"name\"},{\"key\":\"sex\"},{\"key\":\"class\"}]}');
INSERT INTO `menu` VALUES ('2', '1', '商品管理', '0', '{\"schema\":{\"name\":{\"type\":\"string\",\"title\":\"姓名\"},\"class\":{\"type\":\"number\",\"title\":\"班级\"},\"sex\":{\"type\":\"string\",\"title\":\"班级\",\"enum\":[\"unkown\",\"male\",\"female\"]}},\"form\":[{\"key\":\"name\"},{\"key\":\"sex\"},{\"key\":\"class\"}]}');
INSERT INTO `menu` VALUES ('3', '1', '购物车', '1', '{\"schema\":{\"name\":{\"type\":\"string\",\"title\":\"姓名\"},\"class\":{\"type\":\"number\",\"title\":\"班级\"},\"sex\":{\"type\":\"string\",\"title\":\"班级\",\"enum\":[\"unkown\",\"male\",\"female\"]}},\"form\":[{\"key\":\"name\"},{\"key\":\"sex\"},{\"key\":\"class\"}]}');
INSERT INTO `menu` VALUES ('4', '1', '会员管理', '0', '');
INSERT INTO `menu` VALUES ('5', '0', '书城', '0', '');
INSERT INTO `menu` VALUES ('6', '5', '图书管理', '0', '');
INSERT INTO `menu` VALUES ('7', '0', '学生信息管理', '0', '{\"schema\":{\"name\":{\"type\":\"string\",\"title\":\"姓名\"},\"class\":{\"type\":\"number\",\"title\":\"班级\"},\"sex\":{\"type\":\"string\",\"title\":\"性别\",\"enum\":[\"unkown\",\"male\",\"female\"]}},\"form\":[{\"key\":\"name\"},{\"key\":\"sex\"},{\"key\":\"class\"}]}');
INSERT INTO `menu` VALUES ('8', '7', '班级信息', '0', '{\"schema\":{\"name\":{\"type\":\"string\",\"title\":\"姓名\"},\"class\":{\"type\":\"number\",\"title\":\"班级\"},\"sex\":{\"type\":\"string\",\"title\":\"性别\",\"enum\":[\"unkown\",\"male\",\"female\"]}},\"form\":[{\"key\":\"name\"},{\"key\":\"sex\"},{\"key\":\"class\"}]}');
INSERT INTO `menu` VALUES ('9', '7', '学生信息', '0', '{\"schema\":{\"username\":{\"type\":\"string\",\"title\":\"姓名\"},\"class\":{\"type\":\"number\",\"title\":\"班级\"},\"sex\":{\"type\":\"string\",\"title\":\"性别\",\"enum\":[\"unkown\",\"male\",\"female\"]}},\"form\":[{\"key\":\"username\"},{\"key\":\"sex\"},{\"key\":\"class\"}]}');

-- ----------------------------
-- Table structure for page
-- ----------------------------
DROP TABLE IF EXISTS `page`;
CREATE TABLE `page` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '信箱',
  `website` varchar(64) NOT NULL DEFAULT '' COMMENT '網站',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of page
-- ----------------------------
INSERT INTO `page` VALUES ('1', 'myemail_updated', 'www.10php.cn');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_key` varchar(64) NOT NULL DEFAULT '',
  `user_name` varchar(64) NOT NULL DEFAULT '',
  `auth_str` varchar(512) NOT NULL DEFAULT '',
  `pass_word` varchar(128) NOT NULL DEFAULT '',
  `is_admin` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_key` (`user_key`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '123', 'baaj', '[1,5,7]', '123', '1');
INSERT INTO `user` VALUES ('2', 'aa', 'bb', '[1]', 'aa', '1');