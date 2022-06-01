
CREATE DATABASE directory;
USE directory;
DROP TABLE IF EXISTS `person`;
CREATE TABLE `person` (
  
  `name` varchar(100) NOT NULL,
  `aadhar_id` int(64) NOT NULL,  
  PRIMARY KEY (`aadhar_id`)
); 

LOCK TABLES `person` WRITE;

INSERT INTO `person` VALUES 
    ("Amith",12535441),
    ("Amith",1253545541),
    ("Amith",124555341);
	


UNLOCK TABLES;
