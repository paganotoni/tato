CREATE TABLE IF NOT EXISTS games (
    id TEXT PRIMARY KEY,
    home_team TEXT,
    visit_team TEXT,
    game_name TEXT,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);