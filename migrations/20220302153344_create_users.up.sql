CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    username varchar not null,
    encrypted_password varchar not null
)