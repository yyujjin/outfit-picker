-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: outfit-picker
-- ------------------------------------------------------
-- Server version	8.0.34

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
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `userId` varchar(45) NOT NULL,
  `password` varchar(100) NOT NULL,
  `name` varchar(45) NOT NULL,
  `birthday` date NOT NULL,
  `tel` varchar(45) NOT NULL,
  `gender` tinyint unsigned NOT NULL COMMENT '0 male\n1 female',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `userId_UNIQUE` (`userId`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='사용자 테이블  ';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'losgenle','lalland12','박유진','1996-12-04','010-1244-5678',1),(3,'losgenle12','lalland12','바비','1996-12-04','010-1244-5678',1),(4,'lalasweet','nono4545','박혁거세','2004-12-14','010-4444-5555',0),(7,'lalasweet1231','nono4545','박혁거세','2004-12-14','010-4444-5555',0),(9,'lalasweet12314','nono4545','박혁거세','2004-12-14','010-4444-5555',0),(11,'lalasweet12315','nono4545','박혁거세','2004-12-14','010-4444-5555',0),(14,'lalasweet12315d','nono4545','박혁거세','2004-12-14','010-4444-5555',0),(15,'lalasweet12315dd','nono4545','박혁거세','2004-12-14','010-4444-5555',0),(16,'lalaswesqq','$2a$10$IK6E6ydop.EdFZ0D8cmUpuhKZlszypVr9.mxVN4y.zD/3nqMi628a','박혁거세','2004-12-14','010-4444-5555',0),(17,'MUMU','$2a$10$OBIAUYInn5sCFUWR57kYCeiLAX/cS61w9ljXa8jawCdLKoHxuMIyy','세일러문','2004-12-14','010-4444-5555',0),(18,'OMG','$2a$10$SWZ0Ajhp5ZoLoMSw9Iwluu814Lm6Wcc0HqdYalyzPpCH/5GI1S/mm','세일러문','2004-12-14','010-4444-5555',0),(19,'O','$2a$10$sm0UG1vfP/HEO/RN86cc.eyDjWDgSDwns74e2NLulMXSTD8Mso/gy','세일러문','2004-12-14','010-4444-5555',0),(20,'yaho','$2a$10$nn7nO/9IwvEL3F6nNXkw3eaoR6dnt9UHsfL3vEHcp02FXaqOvLKqG','응가','1996-12-04','010-1234-5678',1),(22,'Ogg','$2a$10$5qAs2A73.S/VZnZNg69gV.p03TAhYOsyF6eAr.BN42mTYwxupi5gu','세일러문','2004-12-14','010-4444-5555',0),(26,'Ofg','$2a$10$9YYhYgC07Vyf/6PlfRARAunNDlUA.WjuzA7p.SgiDs2lSVHhkXdt2','세일러문','2004-12-14','010-4444-5555',0),(27,'Ofdd','$2a$10$KIpETLFi2s//s.LzMcuD.ei2PX8p57epUHh/5dL7rFWDs.4mmqkl.','세일러문','2006-01-02','010-4444-5555',0),(28,'123','$2a$10$LfNh4XBcgLNZa.BVdgFJa.2nQtZE9gcR9FxEkKYxOshbyx0x0TDxq','세일러문','2006-01-02','010-4444-5555',0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-04-15  4:07:40
