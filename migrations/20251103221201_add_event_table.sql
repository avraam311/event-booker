-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS event (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    seats_number INTEGER NOT NULL,
    seats_number_left INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS event;
-- +goose StatementEnd
