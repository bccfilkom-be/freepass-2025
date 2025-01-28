const { Router } = require("express");
const {
  getAllSessions,
  leaveFeedback,
  deleteFeedback,
  registerForSession,
  editSession,
  deleteSession,
} = require("../controllers/sessionController.js");
const { authMiddleware } = require("../middlewares/authMiddleware.js");

const router = Router();

router.get("/", getAllSessions);
router.post("/:id/feedback", leaveFeedback);
router.delete("/:id/feedback/:feedbackId", authMiddleware, deleteFeedback);
router.post("/:id/register", registerForSession);
router.put("/:id", editSession);
router.delete("/:id", authMiddleware, deleteSession);

module.exports = router;
