create table if not exists roles
(
    role_id     bigint unsigned auto_increment comment '角色唯一ID',
    created_at  timestamp    not null default current_timestamp comment '创建时间',
    updated_at  timestamp    not null default current_timestamp on update current_timestamp comment '最后更新时间',
    deleted_at  bigint                default 0 comment '删除时间',
    label       varchar(20)  not null comment '角色名称',
    code        varchar(255) not null comment '角色编码',
    status      boolean               default true comment '角色状态',
    description text                  default null comment '菜单描述',

    primary key (role_id),
    index idx_roles_created_at (created_at),
    unique index idx_roles_label (label),
    unique index idx_roles_code (code),
    unique index idx_roles_label_code_deleted_at (label, code, deleted_at)
) engine = InnoDB
  auto_increment = 100000
  default charset = utf8mb4 comment ='角色信息表';