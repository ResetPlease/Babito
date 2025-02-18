INSERT INTO Users (username, hashed_password, balance)
VALUES ($1, $2, $3)
ON CONFLICT (username)
DO UPDATE SET username = EXCLUDED.username
RETURNING id, username, hashed_password
;
