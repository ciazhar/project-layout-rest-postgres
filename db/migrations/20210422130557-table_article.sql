
-- +migrate Up
CREATE TABLE IF NOT EXISTS article
(
    id         uuid,
    title      text,
    content    text,
    author_id  uuid,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES author (id)
);

-- +migrate Down
