create table if not exists users
(
    user_id       bigint unsigned auto_increment comment '用户ID',
    created_at    datetime    not null default current_timestamp comment '创建时间',
    updated_at    datetime    not null default current_timestamp on update current_timestamp comment '最后更新时间',
    deleted_at    bigint               default 0 comment '删除时间',
    last_login_at timestamp            default null comment '最后登陆时间',
    last_login_ip varchar(128)         default null comment '最后登陆IP',
    username      varchar(20) not null comment '用户名',
    email         varchar(64) not null comment '邮箱',
    phone         varchar(20) comment '手机号码',
    status        enum (
        'Active',     # 已激活
        'Unverified', # 未验证
        'Disabled',   # 不可用
        'Deleted',    # 软删除
        'Locked',     # 已锁定
        'Cancelled'   # 已注销
        )                              default 'Active' comment '账户状态',
    first_name    varchar(20) comment '名',
    last_name     varchar(20) comment '性',
    nickname      varchar(20) comment '昵称',
    gender        enum (
        'Female',     # 女
        'Male',       # 男
        'Other'       # 其他
        )                              default 'Other' comment '性别',
    country       varchar(100) comment '国家',
    province      varchar(100) comment '省/州',
    city          varchar(100) comment '城市',
    address       varchar(255) comment '详细地址',
    zip_code      varchar(10) comment '邮编',
    bio           varchar(255) comment '个人简介',
    website       varchar(255) comment '个人网站',
    avatar        TEXT comment '头像',
    date_of_birth datetime comment '出生日期',

    primary key (user_id),
    unique index idx_users_phone (phone),
    unique index idx_users_username (username),
    unique index idx_users_email (email),
    unique index idx_users_created_at (created_at),
    unique index idx_users_username_email_deleted_at (username, email, deleted_at)
) engine = InnoDB
  auto_increment = 100000
  default charset = utf8mb4 comment ='用户基本信息表';