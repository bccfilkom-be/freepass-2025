const response = require("../response")
const { getAllRecords, updateRecord, deleteRecord, insertRecord, getRecordById, createTable } = require("../utils/sqlFunctions")
const sessionSchema = require("../schemas/sessionSchema")
const sessionProposalSchema = require("../schemas/sessionProposalSchema")
const { v4: uuidv4 } = require("uuid")

const getAllSessionProposals = async (req, res) => {
  try {
    await createTable(sessionProposalSchema)
    const proposals = await getAllRecords("session_proposals")
    response(200, proposals, "Session proposals fetched successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const acceptSessionProposal = async (req, res) => {
  const { proposalid } = req.params

  try {
    await createTable(sessionProposalSchema)
    await createTable(sessionSchema)
    const updatedProposal = await updateRecord("session_proposals", "proposalid", proposalid, { status: "accepted" })
    
    const proposal = await getRecordById("session_proposals", "proposalid", proposalid)
    const sessionData = {
      sessionid: uuidv4(),
      title: proposal.title,
      description: proposal.description,
      speaker: proposal.speaker,
      start_time: proposal.start_time,
      end_time: proposal.end_time,
      max_seats: proposal.max_seats,
      created_by: proposal.userid,
      created_at: new Date()
    }
    const newSession = await insertRecord("sessions", sessionData)
    
    response(200, { updatedProposal, newSession }, "Session proposal accepted and session created successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const rejectSessionProposal = async (req, res) => {
  const { proposalid } = req.params

  try {
    await createTable(sessionProposalSchema)
    const updatedProposal = await updateRecord("session_proposals", "proposalid", proposalid, { status: "rejected" })
    response(200, updatedProposal, "Session proposal rejected successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const removeSession = async (req, res) => {
  const { sessionid } = req.params

  try {
    await createTable(sessionSchema)
    await deleteRecord("sessions", "sessionid", sessionid)
    response(200, "", "Session removed successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const removeFeedback = async (req, res) => {
  const { feedbackid } = req.params

  try {
    await createTable("feedback")
    await deleteRecord("feedback", "feedbackid", feedbackid)
    response(200, "", "Feedback removed successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

module.exports = {
  getAllSessionProposals,
  acceptSessionProposal,
  rejectSessionProposal,
  removeSession,
  removeFeedback,
}
