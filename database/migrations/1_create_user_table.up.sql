CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP WITH TIME ZONE,
    updated_at  TIMESTAMP WITH TIME ZONE,
    deleted_at  TIMESTAMP WITH TIME ZONE,
    email       VARCHAR(100) UNIQUE NOT NULL,
    name        VARCHAR(100),
    password    VARCHAR(100),
    profile_url VARCHAR(200)
);