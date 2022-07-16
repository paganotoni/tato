 SELECT
    id,
    player, 
    kind, 
    class, 
    evaluation, 
    starting_zone, 
    ending_zone,
    created_at
FROM
    actions
ORDER BY 
    created_at DESC;
