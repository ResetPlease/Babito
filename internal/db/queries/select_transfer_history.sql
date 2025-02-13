WITH all_user_transfer AS 
(
    SELECT user_id, amount, target_user_id 
    FROM Operations 
    WHERE user_id = $1 AND type = 'transfer'
)

SELECT amount, b.username, b.id
FROM all_user_transfer AS a
JOIN Users AS b 
ON a.target_user_id = b.id
;
