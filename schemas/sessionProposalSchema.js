const sessionProposalSchema = 
  `CREATE TABLE IF NOT EXISTS session_proposals (
    proposalid VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT,
    speaker VARCHAR,
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    max_seats INT NOT NULL,
    proposed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) DEFAULT 'pending',
    userid VARCHAR NOT NULL,
    FOREIGN KEY (userid) REFERENCES users(userid)
  )`

module.exports = sessionProposalSchema
