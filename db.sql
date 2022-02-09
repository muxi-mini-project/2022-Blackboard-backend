CREATE DATABASE IF NOT EXISTS `blackboard`;

USE `blackboard`;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `collections`;
DROP TABLE IF EXISTS `following_organizations`;
DROP TABLE IF EXISTS `organizations`;
DROP TABLE IF EXISTS `groupings`;
DROP TABLE IF EXISTS `announcements`;


CREATE TABLE `users`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` datetime NULL,            
    `updated_at` datetime NULL,             
    `deleted_at` datetime NULL,  
    `student_id` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `nickname` VARCHAR(255) NOT NULL,
    `avatar` VARCHAR(255) NOT NULL,
    `sha` VARCHAR(255) NOT NULL,
    `path` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `collections`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` datetime NULL,            
    `updated_at` datetime NULL,            
    `deleted_at` datetime NULL,  
    `student_id` VARCHAR(100) NOT NULL,
    `announcement_id` VARCHAR(100) NOT NULL,
    `announcement` VARCHAR(500) NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--用户关注的组织
CREATE TABLE `following_organizations`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` datetime NULL,            
    `updated_at` datetime NULL,            
    `deleted_at` datetime NULL,  
    `student_id` VARCHAR(100) NOT NULL,
    `organization_id` VARCHAR(100) NOT NULL,
    `organization_name` VARCHAR(100) NOT NULL,
PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--已存在的组织
CREATE TABLE `organizations`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` datetime NULL,            
    `updated_at` datetime NULL,            
    `deleted_at` datetime NULL, 
     `founder_id` VARCHAR(100) NOT NULL,
    `organization_name` VARCHAR(100) NOT NULL,
    `organization_intro` VARCHAR(100) NOT NULL,
    `avatar` VARCHAR(255) NOT NULL,
    `sha` VARCHAR(255) NOT NULL,
    `path` VARCHAR(255) NOT NULL,
PRIMARY KEY (`id`),
FULLTEXT (`organization_name`) WITH PARSER ngram
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--通告分类
CREATE TABLE `groupings`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` datetime NULL,            
    `updated_at` datetime NULL,            
    `deleted_at` datetime NULL, 
    `organization_id` VARCHAR(100) NOT NULL,
    `organization_name` VARCHAR(100) NOT NULL,
    `group_name` VARCHAR(100) NOT NULL,
PRIMARY KEY(`id`),
FULLTEXT (`group_name`) WITH PARSER ngram
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--通告
CREATE TABLE `announcements`(
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `created_at` datetime NULL,            
    `updated_at` datetime NULL,            
    `deleted_at` datetime NULL, 
    `publisher_id` VARCHAR(255) NOT NULL,
    `organization_id` VARCHAR(255) NOT NULL,
    `organization_name` VARCHAR(255) NOT NULL,
    `group_id` VARCHAR(255) NOT NULL,
    `group_name` VARCHAR(255) NOT NULL,
    `contents` VARCHAR(500) NOT NULL,
PRIMARY KEY (`id`),
FULLTEXT (`organization_name`) WITH PARSER ngram,
FULLTEXT (`group_name`) WITH PARSER ngram,
FULLTEXT (`contents`) WITH PARSER ngram
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
