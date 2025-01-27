const { Pool } = require('pg')
const config = require("../db/config")
const pool = new Pool(config)

const createTable = (schema) => {
  return new Promise((resolve, reject) => {
    pool.query(schema, (err, results) => {
      if (err) {
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
        reject(err)
      } else {
        resolve(results.rows[0])
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
}
