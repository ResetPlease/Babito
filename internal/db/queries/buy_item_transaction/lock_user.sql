SELECT balance
FROM Users
WHERE id = $1
FOR UPDATE
;
 