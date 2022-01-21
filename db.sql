CREATE DATABASE IF NOT EXISTS `blackboard`;

USE `blackboard`;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `collections`;
DROP TABLE IF EXISTS `organizations_created`;
DROP TABLE IF EXISTS `organizations_following`;
DROP TABLE IF EXISTS `organizations`;
DROP TABLE IF EXISTS `groups`;
DROP TABLE IF EXISTS `announcements`;


CREATE TABLE `users`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` VARCHAR(100) NOT NULL,
    `password` VARCHAR(100) NOT NULL,
    `nickname` VARCHAR(100) NOT NULL,
    `headportrait` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `collections`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` VARCHAR(100) NOT NULL,
    `announcement_id` VARCHAR(100) NOT NULL,
    `announcement` VARCHAR(500) NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--用户关注的组织
CREATE TABLE `following_organizations`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` VARCHAR(100) NOT NULL,
    `org_id` VARCHAR(100) NOT NULL,
    `org_name` VARCHAR(100) NOT NULL,
PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--已存在的组织
CREATE TABLE `organizations`(
    `org_id` BIGINT NOT NULL AUTO_INCREMENT,
    `org_logo` VARCHAR(100) NOT NULL,
    `org_name` VARCHAR(100) NOT NULL,
    `intro` VARCHAR(100) NOT NULL,
    `founder_id` VARCHAR(100) NOT NULL,
PRIMARY KEY (`org_id`),
FULLTEXT (`org_name`) WITH PARSER ngram
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--通告分类
CREATE TABLE `groups`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `org_id` VARCHAR(100) NOT NULL,
    `org_name` VARCHAR(100) NOT NULL,
    `group_name` VARCHAR(100) NOT NULL,
PRIMARY KEY(`id`),
FULLTEXT (`group_name`) WITH PARSER ngram
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--通告
CREATE TABLE `announcements`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `publisher_id` BIGINT NOT NULL,
    `org_id` VARCHAR(100) NOT NULL,
    `org_name` VARCHAR(100) NOT NULL,
    `group_id` VARCHAR(100) NOT NULL,
    `group_name` VARCHAR(100) NOT NULL,
    `contents` VARCHAR(500) NOT NULL,
PRIMARY KEY (`id`),
FULLTEXT (`org_name`) WITH PARSER ngram,
FULLTEXT (`group_name`) WITH PARSER ngram,
FULLTEXT (`contents`) WITH PARSER ngram
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
