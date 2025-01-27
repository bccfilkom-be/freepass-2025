const adminSchema = 
  `CREATE TABLE IF NOT EXISTS admins (
    adminid VARCHAR PRIMARY KEY,
    userid VARCHAR NOT NULL,
    FOREIGN KEY (userid) REFERENCES users(userid)
  )`

module.exports = adminSchema
