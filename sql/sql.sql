CREATE DATABASE IF NOT EXISTS rede_social;
USE rede_social;

DROP TABLE IF EXISTS rede_social;

CREATE TABLE rede_social(
    id int auto_increment primary key,
    name varchar(255) not null,
    nick varchar(255) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    createdin timestamp default current_timestamp()
) ENGINE=INNODB;