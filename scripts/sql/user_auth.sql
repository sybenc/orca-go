create table if not exists user_auth
(
    id       bigint unsigned auto_increment comment '用户认证表唯一ID',
    user_id  bigint unsigned comment '用户认证表外键，链接users表',
    password varchar(255) not null comment '用户密码',
    status   enum (
        'Active',     # 已激活
        'Unverified', # 未验证
        'Disabled',   # 不可用
        'Deleted',    # 软删除
        'Locked',     # 已锁定
        'Cancelled'   # 已注销
        ) default 'Active' comment '账户状态',

    primary key (id),
    constraint fk_user_auths_user_id foreign key (user_id)
        references users (user_id) on update cascade on delete cascade
) engine = InnoDB
  default charset = utf8mb4 comment ='用户认证信息表';