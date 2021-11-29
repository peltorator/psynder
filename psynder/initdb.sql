drop type if exists account_kind cascade;
create type account_kind as enum ('person', 'shelter');

drop table if exists accounts cascade;
create table accounts
(
    id            serial primary key,
    email         varchar(255) not null,
    password_hash bytea        not null,
    kind          account_kind,

    unique (email)
);

drop table if exists shelter_info cascade;
create table shelter_info
(
    account_id serial primary key,
    city       varchar(255),
    address    varchar(255),
    phone      varchar(20),

    foreign key (account_id) REFERENCES accounts (id)
);

drop table if exists psynas cascade;
create table psynas
(
    id          serial primary key,
    name        varchar(255) not null,
    description varchar(255) not null,
    photo_link  varchar(255) not null
);

drop table if exists shelter_dogs cascade;
create table shelter_dogs
(
    account_id int not null,
    psyna_id   int not null,

    constraint pk_shelter_dogs primary key (account_id, psyna_id),
    foreign key (account_id) REFERENCES accounts (id),
    foreign key (psyna_id) REFERENCES psynas (id)

);

drop table if exists ratings cascade;
create table ratings
(
    account_id int not null,
    psyna_id   int not null,
    liked      boolean,

    constraint pk_ratings primary key (account_id, psyna_id),
    foreign key (account_id) REFERENCES accounts (id),
    foreign key (psyna_id) REFERENCES psynas (id)

);

insert into accounts(email, password_hash, kind) values ('shelterTest@gmail.com', '123Shelter', 'shelter');

insert into shelter_info values (1, 'Saint Petersburg', 'Nevsky Prospect 14', '+79111234567');

insert into psynas(name, description, photo_link)
values ('Рон', 'Описание2', 'https://thypix.com/wp-content/uploads/lustige-tiere-24.jpg');

insert into shelter_dogs values (1, 1);

insert into psynas(name, description, photo_link)
values ('Биба', 'Описание1', 'https://sun9-10.userapi.com/c830408/v830408596/1e3417/lWKS4Fju4T0.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Иван', 'Описание2', 'https://thypix.com/wp-content/uploads/lustige-tiere-24.jpg');
insert into psynas(name, description, photo_link)
values ('Кобан', 'Описание2', 'https://funik.ru/wp-content/uploads/2018/11/9b2d50675bd5ad956231-700x525.jpg');
insert into psynas(name, description, photo_link)
values ('Буба', 'Описание2', 'https://www.fresher.ru/manager_content/images/sobaki-kotorye-prosto-ne-mogut/big/4.jpg');
insert into psynas(name, description, photo_link)
values ('Добби', 'Описание2', 'https://i.pinimg.com/236x/cf/77/53/cf7753e2bb8398d25868b23975908bf8.jpg');
insert into psynas(name, description, photo_link)
values ('Гарри', 'Описание2', 'https://data.whicdn.com/images/227497769/original.jpg');
insert into psynas(name, description, photo_link)
values ('Гермиона', 'Описание2', 'https://tlgrm.ru/_/stickers/b8a/48c/b8a48c6e-3669-34ce-9895-517e25d4245f/7.jpg');
insert into psynas(name, description, photo_link)
values ('Хагрид', 'Описание2', 'https://pp.userapi.com/tlawq_1bKlfdIHOwc3C9o6AHQulwtKxyr3MiSg/wY4rM4fg-Ww.jpg');
insert into psynas(name, description, photo_link)
values ('Невил', 'Описание2',
        'https://im-01.forfun.com/fetch/w295-ch400-preview/2c/2c19d03f67629d4f1cc2234533b9a192.jpeg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');