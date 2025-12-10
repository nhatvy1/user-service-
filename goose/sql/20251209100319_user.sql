-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  user_base (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL, -- user email is unique
    password VARCHAR(255), -- password can be NULL for OAuth users
    auth_type SMALLINT NOT NULL, -- 0: credential, 1: google
    created_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  user_info (
    id INT PRIMARY KEY REFERENCES user_base (id) ON DELETE CASCADE,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    created_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP
  );

CREATE TABLE
  user_2fa (
    id INT PRIMARY KEY REFERENCES user_base (id) ON DELETE CASCADE,
    enabled BOOLEAN DEFAULT false,
    secret VARCHAR(255),
    created_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP
  );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_2fa;

DROP TABLE IF EXISTS user_info;

DROP TABLE IF EXISTS user_base;

-- +goose StatementEnd