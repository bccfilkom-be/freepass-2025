const express = require("express")
const {
  getAllSessionProposals,
  acceptSessionProposal,
  rejectSessionProposal,
  removeSession,
  removeFeedback,
} = require("../controllers/eventCoordinatorController")
const authMiddleware = require("../middleware/authMiddleware")
const roleMiddleware = require("../middleware/roleMiddleware")
const router = express.Router()

router.get("/proposals", authMiddleware, roleMiddleware("event_coordinator"), getAllSessionProposals)
router.put("/proposals/accept/:proposalid", authMiddleware, roleMiddleware("event_coordinator"), acceptSessionProposal)
router.put("/proposals/reject/:proposalid", authMiddleware, roleMiddleware("event_coordinator"), rejectSessionProposal)
router.delete("/sessions/:sessionid", authMiddleware, roleMiddleware("event_coordinator"), removeSession)
router.delete("/feedback/:feedbackid", authMiddleware, roleMiddleware("event_coordinator"), removeFeedback)

module.exports = router
