/*
 Navicat Premium Dump SQL

 Source Server         : 测试环境
 Source Server Type    : MySQL
 Source Server Version : 50728 (5.7.28-log)
 Source Host           : 10.111.131.208:4000
 Source Schema         : rdc_classroom

 Target Server Type    : MySQL
 Target Server Version : 50728 (5.7.28-log)
 File Encoding         : 65001

 Date: 03/08/2026 17:43:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for rdc_classroom_activity_invitation_user
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_activity_invitation_user`;
CREATE TABLE `rdc_classroom_activity_invitation_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `activity_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '活动id',
  `poster_ct` bigint(20) NOT NULL DEFAULT '0' COMMENT '海报创建时间',
  `screenshot_up_ts` bigint(20) NOT NULL DEFAULT '0' COMMENT '截图上传时间',
  `screenshot_uri` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '截图uri',
  `task_status` int(11) NOT NULL DEFAULT '0' COMMENT '状态',
  `opuser` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `review_info` text COLLATE utf8mb4_bin NOT NULL COMMENT '审核信息',
  `review_ts` bigint(20) NOT NULL DEFAULT '0' COMMENT '审核时间',
  `uploader` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '上传人',
  `reviewer` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '审核人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ua` (`uid`,`activity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='转介绍活动用户表';

-- ----------------------------
-- Table structure for rdc_classroom_announcement
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_announcement`;
CREATE TABLE `rdc_classroom_announcement` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `module_type` int(11) NOT NULL DEFAULT '0' COMMENT '模块类型 1=公告 2=帮助 3=培训',
  `category` int(11) NOT NULL DEFAULT '0' COMMENT '分类 公告=100运营活动 101平台功能 102其他；帮助=200平台政策相关 201平台功能相关 202服务流程相关；培训=300启蒙教学 301进阶教学 302通用教学技巧 303通用教学活动',
  `state` int(11) NOT NULL DEFAULT '0' COMMENT '状态 1=草稿 2=定时发布 3=已发布 4=已撤回',
  `title` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '标题',
  `content` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '消息详情',
  `sort` bigint(20) NOT NULL DEFAULT '1000' COMMENT '排序，置顶=1',
  `opuser` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `published_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '发布时间/定时发布时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_mcs` (`module_type`,`category`,`state`),
  KEY `idx_csp` (`category`,`sort`,`published_time`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='公告';

-- ----------------------------
-- Table structure for rdc_classroom_audit_log
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_audit_log`;
CREATE TABLE `rdc_classroom_audit_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `key_id` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `action` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '事件类型',
  `params` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '详情',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `opts` bigint(20) NOT NULL DEFAULT '0' COMMENT '操作时间；毫秒',
  PRIMARY KEY (`id`),
  KEY `idx_kid` (`key_id`,`ct`),
  KEY `idx_ak` (`action`,`key_id`,`ct`),
  KEY `idx_ao` (`action`,`opts`)
) ENGINE=InnoDB AUTO_INCREMENT=7957 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='操作审记日志';

-- ----------------------------
-- Table structure for rdc_classroom_bigdata_records
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_bigdata_records`;
CREATE TABLE `rdc_classroom_bigdata_records` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `data_type` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '数据类型： workbench-工作台数据',
  `data_unit` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '数据维度： 日20260305、月:202603',
  `data_key` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '数据标识： 老师邮箱、用户id',
  `data` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '格式：{}',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ddd` (`data_type`,`data_unit`,`data_key`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='大数据推送存储数据';

-- ----------------------------
-- Table structure for rdc_classroom_class
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_class`;
CREATE TABLE `rdc_classroom_class` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `class_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '班级类型 1=小班课 2=大班课 3=临时班',
  `class_name` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `teacher_opname` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '班主任',
  `capacity` int(11) NOT NULL DEFAULT '0' COMMENT '班级容量',
  `course_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '课程id',
  `class_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '班级状态 1=正常 2=已结束',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `description` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '备注',
  `used_capacity` int(11) NOT NULL DEFAULT '0' COMMENT '已用容量（包括插班）',
  `extend` json NOT NULL COMMENT '拓展信息',
  `bind_wechat_group` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '绑定企微群id',
  `class_tags` varchar(265) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '班级标签',
  `main_teacher_uids` json DEFAULT NULL COMMENT '主讲老师UID,JSON数组',
  `auto_schedule` text COLLATE utf8mb4_bin COMMENT '自动排课conf',
  `main_teacher_tags` json NOT NULL COMMENT '主讲老师tags',
  PRIMARY KEY (`id`),
  KEY `idx_ct` (`class_type`,`id`),
  KEY `idx_tc` (`id`,`class_type`)
) ENGINE=InnoDB AUTO_INCREMENT=144 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='班级管理';

-- ----------------------------
-- Table structure for rdc_classroom_class_hour_rights_flow
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_class_hour_rights_flow`;
CREATE TABLE `rdc_classroom_class_hour_rights_flow` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `user_id` bigint(20) DEFAULT '0' COMMENT '用户ID',
  `event_type` int(11) NOT NULL DEFAULT '0' COMMENT '事件类型：1-下单发放, 2-取消排课退回, 3-到课消耗, 4-订单退款',
  `event_key` varchar(128) NOT NULL DEFAULT '' COMMENT '事件唯一Key，用于幂等（如订单号、到课记录ID等）',
  `change_value` bigint(20) NOT NULL DEFAULT '0' COMMENT '课时数变化值',
  `gift_change_value` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送课时余额变化值',
  `operator` varchar(64) NOT NULL DEFAULT '' COMMENT '操作人',
  `extra` text COMMENT '备注信息',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_event` (`user_id`,`event_type`,`event_key`) COMMENT '用户流水唯一'
) ENGINE=InnoDB AUTO_INCREMENT=5514 DEFAULT CHARSET=utf8mb4 COMMENT='用户课时权益流水表';

-- ----------------------------
-- Table structure for rdc_classroom_class_schedule
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_class_schedule`;
CREATE TABLE `rdc_classroom_class_schedule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id,排课id,房间id',
  `class_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '班级id',
  `course_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '课程ID',
  `lecture_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '课节id',
  `teacher_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '老师uid',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '排课状态 100=初始化 200=已结束 300=已取消',
  `schedule_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1=小班课 2=直播课',
  `schedule_source` tinyint(4) NOT NULL DEFAULT '1' COMMENT '排课方式 1：手动 2 : 自动',
  `start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始上课时间',
  `end_time` timestamp NULL DEFAULT NULL COMMENT '结束上课时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `cost_hour_amount` int(11) NOT NULL DEFAULT '0' COMMENT '消耗课时数 1课时=100',
  `room_id` bigint(20) DEFAULT NULL COMMENT '房间id',
  PRIMARY KEY (`id`),
  KEY `index_room_id` (`room_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1158 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='班级排课记录';

-- ----------------------------
-- Table structure for rdc_classroom_cmd_history
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_cmd_history`;
CREATE TABLE `rdc_classroom_cmd_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `schedule_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '排课id，房间id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `seq_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '排序id',
  `op_type` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作类型',
  `opts` bigint(20) NOT NULL DEFAULT '0' COMMENT '操作时间,毫秒',
  `op_data` mediumtext COLLATE utf8mb4_bin NOT NULL COMMENT '操作数据',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_ss` (`schedule_id`,`seq_id`),
  KEY `idx_su` (`schedule_id`,`uid`),
  KEY `idx_sot` (`schedule_id`,`op_type`)
) ENGINE=InnoDB AUTO_INCREMENT=1964755 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='房间命令执行历史';

-- ----------------------------
-- Table structure for rdc_classroom_crm_user_group
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_crm_user_group`;
CREATE TABLE `rdc_classroom_crm_user_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id，无意义',
  `sale_teacher` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '助教',
  `group_name` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '人群组名称',
  `cond` text COLLATE utf8mb4_bin NOT NULL COMMENT '筛选条件',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='crm自定义人群组';

-- ----------------------------
-- Table structure for rdc_classroom_crmuser_remark_his
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_crmuser_remark_his`;
CREATE TABLE `rdc_classroom_crmuser_remark_his` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `biz_type` bigint(20) NOT NULL DEFAULT '2' COMMENT '业务线',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `remark_type` int(10) NOT NULL DEFAULT '0' COMMENT '备注类型 1 = 定级,2 = 可上时段,3= 分班,4=排课,5=发货,6=其他',
  `content` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '备注内容',
  `opuser` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `follow_ts` bigint(20) NOT NULL DEFAULT '0' COMMENT '跟进时间',
  `finished_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '完成状态 1=已完成 2=未完成',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=263 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='crm用户备注历史记录';

-- ----------------------------
-- Table structure for rdc_classroom_feedback
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_feedback`;
CREATE TABLE `rdc_classroom_feedback` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `submitter_type` tinyint(4) NOT NULL COMMENT '提交人类型 (1:学生 2:上课老师,3:教务老师)',
  `submitter` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '提交人ID',
  `target_type` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '点评对象类型',
  `target_id` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '点评对象ID ',
  `content` longtext COLLATE utf8mb4_bin COMMENT '评价内容 (JSON格式)',
  `is_anonymous` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否匿名评价',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_sub` (`target_id`,`target_type`,`submitter_type`)
) ENGINE=InnoDB AUTO_INCREMENT=252 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='点评反馈表';

-- ----------------------------
-- Table structure for rdc_classroom_intention_time_metrics
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_intention_time_metrics`;
CREATE TABLE `rdc_classroom_intention_time_metrics` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_level` int(11) NOT NULL DEFAULT '0' COMMENT '用户等级',
  `user_num` int(11) NOT NULL DEFAULT '0' COMMENT '用户人数',
  `intention_time_group` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '可上时段',
  `push_hour_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '版本',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户可上时段统计';

-- ----------------------------
-- Table structure for rdc_classroom_invitation
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_invitation`;
CREATE TABLE `rdc_classroom_invitation` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` bigint(20) NOT NULL DEFAULT '2' COMMENT '业务线',
  `inviter_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '邀请用户uid',
  `invitee_uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '被邀请用户uid',
  `bind_ts` bigint(20) NOT NULL DEFAULT '0' COMMENT '邀请关系绑定时间，毫秒',
  `bind_status` int(11) NOT NULL DEFAULT '0' COMMENT '绑定关系状态 1=正常',
  `utm_source_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '渠道id',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '拓展信息',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_iiuid` (`inviter_uid`,`invitee_uid`),
  UNIQUE KEY `uniq_buu` (`inviter_uid`,`invitee_uid`,`biz_type`),
  KEY `idx_iu` (`invitee_uid`,`ct`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='转介绍邀请表';

-- ----------------------------
-- Table structure for rdc_classroom_order_class_hour_rights
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_order_class_hour_rights`;
CREATE TABLE `rdc_classroom_order_class_hour_rights` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `order_id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单ID',
  `source` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '订单来源',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户',
  `balance` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '课时余额，0.01 课时',
  `total_issued` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '累计发放课时数',
  `total_consumed` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '累计实际消耗课时数',
  `gift_total_consumed` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送课时总消耗',
  `gift_total_issued` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送课时总发放',
  `gift_balance` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送课时余额',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `biz_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '业务标识：1=小班课，2=菲教权益',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_order_source_biz` (`order_id`,`source`,`biz_type`)
) ENGINE=InnoDB AUTO_INCREMENT=1292 DEFAULT CHARSET=utf8mb4 COMMENT='用户订单课时权益表';

-- ----------------------------
-- Table structure for rdc_classroom_order_refund
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_order_refund`;
CREATE TABLE `rdc_classroom_order_refund` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `order_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单id',
  `action_type` int(11) NOT NULL DEFAULT '0' COMMENT '1:退费 2:返现',
  `amount` int(11) NOT NULL DEFAULT '0' COMMENT '退费金额,单位（分）',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `package_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '套餐id',
  `package_name` varchar(128) NOT NULL DEFAULT '' COMMENT '套餐名称',
  `work_order_id` varchar(128) NOT NULL DEFAULT '' COMMENT '工单ID',
  `opuser` varchar(32) NOT NULL DEFAULT '' COMMENT '退款操作人',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
  `refund_id` varchar(64) NOT NULL DEFAULT '' COMMENT '退款id',
  PRIMARY KEY (`id`),
  KEY `idx_o` (`order_id`),
  KEY `idx_u` (`user_id`),
  KEY `idx_woi` (`work_order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8mb4 COMMENT='订单退费表';

-- ----------------------------
-- Table structure for rdc_classroom_outer_order
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_outer_order`;
CREATE TABLE `rdc_classroom_outer_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `order_id` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '订单id',
  `source` int(11) NOT NULL DEFAULT '0' COMMENT '5:抖店',
  `trade_state` int(11) NOT NULL COMMENT '订单状态，如 100-支付成功、200-退款',
  `pay_time` datetime NOT NULL COMMENT '支付时间',
  `msg_str` text COLLATE utf8mb4_bin NOT NULL COMMENT '通用订单消息',
  `teacher` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户&订单归属 order:xx,user:xx',
  `opuser` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '谁分配的老师，readcamp代表自动分配',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '物流手机号关联的user_id',
  `address_info` text COLLATE utf8mb4_bin NOT NULL COMMENT '收货地址信息 json',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间戳',
  `user_assign` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户归属',
  `order_assign` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '订单归属',
  `business` int(10) NOT NULL DEFAULT '0' COMMENT '业务线',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_os` (`order_id`,`source`),
  KEY `idx_pt` (`pay_time`),
  KEY `idx_uid` (`user_id`),
  KEY `idx_ua` (`user_assign`),
  KEY `idx_oa` (`order_assign`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='外部订单表';

-- ----------------------------
-- Table structure for rdc_classroom_outer_order_refund
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_outer_order_refund`;
CREATE TABLE `rdc_classroom_outer_order_refund` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `order_id` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '订单id',
  `refund_id` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '退款id',
  `source` int(11) NOT NULL DEFAULT '0' COMMENT '5:抖店',
  `refund_time` datetime NOT NULL COMMENT '退款时间',
  `msg_str` text COLLATE utf8mb4_bin COMMENT '订单退款消息',
  `pay_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '退款金额，单位：分',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户Id',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间戳',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_osf` (`order_id`,`source`,`refund_id`),
  KEY `idx_pt` (`refund_time`),
  KEY `idx_uid` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17020 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='外部订单退费记录表';

-- ----------------------------
-- Table structure for rdc_classroom_payment_batch
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_payment_batch`;
CREATE TABLE `rdc_classroom_payment_batch` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `batch` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '打款批次',
  `business` int(10) NOT NULL DEFAULT '0' COMMENT '业务线 1=阅读营 2=小班课',
  `real_payamount` int(10) NOT NULL DEFAULT '0' COMMENT '打款总额，人民币分(未扣除手续费)',
  `coursehour_amount` int(10) NOT NULL DEFAULT '0' COMMENT '基础课时费，人民币分',
  `bonus_amount` int(10) NOT NULL DEFAULT '0' COMMENT '奖金，人民币分',
  `other_amount` int(10) NOT NULL DEFAULT '0' COMMENT '其他,人民币分,可能负数',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ubb` (`uid`,`batch`,`business`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='结算-打款批次';

-- ----------------------------
-- Table structure for rdc_classroom_payment_history
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_payment_history`;
CREATE TABLE `rdc_classroom_payment_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `batch` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '批次',
  `business` int(10) NOT NULL DEFAULT '0' COMMENT '业务线 1=阅读营 2=小班课',
  `real_payamount` int(10) NOT NULL DEFAULT '0' COMMENT '打款总金额，人民币分，未扣除手续费',
  `handling_fee` int(10) NOT NULL DEFAULT '0' COMMENT '手续费，人民币分',
  `pay_currency` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '打款币种:CNY,USD',
  `exchange_rate` decimal(5,4) NOT NULL DEFAULT '0.0000' COMMENT '汇率',
  `pay_status` int(10) NOT NULL DEFAULT '0' COMMENT '打款状态 10-待打款，20-打款中，30-打款失败，40-打款成功，50-已驳回，60-已撤销',
  `pay_channel` int(10) NOT NULL DEFAULT '0' COMMENT '打款渠道 1=云账户 2=payoneer 3=空中云汇',
  `remark` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '备注信息',
  `fail_reason` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '失败原因',
  `payment_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '打款id',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_ub` (`uid`,`business`),
  KEY `idx_cb` (`ct`,`batch`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='结算-打款历史记录';

-- ----------------------------
-- Table structure for rdc_classroom_payment_record
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_payment_record`;
CREATE TABLE `rdc_classroom_payment_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `batch` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '批次',
  `business` int(10) NOT NULL DEFAULT '0' COMMENT '业务线 1=阅读营 2=小班课',
  `real_payamount` int(10) NOT NULL DEFAULT '0' COMMENT '打款总金额，人民币分，未扣除手续费',
  `handling_fee` int(10) NOT NULL DEFAULT '0' COMMENT '手续费，人民币分',
  `pay_currency` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '打款币种:CNY,USD',
  `exchange_rate` decimal(5,4) NOT NULL DEFAULT '0.0000' COMMENT '汇率',
  `pay_status` int(10) NOT NULL DEFAULT '0' COMMENT '打款状态 10-待打款，20-打款中，30-打款失败，40-打款成功，50-已驳回，60-已撤销',
  `pay_channel` int(10) NOT NULL DEFAULT '0' COMMENT '打款渠道 1=云账户 2=payoneer 3=空中云汇',
  `remark` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '备注信息',
  `fail_reason` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '失败原因',
  `payment_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '打款id',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ubb` (`uid`,`batch`,`business`),
  KEY `idx_cb` (`ct`,`batch`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='结算-打款记录';

-- ----------------------------
-- Table structure for rdc_classroom_pet
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_pet`;
CREATE TABLE `rdc_classroom_pet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `biz_type` json NOT NULL COMMENT '多业务线数组，[1,2,3]',
  `nickname` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '宠物名',
  `description` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '说明',
  `conf` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '配置信息',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '顺序',
  `opuser` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_sc` (`sort`,`ct`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='宠物';

-- ----------------------------
-- Table structure for rdc_classroom_ph_auto_gift_class_hour
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_ph_auto_gift_class_hour`;
CREATE TABLE `rdc_classroom_ph_auto_gift_class_hour` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `idempotency_key` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '幂等键：uid_process_level',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `process_num` int(10) NOT NULL DEFAULT '0' COMMENT '学习进度',
  `level` int(10) NOT NULL DEFAULT '0' COMMENT '用户等级',
  `order_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送课时关联订单ID',
  `message_id` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '首次处理的消息ID，用作课时流水Rid',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1处理中，2成功，3失败',
  `retry_count` int(10) NOT NULL DEFAULT '0' COMMENT '重试次数',
  `lease_token` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '当前处理租约令牌',
  `last_error` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '最近一次失败原因',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `source` int(10) DEFAULT '0' COMMENT '订单来源',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_idempotency_key` (`idempotency_key`),
  KEY `idx_status_ut` (`status`,`ut`),
  KEY `idx_uid_ct` (`uid`,`ct`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='赠送菲教课时幂等记录';

-- ----------------------------
-- Table structure for rdc_classroom_props_mall
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_props_mall`;
CREATE TABLE `rdc_classroom_props_mall` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `biz_type` json NOT NULL COMMENT '多业务线[1,2,3]',
  `name` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '商品名',
  `rarity_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '稀有度 1=普通',
  `conf` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '配置信息',
  `star_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '星星价格',
  `intimacy_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '亲密度',
  `category` int(11) NOT NULL DEFAULT '0' COMMENT '分类',
  `sub_category` int(11) NOT NULL DEFAULT '0' COMMENT '子分类',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '顺序',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态：1-上架 2-下架',
  `description` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '描述',
  `mount_location` int(11) NOT NULL DEFAULT '0' COMMENT '挂载位置',
  `opuser` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_cs` (`category`,`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=123 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='道具商城';

-- ----------------------------
-- Table structure for rdc_classroom_rate
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_rate`;
CREATE TABLE `rdc_classroom_rate` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `rate` decimal(5,4) NOT NULL DEFAULT '0.0000' COMMENT '汇率',
  `rate_date` int(10) NOT NULL DEFAULT '0' COMMENT '日期 20050102',
  `currency` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '币种 USD',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_rc` (`rate_date`,`currency`)
) ENGINE=InnoDB AUTO_INCREMENT=211 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='结算-打款汇率';

-- ----------------------------
-- Table structure for rdc_classroom_recording_file
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_recording_file`;
CREATE TABLE `rdc_classroom_recording_file` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房间id，排课id',
  `sid` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '声网录制session_id',
  `filepath` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'oss 原始路径',
  `filetype` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '文件类型 audio,video,audio_and_video',
  `convert_filepath` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'oss 转码文件路径',
  `app_type` int(10) NOT NULL COMMENT '2=声网 3=即构',
  `rtcuid` bigint(20) NOT NULL COMMENT '加入频道的uid',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `slice_start_ts` bigint(20) NOT NULL DEFAULT '0' COMMENT '文件开始录制时间',
  `duration` bigint(20) NOT NULL DEFAULT '0' COMMENT '视频时长-毫秒',
  PRIMARY KEY (`id`),
  KEY `idx_rs` (`room_id`,`sid`),
  KEY `idx_fp` (`filepath`)
) ENGINE=InnoDB AUTO_INCREMENT=181 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='录制文件信息';

-- ----------------------------
-- Table structure for rdc_classroom_recording_history
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_recording_history`;
CREATE TABLE `rdc_classroom_recording_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房间id，排课id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '录制状态',
  `app_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '三方类型 2=声网 3=即构',
  `resource_id` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '录制id',
  `record_mode` tinyint(4) NOT NULL DEFAULT '0' COMMENT '录制模式 1=单流 2=合流',
  `sid` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '声网录制session_id',
  `rtcuid` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '录制用户uid',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ra` (`resource_id`,`app_type`),
  KEY `idx_rid` (`room_id`)
) ENGINE=InnoDB AUTO_INCREMENT=546 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='单次录制信息';

-- ----------------------------
-- Table structure for rdc_classroom_recording_log
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_recording_log`;
CREATE TABLE `rdc_classroom_recording_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房间id，排课id',
  `sid` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '声网录制session_id',
  `origin_msg` text COLLATE utf8mb4_bin NOT NULL COMMENT '原始消息',
  `event_type` int(10) NOT NULL COMMENT '消息类型',
  `app_type` int(10) NOT NULL COMMENT '2=声网 3=即构',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_rs` (`room_id`,`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=133 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='录制日志';

-- ----------------------------
-- Table structure for rdc_classroom_replication
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_replication`;
CREATE TABLE `rdc_classroom_replication` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房间id，排课id',
  `cmd_seq_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'cmd 排序id',
  `sync_code` mediumtext COLLATE utf8mb4_bin NOT NULL COMMENT '镜像数据',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_rid` (`room_id`)
) ENGINE=InnoDB AUTO_INCREMENT=203036 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='房间镜像副本';

-- ----------------------------
-- Table structure for rdc_classroom_reward_log
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_reward_log`;
CREATE TABLE `rdc_classroom_reward_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `reward_type` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '奖励类型：rdc_classroom#star',
  `reward_num` int(10) NOT NULL DEFAULT '0' COMMENT '奖励数量',
  `scene` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '业务#场景：rdc_classroom#study_finished',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `opid` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作id，用于处理奖励发放限制',
  `opts` bigint(20) NOT NULL DEFAULT '0' COMMENT '操作时间：毫秒',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_uoo` (`uid`,`opts`,`opid`),
  KEY `idx_sr` (`scene`,`reward_type`)
) ENGINE=InnoDB AUTO_INCREMENT=766 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='奖励发放日志';

-- ----------------------------
-- Table structure for rdc_classroom_room
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_room`;
CREATE TABLE `rdc_classroom_room` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '房间id',
  `start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '开始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `room_type` tinyint(4) NOT NULL DEFAULT '2' COMMENT '房间类型：1=一对一类型,2=一对多类型',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '房间状态：100=初始化 110=上课中 200=已结束 300=已取消',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `opuser` varchar(64) NOT NULL COMMENT '修改人',
  PRIMARY KEY (`id`),
  KEY `idx_start_time` (`start_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1158 DEFAULT CHARSET=utf8mb4 COMMENT='教室房间表';

-- ----------------------------
-- Table structure for rdc_classroom_room_chatmsg
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_room_chatmsg`;
CREATE TABLE `rdc_classroom_room_chatmsg` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房间id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `user_role` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户角色：user,teacher,admin',
  `content` text COLLATE utf8mb4_bin NOT NULL COMMENT '消息内容',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '消息状态 1=正常 2=撤回',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `opts` bigint(20) NOT NULL DEFAULT '0' COMMENT '消息发送时间：毫秒',
  `seq_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送顺序',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_rs` (`room_id`,`seq_id`)
) ENGINE=InnoDB AUTO_INCREMENT=477 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='房间消息表';

-- ----------------------------
-- Table structure for rdc_classroom_room_linshi
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_room_linshi`;
CREATE TABLE `rdc_classroom_room_linshi` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '房间id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '房间状态：100=初始化 110=上课中 200=已结束 300=已取消',
  `start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '开始时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '修改人',
  `room_type` tinyint(4) NOT NULL DEFAULT '2' COMMENT '房间类型：1=一对一类型,2=一对多类型',
  PRIMARY KEY (`id`),
  KEY `idx_start_time` (`start_time`)
) ENGINE=InnoDB AUTO_INCREMENT=533 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='班级排课记录';

-- ----------------------------
-- Table structure for rdc_classroom_room_metrics
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_room_metrics`;
CREATE TABLE `rdc_classroom_room_metrics` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房间id',
  `is_replace_tea` int(10) NOT NULL DEFAULT '0' COMMENT '是否换老师 1=是 0=否',
  `real_join_time` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '实际上课时间',
  `real_leave_time` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '实际下课时间',
  `level1_status` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '老师状态:老师缺席',
  `duration_desc` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '时长:迟到时长<=5分钟',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `evaluation_time` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '点评时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_rid` (`room_id`),
  KEY `idx_rid` (`room_id`,`ct`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='房间指标';

-- ----------------------------
-- Table structure for rdc_classroom_room_module
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_room_module`;
CREATE TABLE `rdc_classroom_room_module` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房间id',
  `module_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '模块实例id',
  `lecture_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '课节id',
  `scene_type` int(10) NOT NULL DEFAULT '0' COMMENT '场景类型',
  `scene_order` bigint(20) DEFAULT '0' COMMENT '场景在阶段内的顺序',
  `module_type` int(10) NOT NULL DEFAULT '0' COMMENT '模块类型',
  `module_template_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '模块模板id',
  `module_name` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '模块名',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '顺序',
  `conf` longtext COLLATE utf8mb4_bin COMMENT '模块配置',
  `study_status` int(10) NOT NULL DEFAULT '0' COMMENT '完成状态，1=已完成，2=部分完成 3=未完成',
  `locked_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '锁定状态，1=已解锁 2=未解锁',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 1=已删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mid` (`module_id`),
  KEY `idx_urs` (`uid`,`room_id`,`sort`),
  KEY `idx_rlu` (`room_id`,`lecture_id`,`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=36202 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户房间模块';

-- ----------------------------
-- Table structure for rdc_classroom_room_module_task
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_room_module_task`;
CREATE TABLE `rdc_classroom_room_module_task` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '房间id',
  `task_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '任务实例id',
  `lecture_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '课节id',
  `module_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '模块实例id',
  `task_template_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '任务模板id',
  `task_type` int(10) NOT NULL DEFAULT '0' COMMENT '任务类型',
  `task_name` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '任务名',
  `conf` longtext COLLATE utf8mb4_bin COMMENT '任务的配置信息',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '顺序',
  `study_status` int(10) NOT NULL DEFAULT '0' COMMENT '任务状态：1=完成，2=部分完成，3=未完成',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 1=已删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_tid` (`task_id`),
  KEY `idx_uls` (`uid`,`room_id`,`sort`),
  KEY `idx_rlu` (`room_id`,`lecture_id`,`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=50415 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户房间模块任务';

-- ----------------------------
-- Table structure for rdc_classroom_sale_order
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_sale_order`;
CREATE TABLE `rdc_classroom_sale_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user_id',
  `order_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单id',
  `package_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'package id',
  `msg_str` text COLLATE utf8mb4_bin NOT NULL COMMENT 'sale_order_detail详细信息，可用于搜索',
  `utm_source` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '渠道',
  `pay_type` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '支付方式',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '实际支付金额，单位:分',
  `source` int(11) NOT NULL DEFAULT '1' COMMENT '订单来源, 1:内部, 2:有赞 3:视频号 5:抖店 6:小红书 7:淘宝',
  `outer_order_id` varchar(32) COLLATE utf8mb4_bin DEFAULT '' COMMENT '外部订单ID，source不为0时有效',
  `order_extra` text COLLATE utf8mb4_bin NOT NULL COMMENT '订单附加信息(json)，如是否回传广告平台',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间戳',
  `teacher` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户&订单归属 order:xx,user:xx',
  `trade_state` smallint(6) NOT NULL DEFAULT '0' COMMENT '订单状态，100-支付成功 200-全部退 210-部分退',
  `user_assign` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '用户归属',
  `order_assign` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '订单归属',
  `business` int(10) NOT NULL DEFAULT '0' COMMENT '业务线',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '拓展信息',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_oi` (`order_id`),
  KEY `idx_uid` (`user_id`),
  KEY `idx_src` (`source`),
  KEY `idx_ua` (`user_assign`),
  KEY `idx_oa` (`order_assign`)
) ENGINE=InnoDB AUTO_INCREMENT=391121 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='站内订单';

-- ----------------------------
-- Table structure for rdc_classroom_sale_order_refund
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_sale_order_refund`;
CREATE TABLE `rdc_classroom_sale_order_refund` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `order_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单id',
  `action_type` int(11) NOT NULL DEFAULT '0' COMMENT '1:退费 2:返现',
  `amount` int(11) NOT NULL DEFAULT '0' COMMENT '退费金额,单位（分）',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `package_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '套餐id',
  `package_name` varchar(128) NOT NULL DEFAULT '' COMMENT '套餐名称',
  `work_order_id` varchar(128) NOT NULL DEFAULT '' COMMENT '工单ID',
  `opuser` varchar(32) NOT NULL DEFAULT '' COMMENT '退款操作人',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
  PRIMARY KEY (`id`),
  KEY `idx_o` (`order_id`),
  KEY `idx_u` (`user_id`),
  KEY `idx_woi` (`work_order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单退费表';

-- ----------------------------
-- Table structure for rdc_classroom_sale_teacher
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_sale_teacher`;
CREATE TABLE `rdc_classroom_sale_teacher` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `biz_type` bigint(20) NOT NULL DEFAULT '2' COMMENT '业务线',
  `teacher_opname` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '顾问名称',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '状态 1=正常 100=离职',
  `polling_switch` tinyint(4) NOT NULL DEFAULT '1' COMMENT '轮训开关 1=开 2=关',
  `capacity` int(11) NOT NULL DEFAULT '0' COMMENT '容量',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `role_tags` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '角色',
  `assign_tags` varchar(1024) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '分配标签',
  `bind_qw` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '绑定公司企微',
  `nickname_en` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '英文昵称',
  `grade_level` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '服务等级 S/A/B...',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='顾问教师表';

-- ----------------------------
-- Table structure for rdc_classroom_sale_teacher_task
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_sale_teacher_task`;
CREATE TABLE `rdc_classroom_sale_teacher_task` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `sale_teacher_opname` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '助教opname',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `task_type` int(10) NOT NULL DEFAULT '0' COMMENT '任务类型, 1=预习未点评 2=复习未点评 3=课后作业未点评',
  `task_source` int(10) NOT NULL DEFAULT '0' COMMENT '任务来源',
  `bind_id` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '绑定业务id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '助教任务状态 1=已完成 3=未完成',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '拓展信息',
  `finished_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '完成时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `finished_num` int(11) NOT NULL DEFAULT '0' COMMENT '完成次数',
  `submit_num` int(11) NOT NULL DEFAULT '0' COMMENT '任务提交次数 已完成-未完成算一次，未完成-未完成不算',
  `submit_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '任务提交时间 已完成-未完成更新，未完成-未完成不更新',
  PRIMARY KEY (`id`),
  KEY `idx_sct` (`sale_teacher_opname`,`ct`,`task_type`),
  KEY `idx_uc` (`uid`,`ct`)
) ENGINE=InnoDB AUTO_INCREMENT=1864 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='助教任务';

-- ----------------------------
-- Table structure for rdc_classroom_schedule_evaluation
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_schedule_evaluation`;
CREATE TABLE `rdc_classroom_schedule_evaluation` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `schedule_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '到课ID',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '学员ID ',
  `submitter` varchar(64) NOT NULL COMMENT '提交人ID',
  `submitter_type` tinyint(4) NOT NULL COMMENT '提交人类型 (1:学生 2:老师)',
  `content` longtext COMMENT '评价内容 (JSON格式)',
  `additional_feedback` text COMMENT '额外意见建议',
  `is_anonymous` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否匿名评价',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `overall` varchar(64) NOT NULL DEFAULT '' COMMENT '整体评价',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_submitter` (`schedule_id`,`uid`,`submitter_type`)
) ENGINE=InnoDB AUTO_INCREMENT=129 DEFAULT CHARSET=utf8mb4 COMMENT='到课评论表';

-- ----------------------------
-- Table structure for rdc_classroom_script
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_script`;
CREATE TABLE `rdc_classroom_script` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `scene` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '场景：classroom',
  `code` text COLLATE utf8mb4_bin NOT NULL COMMENT 'js 代码',
  `version` int(10) NOT NULL DEFAULT '0' COMMENT '版本',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_sv` (`scene`,`version`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='脚本';

-- ----------------------------
-- Table structure for rdc_classroom_shopmall_order
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_shopmall_order`;
CREATE TABLE `rdc_classroom_shopmall_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` bigint(20) NOT NULL COMMENT 'user_id',
  `sale_order_id` bigint(20) NOT NULL COMMENT 'sale order id',
  `root_id` bigint(20) NOT NULL COMMENT '与 sale order 的 trade_no 关联，可知是哪个child',
  `eid` bigint(20) NOT NULL COMMENT '商城订单id',
  `gid` bigint(20) NOT NULL COMMENT '商品id，与 sale order 的 equity_uniqueid 关联',
  `ems` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '物流id',
  `ems_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '物流公司，如中通快递',
  `remark` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '订单备注',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '内部使用的订单状态，1: 待确认地址；2：出库中，3：暂无物流信息；4：在路上，5：已签收，6：问题件，7：已取消',
  `exported_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '导出状态：1: 未导出，2：延缓导出，3：已导出',
  `exported_ts` int(11) NOT NULL DEFAULT '0' COMMENT '导出时间戳',
  `exception_type` int(11) NOT NULL DEFAULT '1' COMMENT '异常类型: 1: 无异常, 2: 发货异常, 3: 签收异常',
  `new_product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '更换货物时的货物id，为0代表无需更换',
  `sku` varchar(128) NOT NULL DEFAULT '' COMMENT 'sku',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间戳',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_eid` (`eid`),
  KEY `idx_ui` (`user_id`),
  KEY `idx_sct` (`status`,`ct`),
  KEY `idx_so` (`sale_order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COMMENT='商城单表';

-- ----------------------------
-- Table structure for rdc_classroom_sku_category
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_sku_category`;
CREATE TABLE `rdc_classroom_sku_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `biz_type` bigint(20) NOT NULL DEFAULT '2' COMMENT '业务线',
  `source` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单来源，1: 内部订单，2: 有赞',
  `sku` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '商品SKU',
  `sku_remark` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'SKU描述信息',
  `category_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '兑换码批次ID',
  `ct` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `creator` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '创建人',
  `gift_conf` text COLLATE utf8mb4_bin NOT NULL COMMENT '赠品配置信息',
  PRIMARY KEY (`id`),
  KEY `idx_sk` (`source`,`sku`)
) ENGINE=InnoDB AUTO_INCREMENT=222129 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='各渠道SKU与兑换码映射关系';

-- ----------------------------
-- Table structure for rdc_classroom_tag
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_tag`;
CREATE TABLE `rdc_classroom_tag` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `tag_name` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '标签名称',
  `scene` bigint(20) NOT NULL DEFAULT '0' COMMENT '标签使用场景',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ts` (`tag_name`,`scene`)
) ENGINE=InnoDB AUTO_INCREMENT=498 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='小班课标签表';

-- ----------------------------
-- Table structure for rdc_classroom_teacher
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_teacher`;
CREATE TABLE `rdc_classroom_teacher` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '老师uid',
  `teacher_name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '教师名称',
  `level_list` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '教师等级[1,2,3]',
  `certificate` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '资格证书',
  `description` longtext COLLATE utf8mb4_bin COMMENT '简介信息',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '状态 1=正常',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `teacher_type` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '教师类型：中教，外教',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=157 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='上课教师表';

-- ----------------------------
-- Table structure for rdc_classroom_teacher_availability
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_teacher_availability`;
CREATE TABLE `rdc_classroom_teacher_availability` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `teacher_uid` bigint(20) NOT NULL COMMENT '老师ID',
  `start_time` time NOT NULL COMMENT '时段开始时间',
  `end_time` time NOT NULL COMMENT '时段结束时间',
  `strategy_start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '策略开始时间',
  `strategy_end_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '策略结束时间',
  `recurrence_type` tinyint(4) NOT NULL COMMENT '循环类型：1-每日循环，2-每周循环',
  `week_days` varchar(20) NOT NULL DEFAULT '' COMMENT '适用于每周循环时，存储周几可用，如"1,2,3,4,5,6,0"表示周一至周日',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1-开启 2-关闭',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_tsr` (`teacher_uid`,`status`,`strategy_start_time`,`strategy_end_time`,`recurrence_type`),
  KEY `idx_sr` (`status`,`strategy_start_time`,`strategy_end_time`,`recurrence_type`),
  KEY `idx_se` (`start_time`,`end_time`)
) ENGINE=InnoDB AUTO_INCREMENT=152 DEFAULT CHARSET=utf8mb4 COMMENT='老师可用开课时间配置表';

-- ----------------------------
-- Table structure for rdc_classroom_teacher_schedule
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_teacher_schedule`;
CREATE TABLE `rdc_classroom_teacher_schedule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '老师id',
  `class_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '班级id',
  `schedule_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '排课id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '排课状态 100=初始化 200=已结束 300=已取消',
  `start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始上课时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束上课时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `lecture_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '课节id',
  `schedule_type` int(11) NOT NULL COMMENT '排课类型 10=小班课 20=直播课 30=临时班小班课',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`,`start_time`),
  KEY `idx_sid` (`schedule_id`,`start_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1112 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='老师排课记录';

-- ----------------------------
-- Table structure for rdc_classroom_user
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user`;
CREATE TABLE `rdc_classroom_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `user_level` int(11) NOT NULL DEFAULT '0' COMMENT '用户定级',
  `biz_type` bigint(20) NOT NULL DEFAULT '0' COMMENT '业务线',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `extend` text COLLATE utf8mb4_bin NOT NULL COMMENT '拓展信息 json',
  `sale_teacher_opname` varchar(255) COLLATE utf8mb4_bin DEFAULT '' COMMENT '归属销售',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_bz` (`uid`,`biz_type`),
  KEY `idx_sc` (`sale_teacher_opname`,`ct`)
) ENGINE=InnoDB AUTO_INCREMENT=369 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

-- ----------------------------
-- Table structure for rdc_classroom_user_bank
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_bank`;
CREATE TABLE `rdc_classroom_user_bank` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id，无意义',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `real_name` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '真实姓名',
  `phone` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '手机号（带区号）',
  `id_card` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '身份证号',
  `bank_cardno` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '银号卡号',
  `pay_channel` int(10) NOT NULL DEFAULT '0' COMMENT '打款渠道 1=云账户',
  `opuser` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `id_type` int(11) NOT NULL DEFAULT '0' COMMENT '证件类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户拓展银行信息';

-- ----------------------------
-- Table structure for rdc_classroom_user_class
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_class`;
CREATE TABLE `rdc_classroom_user_class` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `class_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '班级id',
  `sale_teacher_opname` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '顾问名称',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `assign_type` int(10) NOT NULL DEFAULT '0' COMMENT '分班类型：1=正常分配 2=插班分配',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_cu` (`class_id`,`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=673 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户分班表';

-- ----------------------------
-- Table structure for rdc_classroom_user_class_hour_rights
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_class_hour_rights`;
CREATE TABLE `rdc_classroom_user_class_hour_rights` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `balance` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '课时余额，0.01 课时',
  `total_issued` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '累计发放课时数',
  `total_consumed` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '累计实际消耗课时数',
  `gift_total_consumed` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送课时总消耗',
  `gift_total_issued` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送课时总发放',
  `gift_balance` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送课时余额',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `right_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '权益类型：1=小班课权益，2=菲教权益',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_id` (`user_id`,`right_type`)
) ENGINE=InnoDB AUTO_INCREMENT=1286 DEFAULT CHARSET=utf8mb4 COMMENT='用户课时权益表';

-- ----------------------------
-- Table structure for rdc_classroom_user_home_bills
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_home_bills`;
CREATE TABLE `rdc_classroom_user_home_bills` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `biz_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '业务线',
  `pay_type` int(11) NOT NULL DEFAULT '0' COMMENT '缴费类型 1=电，2=水，3=网，4=学',
  `pay_star_num` int(11) NOT NULL DEFAULT '0' COMMENT '缴费星星数',
  `month_date` int(11) NOT NULL DEFAULT '0' COMMENT '账单月 200601',
  `is_paid` int(11) NOT NULL DEFAULT '2' COMMENT '是否缴费 1=已缴费 2=未缴费',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_ubm` (`uid`,`biz_type`,`month_date`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户家园账单';

-- ----------------------------
-- Table structure for rdc_classroom_user_metrics
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_metrics`;
CREATE TABLE `rdc_classroom_user_metrics` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `field_key` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '指标名',
  `field_val` text COLLATE utf8mb4_bin COMMENT '指标值',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_uk` (`uid`,`field_key`)
) ENGINE=InnoDB AUTO_INCREMENT=13240 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户实时指标';

-- ----------------------------
-- Table structure for rdc_classroom_user_msg
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_msg`;
CREATE TABLE `rdc_classroom_user_msg` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `biz_type` int(11) NOT NULL DEFAULT '0' COMMENT '业务线',
  `is_read` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已读 1=已读',
  `title` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '标题',
  `content` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '消息详情',
  `opuser` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_uc` (`uid`,`ct`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户消息';

-- ----------------------------
-- Table structure for rdc_classroom_user_pet
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_pet`;
CREATE TABLE `rdc_classroom_user_pet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `biz_type` int(11) NOT NULL DEFAULT '0' COMMENT '业务线',
  `pet_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '宠物id',
  `nickname` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '宠物名',
  `conf` longtext COLLATE utf8mb4_bin NOT NULL COMMENT '配置信息',
  `cur_level` bigint(20) NOT NULL DEFAULT '0' COMMENT '当前等级',
  `intimacy_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '亲密度',
  `is_show` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否展示',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_ub` (`uid`,`pet_id`),
  KEY `idx_bu` (`biz_type`,`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户宠物';

-- ----------------------------
-- Table structure for rdc_classroom_user_pet_exp_flow
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_pet_exp_flow`;
CREATE TABLE `rdc_classroom_user_pet_exp_flow` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `biz_type` int(11) NOT NULL DEFAULT '0' COMMENT '业务线',
  `pet_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '宠物id',
  `intimacy_exp_change` int(11) NOT NULL DEFAULT '0' COMMENT '经验变化值，可能为负数',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_upb` (`uid`,`pet_id`,`biz_type`)
) ENGINE=InnoDB AUTO_INCREMENT=596 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户宠物经验';

-- ----------------------------
-- Table structure for rdc_classroom_user_pet_props
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_pet_props`;
CREATE TABLE `rdc_classroom_user_pet_props` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `pet_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '宠物id',
  `biz_type` int(11) NOT NULL DEFAULT '0' COMMENT '业务线',
  `props_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '道具id',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_pp` (`props_id`,`pet_id`),
  KEY `idx_up` (`uid`,`pet_id`)
) ENGINE=InnoDB AUTO_INCREMENT=195 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户道具-宠物绑定关系';

-- ----------------------------
-- Table structure for rdc_classroom_user_props
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_props`;
CREATE TABLE `rdc_classroom_user_props` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `biz_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '业务线',
  `props_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '道具id',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_ubp` (`uid`,`biz_type`,`props_id`)
) ENGINE=InnoDB AUTO_INCREMENT=194 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户道具';

-- ----------------------------
-- Table structure for rdc_classroom_user_reward
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_reward`;
CREATE TABLE `rdc_classroom_user_reward` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `reward_type` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '奖励类型：rdc_classroom#star',
  `reward_num` int(10) NOT NULL DEFAULT '0' COMMENT '奖励数量',
  `scene` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '业务#场景：rdc_classroom#study_finished',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_usr` (`uid`,`scene`,`reward_type`),
  KEY `idx_uoo` (`uid`,`ct`),
  KEY `idx_sr` (`scene`,`reward_type`)
) ENGINE=InnoDB AUTO_INCREMENT=757 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户奖励';

-- ----------------------------
-- Table structure for rdc_classroom_user_room_info
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_room_info`;
CREATE TABLE `rdc_classroom_user_room_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户uid',
  `user_role` varchar(32) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '角色 teacher、student、admin',
  `schedule_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '排课id，房间id',
  `hand_raise_count` int(11) NOT NULL DEFAULT '0' COMMENT '举手次数',
  `star_count` int(11) NOT NULL DEFAULT '0' COMMENT '星星数',
  `brush_color` varchar(16) COLLATE utf8mb4_bin NOT NULL DEFAULT '#000000' COMMENT '画笔颜色',
  `class_duration` int(11) NOT NULL DEFAULT '0' COMMENT '上课时长(秒)',
  `last_ticket_ts` int(11) NOT NULL DEFAULT '0' COMMENT '最后心跳时间',
  `join_time` timestamp NULL DEFAULT NULL COMMENT '进入房间时间',
  `leave_time` timestamp NULL DEFAULT NULL COMMENT '离开房间时间',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 1=在线 2=离线',
  `audio_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '音频状态 1=开麦 2=关麦 3=老师禁用',
  `video_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '视频状态 1=开视频 2=关视频 3=老师禁用',
  `hand_raise_status` int(11) NOT NULL DEFAULT '0' COMMENT '举手状态 1=开视频 2=关视频 3=老师禁用',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `did` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '设备id',
  `platform_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '上台状态 1=未上台 2=上台',
  `brush_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '画笔状态 1=开 2=关 3=老师禁用',
  `courseware_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '课件交互状态 1=开 2=关 3=禁用',
  `chatmsg_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '讨论区发言权限',
  PRIMARY KEY (`id`),
  KEY `idx_lts` (`last_ticket_ts`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1471 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户上课房间信息表';

-- ----------------------------
-- Table structure for rdc_classroom_user_schedule
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_schedule`;
CREATE TABLE `rdc_classroom_user_schedule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `class_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '班级id',
  `schedule_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '排课id',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '排课状态 100=初始化 200=已结束 300=已取消',
  `start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始上课时间',
  `end_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '结束上课时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `opuser` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `user_class_type` int(10) NOT NULL DEFAULT '0' COMMENT '分班类型 1=正常 2=插班',
  `schedule_type` int(11) NOT NULL DEFAULT '0' COMMENT '排课类型 10=小班课 20=直播课 30=临时班小班课',
  PRIMARY KEY (`id`),
  UNIQUE KEY `rdc_classroom_user_schedule_pk` (`uid`,`schedule_id`),
  KEY `idx_ucs` (`uid`,`class_id`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=4936 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户排课记录';

-- ----------------------------
-- Table structure for rdc_classroom_user_schedule_study
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_schedule_study`;
CREATE TABLE `rdc_classroom_user_schedule_study` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id主键',
  `uid` bigint(20) NOT NULL COMMENT '用户id',
  `schedule_id` bigint(20) NOT NULL COMMENT '排课ID',
  `lecture_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '课节ID',
  `course_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '课程ID',
  `finished_picbook_ids` longtext NOT NULL COMMENT '已读完的绘本ID',
  `picbook_vocabulary` bigint(20) NOT NULL DEFAULT '0' COMMENT '完成的绘本词汇量',
  `picbook_nums` bigint(20) NOT NULL DEFAULT '0' COMMENT '完成的绘本数量',
  `homework_take_pic_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '作业拍拍ID',
  `learning_report_view_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '课程报告查看时间',
  `homework_1v1_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '1v1课后练习ID',
  `homework_take_pic_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '完成作业拍照的时间',
  `homework_correct_rate` float(5,2) NOT NULL DEFAULT '0.00' COMMENT '课后练习正确率',
  `homework_finish_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '完成课后作业的时间',
  `picbook_finish_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '绘本完成时间',
  `replay_finish_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '录播学习完成时间',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_us` (`uid`,`schedule_id`),
  KEY `idx_ucl` (`uid`,`course_id`,`lecture_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3665 DEFAULT CHARSET=utf8mb4 COMMENT='用户在课节上的学习情况表';

-- ----------------------------
-- Table structure for rdc_classroom_user_wallet
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_wallet`;
CREATE TABLE `rdc_classroom_user_wallet` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `biz_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '业务线',
  `star_balance` bigint(20) NOT NULL DEFAULT '0' COMMENT '星星余额',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_ub` (`uid`,`biz_type`)
) ENGINE=InnoDB AUTO_INCREMENT=1992 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户钱包';

-- ----------------------------
-- Table structure for rdc_classroom_user_wallet_flow
-- ----------------------------
DROP TABLE IF EXISTS `rdc_classroom_user_wallet_flow`;
CREATE TABLE `rdc_classroom_user_wallet_flow` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT 'uid',
  `biz_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '业务线',
  `flow_change` bigint(20) NOT NULL DEFAULT '0' COMMENT '流水变化值，负数为减少',
  `flow_type` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '流水类型',
  `scene` varchar(256) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '场景',
  `extend` json NOT NULL COMMENT '拓展信息',
  `reason` varchar(512) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '说明',
  `opuser` varchar(128) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '操作人',
  `ct` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ut` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_ubf` (`uid`,`biz_type`),
  KEY `idx_us` (`uid`,`scene`)
) ENGINE=InnoDB AUTO_INCREMENT=2088 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户钱包流水';

SET FOREIGN_KEY_CHECKS = 1;
