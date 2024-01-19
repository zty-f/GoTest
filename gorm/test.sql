CREATE TABLE `user_1`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `stu_id`       bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `value`        int(11) NOT NULL DEFAULT '0' COMMENT '成长值，可正可负',
    `source`       int(11) NOT NULL DEFAULT '-1' COMMENT '成长值更新来源， 10:初始化, 20: 任务 30:激活，',
    `task_id`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '该成长值记录对应的任务id',
    `user_task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '关联user_task表，对应用户任务记录id\n',
    `value_time`   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '获得成长值的时间，要和create_time区分',
    `create_time`  int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `ext_data`     text NOT NULL COMMENT '扩展字段，json格式',
    PRIMARY KEY (`id`) USING BTREE,
    KEY            `idx_uid_src` (`stu_id`,`source`),
    KEY            `idx_uid_vtime` (`stu_id`,`value_time`),
    KEY            `idx_utaskid` (`user_task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成长值记录表';
CREATE TABLE `user_2`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `stu_id`       bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `value`        int(11) NOT NULL DEFAULT '0' COMMENT '成长值，可正可负',
    `source`       int(11) NOT NULL DEFAULT '-1' COMMENT '成长值更新来源， 10:初始化, 20: 任务 30:激活，',
    `task_id`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '该成长值记录对应的任务id',
    `user_task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '关联user_task表，对应用户任务记录id\n',
    `value_time`   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '获得成长值的时间，要和create_time区分',
    `create_time`  int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `ext_data`     text NOT NULL COMMENT '扩展字段，json格式',
    PRIMARY KEY (`id`) USING BTREE,
    KEY            `idx_uid_src` (`stu_id`,`source`),
    KEY            `idx_uid_vtime` (`stu_id`,`value_time`),
    KEY            `idx_utaskid` (`user_task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成长值记录表';
CREATE TABLE `user_3`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `stu_id`       bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `value`        int(11) NOT NULL DEFAULT '0' COMMENT '成长值，可正可负',
    `source`       int(11) NOT NULL DEFAULT '-1' COMMENT '成长值更新来源， 10:初始化, 20: 任务 30:激活，',
    `task_id`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '该成长值记录对应的任务id',
    `user_task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '关联user_task表，对应用户任务记录id\n',
    `value_time`   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '获得成长值的时间，要和create_time区分',
    `create_time`  int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `ext_data`     text NOT NULL COMMENT '扩展字段，json格式',
    PRIMARY KEY (`id`) USING BTREE,
    KEY            `idx_uid_src` (`stu_id`,`source`),
    KEY            `idx_uid_vtime` (`stu_id`,`value_time`),
    KEY            `idx_utaskid` (`user_task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成长值记录表';
CREATE TABLE `user_4`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `stu_id`       bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `value`        int(11) NOT NULL DEFAULT '0' COMMENT '成长值，可正可负',
    `source`       int(11) NOT NULL DEFAULT '-1' COMMENT '成长值更新来源， 10:初始化, 20: 任务 30:激活，',
    `task_id`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '该成长值记录对应的任务id',
    `user_task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '关联user_task表，对应用户任务记录id\n',
    `value_time`   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '获得成长值的时间，要和create_time区分',
    `create_time`  int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `ext_data`     text NOT NULL COMMENT '扩展字段，json格式',
    PRIMARY KEY (`id`) USING BTREE,
    KEY            `idx_uid_src` (`stu_id`,`source`),
    KEY            `idx_uid_vtime` (`stu_id`,`value_time`),
    KEY            `idx_utaskid` (`user_task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成长值记录表';
CREATE TABLE `user_5`
(
    `id`           int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `stu_id`       bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `value`        int(11) NOT NULL DEFAULT '0' COMMENT '成长值，可正可负',
    `source`       int(11) NOT NULL DEFAULT '-1' COMMENT '成长值更新来源， 10:初始化, 20: 任务 30:激活，',
    `task_id`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '该成长值记录对应的任务id',
    `user_task_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '关联user_task表，对应用户任务记录id\n',
    `value_time`   int(11) unsigned NOT NULL DEFAULT '0' COMMENT '获得成长值的时间，要和create_time区分',
    `create_time`  int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `ext_data`     text NOT NULL COMMENT '扩展字段，json格式',
    PRIMARY KEY (`id`) USING BTREE,
    KEY            `idx_uid_src` (`stu_id`,`source`),
    KEY            `idx_uid_vtime` (`stu_id`,`value_time`),
    KEY            `idx_utaskid` (`user_task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成长值记录表';
