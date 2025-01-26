const { Router } = require("express");
const {
  getAllSessions,
  leaveFeedback,
  registerForSession,
} = require("../controllers/sessionController.js");

const router = Router();

router.get("/", getAllSessions);
router.post("/:id/feedback", leaveFeedback);
router.post("/:id/register", registerForSession);

module.exports = router;
