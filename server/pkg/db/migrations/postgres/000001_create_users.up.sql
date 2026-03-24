CREATE TABLE IF NOT EXIST users (
    id VARCHAR(26) PRIMARY KEY,
    nickname VARCHAR(64),
    firstname VARCHAR(64),
    lastname VARCHAR(64),
    email VARCHAR(128) UNIQUE,
    createat bigint,
    updateat bigint,
    deleteat bigint,
    locale VARCHAR(5)
)