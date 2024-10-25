create table if not exists role_menu
(
    role_id varchar(21) comment '关联的角色ID',
    menu_id varchar(21) comment '关联的菜单ID',

    primary key (role_id, menu_id),
    constraint fk_role_menu_role_id foreign key (role_id)
        references roles (role_id) on delete cascade on update cascade,
    constraint fk_role_menu_menu_id foreign key (menu_id)
        references menu (menu_id) on delete cascade on update cascade
)engine = InnoDB
 default charset = utf8mb4 comment ='角色菜单关联表';