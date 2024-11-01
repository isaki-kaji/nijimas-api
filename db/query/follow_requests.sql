-- name: CreateFollowRequest :one
INSERT INTO follow_requests (
  request_id,
  uid,
  following_uid,
  status
) VALUES (
  $1, $2, $3, '0'
) RETURNING *;

-- name: GetFollowRequest :one
SELECT * 
FROM follow_requests
WHERE uid = $1 AND following_uid = $2 AND status = '0';
  
-- name: GetFollowRequests :many
SELECT
  fr.request_id,
  u.uid,
  u.username,
  u.profile_image_url
FROM users u
JOIN follow_requests fr
ON u.uid = fr.following_uid
WHERE fr.following_uid = $1 AND fr.status = '0';

-- name: UpdateRequestToApproved :one
UPDATE follow_requests
SET status = '1'
WHERE request_id = $1
RETURNING *;

-- name: UpdateRequestToRejected :one
UPDATE follow_requests
SET status = '2'
WHERE request_id = $1
RETURNING *;

-- name: DeleteFollowRequest :one
DELETE FROM follow_requests
WHERE following_uid = $2 AND uid = $1 AND status = '0'
RETURNING *;
