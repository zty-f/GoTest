/*
 * 测试sql时间类型
 * timestamp和datetime区别
1.时间范围不同，通过我讲的这个生产问题，我们第一个得出的结论就是时间范围不同。timestamp的范围是有限制的，而datetime的时间是没有限制的。
timestamp的范围是从1970-01-01 到2038-01-19，这也就是为什么上面的生产问题会报错的原因。2018年，买了20年的会员就到了2038年，只要过了1-19号，那么用timestamp就会出现错误。
2.timestamp和datetime占用大小不一样。
在v5.6.4之前的版本。datetime占8个字节，timestamp占4个字节。
在这个版本之后datetime和timestamp占用的字节数都有一定程度的减少，并且根据是否有秒的小数的情况，相对应生成浮动大小。
3.timestamp会根据时区的情况进行时间转换，假设当前存储的时区和检索的时区有差异，那么timestamp会根据检索的时区进行转换。
而datetime就不会，datetime是属于你给他什么他就拿到什么，没有timestamp得这个功能。
4.二者存储null时有不同的效果，如果你的时间是datetime那么datetime就会把null存入进去，如果你用的是timestamp，那么timestamp会根据现有的时间帮你插入。
对于二者的区别就介绍的差不多了，在我们开发的实际情况中，要根据实际的业务情况去选择时使用timestamp还是datetime，如果没有什么特殊的需求，那就用timedate即可
建议使用timestamp的情况
1.当我们存储时间范围是比较小的时候。
2.当时间精度要求比较高的时候，timestamp可以达到的精度是纳秒级别。
3.需要使用默认值为当前时间的。
建议使用datetime 的情况
1.存储时间范围较大的时间。
2.需要存储时间与时区无关的情况。
 */
CREATE TABLE `type_test`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `year`              year(4)             NOT NULL COMMENT '年份-2023',
    `date`              date                NOT NULL COMMENT '日期-2023-08-01',
    `time`              time                NOT NULL COMMENT '时间-12:00:00',
    `timestamp_created` timestamp COMMENT '创建时间-2023-08-01 12:00:00', # DEFAULT CURRENT_TIMESTAMP 两种类型均可以使用
    `timestamp_updated` timestamp COMMENT '更新时间-2023-08-01 12:00:00', # DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    `datetime_created`  datetime COMMENT '创建时间-2023-08-01 12:00:00', # DEFAULT CURRENT_TIMESTAMP
    `datetime_updated`  datetime COMMENT '更新时间-2023-08-01 12:00:00', # DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='sql时间类型测试表';

# 测试数据
INSERT INTO `type_test` (`year`, `date`, `time`)
VALUES (2023, '2023-08-01', '12:00:00');
INSERT INTO `type_test` (`year`, `date`, `time`, `timestamp_created`, `timestamp_updated`, `datetime_created`, `datetime_updated`)
VALUES (2023, '2023-08-01', '12:00:00',null,null,null,null);
