create table if not exists menu
(
    menu_id     bigint unsigned auto_increment comment '菜单唯一ID',
    created_at  timestamp    not null default current_timestamp comment '创建时间',
    updated_at  timestamp    not null default current_timestamp on update current_timestamp comment '最后更新时间',
    deleted_at  bigint                default 0 comment '删除时间',
    label       varchar(20)  not null comment '菜单名称',
    code        varchar(255) not null comment '菜单编码',
    parent_id   bigint unsigned       default null comment '父级菜单ID',
    type        enum (
        'Menu',      # 菜单
        'Directory', # 目录
        'Button'     # 按钮
        )                    not null comment '菜单类型',
    route       text         not null comment '菜单路由路径',
    component   text         not null comment '菜单组件路径',
    icon_name   varchar(255)          default null comment '图标名称',
    `order`     int                   default 0 comment '菜单组内排序',
    keep_alive  boolean               default false comment '是否持久化菜单',
    `show`      boolean               default true comment '是否展示该菜单',
    status      boolean               default true comment '是否启用菜单',
    description text                  default null comment '菜单描述',

    primary key (menu_id),
    index idx_menu_created_at (created_at),
    index idx_menu_parent_id (parent_id),
    unique index idx_menu_label (label),
    unique index idx_menu_code (code),
    unique index idx_menu_label_code_deleted_at (label, code, deleted_at),

    constraint fk_menu_parent_id foreign key (parent_id)
        references menu (menu_id) on delete cascade on update cascade
) engine = InnoDB
  default charset = utf8mb4 comment ='菜单信息表';
