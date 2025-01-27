const eventCoordinatorSchema = 
  `CREATE TABLE IF NOT EXISTS event_coordinators (
    coordinatorid VARCHAR PRIMARY KEY,
    userid VARCHAR NOT NULL,
    FOREIGN KEY (userid) REFERENCES users(userid)
  )`

module.exports = eventCoordinatorSchema
