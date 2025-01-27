const feedbackSchema = 
  `CREATE TABLE IF NOT EXISTS feedback (
    feedbackid VARCHAR PRIMARY KEY,
    sessionid VARCHAR NOT NULL,
    userid VARCHAR NOT NULL,
    comment TEXT,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sessionid) REFERENCES sessions(sessionid),
    FOREIGN KEY (userid) REFERENCES users(userid)
  )`

module.exports = feedbackSchema
