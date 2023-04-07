CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR(100),
    fullname VARCHAR(100),
    email text UNIQUE,
    phone VARCHAR(100),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
