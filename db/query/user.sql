-- name: CreateUser :one
INSERT INTO "user"."core" (email, hash, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING *;
