##
```mysql

drop table if exists income;
create table income
(
    id          int auto_increment
        primary key,
    user_id     int            default 0    not null comment '用户id',
    amount      decimal(15, 2) default 0.00 not null comment '收益金额',
    type        varchar(32)                 not null comment '收益类型:积分POINT,余额BALANCE',
    biz_type    varchar(32)                 not null comment '业务类型:投资INVEST,推荐RECOMMEND,团队TEAM,俱乐部CLUB',
    biz_id      varchar(255)                not null comment '业务id',
    relation_id int            default 0    not null comment '关联id',
    reason      varchar(255)   default ''   not null comment '原因',
    status      varchar(16)    default ''   not null comment '状态:未领取WAIT,已领取RECEIVED',
    created_at  datetime                    not null comment '创建时间',
    updated_at  datetime                    not null comment '更新时间',
    constraint un_biz_id
        unique (biz_id)
)
    comment '收益表';

```