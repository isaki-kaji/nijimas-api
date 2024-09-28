-- name: CreateUser :one
INSERT INTO users (
  uid,
  username,
  country_code
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE uid = $1;

-- name: UpdateUser :one
UPDATE users SET
  username = COALESCE(sqlc.narg(username), username),
  self_intro = COALESCE(sqlc.narg(self_intro), self_intro),
  profile_image_url = COALESCE(sqlc.narg(profile_image_url), profile_image_url)
WHERE uid = sqlc.arg(uid)
RETURNING *;

-- name: GetFollowUsers :many
SELECT
  users.uid, 
  users.username
FROM users
JOIN follows
ON users.uid = follows.follow_user_id
WHERE follows.uid = $1;
