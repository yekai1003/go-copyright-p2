drop database if exists copyright;
create database copyright character set utf8;
use copyright

drop table if exists vote;
drop table if exists account_content;
drop table if exists aution;
drop table if exists account;
drop table if exists content;
create table account
(
   account_id           int not null primary key auto_increment,
   email                 varchar(50),
   username             varchar(30),
   identity_id          varchar(100),
   address              varchar(256)
);
CREATE UNIQUE INDEX account_email_uindex ON copyright.account (email);
CREATE UNIQUE INDEX account_name_uindex ON copyright.account (username);
alter table account comment '账户表';


create table content
(
   content_id           int not null primary key auto_increment,
   title                varchar(100),
   content              varchar(256),
   content_hash         varchar(100),
   ts                   timestamp
);

create table account_content
(
   content_hash         varchar(100),
   token_id             int,
   address              varchar(100),
   ts                   timestamp
);


create table auction
(
   content_hash         varchar(256),
   address              varchar(100),
   token_id             int,
   percent              int,
   price                int,
   status               int,
   ts                   timestamp
);


create table vote
(
   vote_id              int primary key auto_increment,
   address              varchar(100),   
   content_hash         varchar(256),
   vote_time            timestamp,
   comment              varchar(100)
);

alter table vote comment '投票表，一个账户一个图片，只能投一票，一票代表50pxc';
CREATE UNIQUE INDEX vote_uindex ON copyright.vote (address,content_hash);

delete from vote;
delete from aution;
delete from account_content;
delete from content;
delete from account;
