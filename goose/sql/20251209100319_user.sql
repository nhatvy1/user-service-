-- +goose Up
-- +goose StatementBegin
CREATE TABLE
user_base (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL, -- user email is unique
    "password" VARCHAR(255), -- password can be NULL for OAuth users
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
  firstname VARCHAR(50),
  lastname VARCHAR(50),
  avatar_url VARCHAR(255),
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

CREATE TABLE provider_account (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES user_base (id) ON DELETE CASCADE,
  provider_name SMALLINT NOT NULL, -- e.g., 1 for Google, 2 for Facebook
  provider_user_id VARCHAR(100) NOT NULL,
  created_at TIMESTAMP
  WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP
  WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP,

  UNIQUE (provider_name, provider_user_id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_2fa;

DROP TABLE IF EXISTS user_info;

DROP TABLE IF EXISTS user_base;

-- +goose StatementEnd