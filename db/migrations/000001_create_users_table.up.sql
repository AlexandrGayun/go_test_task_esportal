CREATE TABLE IF NOT EXISTS users(
    id serial primary key,
    username varchar(100) not null,
    password text not null,
    email varchar(100) not null,
    last_name varchar(200) not null,
    first_name varchar(200) not null,
    age tinyint unsigned not null);