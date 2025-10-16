CREATE DATABASE godo_dev;

CREATE USER godo_dev_user WITH PASSWORD 'godo_dev_pass';

GRANT ALL PRIVILEGES ON DATABASE godo_dev TO godo_dev_user;

\connect godo_dev;

CREATE TABLE users (
    id         UUID         PRIMARY KEY,
    username   VARCHAR(255) NOT NULL,
    email      VARCHAR(255) UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE todo_items (
    id          UUID         PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) UNIQUE,
    is_done     BOOLEAN      DEFAULT false,
    created_at  TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    user_id     UUID         NOT NULL,
    
    FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE users OWNER TO godo_dev_user;
ALTER TABLE todo_items OWNER TO godo_dev_user;