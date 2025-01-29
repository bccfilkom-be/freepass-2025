const { Pool } = require('pg')
const config = require("../db/config")
const pool = new Pool(config)

const createTable = (schema) => {
  return new Promise((resolve, reject) => {
    pool.query(schema, (err, results) => {
      if (err) {
        console.error("Error creating table:", err)
        reject(err)
      } else {
        resolve(results)
      }
    })
  })
}

const checkRecordExists = (tableName, column, value) => {
  return new Promise((resolve, reject) => {
    const query = `SELECT * FROM ${tableName} WHERE ${column} = $1`

    pool.query(query, [value], (err, results) => {
      if (err) {
        console.error("Error checking record existence:", err)
        reject(err)
      } else {
        resolve(results.rows.length ? results.rows[0] : null)
      }
    })
  })
}

const insertRecord = (tableName, record) => {
  return new Promise((resolve, reject) => {
    const columns = Object.keys(record).join(', ')
    const placeholders = Object.keys(record).map((_, index) => `$${index + 1}`).join(', ')
    const values = Object.values(record)

    const query = `INSERT INTO ${tableName} (${columns}) VALUES (${placeholders}) RETURNING *`
    pool.query(query, values, (err, results) => {
      if (err) {
        console.error("Error inserting record:", err)
        reject(err)
      } else {
        resolve(results.rows[0])
      }
    })
  })
}

const updateRecord = (tableName, column, value, updatedData) => {
  return new Promise((resolve, reject) => {
    const setClause = Object.keys(updatedData).map((key, index) => `${key} = $${index + 2}`).join(', ')
    const values = [value, ...Object.values(updatedData)]

    const query = `UPDATE ${tableName} SET ${setClause} WHERE ${column} = $1 RETURNING *`
    pool.query(query, values, (err, results) => {
      if (err) {
        console.error("Error updating record:", err)
        reject(err)
      } else {
        resolve(results.rows[0])
      }
    })
  })
}

const getRecordById = (tableName, column, value) => {
  return new Promise((resolve, reject) => {
    const query = `SELECT * FROM ${tableName} WHERE ${column} = $1`
    pool.query(query, [value], (err, results) => {
      if (err) {
        console.error("Error getting record by ID:", err)
        reject(err)
      } else {
        resolve(results.rows[0])
      }
    })
  })
}

const deleteRecord = (tableName, column, value) => {
  return new Promise((resolve, reject) => {
    const query = `DELETE FROM ${tableName} WHERE ${column} = $1`
    pool.query(query, [value], (err, results) => {
      if (err) {
        console.error("Error deleting record:", err)
        reject(err)
      } else {
        resolve(results)
      }
    })
  })
}

const getAllRecords = (tableName) => {
  return new Promise((resolve, reject) => {
    const query = `SELECT * FROM ${tableName}`
    pool.query(query, (err, results) => {
      if (err) {
        console.error("Error getting all records:", err)
        reject(err)
      } else {
        resolve(results.rows)
      }
    })
  })
}

const checkSessionAvailability = (sessionid) => {
  return new Promise((resolve, reject) => {
    const query = `
      SELECT s.max_seats, COALESCE(COUNT(sr.sessionid), 0) as registration_count
      FROM sessions s
      LEFT JOIN session_registrations sr ON sr.sessionid = s.sessionid
      WHERE s.sessionid = $1
      GROUP BY s.max_seats
    `
    pool.query(query, [sessionid], (err, results) => {
      if (err) {
        console.error("Error checking session availability:", err)
        reject(err)
      } else {
        if (results.rows.length > 0) {
          const { registration_count, max_seats } = results.rows[0]
          resolve(registration_count >= max_seats)
        } else {
          resolve(false)
        }
      }
    })
  })
}

const checkOverlappingRegistrations = (userid, sessionid) => {
  return new Promise((resolve, reject) => {
    const query = `
      SELECT 1
      FROM session_registrations sr
      JOIN sessions s1 ON sr.sessionid = s1.sessionid
      JOIN sessions s2 ON s2.sessionid = $1
      WHERE sr.userid = $2
      AND s1.start_time < s2.end_time
      AND s1.end_time > s2.start_time
    `
    pool.query(query, [sessionid, userid], (err, results) => {
      if (err) {
        console.error("Error checking overlapping registrations:", err)
        reject(err)
      } else {
        resolve(results.rows.length > 0)
      }
    })
  })
}

const getFeedbackWithUsername = (sessionid) => {
  return new Promise((resolve, reject) => {
    const query = `
      SELECT f.feedbackid, f.sessionid, f.userid, u.username, f.comment, f.rating, f.created_at
      FROM feedback f
      JOIN users u ON f.userid = u.userid
      WHERE f.sessionid = $1
    `
    pool.query(query, [sessionid], (err, results) => {
      if (err) {
        console.error("Error retrieving feedback with username:", err)
        reject(err)
      } else {
        resolve(results.rows)
      }
    })
  })
}

module.exports = {
  createTable,
  checkRecordExists,
  insertRecord,
  updateRecord,
  getRecordById,
  deleteRecord,
  getAllRecords,
  checkSessionAvailability,
  checkOverlappingRegistrations,
  getFeedbackWithUsername,
}
