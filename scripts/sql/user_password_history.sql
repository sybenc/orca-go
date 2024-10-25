create table if not exists user_password_history
(
    id              bigint unsigned comment '用户密码历史表唯一ID',
    user_id         bigint unsigned comment '用户密码历史表外键，链接users表',
    password        varchar(255) not null comment '用户密码Argon2id哈希',
    last_changed_at datetime     default null comment '最后更改时间',
    last_used_at    datetime     default null comment '最后使用时间',
    last_used_ip    varchar(128) default null comment '最后使用IP',

    primary key (id),
    constraint fk_user_password_history_user_id foreign key (user_id)
        references users (user_id) on update cascade on delete cascade
) engine = InnoDB
  default charset = utf8mb4 comment ='用户密码历史表';