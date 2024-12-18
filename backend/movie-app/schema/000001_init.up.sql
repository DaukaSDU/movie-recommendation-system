CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE movies
(
    id serial not null unique,
    title varchar(255) not null
);

CREATE TABLE users_movies
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    movie_id int references movies (id) on delete cascade not null
);