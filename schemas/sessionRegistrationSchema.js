const sessionRegistrationSchema = 
  `CREATE TABLE IF NOT EXISTS session_registrations (
    registrationid VARCHAR PRIMARY KEY,
    sessionid VARCHAR NOT NULL,
    userid VARCHAR NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sessionid) REFERENCES sessions(sessionid),
    FOREIGN KEY (userid) REFERENCES users(userid),
    CONSTRAINT no_overlap_registration CHECK (
      NOT EXISTS (
        SELECT 1 FROM session_registrations sr
        JOIN sessions s1 ON sr.sessionid = s1.sessionid
        JOIN sessions s2 ON session_registrations.sessionid = s2.sessionid
        WHERE sr.userid = session_registrations.userid
        AND s1.start_time < s2.end_time
        AND s1.end_time > s2.start_time
      )
    )
  )`

module.exports = sessionRegistrationSchema
