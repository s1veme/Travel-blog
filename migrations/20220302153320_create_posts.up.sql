CREATE TABLE posts (
    id bigserial not null primary key,
    title text NOT NULL,
    content text NOT NULL,
    owner int NOT NULL,
    FOREIGN KEY (owner) REFERENCES users(id) ON DELETE CASCADE
)