const { Router } = require("express");
const authRoutes = require("./authRoutes.js");
const userRoutes = require("./userRoutes.js");
const sessionRoutes = require("./sessionRoutes.js");
const sessionProposalRoutes = require("./sessionProposalRoutes.js");
const adminRoutes = require("./adminRoutes.js");

const router = Router();

router.use("/auth", authRoutes);
router.use("/users", userRoutes);
router.use("/sessions", sessionRoutes);
router.use("/session-proposals", sessionProposalRoutes);
router.use("/admin", adminRoutes);

module.exports = router;
