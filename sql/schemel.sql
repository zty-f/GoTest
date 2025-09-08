-- 老师可用开课时间配置表
CREATE TABLE rdc_classroom_teacher_availability
(
    id                  BIGINT      NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    teacher_uid         BIGINT      NOT NULL COMMENT '老师ID',
    start_time          TIME        NOT NULL COMMENT '时段开始时间',
    end_time            TIME        NOT NULL COMMENT '时段结束时间',
    strategy_start_time TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '策略开始时间',
    strategy_end_time   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '策略结束时间',
    recurrence_type     TINYINT     NOT NULL COMMENT '循环类型：1-每日循环，2-每周循环',
    week_days           VARCHAR(20) NOT NULL DEFAULT '' COMMENT '适用于每周循环时，存储周几可用，如"1,2,3,4,5,6,0"表示周一至周日',
    status              TINYINT     NOT NULL DEFAULT 1 COMMENT '状态：1-开启 2-关闭',
    ct                  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    ut                  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX idx_tsr (teacher_uid, status, strategy_start_time, strategy_end_time, recurrence_type),
    INDEX idx_sr (status, strategy_start_time, strategy_end_time, recurrence_type),
    INDEX idx_se (start_time, end_time)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='老师可用开课时间配置表';

CREATE TABLE `readcamp_relations`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `uid`         bigint(20) unsigned NOT NULL COMMENT '关注者id',
    `followed_id` bigint(20) unsigned NOT NULL COMMENT '被关注者id',
    `status`      int                 NOT NULL DEFAULT '1' COMMENT '关注状态 1: 关注, 2:取关',
    `ct`          datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间-关注最新',
    `ut`          datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_uf` (`uid`, `followed_id`) USING BTREE,
    KEY `idx_f` (`followed_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户记录表';

CREATE TABLE `readcamp_reports`
(
    `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `uid`           bigint(20) unsigned NOT NULL COMMENT '举报人id',
    `target_id`     bigint(20) unsigned NOT NULL COMMENT '被举报目标id',
    `target_type`   tinyint(4)          NOT NULL COMMENT '举报类型 1:用户, 2:帖子, 3:评论, 4:其他',
    `reason`        varchar(255)        NOT NULL COMMENT '举报原因',
    `evidence`      varchar(500)        NOT NULL DEFAULT '' COMMENT '证据(图片/视频URL)',
    `status`        tinyint(4)          NOT NULL DEFAULT '1' COMMENT '处理状态 1:待处理, 2:处理中, 3:已处理, 4:已驳回',
    `handler_id`    bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '处理人id(管理员)',
    `handle_result` varchar(255)        NOT NULL DEFAULT '' COMMENT '处理结果描述',
    `ct`            datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `ut`            datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_u` (`uid`) USING BTREE COMMENT '举报人索引',
    KEY `idx_t` (`target_id`) USING BTREE COMMENT '被举报目标索引',
    KEY `idx_s` (`status`) USING BTREE COMMENT '处理状态索引'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户举报记录表';

CREATE TABLE `readcamp_user_privacy`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `uid`          bigint(20) unsigned NOT NULL COMMENT '用户id',
    `privacy_type` tinyint(4)          NOT NULL COMMENT '隐私配置类型 1-个人主页相关',
    `privacy_id`   int unsigned        NOT NULL COMMENT '隐私配置id',
    `visible`      int                 NOT NULL COMMENT '当前隐私配置是否可见 1可见 2不可见',
    `ct`           datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `ut`           datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_upp` (`uid`, `privacy_type`, `privacy_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户隐私设置记录表';

CREATE TABLE `readcamp_chat_messages`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `from_account` varchar(128)        NOT NULL COMMENT '消息发送者',
    `to_account`   varchar(128)        NOT NULL COMMENT '消息接收者',
    `msg`          text                NOT NULL COMMENT '用户发的消息',
    `reply`        text COMMENT '对消息的回复',
    `tread_time`   int    NOT NULL DEFAULT 0 COMMENT '消息踩时间戳',
    `copy_time`    int    NOT NULL DEFAULT 0 COMMENT '消息复制时间戳',
    `share_time`   int    NOT NULL DEFAULT 0 COMMENT '消息转发时间戳',
    `ct`           datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `ut`           datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_ftc` (`from_account`, `to_account`, `ct`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户聊天消息表';