CREATE TABLE users
(
    id serial not null unique,
    email varchar(255) not null,
    password_hash varchar(255) not null
);