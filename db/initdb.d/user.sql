-- Use SQL Formatter Extension
-- drop schema if exists cyberdb;
-- create schema cyberdb;
drop table if exists user;

drop table if exists user_update;

-- リソース
CREATE TABLE `user` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `token` varchar(255) NOT NULL
);

-- イベント
CREATE TABLE `update_user` (
    `updated_id` int PRIMARY KEY NOT NULL,
    `user_id` int,
    `user_name` varchar(255),
    `updated_at` datetime NOT NULL DEFAULT (now())
);

-- イベント
CREATE TABLE `update_auth_token` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `user_id` int,
    `old_auth_token` varchar(255) NOT NULL,
    `new_auth_token` varchar(255) NOT NULL,
    `updated_at` datetime NOT NULL DEFAULT (now())
);

-- 外部キー追加
ALTER TABLE
    `update_user`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE
    `update_auth_token`
ADD
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

-- 適当に値を追加
insert into
    user (name, token)
values
    ("shunpei", "token_shunpei")