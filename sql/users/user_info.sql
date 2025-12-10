-- name: FinduserInfoById :one
SELECT firstname, lastname
FROM user_info
WHERE id = $1;



