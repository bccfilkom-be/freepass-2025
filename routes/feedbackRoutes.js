const express = require("express")
const { leaveFeedback, deleteFeedback, getAllFeedback } = require("../controllers/feedbackController")
const authMiddleware = require("../middleware/authMiddleware")
const router = express.Router()

router.post("/", authMiddleware, leaveFeedback)
router.delete("/:feedbackid", authMiddleware, deleteFeedback)
router.get("/", authMiddleware, getAllFeedback)

module.exports = router
