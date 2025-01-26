const { Router } = require("express");
const { getAllSessions } = require("../controllers/sessionController.js");
const { leaveFeedback } = require("../controllers/sessionController.js");

const router = Router();

router.get("/", getAllSessions);

router.post("/:id/feedback", leaveFeedback);

module.exports = router;
