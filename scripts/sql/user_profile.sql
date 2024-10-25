create table if not exists user_profile
(
    id            bigint unsigned comment '用户简介唯一ID',
    user_id       bigint unsigned comment '外键，链接users表',
    email         varchar(64) not null comment '邮箱',
    phone         varchar(20) comment '手机号码',
    first_name    varchar(20) comment '名',
    last_name     varchar(20) comment '性',
    nickname      varchar(20) comment '昵称',
    gender        enum (
        'Female', # 女
        'Male',   # 男
        'Other'   # 其他
        )                      default 'Other' comment '性别',
    country       varchar(100) comment '国家',
    province      varchar(100) comment '省/州',
    city          varchar(100) comment '城市',
    address       varchar(255) comment '详细地址',
    zip_code      varchar(10) comment '邮编',
    bio           varchar(255) comment '个人简介',
    website       varchar(255) comment '个人网站',
    avatar        TEXT comment '头像',
    date_of_birth datetime comment '出生日期',
    last_login_at datetime     default null comment '最后登陆时间',
    last_login_ip varchar(128) default null comment '最后登陆IP',

    primary key (id),
    unique index idx_users_phone (phone),
    unique index idx_users_email (email),
    constraint fk_user_profile_user_id foreign key (user_id)
        references users (user_id) on update cascade on delete cascade
) engine = InnoDB
  default charset = utf8mb4 comment ='用户基本信息表';