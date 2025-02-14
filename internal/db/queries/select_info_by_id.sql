SELECT
    o.user_id,
    u1.username,
    o.type,
    o.amount,
    o.target_user_id,
    u2.username AS target_username,
    o.item
FROM
    Operations AS o
JOIN
    Users AS u1 ON o.user_id = u1.id
LEFT JOIN
    Users AS u2 ON o.target_user_id = u2.id
WHERE
    $1 IN (o.user_id, o.target_user_id)
;
