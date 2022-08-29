drop schema if exists users;
create schema users;
USE users;

drop table if exists user;

create table user (
    id int not null auto_increment primary key
    name varchar(255) not null
);

insert int user (name) values ("shunpei")