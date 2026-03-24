CREATE TABLE IF NOT EXISTS sessions (
    token VARCHAR(64) PRIMARY KEY,
    createat bigint,
    type VARCHAR(64)
)