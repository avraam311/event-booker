-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS book (
    id SERIAL PRIMARY KEY,
    person_name TEXT NOT NULL,
    book VARCHAR(15) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    event_id INTEGER NOT NULL REFERENCES event(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS book;
-- +goose StatementEnd
