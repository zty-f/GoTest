-- rdcsaas_sale_order 站内订单表
CREATE TABLE `rdcsaas_sale_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `order_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '站内订单ID',
  `package_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '套餐ID',
  `msg_str` text COLLATE utf8mb4_bin NOT NULL COMMENT '订单详情，可用于搜索',
  `utm_source` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '渠道',
  `pay_type` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '支付方式',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '实际支付金额，单位：分',
  `source` int(11) NOT NULL DEFAULT '0' COMMENT '订单来源',
  `outer_order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '外部订单ID，非内部订单时有效',
  `order_extra` text COLLATE utf8mb4_bin NOT NULL COMMENT '订单附加信息，JSON格式',
  `trade_state` smallint(6) NOT NULL DEFAULT '0' COMMENT '订单状态：100=支付成功，200=全部退款，210=部分退款',
  `user_assign` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户归属顾问',
  `order_assign` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '订单归属顾问',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '扩展信息，JSON格式',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_biz_order` (`biz_type`,`order_id`),
  KEY `idx_biz_user_ct` (`biz_type`,`user_id`,`ct`),
  KEY `idx_biz_source_outer` (`biz_type`,`source`,`outer_order_id`),
  KEY `idx_biz_user_assign` (`biz_type`,`user_assign`),
  KEY `idx_biz_order_assign` (`biz_type`,`order_assign`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='站内订单表';

-- rdcsaas_sale_order_refund 站内订单退款表
CREATE TABLE `rdcsaas_sale_order_refund` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `order_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '站内订单ID',
  `refund_id` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '退款ID',
  `action_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '退款类型：1=退费，2=返现',
  `amount` int(11) NOT NULL DEFAULT '0' COMMENT '退款金额，单位：分',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `package_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '套餐ID',
  `package_name` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '套餐名称',
  `work_order_id` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '工单ID',
  `refund_reason` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '退款原因',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '退款操作人',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '扩展信息，JSON格式',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_biz_order_refund` (`biz_type`,`order_id`,`refund_id`),
  KEY `idx_biz_user_ct` (`biz_type`,`user_id`,`ct`),
  KEY `idx_biz_work_order` (`biz_type`,`work_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='站内订单退款表';

