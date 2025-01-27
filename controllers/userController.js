const { v4: uuidv4 } = require("uuid")
const response = require("../response")
const { updateRecord, getRecordById, getAllRecords, insertRecord } = require("../utils/sqlFunctions")

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

const viewAllSessions = async (req, res) => {
  try {
    const sessions = await getAllRecords("sessions")
    response(200, sessions, "Sessions fetched successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const leaveFeedback = async (req, res) => {
  const feedbackData = {
    feedbackid: uuidv4(),
    userid: req.user.userid,
    ...req.body
  }

  try {
    const newFeedback = await insertRecord("feedback", feedbackData)
    response(201, newFeedback, "Feedback submitted successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

module.exports = {
  updateUserProfile,
  getUserProfile,
  viewAllSessions,
  leaveFeedback,
}
