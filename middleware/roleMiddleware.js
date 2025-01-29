const response = require("../response")
const { getRecordById } = require("../utils/sqlFunctions")

const roleMiddleware = (requiredRole) => {
  return async (req, res, next) => {
    const user = await getRecordById("users", "userid", req.user.userid)
    const userRole = user.role

    if (userRole !== requiredRole) {
      return response(403, " ", "you are not authorized for this action", res)
    }

    next()
  }
}

module.exports = roleMiddleware
