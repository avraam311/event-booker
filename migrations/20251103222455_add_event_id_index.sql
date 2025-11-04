-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_book_event_id ON book(event_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_book_event_id;
-- +goose StatementEnd
