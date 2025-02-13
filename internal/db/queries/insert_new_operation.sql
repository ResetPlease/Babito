INSERT INTO Operations
(
    user_id,
    type,
    amount,
    target_user_id,
    item
)
VALUES($1, $2, $3, $4, $5);