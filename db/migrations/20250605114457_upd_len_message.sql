-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE chat_history ALTER COLUMN content TYPE TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE chat_history ALTER COLUMN content TYPE VARCHAR(25);
-- +goose StatementEnd
