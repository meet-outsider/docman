/*
 Navicat Premium Data Transfer

 Source Server         : root@localhost
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3306
 Source Schema         : docman

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 22/03/2023 22:35:21
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (3, 'p', 'guest', '/info', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (1, 'p', 'root', '/info', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for file
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `path` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `is_dir` tinyint DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of file
-- ----------------------------
BEGIN;
INSERT INTO `file` (`id`, `name`, `path`, `is_dir`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '1.png', 'images/1.png', 0, '2023-03-20 14:56:07', '2023-03-20 14:56:07', '2023-03-20 15:03:17');
INSERT INTO `file` (`id`, `name`, `path`, `is_dir`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '1.png', 'images/1.png', 0, '2023-03-20 14:56:26', '2023-03-20 14:56:26', '2023-03-20 16:35:28');
INSERT INTO `file` (`id`, `name`, `path`, `is_dir`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, '1.png', 'images/1.png', 0, '2023-03-20 16:53:11', '2023-03-20 16:53:11', NULL);
INSERT INTO `file` (`id`, `name`, `path`, `is_dir`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, '1.png', 'images/1.png', 0, '2023-03-20 17:02:41', '2023-03-20 17:02:41', NULL);
INSERT INTO `file` (`id`, `name`, `path`, `is_dir`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, '1.png', 'images/1.png', 0, '2023-03-20 17:03:10', '2023-03-20 17:03:10', NULL);
INSERT INTO `file` (`id`, `name`, `path`, `is_dir`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, '1.png', 'images/1.png', 0, '2023-03-20 17:03:55', '2023-03-20 17:03:55', NULL);
COMMIT;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `type` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of permission
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'tourists', '2023-03-22 16:22:15', '2023-03-22 16:22:15', NULL);
COMMIT;

-- ----------------------------
-- Table structure for role_permission
-- ----------------------------
DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission` (
  `role_id` int NOT NULL,
  `permission_id` int NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of role_permission
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `nickname` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `status` tinyint DEFAULT NULL,
  `phone` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'outsider', 'outsider@12', NULL, 'meet.outsider@gmail.com', 0, '15538854641', '2023-03-20 13:44:46', NULL, NULL);
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'root', 'root', '', '', NULL, NULL, '2023-03-22 15:16:19', '2023-03-22 15:16:19', NULL);
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'guest', 'guest', '', '', NULL, NULL, '2023-03-22 15:16:42', '2023-03-22 15:16:42', NULL);
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `user_id` int NOT NULL,
  `role_id` int NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
