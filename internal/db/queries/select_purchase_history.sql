SELECT amount, item
FROM Operations 
WHERE user_id = $1 AND type = 'purchase'
;