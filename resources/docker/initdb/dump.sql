
CREATE DATABASE directory;
USE directory;
DROP TABLE IF EXISTS `person`;
CREATE TABLE `person` (
  `customer_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `aadhar_id` int(64) NOT NULL,  
  PRIMARY KEY (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2009 DEFAULT CHARSET=latin1; 

LOCK TABLES `person` WRITE;

INSERT INTO `person` VALUES
 	      (2000,'Steve',27612),
	      (2001,'Arian', 456561),
	      (2002,'Hadley', 846486),
	      (2003,'Ben', 513441),
      	(2004,'Nina', 475165),
	      (2005,'Osman',346841 ),
        (2006,'Amith',355441),
        (2007,'Rob',6541541),
        (2008,'Matt',89212441);


UNLOCK TABLES;
