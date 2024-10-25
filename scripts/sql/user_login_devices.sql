create table if not exists user_login_devices
(
    id           bigint unsigned comment '用户登陆设备历史表唯一ID',
    user_id      bigint unsigned comment '用户登陆设备历史表外键，链接users表',
    os           varchar(64)  not null comment '登陆的操作系统',
    device_name  varchar(100) not null comment '登陆设备',
    device_type  varchar(50)  not null comment '登陆设备类型',
    browser      varchar(50) comment '登陆使用的浏览器',
    last_used_at datetime     default null comment '最后使用时间',
    last_used_ip varchar(128) default null comment '最后使用IP',

    primary key (id),
    constraint fk_user_login_devices_user_id foreign key (user_id)
        references users (user_id) on update cascade on delete cascade
) engine = InnoDB
  default charset = utf8mb4 comment ='用户登陆设备历史表';