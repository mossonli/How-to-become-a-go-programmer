/*
 Navicat Premium Data Transfer

 Source Server         : mysqlDocker
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : localhost:3306
 Source Schema         : blogger

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 16/06/2020 10:36:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint(20) NOT NULL COMMENT '分类id',
  `content` longtext NOT NULL COMMENT '文章内容',
  `title` varchar(1024) NOT NULL COMMENT '文章标题',
  `view_count` int(255) NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) NOT NULL COMMENT '评论次数',
  `username` varchar(128) NOT NULL COMMENT '作者',
  `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1',
  `summary` varchar(256) NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article
-- ----------------------------
BEGIN;
INSERT INTO `article` VALUES (1, 1, 'this a test ak dkdkdkddkddkd111', '我是标题1', 1, 0, 'Mr.Sun1', 1, '我是\n									很多的\n									内容1', '2019-10-04 23:34:06', NULL);
INSERT INTO `article` VALUES (2, 2, 'this a test ak dkdkdkddkddkd222', '我是标题2', 1, 0, 'Mr.Sun2', 1, '我是\n									很多的\n									内容2', '2019-10-04 23:34:39', NULL);
INSERT INTO `article` VALUES (3, 2, 'this a test ak dkdkdkddkddkd333', '我是标题3', 1, 1, 'Mr.Sun3', 1, '我是\n									很多的\n									内容3', '2019-10-04 23:34:55', NULL);
INSERT INTO `article` VALUES (4, 1, '这里是文章的内容', '文章标题', 1, 0, 'Mosson', 1, '文章总结', '2020-06-11 13:09:24', NULL);
INSERT INTO `article` VALUES (5, 1, '这里是文章的内容', '文章标题', 1, 0, 'Mosson', 1, '文章总结', '2020-06-11 13:17:55', NULL);
INSERT INTO `article` VALUES (6, 1, '<p>我是测试内容</p>', '测试', 0, 1, 'mosson', 1, '<p>我是测试内容</p>', '2020-06-15 17:39:40', NULL);
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `category_name` varchar(255) NOT NULL COMMENT '分类名字',
  `category_no` int(10) NOT NULL COMMENT '分类排序',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1, 'css/html', 1, '2019-08-12 10:55:45', '2019-08-12 10:59:00');
INSERT INTO `category` VALUES (2, '后端开发', 2, '2019-08-12 10:56:07', '2019-08-12 10:59:03');
INSERT INTO `category` VALUES (3, 'Java开发', 3, '2019-08-12 10:56:16', '2019-08-12 10:59:05');
INSERT INTO `category` VALUES (4, 'C++开发', 4, '2019-08-12 10:56:24', '2019-08-12 10:59:08');
INSERT INTO `category` VALUES (5, '架构剖析', 5, '2019-08-12 10:56:36', '2019-08-12 10:59:10');
INSERT INTO `category` VALUES (6, 'Golang开发', 6, '2019-08-12 10:56:45', '2019-08-12 10:59:14');
COMMIT;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text NOT NULL COMMENT '评论内容',
  `username` varchar(64) NOT NULL COMMENT '评论作者',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) NOT NULL DEFAULT '1' COMMENT '评论状态: 0, 删除；1， 正常',
  `article_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
INSERT INTO `comment` VALUES (1, '我是评论1', '小李', '2020-06-15 17:40:42', 1, 6);
COMMIT;

-- ----------------------------
-- Table structure for leave
-- ----------------------------
DROP TABLE IF EXISTS `leave`;
CREATE TABLE `leave` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
