CREATE TABLE `go_account` (
      `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
      `username` varchar(50) NOT NULL COMMENT '用户名',
      `password` varchar(255) DEFAULT NULL COMMENT '密码',
      `level` int(2) DEFAULT '2' COMMENT '账号类型，1为管理员，2为商户',
      `role_id` int(11) NOT NULL DEFAULT '2' COMMENT '(角色id)',
      `mobile` varchar(14) DEFAULT NULL COMMENT '手机号码',
      `status` int(2) NOT NULL DEFAULT '1' COMMENT '状态1：为正常 -1：为冻结',
      `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
      `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
      `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;



CREATE TABLE `go_app_pass_limit_host` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '服务器访问白名单表主键自增ID',
      `pass_host` varchar(64) DEFAULT NULL COMMENT '活跃的域名',
      `create_date` datetime DEFAULT NULL COMMENT '创建日期',
      `update_date` datetime DEFAULT NULL COMMENT '修改日期',
      `creator` varchar(64) DEFAULT NULL COMMENT '添加人ID',
      `updater` varchar(64) DEFAULT NULL COMMENT '更新人ID',
      `is_show` int(3) DEFAULT NULL COMMENT '是否展示 1是 0否',
      `is_delete` int(3) DEFAULT NULL COMMENT '是否删除 1是 0否',
      `pass_level` int(3) DEFAULT '0' COMMENT '活跃等级',
      `release_source` varchar(64) DEFAULT NULL COMMENT '放行源 (放行ip来源标注）',
      `release_module` varchar(64) NOT NULL COMMENT '放行模块 (对应系统模块)',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COMMENT='app服务器访问白名单表';


CREATE TABLE `go_email_log` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '邮箱发送记录表ID',
    `sending_mailbox` varchar(64) DEFAULT '' COMMENT '发件邮箱账号',
    `receive_email` varchar(64) DEFAULT NULL COMMENT '收件箱账号',
    `send_total` int(6) DEFAULT '0' COMMENT '发送数量',
    `available_number` int(6) DEFAULT '300' COMMENT '剩余可用条数',
    `status` int(1) DEFAULT NULL COMMENT '发送状态：1成功 0失败',
    `creat_date` datetime DEFAULT NULL COMMENT '创建时间',
    `update_date` datetime DEFAULT NULL COMMENT '修改时间',
    `creator` varchar(64) DEFAULT '' COMMENT '创建人',
    `updater` varchar(64) DEFAULT '' COMMENT '修改人',
    `remarks` varchar(255) DEFAULT '' COMMENT '备注',
    `body` text COMMENT '发送内容',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `go_email_routing` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '邮箱路由节点配置表ID',
    `email` varchar(30) DEFAULT NULL COMMENT '发信人邮箱',
    `channel` varchar(64) DEFAULT '' COMMENT '三方渠道',
    `channel_username` varchar(30) DEFAULT NULL COMMENT '渠道用户名',
    `channel_password` varchar(125) DEFAULT NULL COMMENT '渠道密码',
    `channel_host` varchar(50) DEFAULT NULL COMMENT '渠道服务器',
    `channel_port` int(6) DEFAULT NULL COMMENT '渠道端口号',
    `available_number` int(6) DEFAULT '300' COMMENT '剩余可用条数',
    `max_number` int(6) unsigned DEFAULT '300' COMMENT '每日最多使用条数',
    `rate_success` double(10,6) DEFAULT NULL COMMENT '成功率',
      `rate_fail` double(10,6) DEFAULT NULL COMMENT '失败率',
      `sort` int(5) unsigned DEFAULT '1' COMMENT '优先排序',
      `valid` int(3) unsigned DEFAULT '1' COMMENT '有效状态 1有效 0失效',
      `creat_date` datetime DEFAULT NULL COMMENT '创建时间',
      `update_date` datetime DEFAULT NULL COMMENT '修改时间',
      `creator` varchar(64) DEFAULT '' COMMENT '创建人',
      `updater` varchar(64) DEFAULT '' COMMENT '修改人',
      `remarks` text COMMENT '备注',
      `type` int(2) DEFAULT '1' COMMENT '类型：1验证码通知，2营销邮件',
      PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `go_fund` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT,
   `product_id` varchar(255) DEFAULT NULL COMMENT '产品id',
   `code` varchar(20) DEFAULT NULL COMMENT '代码',
   `name` varchar(255) DEFAULT NULL COMMENT '名称',
   `amount` double(10,2) DEFAULT NULL COMMENT '金额',
  `nav` double(10,5) DEFAULT NULL COMMENT '最新净值',
  `status` int(1) DEFAULT NULL COMMENT '状态：0-未启用，1-已启用',
  `create_at` bigint(14) DEFAULT NULL COMMENT '创建时间',
  `update_at` bigint(14) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `go_fund_day` (
   `id` bigint(20) NOT NULL AUTO_INCREMENT,
   `code` varchar(20) DEFAULT NULL COMMENT '代码',
   `name` varchar(255) DEFAULT NULL COMMENT '名称',
   `amount` double(10,2) DEFAULT NULL COMMENT '金额',
  `nav` double(10,5) DEFAULT NULL COMMENT '最新净值',
  `day_ts` bigint(14) DEFAULT NULL COMMENT '当天的时间戳',
  `day_at` timestamp NULL DEFAULT NULL COMMENT '当天的时间',
  `create_at` bigint(14) DEFAULT NULL COMMENT '创建时间',
  `update_at` bigint(14) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2049 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `go_login_log` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `account_id` int(11) NOT NULL COMMENT '账号id',
    `type` int(2) NOT NULL COMMENT '类型1：为总后台用户，2：为合作商用户',
    `username` varchar(50) NOT NULL COMMENT '用户名',
    `role_id` int(11) NOT NULL COMMENT '角色id',
    `ip` varchar(20) NOT NULL COMMENT '登录ip',
    `address` varchar(20) DEFAULT NULL COMMENT '登录地址',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_go_login_log_created_at` (`created_at`) USING BTREE,
    KEY `idx_go_login_log_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=492 DEFAULT CHARSET=utf8mb4;




CREATE TABLE `go_operation_log` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `type` int(2) NOT NULL DEFAULT '2' COMMENT '类型1：为总后台用户，2：为合作商用户',
    `account_id` int(11) NOT NULL COMMENT '账号id',
    `content` varchar(255) NOT NULL COMMENT '操作内容',
    `ip` varchar(20) NOT NULL COMMENT 'ip地址',
    `address` varchar(20) DEFAULT NULL COMMENT '操作地址',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_go_operation_log_created_at` (`created_at`) USING BTREE,
    KEY `idx_go_operation_log_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=326 DEFAULT CHARSET=utf8 COMMENT='操作记录表';




CREATE TABLE `go_rapidly_tbl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) DEFAULT NULL COMMENT '股票代码',
  `tag` int(11) DEFAULT NULL COMMENT '1:上涨，0:下跌',
  `key` varchar(255) DEFAULT NULL COMMENT '消息类型',
  `desc` varchar(255) DEFAULT NULL COMMENT '消息描述',
  `old` double(10,2) DEFAULT NULL COMMENT '之前价格',
  `new` double(10,2) DEFAULT NULL COMMENT '当前价格',
  `percent` double(10,2) DEFAULT NULL COMMENT '百分比',
  `offset_percent` double(10,2) DEFAULT NULL COMMENT '相对百分比',
  `day` date DEFAULT NULL COMMENT '当日0点时间戳',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`,`day`,`tag`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='股票急速上涨下跌提醒';


CREATE TABLE `go_role` (
   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
   `name` varchar(100) NOT NULL COMMENT '角色名称',
   `routes` varchar(255) DEFAULT NULL COMMENT '路由id,该角色所具有的路由',
   `desc` varchar(255) DEFAULT NULL COMMENT '角色描述',
   `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
   `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `idx_go_role_created_at` (`created_at`) USING BTREE,
   KEY `idx_go_role_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;




CREATE TABLE `go_routes` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
     `sort` int(11) DEFAULT '0' COMMENT '排序',
     `type` varchar(10) NOT NULL DEFAULT 'page' COMMENT 'page-页面   api-接口',
     `is_menu` int(1) NOT NULL DEFAULT '1' COMMENT '是否根菜单1-是 0-否',
     `route` varchar(255) DEFAULT NULL COMMENT '访问路由地址',
     `component` varchar(255) DEFAULT NULL COMMENT '页面组件地址',
     `name` varchar(100) NOT NULL COMMENT '路由名称',
     `icon` varchar(255) DEFAULT NULL COMMENT 'icon图标',
     `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级id',
     `create_by` int(11) DEFAULT NULL COMMENT '创建者',
     `status` char(1) DEFAULT NULL COMMENT '1-已启用   0-未启用',
     `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
     `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
     `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
     PRIMARY KEY (`id`) USING BTREE,
     KEY `idx_go_routes_created_at` (`created_at`) USING BTREE,
     KEY `idx_go_routes_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;




CREATE TABLE `go_stock` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `code` varchar(20) DEFAULT NULL COMMENT '代码',
    `name` varchar(255) DEFAULT NULL COMMENT '名称',
    `amount` double(10,2) DEFAULT NULL COMMENT '金额',
  `nav` double(10,5) DEFAULT NULL COMMENT '最新净值',
  `status` int(1) DEFAULT NULL COMMENT '状态：0-未启用，1-已启用',
  `create_at` bigint(14) DEFAULT NULL COMMENT '创建时间',
  `update_at` bigint(14) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5429 DEFAULT CHARSET=utf8mb4;



CREATE TABLE `go_stock_day` (
                                `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                `code` varchar(20) DEFAULT NULL COMMENT '代码',
                                `name` varchar(255) DEFAULT NULL COMMENT '名称',
                                `amount` double(10,2) DEFAULT NULL COMMENT '金额',
  `nav` double(10,5) DEFAULT NULL COMMENT '最新净值',
  `day_ts` bigint(14) DEFAULT NULL COMMENT '当天的时间戳',
  `day_at` timestamp NULL DEFAULT NULL COMMENT '当天的时间',
  `create_at` bigint(14) DEFAULT NULL COMMENT '创建时间',
  `update_at` bigint(14) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5338 DEFAULT CHARSET=utf8mb4;




CREATE TABLE `go_user` (
   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
   `openid` varchar(50) DEFAULT NULL COMMENT 'openid',
   `nickname` varchar(50) NOT NULL DEFAULT '2' COMMENT '昵称',
   `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
   `username` varchar(50) NOT NULL COMMENT '用户名',
   `password` varchar(255) DEFAULT NULL COMMENT '密码',
   `mobile` bigint(15) DEFAULT '2' COMMENT '手机号',
   `status` int(2) NOT NULL DEFAULT '1' COMMENT '状态1：为正常 -1：为冻结',
   `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
   `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
   `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `idx_go_account_created_at` (`created_at`) USING BTREE,
   KEY `idx_go_account_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;