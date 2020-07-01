CREATE DATABASE `configurator` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `configurator`;

DROP TABLE IF EXISTS `account`;

CREATE TABLE `account`
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

DROP TABLE IF EXISTS `app`;

CREATE TABLE `app`
(
    `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT,
    `env`        char(32)            NOT NULL DEFAULT '',
    `group`      char(64)            NOT NULL DEFAULT '',
    `project`    char(64)            NOT NULL DEFAULT '',
    `version`    char(64)            NOT NULL DEFAULT '',
    `created_at` timestamp           NULL     DEFAULT NULL,
    `deleted_at` timestamp           NULL     DEFAULT NULL,
    `updated_at` timestamp           NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `app_idx` (`env`, `group`, `project`, `version`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `message_box`;

CREATE TABLE `message_box`
(
    `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`   char(32)            NOT NULL DEFAULT '',
    `message`    varchar(256)        NOT NULL DEFAULT '',
    `status`     int(11) unsigned    NOT NULL,
    `created_at` timestamp           NULL     DEFAULT NULL,
    `deleted_at` timestamp           NULL     DEFAULT NULL,
    `updated_at` timestamp           NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `username_idx` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

DROP TABLE IF EXISTS `operation_log`;

CREATE TABLE `operation_log`
(
    `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`   char(32)            NOT NULL DEFAULT '',
    `message`    varchar(256)        NOT NULL DEFAULT '',
    `created_at` timestamp           NULL     DEFAULT NULL,
    `deleted_at` timestamp           NULL     DEFAULT NULL,
    `updated_at` timestamp           NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `username_idx` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
