-- name: UserRegister :one
WITH new_user AS (
  INSERT INTO
    user_base (email, "password", auth_type)
  VALUES
    ($1, $2, $3) RETURNING id
),
insert_user_info AS (
  INSERT INTO
    user_info (id, firstname, lastname)
  SELECT id, $4, $5
  FROM new_user
),
insert_user_2fa AS (
  INSERT INTO user_2fa (id, enabled, secret)
  SELECT id, $6, $7
  FROM new_user
)
SELECT id
FROM new_user;

-- name: CheckUserExists :one
SELECT
  EXISTS (
    SELECT  id, email
    FROM  user_base
    WHERE email = $1
  );

-- name: GetUserLoginInfo :one
SELECT id, email, password
FROM user_base
WHERE email = $1;