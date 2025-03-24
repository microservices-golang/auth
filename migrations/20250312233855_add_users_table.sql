-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NUll,
    password TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at timestamp with time zone default current_timestamp NOT NULL,
    updated_at timestamp with time zone default current_timestamp NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users"
-- +goose StatementEnd
