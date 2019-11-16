create table detail
(
    id    serial not null
        constraint detail_pkey
            primary key,
    alias varchar(64),
    type  text
);

alter table detail
    owner to postgres;

INSERT INTO public.detail (id, alias, type) VALUES (11, 'Победившая Команда', '[''TEAM'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (5, 'Цвет карточки', '[''ENUM'', [''Желтая'', ''Красная'']]');
INSERT INTO public.detail (id, alias, type) VALUES (6, 'Причина', '[''STR'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (1, 'Кто', '[''PERS'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (3, 'Вратарь', '[''PERS'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (12, 'Счет', '[''STR'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (4, 'Кто бьёт', '[''PERS'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (7, 'Кому', '[''PERS'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (9, 'Кто встает на замену', '[''PERS'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (8, 'Кто уходит', '[''PERS'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (2, 'Количество очков', '[''INT'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (10, 'Команда', '[''TEAM'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (13, 'Кто бросает', '[''PERS'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (14, 'Чей фол', '[''PERS'', ]');
INSERT INTO public.detail (id, alias, type) VALUES (15, 'Результат', '[''ENUM'', [''Попал'', ''Промахнулся'']]');
create table event
(
    id        serial not null
        constraint event_pkey
            primary key,
    name      varchar(20),
    kindsport integer
        constraint event_kindsport_fkey
            references kindsport
);

alter table event
    owner to postgres;

INSERT INTO public.event (id, name, kindsport) VALUES (1, 'Гол', 1);
INSERT INTO public.event (id, name, kindsport) VALUES (2, 'Бросок', 2);
INSERT INTO public.event (id, name, kindsport) VALUES (3, 'Пенальти', 1);
INSERT INTO public.event (id, name, kindsport) VALUES (4, 'Карточка', 1);
INSERT INTO public.event (id, name, kindsport) VALUES (5, 'Замена', 1);
INSERT INTO public.event (id, name, kindsport) VALUES (6, 'Замена', 2);
INSERT INTO public.event (id, name, kindsport) VALUES (7, 'Победа', 1);
INSERT INTO public.event (id, name, kindsport) VALUES (8, 'Победа', 2);
INSERT INTO public.event (id, name, kindsport) VALUES (9, 'Штрафной бросок', 2);
create table game
(
    id        serial not null
        constraint game_pkey
            primary key,
    date      date,
    kindsport integer
        constraint game_kindsport_fkey
            references kindsport
);

alter table game
    owner to postgres;

INSERT INTO public.game (id, date, kindsport) VALUES (68, '2019-11-14', 1);
INSERT INTO public.game (id, date, kindsport) VALUES (69, '2019-11-13', 1);
INSERT INTO public.game (id, date, kindsport) VALUES (70, '2019-11-16', 1);
create table historydetail
(
    id           serial not null
        constraint historydetail_pkey
            primary key,
    historyevent integer
        constraint historydetail_historyevent_fkey
            references historyevent,
    item         integer
        constraint historydetail_item_fkey
            references item,
    value        text
);

alter table historydetail
    owner to postgres;

INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (12, 9, 1, 'Иванов');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (13, 9, 2, 'Сидоров');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (14, 10, 5, 'Петров');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (15, 10, 6, 'Иванов');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (16, 11, 1, 'Сидоров');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (17, 11, 2, 'Иванов');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (18, 12, 12, 'Булгаков');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (19, 12, 13, 'Булгаков');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (20, 12, 14, 'Только вперед');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (21, 13, 5, 'Иванов');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (22, 13, 6, 'Иванов');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (23, 14, 9, 'Красная');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (24, 14, 10, 'iyiyfyifvuh');
INSERT INTO public.historydetail (id, historyevent, item, value) VALUES (25, 14, 11, 'Иванов');
create table historyevent
(
    id    serial not null
        constraint historyevent_pkey
            primary key,
    game  integer
        constraint historyevent_game_fkey
            references game,
    time  interval,
    event integer
        constraint historyevent_event_fkey
            references event
);

alter table historyevent
    owner to postgres;

INSERT INTO public.historyevent (id, game, time, event) VALUES (9, 68, '0 years 0 mons 0 days 0 hours 2 mins 0.00 secs', 1);
INSERT INTO public.historyevent (id, game, time, event) VALUES (10, 69, '0 years 0 mons 0 days 0 hours 2 mins 0.00 secs', 3);
INSERT INTO public.historyevent (id, game, time, event) VALUES (11, 69, '0 years 0 mons 0 days 0 hours 4 mins 0.00 secs', 1);
INSERT INTO public.historyevent (id, game, time, event) VALUES (12, 69, '0 years 0 mons 0 days 0 hours 6 mins 0.00 secs', 5);
INSERT INTO public.historyevent (id, game, time, event) VALUES (13, 70, '0 years 0 mons 0 days 0 hours 0 mins 50.00 secs', 3);
INSERT INTO public.historyevent (id, game, time, event) VALUES (14, 70, '0 years 0 mons 0 days 342935 hours 31 mins 30.00 secs', 4);
create table item
(
    id     serial not null
        constraint item_pkey
            primary key,
    detail integer
        constraint item_detail_fkey
            references detail,
    event  integer
        constraint item_event_fkey
            references event
);

alter table item
    owner to postgres;

INSERT INTO public.item (id, detail, event) VALUES (1, 1, 1);
INSERT INTO public.item (id, detail, event) VALUES (2, 3, 1);
INSERT INTO public.item (id, detail, event) VALUES (3, 1, 2);
INSERT INTO public.item (id, detail, event) VALUES (4, 2, 2);
INSERT INTO public.item (id, detail, event) VALUES (5, 4, 3);
INSERT INTO public.item (id, detail, event) VALUES (6, 3, 3);
INSERT INTO public.item (id, detail, event) VALUES (9, 5, 4);
INSERT INTO public.item (id, detail, event) VALUES (10, 6, 4);
INSERT INTO public.item (id, detail, event) VALUES (11, 7, 4);
INSERT INTO public.item (id, detail, event) VALUES (12, 8, 5);
INSERT INTO public.item (id, detail, event) VALUES (13, 9, 5);
INSERT INTO public.item (id, detail, event) VALUES (14, 10, 5);
INSERT INTO public.item (id, detail, event) VALUES (15, 8, 6);
INSERT INTO public.item (id, detail, event) VALUES (16, 9, 6);
INSERT INTO public.item (id, detail, event) VALUES (17, 10, 6);
INSERT INTO public.item (id, detail, event) VALUES (18, 11, 7);
INSERT INTO public.item (id, detail, event) VALUES (19, 12, 7);
INSERT INTO public.item (id, detail, event) VALUES (20, 11, 8);
INSERT INTO public.item (id, detail, event) VALUES (21, 12, 8);
INSERT INTO public.item (id, detail, event) VALUES (22, 13, 9);
INSERT INTO public.item (id, detail, event) VALUES (23, 14, 9);
INSERT INTO public.item (id, detail, event) VALUES (24, 15, 9);
create table kindsport
(
    id   serial not null
        constraint kindsport_pkey
            primary key,
    name text
);

alter table kindsport
    owner to postgres;

INSERT INTO public.kindsport (id, name) VALUES (1, 'Футбол');
INSERT INTO public.kindsport (id, name) VALUES (2, 'Баскетбол');
create table participant
(
    id   serial not null
        constraint participant_pkey
            primary key,
    name varchar(32)
        constraint participant_name_key
            unique,
    team integer
        constraint participant_team_fkey
            references team
);

alter table participant
    owner to postgres;

INSERT INTO public.participant (id, name, team) VALUES (1, 'Иванов', 1);
INSERT INTO public.participant (id, name, team) VALUES (2, 'Петров', 1);
INSERT INTO public.participant (id, name, team) VALUES (3, 'Сидоров', 1);
INSERT INTO public.participant (id, name, team) VALUES (4, 'Булгаков', 2);
INSERT INTO public.participant (id, name, team) VALUES (5, 'Пушкин', 2);
INSERT INTO public.participant (id, name, team) VALUES (6, 'Ахматова', 2);
INSERT INTO public.participant (id, name, team) VALUES (7, 'Гагарин', 3);
INSERT INTO public.participant (id, name, team) VALUES (8, 'Титов', 3);
INSERT INTO public.participant (id, name, team) VALUES (9, 'Николаев', 3);
INSERT INTO public.participant (id, name, team) VALUES (10, 'Жуков', 4);
INSERT INTO public.participant (id, name, team) VALUES (11, 'Рокосовский', 4);
INSERT INTO public.participant (id, name, team) VALUES (12, 'Малиновкий', 4);
create table team
(
    id        serial not null
        constraint team_pkey
            primary key,
    name      varchar(32)
        constraint team_name_key
            unique,
    kindsport integer
        constraint team_kindsport_fkey
            references kindsport
);

alter table team
    owner to postgres;

INSERT INTO public.team (id, name, kindsport) VALUES (1, 'Только вперед', 1);
INSERT INTO public.team (id, name, kindsport) VALUES (2, 'Непобедимые', 1);
INSERT INTO public.team (id, name, kindsport) VALUES (3, 'Солянка по француски', 2);
INSERT INTO public.team (id, name, kindsport) VALUES (4, 'Кружка соли', 2);
create table teamsingames
(
    id   serial not null
        constraint teamsingames_pkey
            primary key,
    game integer
        constraint teamsingames_game_fkey
            references game,
    team integer
        constraint teamsingames_team_fkey
            references team
);

alter table teamsingames
    owner to postgres;

INSERT INTO public.teamsingames (id, game, team) VALUES (45, 68, 2);
INSERT INTO public.teamsingames (id, game, team) VALUES (46, 68, 1);
INSERT INTO public.teamsingames (id, game, team) VALUES (47, 69, 2);
INSERT INTO public.teamsingames (id, game, team) VALUES (48, 69, 1);
INSERT INTO public.teamsingames (id, game, team) VALUES (49, 70, 1);
INSERT INTO public.teamsingames (id, game, team) VALUES (50, 70, 2);
