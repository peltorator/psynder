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

-- insert into accounts(email, password_hash, kind) values ('shelterTest@gmail.com', '123Shelter', 'shelter');
--
-- insert into shelter_info values (1, 'Saint Petersburg', 'Nevsky Prospect 14', '+79111234567');
--
-- insert into psynas(name, description, photo_link)
-- values ('Рон', 'Описание2', 'https://thypix.com/wp-content/uploads/lustige-tiere-24.jpg');
--
-- insert into shelter_dogs values (1, 1);

insert into psynas(name, description, photo_link)
values ('Гарри', 'Поттер?', 'https://lh3.googleusercontent.com/proxy/3AOe2HRN5BZwbXVaa4waNtuiYvI0h20pyusC6RoA9thb3YO6eEN3sDuouvwhXahro3oClRwSMW1eKsGh884VlLJoipT6Hr6kEI6nwQ_M6OuJ-NoE3aDELAWS-E8yY_mQChPIu18Sklvigb9sbvD6dQ');
insert into psynas(name, description, photo_link)
values ('Гермиона', '---* (типа палочка)', 'https://www.rulez-t.info/uploads/posts/2013-04/1365414252_1365364110_33-dogs-that-cannot7.jpg');
insert into psynas(name, description, photo_link)
values ('Хагрид', 'съел Кобана(', 'https://sobaky.info/wp-content/uploads/2019/06/001-1.jpg');
insert into psynas(name, description, photo_link)
values ('Невил', 'ка, а ложка',
        'https://image2.thematicnews.com/uploads/images/00/00/39/2016/10/03/4a8c5e862b.jpg');
insert into psynas(name, description, photo_link)
values ('Биба', 'брат Бобы', 'https://sun9-10.userapi.com/c830408/v830408596/1e3417/lWKS4Fju4T0.jpg');
insert into psynas(name, description, photo_link)
values ('Боба', 'боб Бибы', 'https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg');
insert into psynas(name, description, photo_link)
values ('Иван', 'абобус', 'https://thypix.com/wp-content/uploads/lustige-tiere-24.jpg');
insert into psynas(name, description, photo_link)
values ('Кобан', 'чик бегает по полю (весело)', 'https://funik.ru/wp-content/uploads/2018/11/9b2d50675bd5ad956231-700x525.jpg');
insert into psynas(name, description, photo_link)
values ('Буба', 'нет блин, черва', 'https://www.fresher.ru/manager_content/images/sobaki-kotorye-prosto-ne-mogut/big/4.jpg');
insert into psynas(name, description, photo_link)
values ('Добби', 'свободен!', 'https://i.pinimg.com/236x/cf/77/53/cf7753e2bb8398d25868b23975908bf8.jpg');
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