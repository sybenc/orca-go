create table if not exists user_auth
(
    user_auth_id bigint unsigned auto_increment comment '用户认证表唯一ID',
    user_id      bigint unsigned comment '用户ID',
    created_at   datetime     not null default current_timestamp comment '创建时间',
    updated_at   datetime     not null default current_timestamp on update current_timestamp comment '最后更新时间',
    deleted_at   bigint                default 0 comment '删除时间',
    password     varchar(255) not null comment '用户密码',
    mfa_enable   boolean               default false comment '是否开启mfa认证',
    mfa_backup_key varchar(255) comment 'mfa备用密钥',

    primary key (user_auth_id),
    constraint fk_user_auths_user_id foreign key (user_id)
        references users (user_id) on update cascade on delete cascade
)