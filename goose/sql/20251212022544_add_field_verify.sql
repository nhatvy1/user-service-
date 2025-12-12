-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_base ADD COLUMN verify BOOLEAN DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
