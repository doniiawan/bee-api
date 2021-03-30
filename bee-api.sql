/*
 Navicat Premium Data Transfer

 Source Server         : Local MySQL
 Source Server Type    : MySQL
 Source Server Version : 50726
 Source Host           : localhost:3306
 Source Schema         : bee-api

 Target Server Type    : MySQL
 Target Server Version : 50726
 File Encoding         : 65001

 Date: 26/03/2021 10:59:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'coba', '$2b$10$lIGjNgnnLPfC9Sk9Q1SwRehvpaSCV.tNSXpChlJHic.DmtI/gpRcC');
INSERT INTO `user` VALUES (2, 'qwe', 'qwe');
INSERT INTO `user` VALUES (3, 'asd', 'asd');
INSERT INTO `user` VALUES (4, 'admin', 'admin');
INSERT INTO `user` VALUES (5, 'admin', 'admin');
INSERT INTO `user` VALUES (6, 'admin', 'admin');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
