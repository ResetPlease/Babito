-- $1 - from_user_id, $2 - to_user_username

WITH to_user AS (
    SELECT id FROM Users WHERE username = $2
)
SELECT id, balance 
FROM Users
WHERE id IN ($1, (SELECT id FROM to_user))
ORDER BY id 
FOR UPDATE;