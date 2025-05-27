-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS chat_history (
    user_id BIGINT,
	content VARCHAR(25) NOT NULL,
	role VARCHAR(25) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS chat_history;
-- +goose StatementEnd
