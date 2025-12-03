-- name: CheckUserExists :one
SELECT EXISTS (
  SELECT 1 FROM users WHERE EMAIL = $1
);

-- name: GetUserLoginInfo :one
SELECT id, email, password
FROM users 
WHERE email = $1 and password = $2;