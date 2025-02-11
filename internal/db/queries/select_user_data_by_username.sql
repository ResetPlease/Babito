SELECT
    id,
    username,
    hashed_password,
    balance
FROM Users
WHERE
    username = $1
;
