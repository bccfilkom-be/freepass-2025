const jwt = require("jsonwebtoken")


const authMiddleware = (req, res, next) => {
  const token = req.cookies.token
  try {
    const user = jwt.verify(token, process.env.JWT_SECRET)
    req.user = user
    next()
  } catch (err) {
    console.log(err)
    res.clearCookie("token")
  }
}

module.exports = authMiddleware
