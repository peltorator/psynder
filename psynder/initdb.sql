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
insert into psynas(name, description, photolink) values ('Иван', 'Описание2', 'https://thypix.com/wp-content/uploads/lustige-tiere-24.jpg') ;
insert into psynas(name, description, photolink) values ('Кобан', 'Описание2', 'https://funik.ru/wp-content/uploads/2018/11/9b2d50675bd5ad956231-700x525.jpg') ;
insert into psynas(name, description, photolink) values ('Буба', 'Описание2', 'https://www.fresher.ru/manager_content/images/sobaki-kotorye-prosto-ne-mogut/big/4.jpg') ;
insert into psynas(name, description, photolink) values ('Добби', 'Описание2', 'https://i.pinimg.com/236x/cf/77/53/cf7753e2bb8398d25868b23975908bf8.jpg') ;
insert into psynas(name, description, photolink) values ('Гарри', 'Описание2', 'https://data.whicdn.com/images/227497769/original.jpg') ;
insert into psynas(name, description, photolink) values ('Гермиона', 'Описание2', 'https://tlgrm.ru/_/stickers/b8a/48c/b8a48c6e-3669-34ce-9895-517e25d4245f/7.jpg') ;
insert into psynas(name, description, photolink) values ('Хагрид', 'Описание2', 'https://pp.userapi.com/tlawq_1bKlfdIHOwc3C9o6AHQulwtKxyr3MiSg/wY4rM4fg-Ww.jpg') ;
insert into psynas(name, description, photolink) values ('Невил', 'Описание2', 'https://im-01.forfun.com/fetch/w295-ch400-preview/2c/2c19d03f67629d4f1cc2234533b9a192.jpeg') ;
insert into psynas(name, description, photolink) values ('Рон', 'Описание2', 'https://thypix.com/wp-content/uploads/lustige-tiere-24.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
insert into psynas(name, description, photolink) values ('Боба', 'Описание2', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg') ;
