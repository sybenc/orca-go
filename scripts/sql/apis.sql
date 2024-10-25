create table if not exists apis
(
    api_id      varchar(21) comment 'api信息表唯一id，由nanoid生成',
    created_at  datetime                             not null default current_timestamp comment '创建时间',
    updated_at  datetime                             not null default current_timestamp on update current_timestamp comment '最后更新时间',
    deleted_at  bigint                                        default 0 comment '删除时间',
    name        varchar(64)                          not null comment 'api名称',
    end_point   varchar(255)                         not null comment 'api路径',
    http_method enum ('GET', 'POST', 'PUT','DELETE') not null comment 'http请求方法',
    `group`     varchar(64)                                   default 'default' comment 'api分组',
    version     varchar(64)                          not null comment 'api版本',
    status      boolean                                       default true comment 'api状态',

    primary key (api_id),
    unique index idx_apis_name (name),
    unique index idx_apis_group (`group`)
) engine = InnoDB
  default charset = utf8mb4 comment ='api信息表';