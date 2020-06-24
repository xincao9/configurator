CREATE DATABASE `configurator` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `configurator`;

DROP TABLE IF EXISTS `session`;

CREATE TABLE `session`
(
    `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`   char(32)            NOT NULL DEFAULT '',
    `password`   char(32)            NOT NULL DEFAULT '',
    `expire`     timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `token`      char(128)           NOT NULL DEFAULT '',
    `created_at` timestamp           NULL     DEFAULT NULL,
    `deleted_at` timestamp           NULL     DEFAULT NULL,
    `updated_at` timestamp           NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username_idx` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
