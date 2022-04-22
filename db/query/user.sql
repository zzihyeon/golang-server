-- name: CreateUser :one
INSERT INTO users (
  email,
	name,
	phone,
  gender,
  birth_date
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE uid = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET email=$2, name=$3, phone=$4, gender=$5, birth_date=$6
WHERE uid = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE uid = $1;