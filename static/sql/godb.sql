/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : godb

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-07-03 21:20:40
*/

SET FOREIGN_KEY_CHECKS=0;

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
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES ('1', '0', '商城', '0', '{\"schema\": {\"field\":{\"type\":\"string\",\"title\":\"name\"}},\"form\":[{\"key\":\"field\"},{\"type\":\"submit\",\"title\":\"submit\"}]}');
INSERT INTO `menu` VALUES ('2', '1', '商品管理', '0', '{\"a\":\"b\"}');
INSERT INTO `menu` VALUES ('3', '1', '购物车', '1', '');
INSERT INTO `menu` VALUES ('4', '1', '会员管理', '0', '');
INSERT INTO `menu` VALUES ('5', '0', '书城', '0', '');
INSERT INTO `menu` VALUES ('6', '5', '图书管理', '0', '');

-- ----------------------------
-- Table structure for page
-- ----------------------------
DROP TABLE IF EXISTS `page`;
CREATE TABLE `page` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
  `website` varchar(64) NOT NULL DEFAULT '' COMMENT '网站',
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
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'aa', 'leixiaotian', '15', '12345', '0');
