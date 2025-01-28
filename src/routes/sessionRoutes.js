const { Router } = require("express");
const {
  getAllSessions,
  leaveFeedback,
  registerForSession,
  editSession,
  deleteSession,
} = require("../controllers/sessionController.js");

const router = Router();

router.get("/", getAllSessions);
router.post("/:id/feedback", leaveFeedback);
router.post("/:id/register", registerForSession);
router.put("/:id", editSession);
router.delete("/:id", deleteSession);

module.exports = router;
