
-- +migrate Up
CREATE TABLE IF NOT EXISTS author
(
    id         uuid,
    name       text unique,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    PRIMARY KEY (id)
);

-- +migrate Down
