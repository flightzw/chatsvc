CREATE DATABASE IF NOT EXISTS `chatsvc`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像url',
  `nickname` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
  `gender` tinyint NOT NULL DEFAULT '0' COMMENT '性别 0:未知 1:男 2:女',
  `signature` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '个性签名',
  `type` tinyint unsigned NOT NULL COMMENT '类型 1:普通用户 2:AI模型 3管理员',
  `status` tinyint unsigned NOT NULL COMMENT '状态 1:正常 2:封禁',
  `last_login_at` timestamp NULL DEFAULT NULL COMMENT '最后上线时间',
  `last_login_ip` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '最后上线ip',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=10008 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户';

CREATE TABLE IF NOT EXISTS `friends` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `friend_id` int unsigned NOT NULL COMMENT '好友id',
  `friend_nickname` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
  `friend_avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '头像url',
  `remark` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '备注',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_id_friend_id` (`user_id`,`friend_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='好友';

CREATE TABLE IF NOT EXISTS `private_messages` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `send_id` int unsigned NOT NULL COMMENT '用户uid',
  `recv_id` int unsigned NOT NULL COMMENT '好友uid',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '发送内容',
  `type` tinyint NOT NULL COMMENT '消息类型',
  `status` tinyint NOT NULL COMMENT '状态 0:未送达 1:已送达 2:撤回 3:已读',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='私聊消息';

CREATE TABLE IF NOT EXISTS `sensitive_words` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(64) NOT NULL COMMENT '敏感词内容',
  `enabled` tinyint NOT NULL DEFAULT 0 COMMENT '是否启用 1:是 0:否',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_content` (`content`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='敏感词';

CREATE TABLE IF NOT EXISTS `configs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int unsigned NOT NULL COMMENT '上级配置 id',
  `key` varchar(20) NOT NULL COMMENT '配置名称',
  `value` varchar(1024) NOT NULL COMMENT '配置参数',
  `comment` varchar(20) NOT NULL COMMENT '备注',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_parent_id_key` (`parent_id`, `key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='配置项';