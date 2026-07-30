create table users
(
    id         varchar(100) not null,
    name       varchar(100) not null,
    password   varchar(100) not null,
    token      varchar(100) null,
    created_at timestamptz  not null,
    updated_at timestamptz  not null,
    primary key (id)
);