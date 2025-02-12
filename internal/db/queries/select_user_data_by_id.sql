SELECT
    id,
    username,
    hashed_password,
    balance
FROM Users
WHERE
    id = $1
;
