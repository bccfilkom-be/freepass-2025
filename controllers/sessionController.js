const { v4: uuidv4 } = require("uuid")
const response = require("../response")
const { insertRecord, updateRecord, deleteRecord, getAllRecords, getRecordById, checkSessionAvailability, checkOverlappingRegistrations, createTable } = require("../utils/sqlFunctions")
const sessionSchema = require("../schemas/sessionSchema")
const sessionRegistrationSchema = require("../schemas/sessionRegistrationSchema")

const updateSession = async (req, res) => {
  const { sessionid } = req.params
  const updatedData = req.body

  try {
    await createTable(sessionSchema)
    const session = await getRecordById("sessions", "sessionid", sessionid)
    if (session.created_by !== req.user.userid) {
      response(403, "", "You can only update your own sessions", res)
      return
    }

    const updatedSession = await updateRecord("sessions", "sessionid", sessionid, updatedData)
    response(200, updatedSession, "Session updated successfully", res)
  } catch (error) {
    console.error("Error updating session:", error)
    response(500, "", error, res)
  }
}

const deleteSession = async (req, res) => {
  const { sessionid } = req.params

  try {
    await createTable(sessionSchema)
    const session = await getRecordById("sessions", "sessionid", sessionid)
    if (session.created_by !== req.user.userid) {
      response(403, "", "You can only delete your own sessions", res)
      return
    }

    await deleteRecord("sessions", "sessionid", sessionid)
    response(200, "", "Session deleted successfully", res)
  } catch (error) {
    console.error("Error deleting session:", error)
    response(500, "", error, res)
  }
}

const getAllSessions = async (req, res) => {
  try {
    await createTable(sessionSchema)
    const sessions = await getAllRecords("sessions")
    response(200, sessions, "Sessions fetched successfully", res)
  } catch (error) {
    console.error("Error fetching sessions:", error)
    response(500, "", error, res)
  }
}

const registerForSession = async (req, res) => {
  const { sessionid } = req.body
  const userid = req.user.userid

  try {
    await createTable(sessionRegistrationSchema)
    const hasOverlap = await checkOverlappingRegistrations(userid, sessionid)
    const isFull = await checkSessionAvailability(sessionid)

    if (hasOverlap) {
      response(400, "", "Cannot register for overlapping sessions", res)
    } else if (isFull) {
      response(400, "", "No available seats for this session", res)
    } else {
      const sessionRegistrationsData = {
        registrationid: uuidv4(),
        sessionid,
        userid 
      }
      const newRegistration = await insertRecord("session_registrations", sessionRegistrationsData)
      response(201, newRegistration, "Registered for session successfully", res)
    }
  } catch (error) {
    console.error("Error registering for session:", error)
    response(500, "", error, res)
  }
}

module.exports = {
  updateSession,
  deleteSession,
  getAllSessions,
  registerForSession,
}