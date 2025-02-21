CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    discord_id TEXT UNIQUE NOT NULL,
    username TEXT NOT NULL
);

CREATE TABLE inhouses (
    id TEXT PRIMARY KEY, -- Changed from SERIAL to TEXT for ULID
    token TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    status TEXT CHECK (status IN ('waiting', 'ongoing', 'completed')) NOT NULL
);

CREATE TABLE inhouse_players (
    inhouse_id TEXT REFERENCES inhouses(id) ON DELETE CASCADE, -- Updated reference type
    player_id INT REFERENCES players(id) ON DELETE CASCADE,
    PRIMARY KEY (inhouse_id, player_id)
);

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    inhouse_id TEXT REFERENCES inhouses(id) ON DELETE CASCADE, -- Updated reference type
    name TEXT NOT NULL
);

CREATE TABLE team_players (
    team_id INT REFERENCES teams(id) ON DELETE CASCADE,
    player_id INT REFERENCES players(id) ON DELETE CASCADE,
    PRIMARY KEY (team_id, player_id)
);

CREATE TABLE game_results (
    id SERIAL PRIMARY KEY,
    inhouse_id TEXT REFERENCES inhouses(id) ON DELETE CASCADE, -- Updated reference type
    winning_team TEXT NOT NULL,
    completed_at TIMESTAMP DEFAULT NOW()
);
CREATE TABLE revoked_tokens (
    id SERIAL PRIMARY KEY,
    token TEXT UNIQUE NOT NULL,
    revoked_at TIMESTAMP DEFAULT NOW()
);
