CREATE TABLE IF NOT EXIST sessinos (
    id VARCHAR(26) PRIMARY KEY,
    token VARCHAR(26),
    createat bigint,
    expiresat bigint,
    lastactivityat bigint,
    userid VARCHAR(26)
    deviceid VARCHAR(512)
    isoath boolean
)