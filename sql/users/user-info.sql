-- name: CheckUserExists :one
SELECT EXISTS (
  SELECT 1 FROM user_base WHERE EMAIL = $1
);

-- name: GetUserLoginInfo :one
SELECT id, email, password
FROM user_base 
WHERE email = $1 and password = $2;