UPDATE Users
SET balance = balance - $2
WHERE id = $1
;
