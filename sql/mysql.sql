
/*Table structure for table `greet_member` */

DROP TABLE IF EXISTS `greet_member`;

CREATE TABLE `greet_member` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'Id',
    `mobile` char(11) NOT NULL DEFAULT '' COMMENT '登陆手机号',
    `nickname` varchar(50) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `uk_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;