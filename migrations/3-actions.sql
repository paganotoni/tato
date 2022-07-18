CREATE TABLE IF NOT EXISTS actions (
    id TEXT PRIMARY KEY,
    
    player TEXT,
    kind TEXT,
    class TEXT,
    evaluation TEXT,
    starting_zone TEXT,
    ending_zone TEXT,
    game_id TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);