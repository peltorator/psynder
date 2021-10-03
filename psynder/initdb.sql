drop table if exists accounts cascade;
create table accounts
(
    id             serial primary key,
    email          varchar(255) not null,
    password_hash  bytea not null,

    unique (email)
);

drop table if exists psynas cascade;
create table psynas
(
    id          serial primary key,
    name        varchar(255) not null,
    description varchar(255) not null,
    PhotoLink   varchar(255) not null
);

drop table if exists likes cascade;
create table likes
(
    accountId int not null,
    psynaId   int not null,

    foreign key (accountId) REFERENCES accounts(id),
    foreign key (psynaId) REFERENCES psynas(id)

);

drop table if exists lastView cascade;
create table lastView
(
    accountId int not null primary key,
    psynaId   int not null,

    foreign key (accountId) REFERENCES accounts(id)
);

insert into psynas(name, description, photolink) values ('Биба', 'Описание1', 'https://sun9-10.userapi.com/c830408/v830408596/1e3417/lWKS4Fju4T0.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;