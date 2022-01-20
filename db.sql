CREATE DATABASE IF NOT EXISTS `blackboard`;

USE `blackboard`;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `collections`;
DROP TABLE IF EXISTS `organizations_created`;
DROP TABLE IF EXISTS `organizations_following`;
DROP TABLE IF EXISTS `organizations`;
DROP TABLE IF EXISTS `groups`;
DROP TABLE IF EXISTS `announcements`;

--用户信息
CREATE TABLE `users`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` BIGINT NOT NULL,
    `password` varchar(100) NOT NULL,
    `nickname` VARCHAR(100) NOT NULL,
    `headportrait` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `collections`{
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` BIGINT NOT NULL,
    `announcement_id` BIGINT NOT NULL,
    `announcement` VARCHAR NULL,
    PRIMARY KEY (`id`)
}ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `organizations_created`{
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` BIGINT NOT NULL,
    `org_id` BIGINT NOT NULL,
    `org_name` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
}ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `organizations_following`{
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `student_id` BIGINT NOT NULL,
    `org_id` VARCHAR(100) NOT NULL,
    `org_name` VARCHAR(100) NOT NULL,
    PRIMARY KEY (`id`)
}ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `organizations`{
    `org_id` BIGINT NOT NULL AUTO_INCREMENT,
    `org_logo` VARCHAR(100) NOT NULL,
    `org_name` VARCHAR(100) NOT NULL,
    `intro` VARCHAR(100) NOT NULL,
    `org_groups` VARCHAR NOT NULL,
    PRIMARY KEY (`org_id`),
    FULLTEXT (`org_name`) WITH PARSER ngram
}ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `groups`{
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `org_name` BIGINT NOT NULL,
    'group_name' VARCHAR NOT NULL,
    PRIMARY KEY(`id`),
    FULLTEXT (`group_name`) WITH PARSER ngram
}ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `announcements`{
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `contents` varchar(500) NOT NULL,
    PRIMARY KEY (`id`),
    FULLTEXT (`contents`) WITH PARSER ngram
}ENGINE=InnoDB DEFAULT CHARSET=utf8;
