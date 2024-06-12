CREATE DATABASE  IF NOT EXISTS `harmoniq-dev` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `harmoniq-dev`;
-- MySQL dump 
--
-- Host: localhost    Database: harmoniq-dev
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product`  (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(32) NULL DEFAULT NULL,
  `name` varchar(100) NULL DEFAULT NULL,
  `description` text NULL,
  `category_id` bigint(11) NULL DEFAULT NULL,
  `weight` decimal(10, 2) NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  `stock` int NULL DEFAULT NULL,
  `status` int NULL DEFAULT NULL COMMENT '1:active,2:inactive',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (1, 'PRD0001', 'Product 1', 'Lorem Ipsum is simply dummy text of the printing and typesetting industry.', 2, 3600.00, 34000.00, 100, 1, '2023-10-14 23:48:51', '2023-10-14 23:50:02');
INSERT INTO `product` VALUES (2, 'PRD0002', 'Product 2', 'Lorem Ipsum is simply dummy text of the printing and typesetting industry.', 3, 9100.00, 72000.00, 100, 1, '2023-10-14 23:48:53', '2023-10-14 23:50:02');
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_image`
--

DROP TABLE IF EXISTS `product_image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_image`  (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `product_id` bigint(11) NULL DEFAULT NULL,
  `image_url` text NULL,
  `main_image` int NULL DEFAULT NULL COMMENT '1:yes,2:no',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product_image` WRITE;
/*!40000 ALTER TABLE `product_image` DISABLE KEYS */;
INSERT INTO `product_image` 
VALUES (1, 1, 'https://api.dicebear.com/7.x/icons/svg?seed=Lilly', 1, '2023-10-14 23:51:50', NULL);
/*!40000 ALTER TABLE `product_image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category`  (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NULL DEFAULT NULL,
  `icon_url` text NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1, 'Makanan', 'https://api.dicebear.com/5.x/icons/svg?icon=box', '2023-10-14 17:45:00', '2023-10-14 17:49:42');
INSERT INTO `category` VALUES (2, 'Peralatan Dapur', 'https://api.dicebear.com/5.x/icons/svg?icon=box', '2023-10-14 17:45:09', '2023-10-14 17:49:42');
INSERT INTO `category` VALUES (3, 'Minuman', 'https://api.dicebear.com/5.x/icons/svg?icon=box', '2023-10-14 17:45:13', '2023-10-14 17:49:42');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(100) DEFAULT '',
  `phone` varchar(15) DEFAULT '',
  `password` varchar(250) DEFAULT '',
  `name` varchar(100) DEFAULT '',
  `status` tinyint(1) DEFAULT '0',
  `last_login_at` datetime NULL DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user`(`email`,`phone`,`password`,`name`,`status`,`created_at`)
VALUES('ardi@example.com','+6281122334444','$2a$14$taBX9l6UqoiQBT2oi0AM3eqDYO2CFBqIYPQY1AomrT0MQzkbt7Rmy','ardi',1,NOW()),
('atun@example.com','+6281122334444','$2a$14$taBX9l6UqoiQBT2oi0AM3eqDYO2CFBqIYPQY1AomrT0MQzkbt7Rmy','atun wati',1,NOW()),
('joy@example.com','+6281122334444','$2a$14$taBX9l6UqoiQBT2oi0AM3eqDYO2CFBqIYPQY1AomrT0MQzkbt7Rmy','joy boy',1,NOW());
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cart`
--

DROP TABLE IF EXISTS `cart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cart`  (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(11) NULL DEFAULT NULL,
  `product_id` bigint(11) NULL DEFAULT NULL,
  `qty` int NULL DEFAULT NULL,
  `status` int NULL DEFAULT NULL COMMENT '1:active,2:inactive',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cart`
--

LOCK TABLES `cart` WRITE;
/*!40000 ALTER TABLE `cart` DISABLE KEYS */;
INSERT INTO `cart` VALUES (1, 4, 1, 11, 1, '2023-10-14 23:56:29', '2023-10-14 23:56:29');
/*!40000 ALTER TABLE `cart` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-13 17:17:00
