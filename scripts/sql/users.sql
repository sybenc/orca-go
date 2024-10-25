create table if not exists users
(
    user_id    bigint unsigned auto_increment comment '用户唯一ID，雪花算法生成',
    created_at datetime    not null default current_timestamp comment '创建时间',
    updated_at datetime    not null default current_timestamp on update current_timestamp comment '最后更新时间',
    deleted_at bigint               default 0 comment '删除时间',
    username   varchar(20) not null comment '用户名',

    primary key (user_id),
    unique index idx_users_username (username),
    unique index idx_users_created_at (created_at),
    unique index idx_users_username_email_deleted_at (username, deleted_at)
) engine = InnoDB
  default charset = utf8mb4 comment ='用户表';