-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens(
    token,
    expires_at,
    user_id
)
VALUES(
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT users.*,
    refresh_tokens.token,
    refresh_tokens.created_at AS token_created_at,
    refresh_tokens.expires_at,
    refresh_tokens.revoked_at
FROM users
JOIN refresh_tokens
ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = $1
AND refresh_tokens.revoked_at IS NULL
AND refresh_tokens.expires_at > NOW();

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at=$2,
    updated_at=$3
WHERE token=$1;