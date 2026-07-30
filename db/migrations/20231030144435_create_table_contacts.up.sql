create table contacts
(
    id         varchar(100) not null,
    first_name varchar(100) not null,
    last_name  varchar(100) null,
    email      varchar(100) null,
    phone      varchar(100) null,
    user_id    varchar(100) not null,
    created_at timestamptz  not null,
    updated_at timestamptz  not null,
    primary key (id),
    CONSTRAINT fk_contacts_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);