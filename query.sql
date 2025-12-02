-- name: GetAllUsers :many
SELECT uuid, username, email_address, metadata, created_at, edited_at, deleted_at
	FROM auth.users;