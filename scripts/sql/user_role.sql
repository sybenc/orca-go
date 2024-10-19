create table if not exists user_role
(
    user_id bigint unsigned comment '关联的用户ID',
    role_id bigint unsigned comment '关联的角色ID',

    primary key (user_id, role_id),
    constraint fk_user_role_user_id foreign key (user_id)
        references users (user_id) on delete cascade on update cascade,
    constraint fk_user_role_role_id foreign key (role_id)
        references roles (role_id) on delete cascade on update cascade
)