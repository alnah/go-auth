-- name: CreateUser :one
INSERT INTO "user"."core" (email, hash, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"."core" WHERE email = $1;
