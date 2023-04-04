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

 Date: 04/04/2023 16:30:08
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
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (1, 'p', 'superuser', '*', '*', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (3, 'p', 'tourist', '/api/v1/users', 'GET', NULL, NULL, NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of permission
-- ----------------------------
BEGIN;
INSERT INTO `permission` (`id`, `name`, `path`, `type`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'file', 'file', NULL, NULL, NULL, NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=889 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'tourist', '2023-03-23 15:57:07', '2023-03-23 15:57:07', NULL);
INSERT INTO `role` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES (333, 'superuser', '2023-03-23 15:51:55', '2023-03-23 15:51:55', NULL);
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
INSERT INTO `role_permission` (`role_id`, `permission_id`) VALUES (1, 1);
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
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'outsider', '$2a$10$fdu/v4gJ/rCD4i5PYqWNa.V7CCgqykcmxDZRjwttr9sUOBrlx/Y8O', 'outsider', 'meet.outsider@gmail.com', 0, '13333333333', '2023-03-24 15:37:43', '2023-03-24 15:37:43', NULL);
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'guest', '$2a$10$YYlnsXijeksDHh/X71W.Gu2JEaTVc.hFKBLYiAOn89nD39S65RHdi', 'gguest', 'example@email.com', 0, '12345678911', '2023-03-25 10:04:14', '2023-03-25 10:04:14', NULL);
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, '叶桂英s', '$2a$10$P28/gedFW4wAEMFX3a.kNeeEGhjMv0/rKqCRv4AAvp4S85/V9sSFy', '冯平', 'l.qcerwhr@qq.com', 0, '18693207699', '2023-03-28 09:26:29', '2023-03-28 09:26:29', '2023-03-28 14:13:51');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, '冯秀兰s', '$2a$10$vBlEqfCnKY4fpI881a/FAe3kln1wrkfk.5ZjZ6yQv5OoJwa6clWRy', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:27:34', '2023-03-28 09:27:34', '2023-03-28 14:13:59');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, '冯秀兰aas', '$2a$10$liVV.PKCqP.CiYYR6S9P4OvkOkLQE2tGQLRFLiHMkmn5liBYL4Vdm', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:27:39', '2023-03-28 09:27:39', '2023-03-28 14:00:35');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, '123123123', '$2a$10$D2xD56GBgewtYIg.RI0l/exENzEZr32K0LfFAAZiUzVUnFxABoWfq', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:27:43', '2023-03-28 09:27:43', '2023-03-28 14:00:35');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, '12312323', '$2a$10$25nS8CHQoWG.crFK3djbQ.6Rj2RHHGfQQWROJ08Y0u8.u/VDd86DS', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:27:46', '2023-03-28 09:27:46', '2023-03-28 14:15:05');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, '12313', '$2a$10$1t/fyjLr2KvDWq1kxWBOAeNAfGBO1zPyj4kl3jcv1RVBIy5uG9.4i', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:27:48', '2023-03-28 09:27:48', '2023-03-28 14:15:20');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, '121233', '$2a$10$XAxesqXu2GIqe6gDHjmYs.05Yzx.sSzqPrn/CryWUMxRUzxtC6tWm', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:27:51', '2023-03-28 09:27:51', '2023-03-28 14:15:22');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (17, 'asdadas', '$2a$10$/aGfsPK5yoru8fEIyUdjwuO1BV3funcjWqV1P/B.RSnkI47VDuXO6', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:27:53', '2023-03-28 09:27:53', '2023-03-28 14:15:24');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, 'asdadasd', '$2a$10$0vFzvU.erYIUmN73E2EMrO0Eb/fofDdyNqWSUcdgrHmg7VK5E/8ZS', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:27:57', '2023-03-28 09:27:57', '2023-03-28 14:15:24');
INSERT INTO `user` (`id`, `username`, `password`, `nickname`, `email`, `status`, `phone`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, 'asda4', '$2a$10$kVLv4bp4V7.jqTKBB8EgruXfTG6713XvlaMc/YPj/ImdAYEDi9xsi', '朱静', 'm.lhx@qq.com', 0, '18135209994', '2023-03-28 09:28:00', '2023-03-28 09:28:00', '2023-03-28 14:15:24');
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
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (1, 333);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (9, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (10, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (11, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (12, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (13, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (14, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (15, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (16, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (17, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (18, 1);
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (19, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
