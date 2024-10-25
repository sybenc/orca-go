create table if not exists role_api
(
    role_id varchar(21) comment '关联的角色ID',
    api_id  varchar(21) comment '关联的API ID',

    primary key (role_id, api_id),
    constraint fk_role_api_role_id foreign key (role_id)
        references roles (role_id) on update cascade on delete cascade,
    constraint fk_role_api_api_id foreign key (api_id)
        references apis (api_id) on update cascade on delete cascade
) engine = InnoDB
  default charset = utf8mb4 comment ='角色API关联表';