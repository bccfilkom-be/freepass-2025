const sessionSchema = 
  `CREATE TABLE IF NOT EXISTS sessions (
    sessionid VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT,
    speaker VARCHAR,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    max_seats INT NOT NULL,
    created_by VARCHAR ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users(userid)
  )`

module.exports = sessionSchema
