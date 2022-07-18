CREATE TABLE IF NOT EXISTS games (
    id TEXT PRIMARY KEY,
    
    home_team TEXT,
    visitor_team TEXT,
    game_name TEXT,

    service_team TEXT,
    home_team_setter_position INTEGER,
    visitor_team_setter_position INTEGER,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);