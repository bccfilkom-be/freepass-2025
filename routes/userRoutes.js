const express = require("express")
const {
  updateUserProfile,
  getUserProfile,
  viewUserInfoByUsername,
} = require("../controllers/userController")
const authMiddleware = require("../middleware/authMiddleware")
const router = express.Router()

router.put("/profile", authMiddleware, updateUserProfile)
router.get("/profile/:userid", authMiddleware, getUserProfile)
router.get("/user/:username", authMiddleware, viewUserInfoByUsername)

module.exports = router
