-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_info
ADD COLUMN age INT;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_info
DROP COLUMN IF EXISTS age;

-- +goose StatementEnd