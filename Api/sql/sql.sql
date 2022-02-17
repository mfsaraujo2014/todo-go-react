CREATE DATABASE IF NOT EXISTS todo;
USE todo;

DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks(
    id int auto_increment primary key,
    title varchar(150) not null,
    completed TINY INT (1)
) ENGINE=INNODB;
