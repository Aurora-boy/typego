--
-- 数据库: `typego`
--

-- --------------------------------------------------------

--
-- 表的结构 `typego_users`
--

CREATE TABLE `typego_users` (
  `uid` int(10) unsigned NOT NULL auto_increment,
  `username` varchar(32) default NULL,
  `password` varchar(64) default NULL,
  `mail` varchar(150) default NULL,
  `created` int(10) unsigned default '0',
  `logged` int(10) unsigned default '0',
  PRIMARY KEY  (`uid`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `mail` (`mail`)
);
