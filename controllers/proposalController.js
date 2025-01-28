const { v4: uuidv4 } = require("uuid")
const response = require("../response")
const { insertRecord, updateRecord, deleteRecord, getRecordById, checkRecordExists, createTable } = require("../utils/sqlFunctions")
const sessionProposalSchema = require("../schemas/sessionProposalSchema")

const createProposal = async (req, res) => {
  const proposalData = {
    proposalid: uuidv4(),
    userid: req.user.userid,
    ...req.body,
    proposed_at: new Date(),
  }

  try {
    await createTable(sessionProposalSchema)
    const existingProposal = await checkRecordExists(
      "session_proposals",
      "userid",
      req.user.userid
    )

    if (existingProposal && existingProposal.status === 'pending') {
      response(400, "", "You can only have one pending session proposal at a time", res)
    } else {
      const newProposal = await insertRecord("session_proposals", proposalData)
      response(201, newProposal, "Session proposal created successfully", res)
    }
  } catch (error) {
    response(500, "", error, res)
  }
}

const getProposal = async (req, res) => {
  const { proposalid } = req.params

  try {
    await createTable(sessionProposalSchema)
    const proposal = await getRecordById("session_proposals", "proposalid", proposalid)
    if (proposal.userid !== req.user.userid) {
      response(403, "", "You can only view your own session proposals", res)
      return
    }
    response(200, proposal, "Session proposal details fetched successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const updateProposal = async (req, res) => {
  const { proposalid } = req.params
  const updatedData = req.body

  try {
    await createTable(sessionProposalSchema)
    const proposal = await getRecordById("session_proposals", "proposalid", proposalid)
    if (proposal.userid !== req.user.userid) {
      response(403, "", "You can only update your own session proposals", res)
      return
    }

    const updatedProposal = await updateRecord("session_proposals", "proposalid", proposalid, updatedData)
    response(200, updatedProposal, "Session proposal updated successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

const deleteProposal = async (req, res) => {
  const { proposalid } = req.params

  try {
    await createTable(sessionProposalSchema)
    const proposal = await getRecordById("session_proposals", "proposalid", proposalid)
    if (proposal.userid !== req.user.userid) {
      response(403, "", "You can only delete your own session proposals", res)
      return
    }

    await deleteRecord("session_proposals", "proposalid", proposalid)
    response(200, "", "Session proposal deleted successfully", res)
  } catch (error) {
    response(500, "", error, res)
  }
}

module.exports = {
  createProposal,
  getProposal,
  updateProposal,
  deleteProposal,
}
