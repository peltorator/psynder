drop table if exists accounts cascade;
create table accounts
(
    id        serial primary key,
    email     varchar(255) not null,
    password_hash  bytea not null,

    unique (email)
);