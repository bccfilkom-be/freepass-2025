const { Pool } = require("pg")
const config = require("./config")

const connectDB = async () => {
  const pool = new Pool(config)

  try {
    const client = await pool.connect() 
    console.log("Connected to database")
    client.release() 
  } catch (err) {
    console.error({ err }) 
  }
}

module.exports = connectDB
