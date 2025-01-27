const express = require("express")
const dotenv = require("dotenv")
dotenv.config()
const cors = require("cors")
const connectDB = require("./db/connectDb")
const authRoutes = require("./routes/authRoutes")
const userRoutes = require("./routes/userRoutes")
const cookiesParser = require("cookie-parser")

const port = process.env.PORT
const app = express()

app.use(cors())
app.use(express.json())
app.use(express.urlencoded({ extended: false }))
app.use(cookiesParser())

app.use("/", authRoutes)
app.use("/user", userRoutes)

connectDB()

app.listen(port, () => {
  console.log(`Server running on Port: ${port}`)
})