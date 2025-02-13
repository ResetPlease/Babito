UPDATE Users SET balance = CASE 
    WHEN id = $1 THEN balance - $3
    WHEN id = $2 THEN balance + $3
END
WHERE id IN ($1,$2);
-- checking 'balance >= amount' in go code
