
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "hstore";

CREATE TYPE user_role AS ENUM ('admin', 'staff', 'user', 'disable');

CREATE TABLE users (
    id serial PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100),
    display_name VARCHAR(50),
    picture TEXT,
    token_version BIGINT,
    role user_role NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

CREATE TABLE user_providers (
    user_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE ON UPDATE CASCADE,
    provider VARCHAR(16) NOT NULL,
    provider_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP,
    PRIMARY KEY(provider, provider_id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS user_providers;
DROP TABLE IF EXISTS users;
DROP TYPE user_role;
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS "hstore";