-- rdcsaas_outer_order 站外订单表
CREATE TABLE `rdcsaas_outer_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '站外订单ID',
  `source` int(11) NOT NULL DEFAULT '0' COMMENT '订单来源',
  `trade_state` smallint(6) NOT NULL DEFAULT '0' COMMENT '订单状态：100=支付成功，200=全部退款，210=部分退款',
  `pay_amount` int(11) NOT NULL DEFAULT '0' COMMENT '实际支付金额，单位：分',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `msg_str` text COLLATE utf8mb4_bin NOT NULL COMMENT '平台原始订单消息',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联用户ID',
  `address_info` text COLLATE utf8mb4_bin NOT NULL COMMENT '收货地址信息，JSON格式',
  `user_assign` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户归属顾问',
  `order_assign` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '订单归属顾问',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '分配操作人，系统自动分配时写服务名',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '扩展信息，JSON格式',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_biz_source_order` (`biz_type`,`source`,`order_id`),
  KEY `idx_biz_pay_time` (`biz_type`,`pay_time`),
  KEY `idx_biz_user_ct` (`biz_type`,`user_id`,`ct`),
  KEY `idx_biz_user_assign` (`biz_type`,`user_assign`),
  KEY `idx_biz_order_assign` (`biz_type`,`order_assign`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='站外订单表';

-- rdcsaas_outer_order_refund 站外订单退款表
CREATE TABLE `rdcsaas_outer_order_refund` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '站外订单ID',
  `refund_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '站外退款ID',
  `source` int(11) NOT NULL DEFAULT '0' COMMENT '订单来源',
  `refund_time` datetime DEFAULT NULL COMMENT '退款时间',
  `msg_str` text COLLATE utf8mb4_bin NOT NULL COMMENT '平台原始退款消息',
  `refund_amount` int(11) NOT NULL DEFAULT '0' COMMENT '退款金额，单位：分',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联用户ID',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '扩展信息，JSON格式',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_biz_source_refund` (`biz_type`,`source`,`refund_id`),
  KEY `idx_biz_order` (`biz_type`,`source`,`order_id`),
  KEY `idx_biz_refund_time` (`biz_type`,`refund_time`),
  KEY `idx_biz_user_ct` (`biz_type`,`user_id`,`ct`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='站外订单退款表';

-- rdcsaas_user 用户表
CREATE TABLE `rdcsaas_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户UID',
  `user_level` int(11) NOT NULL DEFAULT '0' COMMENT '用户等级',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1=正常，2=停用',
  `sale_teacher_opname` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '归属顾问账号',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '扩展信息，JSON格式',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_biz_uid` (`biz_type`,`uid`),
  KEY `idx_biz_teacher_ct` (`biz_type`,`sale_teacher_opname`,`ct`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

-- rdcsaas_sale_teacher 顾问表
CREATE TABLE `rdcsaas_sale_teacher` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `teacher_opname` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '顾问账号',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '状态：1=正常，100=离职',
  `polling_switch` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '轮询分配开关：1=开，2=关',
  `capacity` int(11) NOT NULL DEFAULT '0' COMMENT '可服务用户容量',
  `role_tags` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '角色标签',
  `assign_tags` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '分配标签',
  `bind_qw` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '绑定企业微信账号',
  `nickname_en` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '英文昵称',
  `grade_level` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '服务等级，如S/A/B',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '扩展信息，JSON格式',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_biz_opname` (`biz_type`,`teacher_opname`),
  KEY `idx_biz_status_polling` (`biz_type`,`status`,`polling_switch`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='顾问表';

-- rdcsaas_auditlog 审计日志表
CREATE TABLE `rdcsaas_auditlog` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `target_type` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作对象类型，如sale_order、user、course',
  `target_id` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作对象ID',
  `action` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作类型',
  `old_value` text COLLATE utf8mb4_bin NOT NULL COMMENT '变更前内容，JSON格式',
  `new_value` text COLLATE utf8mb4_bin NOT NULL COMMENT '变更后内容，JSON格式',
  `params` text COLLATE utf8mb4_bin NOT NULL COMMENT '操作参数或补充详情，JSON格式',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `opts` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '操作时间，Unix毫秒时间戳',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_biz_target_ct` (`biz_type`,`target_type`,`target_id`,`ct`),
  KEY `idx_biz_action_ct` (`biz_type`,`action`,`ct`),
  KEY `idx_biz_opuser_ct` (`biz_type`,`opuser`,`ct`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='审计日志表';

-- rdcsaas_course 课程表
CREATE TABLE `rdcsaas_course` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '课程ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `course_name` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '课程名称',
  `course_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '课程类型，由业务线自行定义',
  `course_level` int(11) NOT NULL DEFAULT '0' COMMENT '课程等级',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1=草稿，2=启用，3=停用',
  `description` text COLLATE utf8mb4_bin NOT NULL COMMENT '课程描述',
  `cover_uri` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '课程封面URI',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序值，值越小越靠前',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '扩展信息，JSON格式',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_biz_status_sort` (`biz_type`,`status`,`sort`),
  KEY `idx_biz_name` (`biz_type`,`course_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='课程表';

-- rdcsaas_course_user 用户开课记录表
CREATE TABLE `rdcsaas_course_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '开课记录ID',
  `biz_type` tinyint(3) unsigned NOT NULL COMMENT '业务线：1=飞跃英语启蒙，2=畅快英语启蒙',
  `course_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '课程ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `order_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联站内订单ID，手工开课时为0',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态：1=学习中，2=已完成，3=已失效',
  `studied_rate` int(11) NOT NULL DEFAULT '0' COMMENT '学习进度，0-100',
  `open_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开课时间',
  `expire_time` datetime DEFAULT NULL COMMENT '课程到期时间，NULL表示不限制',
  `finish_time` datetime DEFAULT NULL COMMENT '完成时间',
  `remark_info` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '开课备注',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '扩展信息，JSON格式',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_biz_user_course` (`biz_type`,`user_id`,`course_id`),
  KEY `idx_biz_course_status` (`biz_type`,`course_id`,`status`),
  KEY `idx_biz_user_status` (`biz_type`,`user_id`,`status`),
  KEY `idx_biz_order` (`biz_type`,`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户开课记录表';
