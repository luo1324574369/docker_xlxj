CREATE TABLE `t_user`
(
    `id`          int(10) NOT NULL AUTO_INCREMENT COMMENT '自增',
    `open_id`     varchar(128) NOT NULL DEFAULT '' COMMENT '微信openid',
    `session_key` varchar(128) NOT NULL DEFAULT '' COMMENT '微信session_key',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_open_id` (`open_id`)
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8

CREATE TABLE `t_user_setting`
(
    `id`          int(10) NOT NULL AUTO_INCREMENT COMMENT '自增',
    `user_id`     int(10) NOT NULL DEFAULT '0' COMMENT '用户id',
    `bg_url`      varchar(255) NOT NULL DEFAULT '' COMMENT '背景图片',
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8