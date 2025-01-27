const express = require("express")
const {
  updateSession,
  deleteSession,
  getAllSessions,
  registerForSession,
} = require("../controllers/sessionController")
const authMiddleware = require("../middleware/authMiddleware")
const router = express.Router()


router.put("/:sessionid", authMiddleware, updateSession)
router.delete("/:sessionid", authMiddleware, deleteSession)
router.get("/", authMiddleware, getAllSessions)
router.post("/register", authMiddleware, registerForSession)

module.exports = router
