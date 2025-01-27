const sessionRegistrationSchema = 
  `CREATE TABLE IF NOT EXISTS session_registrations (
    registrationid VARCHAR PRIMARY KEY,
    sessionid VARCHAR NOT NULL,
    userid VARCHAR NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sessionid) REFERENCES sessions(sessionid),
    FOREIGN KEY (userid) REFERENCES users(userid)
  )`

module.exports = sessionRegistrationSchema
