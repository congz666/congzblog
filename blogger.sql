/*
SQLyog Ultimate v10.00 Beta1
MySQL - 5.5.58 : Database - congzblog
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`congzblog` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `congzblog`;

/*Table structure for table `article` */

DROP TABLE IF EXISTS `article`;

CREATE TABLE `article` (
  `article_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint(20) NOT NULL COMMENT '分类id',
  `content` longtext NOT NULL COMMENT '文章内容',
  `title` varchar(1024) NOT NULL COMMENT '文章标题',
  `view_count` int(255) NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) NOT NULL COMMENT '评论次数',
  `username` varchar(128) NOT NULL COMMENT '作者',
  `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1',
  `summary` varchar(256) NOT NULL COMMENT '文章摘要',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`article_id`),
  FOREIGN KEY(`category_id`) REFERENCES category(`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

/*Data for the table `article` */



/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `category_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `category_name` varchar(255) NOT NULL COMMENT '分类名字',
  `category_no` int(10) NOT NULL COMMENT '分类排序',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

/*Data for the table `category` */

insert  into `category`(`category_id`,`category_name`,`category_no`,`create_time`,`update_time`) values (1,'golang开发',0,'2019-08-12 10:55:45','2019-08-12 10:59:00'),(2,'数据库',0,'2019-08-12 10:56:07','2019-08-12 10:59:03'),(3,'leetcode题目',0,'2019-08-12 10:56:16','2019-08-12 10:59:05'),(4,'前端开发',0,'2019-08-12 10:56:24','2019-08-12 10:59:08'),(5,'架构剖析',0,'2019-08-12 10:56:36','2019-08-12 10:59:10');

/*Table structure for table `comment` */

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `comment_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text NOT NULL COMMENT '评论内容',
  `summary` varchar(256) NOT NULL COMMENT '评论摘要',
  `username` varchar(64) NOT NULL COMMENT '评论作者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `email` varchar(255) NOT NULL,
  `article_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`comment_id`),
  FOREIGN KEY(`article_id`) REFERENCES article(`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `comment` */

/*Table structure for table `leave` */

DROP TABLE IF EXISTS `leave`;

CREATE TABLE `leave` (
  `leave_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `summary` varchar(256) NOT NULL COMMENT '留言摘要',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`leave_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `leave` */

/*Table structure for table `administrator` */

DROP TABLE IF EXISTS `administrator`;

CREATE TABLE `administrator` (
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `administrator` */
insert into `administrator`(`name`,`password`)values('congz','4636d609d7c4ce291abb8ebcbbe8ca5117df735b');

DROP TABLE IF EXISTS `session`;

CREATE TABLE `session` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
