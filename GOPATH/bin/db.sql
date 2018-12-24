create schema lightpan;
use lightpan;
CREATE TABLE `groupuser` (
  `id` varchar(45) NOT NULL,
  `user` varchar(45) NOT NULL,
  `jointime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `groupuser` (`id`,`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
  `id` varchar(45) NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT '用户',
  `password` varchar(256) NOT NULL DEFAULT '' COMMENT ' ',
  `face` varchar(256) DEFAULT NULL,
  `registetime` datetime NOT NULL,
  `type` varchar(45) NOT NULL DEFAULT '',
  `parent` varchar(45) NOT NULL DEFAULT '' COMMENT 'if type=group,parent = group owner',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `file` (
  `user` varchar(45) NOT NULL,
  `folder` varchar(200) DEFAULT NULL,
  `name` varchar(260) NOT NULL,
  `pub` tinyint(1) NOT NULL DEFAULT '1',
  `del` tinyint(1) NOT NULL DEFAULT '1',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `version` int(11) NOT NULL DEFAULT '0',
  `object4d` varchar(260) NOT NULL,
  UNIQUE KEY `furl` (`user`,`folder`,`name`,`del`,`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


