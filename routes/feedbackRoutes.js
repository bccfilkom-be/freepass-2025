const express = require("express")
const { leaveFeedback, deleteFeedback, getAllFeedback, getFeedbackBySession } = require("../controllers/feedbackController")
const authMiddleware = require("../middleware/authMiddleware")
const router = express.Router()

router.post("/", authMiddleware, leaveFeedback)
router.delete("/:feedbackid", authMiddleware, deleteFeedback)
router.get("/", getAllFeedback)
router.get("/session/:sessionid",  getFeedbackBySession)

module.exports = router
