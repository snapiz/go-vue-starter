
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "hstore";

CREATE TYPE user_role AS ENUM ('admin', 'staff', 'user');
CREATE TYPE user_state AS ENUM ('enable', 'disable', 'maintenance');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(100) UNIQUE NOT NULL,
    email_hash VARCHAR(32) NOT NULL,
    username VARCHAR(50) UNIQUE,
    password VARCHAR(100),
    display_name VARCHAR(50),
    picture TEXT,
    token_version BIGINT,
    role user_role NOT NULL DEFAULT 'user',
    state user_state NOT NULL DEFAULT 'enable',
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

CREATE TABLE user_providers (
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE ON UPDATE CASCADE,
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
DROP TYPE user_state;
DROP TYPE user_role;
DROP EXTENSION IF EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS "hstore";