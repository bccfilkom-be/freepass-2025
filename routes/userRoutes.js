const express = require("express")
const {
  updateUserProfile,
  getUserProfile,
  viewAllSessions,
  leaveFeedback,
  viewUserInfoByUsername,
} = require("../controllers/userController")
const authMiddleware = require("../middleware/authMiddleware")
const router = express.Router()

router.put("/profile", authMiddleware, updateUserProfile)
router.get("/profile/:userid", authMiddleware, getUserProfile)
router.get("/sessions", authMiddleware, viewAllSessions)
router.post("/feedback", authMiddleware, leaveFeedback)
router.get("/user/:username", authMiddleware, viewUserInfoByUsername)

module.exports = router
