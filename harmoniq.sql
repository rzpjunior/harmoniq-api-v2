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


DROP TABLE IF EXISTS `record_label`;

CREATE TABLE record_label (
  id int unsigned not null,
  name varchar(50) not null,
  PRIMARY KEY (id),
  UNIQUE KEY uk_name_in_record_label (name)
);

LOCK TABLES `record_label` WRITE;

-- Record Label data
INSERT INTO record_label VALUES(1,'Blackened');
INSERT INTO record_label VALUES(2,'Warner Bros');
INSERT INTO record_label VALUES(3,'Universal');
INSERT INTO record_label VALUES(4,'MCA');
INSERT INTO record_label VALUES(5,'Elektra');
INSERT INTO record_label VALUES(6,'Capitol');

UNLOCK TABLES;

DROP TABLE IF EXISTS `artist`;

-- Artist table
CREATE TABLE artist (
  id int unsigned not null,
  record_label_id int unsigned not null,
  name varchar(50) not null,
  PRIMARY KEY (id),
  KEY fk_record_label_in_artist (record_label_id),
  CONSTRAINT fk_record_label_in_artist FOREIGN KEY (record_label_id) REFERENCES record_label (id),
  UNIQUE KEY uk_name_in_artist (record_label_id, name)
);

LOCK TABLES `artist` WRITE;

-- Artist data
INSERT INTO artist VALUES(1, 1,'Metallica');
INSERT INTO artist VALUES(2, 1,'Megadeth');
INSERT INTO artist VALUES(3, 1,'Anthrax');
INSERT INTO artist VALUES(4, 2,'Eric Clapton');
INSERT INTO artist VALUES(5, 2,'ZZ Top');
INSERT INTO artist VALUES(6, 2,'Van Halen');
INSERT INTO artist VALUES(7, 3,'Lynyrd Skynyrd');
INSERT INTO artist VALUES(8, 3,'AC/DC');
INSERT INTO artist VALUES(9, 6,'The Beatles');

UNLOCK TABLES;

DROP TABLE IF EXISTS `album`;

-- Album Table
CREATE TABLE album (
  id int unsigned not null,
  artist_id int unsigned not null,
  name varchar(50) not null,
  year int unsigned not null,
  PRIMARY KEY (id),
  KEY fk_artist_in_album (artist_id),
  CONSTRAINT fk_artist_in_album FOREIGN KEY (artist_id) REFERENCES artist (id),
  UNIQUE KEY uk_name_in_album (artist_id, name)
);

LOCK TABLES `album` WRITE;

-- Album data
INSERT INTO album VALUES(1, 1, '...And Justice For All',1988);
INSERT INTO album VALUES(2, 1, 'Black Album',1991);
INSERT INTO album VALUES(3, 1, 'Master of Puppets',1986);
INSERT INTO album VALUES(4, 2, 'Endgame',2009);
INSERT INTO album VALUES(5, 2, 'Peace Sells',1986);
INSERT INTO album VALUES(6, 3, 'The Greater of 2 Evils',2004);
INSERT INTO album VALUES(7, 4, 'Reptile',2001);
INSERT INTO album VALUES(8, 4, 'Riding with the King',2000);
INSERT INTO album VALUES(9, 5, 'Greatest Hits',1992);
INSERT INTO album VALUES(10, 6, 'Greatest Hits',2004);
INSERT INTO album VALUES(11, 7, 'All-Time Greatest Hits',1975);
INSERT INTO album VALUES(12, 8, 'Greatest Hits',2003);
INSERT INTO album VALUES(13, 9, 'Sgt. Pepper''s Lonely Hearts Club Band', 1967);

UNLOCK TABLES;

DROP TABLE IF EXISTS `song`;

-- Song table
CREATE TABLE song (
  id int unsigned not null,
  album_id int unsigned not null,
  name varchar(50) not null,
  duration real not null,
  PRIMARY KEY (id),
  KEY fk_album_in_song (album_id),
  CONSTRAINT fk_album_in_song FOREIGN KEY (album_id) REFERENCES album (id),
  UNIQUE KEY uk_name_in_song (album_id, name)
);

LOCK TABLES `song` WRITE;

-- Song data
INSERT INTO song VALUES(1,1,'One',7.25);
INSERT INTO song VALUES(2,1,'Blackened',6.42);
INSERT INTO song VALUES(3,2,'Enter Sandman',5.3);
INSERT INTO song VALUES(4,2,'Sad But True',5.29);
INSERT INTO song VALUES(5,3,'Master of Puppets',8.35);
INSERT INTO song VALUES(6,3,'Battery',5.13);
INSERT INTO song VALUES(7,4,'Dialectic Chaos',2.26);
INSERT INTO song VALUES(8,4,'Endgame',5.57);
INSERT INTO song VALUES(9,5,'Peace Sells',4.09);
INSERT INTO song VALUES(10,5,'The Conjuring',5.09);
INSERT INTO song VALUES(11,6,'Madhouse',4.26);
INSERT INTO song VALUES(12,6,'I am the Law',6.03);
INSERT INTO song VALUES(13,7,'Reptile',3.36);
INSERT INTO song VALUES(14,7,'Modern Girl',4.49);
INSERT INTO song VALUES(15,8,'Riding with the King',4.23);
INSERT INTO song VALUES(16,8,'Key to the Highway',3.39);
INSERT INTO song VALUES(17,9,'Sharp Dressed Man',4.15);
INSERT INTO song VALUES(18,9,'Legs',4.32);
INSERT INTO song VALUES(19,10,'Eruption',1.43);
INSERT INTO song VALUES(20,10,'Hot For Teacher',4.43);
INSERT INTO song VALUES(21,11,'Sweet Home Alabama',4.45);
INSERT INTO song VALUES(22,11,'Free Bird',14.23);
INSERT INTO song VALUES(23,12,'Thunderstruck',4.52);
INSERT INTO song VALUES(24,12,'T.N.T',3.35);
INSERT INTO song VALUES(25,13,'Sgt. Pepper''s Lonely Hearts Club Band', 2.0333);
INSERT INTO song VALUES(26,13,'With a Little Help from My Friends', 2.7333);
INSERT INTO song VALUES(27,13,'Lucy in the Sky with Diamonds', 3.4666);
INSERT INTO song VALUES(28,13,'Getting Better', 2.80);
INSERT INTO song VALUES(29,13,'Fixing a Hole', 2.60);
INSERT INTO song VALUES(30,13,'She''s Leaving Home', 3.5833);
INSERT INTO song VALUES(31,13,'Being for the Benefit of Mr. Kite!',2.6166);
INSERT INTO song VALUES(32,13,'Within You Without You',5.066);
INSERT INTO song VALUES(33,13,'When I''m Sixty-Four',2.6166);
INSERT INTO song VALUES(34,13,'Lovely Rita', 2.7);
INSERT INTO song VALUES(35,13,'Good Morning Good Morning', 2.6833);
INSERT INTO song VALUES(36,13,'Sgt. Pepper''s Lonely Hearts Club Band (Reprise)', 1.3166);
INSERT INTO song VALUES(37,13,'A Day in the Life', 5.65);

UNLOCK TABLES

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-12-13 17:17:00
