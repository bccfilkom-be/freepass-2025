const { Router } = require("express");
const { getAllSessions } = require("../controllers/sessionController.js");

const router = Router();

router.get("/", getAllSessions);

module.exports = router;
