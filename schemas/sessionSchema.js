const sessionSchema = 
  `CREATE TABLE IF NOT EXISTS sessions (
    sessionid VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT,
    speaker VARCHAR,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    max_seats INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  )`

module.exports = sessionSchema
