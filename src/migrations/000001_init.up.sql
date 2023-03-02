CREATE TABLE songs (
    uid serial primary key not null,
    title varchar(255),
    duration bigint
);

insert into songs (title, duration) values ('test0', 4000);
insert into songs (title, duration) values ('test1', 4000);
insert into songs (title, duration) values ('test2', 4000);
insert into songs (title, duration) values ('test3', 4000);
insert into songs (title, duration) values ('test4', 4000);
insert into songs (title, duration) values ('test5', 4000);
insert into songs (title, duration) values ('test6', 4000);
insert into songs (title, duration) values ('test7', 4000);
insert into songs (title, duration) values ('test8', 4000);