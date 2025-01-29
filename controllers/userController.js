const response = require("../response")
const { updateRecord, getRecordById, checkRecordExists } = require("../utils/sqlFunctions")

const updateUserProfile = async (req, res) => {
  const { userid } = req.user
  const updatedData = req.body

  try {
    const updatedUser = await updateRecord("users", "userid", userid, updatedData)
    response(200, updatedUser, "Profile updated successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const getUserProfile = async (req, res) => {
  const { userid } = req.params

  try {
    const user = await getRecordById("users", "userid", userid)
    response(200, user, "User profile fetched successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}


const viewUserInfoByUsername = async (req, res) => {
  const { username } = req.params

  try {
    const user = await checkRecordExists("users", "username", username)
    if (user) {
      response(200, user, "User information fetched successfully", res)
    } else {
      response(404, "", "User not found", res)
    }
  } catch (error) {
    response(500, "", error, res)
  }
}


module.exports = {
  updateUserProfile,
  getUserProfile,
  viewUserInfoByUsername,
}
