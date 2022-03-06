-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: localhost    Database: pustaka-api
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `jenis` varchar(100) DEFAULT NULL,
  `name` longtext,
  `description` longtext,
  `rating` bigint DEFAULT NULL,
  `price_before` bigint DEFAULT NULL,
  `price_today` bigint DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'Makanan Pokok','Minyak GOReng','Minyak langka ga iso di gae goreng',5,14000,19000,'2022-02-28 07:53:58.332','2022-02-28 07:53:58.332'),(2,'Sembako','Beras PKH','Beras terbaik untuk rakyat yang paling baik',5,14000,12000,'2022-02-28 07:55:39.973','2022-02-28 07:55:39.973'),(3,'Sembako','Program Sedekah Bumi','Tiap 2 minggu sekali dapat bantuan',4,20000,10000,'2022-02-28 07:59:09.670','2022-02-28 20:37:04.093'),(4,'Miras','Gedhang Klutuk','Jamu mujarab bagi orang sakit hati',5,25000,40000,'2022-02-28 08:00:34.549','2022-02-28 08:00:34.549'),(6,'Miras','Anggur Merah Orang Tua','Rasa enak lah',5,45000,65000,'2022-02-28 08:50:21.237','2022-02-28 08:50:21.237'),(8,'Sembako','Mie GOreng','rasane mantep pokok',5,2500,3000,'2022-03-02 01:04:49.069','2022-03-02 01:04:49.069'),(9,'Sembako','Mie GOreng','rasane mantep pokok',5,2500,3000,'2022-03-02 07:50:09.763','2022-03-02 07:50:09.763'),(10,'Sembako','Mie Dog-dpg','rasane joss gandos',5,2500,3000,'2022-03-02 07:53:44.690','2022-03-02 07:53:44.690'),(11,'Sembako','Mie Geprek','rasane joss pedes poll',5,10000,15000,'2022-03-02 08:40:28.489','2022-03-02 08:40:28.489'),(12,'Sembako','Mie Geprek','rasane joss pedes poll',5,10000,15000,'2022-03-05 12:59:50.844','2022-03-05 12:59:50.844'),(13,'Sembako','Beras SKTM','Pemberian beras untuk PKH dilakukan 10 harin',10,12500,65000,'2022-03-05 14:29:48.287','2022-03-05 14:29:48.287'),(14,'Sembako','Beras Berkutu','Pemberian beras untuk PKH dilakukan 10 harin',10,12500,65000,'2022-03-05 14:30:04.993','2022-03-05 14:30:04.993');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-03-06 16:39:34
