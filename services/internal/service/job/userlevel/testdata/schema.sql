CREATE TABLE `user` (
  `user_id` int NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `level` int NOT NULL DEFAULT '0' COMMENT '代理级别',
  `customer_id` varchar(128) NOT NULL DEFAULT '0' COMMENT '客户id',
  `pid` int NOT NULL DEFAULT '0' COMMENT '父级id',
  `invite_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邀请码',
  `is_club_owner` varchar(16) DEFAULT 'N' COMMENT '是否俱乐部拥有者',
  `status` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '状态:启用ON,禁用OFF',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `sex` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '性别',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `pay_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付密码',
  `secret_two_fa` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '二次验证密钥',
  `bsc_address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'bsc地址',
  `bsc_uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'bsc uid',
  `trc_address` varchar(255) DEFAULT '' COMMENT 'trc20 address',
  `trc_uid` varchar(255) DEFAULT '' COMMENT 'trc20 uid',
  `balance` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '余额',
  `income` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '收益',
  `point` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '积分',
  `free_fee_withdraw_amount` decimal(15,2) DEFAULT '0.00' COMMENT '免手续费提现额度',
  `is_valid` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'N' COMMENT '是否有效',
  `kyc_status` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'KYC状态:未认证UNAUTH,已认证AUTH',
  `enable_level_grade` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'Y' COMMENT '是否开启用户等级定级',
  `invest_amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '投资金额',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `un_username` (`username`) COMMENT '用户名唯一',
  UNIQUE KEY `un_invite_code` (`invite_code`) COMMENT '邀请码唯一',
  UNIQUE KEY `un_customer_id` (`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1538 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

CREATE TABLE `user_tree` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户id',
  `parent_id` int NOT NULL COMMENT '父节点id',
  `distance` int NOT NULL COMMENT '层级距离',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_user_id_parent_id` (`user_id`,`parent_id`) COMMENT '用户id和父id唯一',
  KEY `idx_user_id` (`user_id`),
  KEY `idx_parent_id_distance` (`parent_id`,`distance`)
) ENGINE=InnoDB AUTO_INCREMENT=16045 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户树';

CREATE TABLE `setting` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_key` (`key`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `fund` (
  `fund_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户id',
  `from_user_id` int NOT NULL DEFAULT '0' COMMENT '来源用户id',
  `amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '金额',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型:INCOME,EXPEND',
  `account_type` varchar(32) NOT NULL COMMENT '账户类型:余额BALANCE,收益余额INCOME',
  `biz_type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务类型:领取收益INCOME_RECEIVE,投资INVEST,提现WITHDRAW,投资完成INVEST_FINISH',
  `order_id` int NOT NULL DEFAULT '0' COMMENT '订单id',
  `withdraw_id` int NOT NULL DEFAULT '0' COMMENT '提现id',
  `biz_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '业务用于幂等处理',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资金变动原因',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`fund_id`),
  UNIQUE KEY `un_biz_id` (`biz_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=20086 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户资金流水表';

CREATE TABLE `user_level` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名称:T1-T9',
  `level` int unsigned NOT NULL DEFAULT '0' COMMENT '级别:1-9',
  `direct_valid_member_count` int unsigned NOT NULL DEFAULT '0' COMMENT '直推有效会员数量',
  `agent_count` int unsigned NOT NULL DEFAULT '0' COMMENT '代理商数量',
  `invest_amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '会员投资要求',
  `team_invest_amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '业绩要求',
  `commission_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '佣金收益率',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `un_level` (`level`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='代理级别表';

CREATE TABLE `user_level_change_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户id',
  `from_level` int NOT NULL DEFAULT '0' COMMENT '变更前等级',
  `to_level` int NOT NULL DEFAULT '0' COMMENT '变更后等级',
  `reason` varchar(255) NOT NULL DEFAULT '' COMMENT '变更原因',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=98 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户等级变更记录表';

CREATE TABLE `invest_order` (
  `invest_id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL DEFAULT '0' COMMENT '订单id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `contract_id` int NOT NULL DEFAULT '0' COMMENT '合约id',
  `status` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '状态:释放中RELEASING,已结束FINISHED',
  `amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '投资金额',
  `balance_amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '余额投资金额',
  `income_amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '收益投资金额',
  `release_amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '已释放金额',
  `release_point` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '已释放积分',
  `contract_title` json NOT NULL COMMENT '合约标题',
  `contract_days` int NOT NULL DEFAULT '0' COMMENT '合约天数',
  `contract_max_amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '合约最大投资额度',
  `contract_min_amount` decimal(15,2) NOT NULL DEFAULT '0.00' COMMENT '合约最小投资额度',
  `contract_income_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '合约收益率',
  `contract_point_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '合约积分日化率',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`invest_id`),
  UNIQUE KEY `un_order_id` (`order_id`) COMMENT '订单id唯一',
  KEY `idx_user_id` (`user_id`) COMMENT '用户id索引',
  KEY `idx_order_id` (`order_id`) COMMENT '订单id索引',
  KEY `idx_created_at` (`created_at`) COMMENT '创建时间索引'
) ENGINE=InnoDB AUTO_INCREMENT=1476 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单投资表';

