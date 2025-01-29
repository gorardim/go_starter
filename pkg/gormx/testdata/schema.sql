-- auto-generated definition
create table user
(
    id         int(11) unsigned auto_increment primary key,
    username       varchar(32) default '' not null comment '登录账号',
    balance        decimal(10,2) default 0.00 not null comment '余额',
    created_at      datetime               not null comment '创建时间',
    updated_at      datetime               not null comment '更新时间'
) charset = utf8mb4;

