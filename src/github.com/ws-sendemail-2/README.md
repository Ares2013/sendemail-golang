####创建数据库
```
    CREATE TABLE IF NOT EXISTS `send_email_info`(
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
      `from_email` varchar(1000) NOT NULL DEFAULT '' COMMENT '发邮件的来源',
        `to_email` varchar(1000) NOT NULL DEFAULT '' COMMENT '发邮件的收件',
        `cc_email` varchar(1000) NOT NULL DEFAULT '' COMMENT '发邮件的抄送',
      `subject` text not null	COMMENT '邮件的subject',
      `body` text NOT NULL COMMENT '邮件的Body',
      `is_html` TINYINT(2) UNSIGNED NOT NULL DEFAULT 2 COMMENT '邮件是否是html格式,1是html,2是文字',
        `is_attach` TINYINT(2) UNSIGNED NOT NULL DEFAULT 2 COMMENT '邮件是否含有附件,1是含有附件,2是不含附件',
        `attach` VARCHAR(1000) NOT NULL DEFAULT 2 COMMENT '邮件附件',
      `ctime` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '数据创建时间',
        `etime` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '数据创建时间',
      `extension` text COMMENT '收集用户信息扩展字段,如果不是邮箱，可以添加字段',
      `unique_key` VARCHAR(500) NOT NULL DEFAULT '0' COMMENT '唯一加密key',
        `user_ip` varchar(25) NOT NULL DEFAULT '' COMMENT '用户的ip地址转换成长整形',
        `status` TINYINT(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '发送邮件的状态，1是已经接收并发送,2是重新发送,3是发送失败,4是已取消',
        `weights` TINYINT(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '发送邮件的优先级，1普通通知型邮件,2重要邮件,3特殊邮件,4紧急邮件',
      PRIMARY KEY (`id`),
      KEY `ix_fe_sw` (`from_email`,`status`,`weights`),
      KEY `ix_uk` (`unique_key`),
        key `ix_toe`(`to_email`)
    )ENGINE=Innodb AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='发布邮件记录表';
```

