const response = require("../response")
const { v4: uuidv4 } = require("uuid")
const { insertRecord, deleteRecord, getAllRecords, getFeedbackWithUsername, createTable } = require("../utils/sqlFunctions")
const feedbackSchema = require("../schemas/feedbackSchema")

const leaveFeedback = async (req, res) => {
  const feedbackData = {
    feedbackid: uuidv4(),
    userid: req.user.userid,
    ...req.body
  }

  try {
    await createTable(feedbackSchema)
    const newFeedback = await insertRecord("feedback", feedbackData)
    response(201, newFeedback, "Feedback submitted successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const deleteFeedback = async (req, res) => {
  const { feedbackid } = req.params

  try {
    await createTable(feedbackSchema)
    await deleteRecord("feedback", "feedbackid", feedbackid)
    response(200, "", "Feedback deleted successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const getAllFeedback = async (req, res) => {
  try {
    await createTable(feedbackSchema)
    const feedback = await getAllRecords("feedback")
    response(200, feedback, "Feedback fetched successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const getFeedbackBySession = async (req, res) => {
  const { sessionid } = req.params

  try {
    await createTable(feedbackSchema)
    const feedbackList = await getFeedbackWithUsername(sessionid)
    response(200, feedbackList, "Feedback for session fetched successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

module.exports = {
  leaveFeedback,
  deleteFeedback,
  getAllFeedback,
  getFeedbackBySession,
}
