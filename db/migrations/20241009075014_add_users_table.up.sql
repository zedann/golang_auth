CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR NOT NULl,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL
)