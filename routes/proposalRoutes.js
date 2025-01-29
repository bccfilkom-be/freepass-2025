const express = require("express")
const {
  createProposal,
  getProposal,
  updateProposal,
  deleteProposal,
} = require("../controllers/proposalController")
const authMiddleware = require("../middleware/authMiddleware")
const router = express.Router()

router.post("/", authMiddleware, createProposal)
router.get("/:proposalid", authMiddleware, getProposal)
router.put("/:proposalid", authMiddleware, updateProposal)
router.delete("/:proposalid", authMiddleware, deleteProposal)

module.exports = router
